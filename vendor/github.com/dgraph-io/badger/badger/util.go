/*
 * Copyright 2017 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package badger

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"sync/atomic"

	"github.com/dgraph-io/badger/table"
	"github.com/dgraph-io/badger/y"
	"github.com/pkg/errors"
)

// summary is produced when DB is closed. Currently it is used only for testing.
type summary struct {
	fileIDs map[uint64]bool
}

func (s *levelsController) getSummary() *summary {
	out := &summary{
		fileIDs: make(map[uint64]bool),
	}
	for _, l := range s.levels {
		l.getSummary(out)
	}
	return out
}

func (s *levelHandler) getSummary(sum *summary) {
	s.RLock()
	defer s.RUnlock()
	for _, t := range s.tables {
		sum.fileIDs[t.ID()] = true
	}
}

func (s *KV) validate() { s.lc.validate() }

func (s *levelsController) validate() error {
	for _, l := range s.levels {
		if err := l.validate(); err != nil {
			return errors.Wrap(err, "Levels Controller")
		}
	}
	return nil
}

// Check does some sanity check on one level of data or in-memory index.
func (s *levelHandler) validate() error {
	if s.level == 0 {
		return nil
	}

	s.RLock()
	defer s.RUnlock()
	numTables := len(s.tables)
	for j := 1; j < numTables; j++ {
		if j >= len(s.tables) {
			return errors.Errorf("Level %d, j=%d numTables=%d", s.level, j, numTables)
		}

		if bytes.Compare(s.tables[j-1].Biggest(), s.tables[j].Smallest()) >= 0 {
			return errors.Errorf(
				"Inter: %s vs %s: level=%d j=%d numTables=%d",
				string(s.tables[j-1].Biggest()), string(s.tables[j].Smallest()), s.level, j, numTables)
		}

		if bytes.Compare(s.tables[j].Smallest(), s.tables[j].Biggest()) > 0 {
			return errors.Errorf(
				"Intra: %q vs %q: level=%d j=%d numTables=%d",
				s.tables[j].Smallest(), s.tables[j].Biggest(), s.level, j, numTables)
		}
	}
	return nil
}

// func (s *KV) debugPrintMore() { s.lc.debugPrintMore() }

// // debugPrintMore shows key ranges of each level.
// func (s *levelsController) debugPrintMore() {
// 	s.Lock()
// 	defer s.Unlock()
// 	for i := 0; i < s.kv.opt.MaxLevels; i++ {
// 		s.levels[i].debugPrintMore()
// 	}
// }

// func (s *levelHandler) debugPrintMore() {
// 	s.RLock()
// 	defer s.RUnlock()
// 	s.elog.Printf("Level %d:", s.level)
// 	for _, t := range s.tables {
// 		y.Printf(" [%s, %s]", t.Smallest(), t.Biggest())
// 	}
// 	y.Printf("\n")
// }

// reserveFileIDs reserve k fileIDs. Returns pair is a half-interval.
// If we return [3, 6), it means use 3, 4, 5.
func (s *levelsController) reserveFileIDs(k int) (uint64, uint64) {
	id := atomic.AddUint64(&s.maxFileID, uint64(k))
	return id - uint64(k), id
}

func getIDMap(dir string) map[uint64]struct{} {
	fileInfos, err := ioutil.ReadDir(dir)
	y.Check(err)
	idMap := make(map[uint64]struct{})
	for _, info := range fileInfos {
		if info.IsDir() {
			continue
		}
		fileID, ok := table.ParseFileID(info.Name())
		if !ok {
			continue
		}
		idMap[fileID] = struct{}{}
	}
	return idMap
}

// mod65535 mods by 65535 fast.
func mod65535(a uint32) uint32 {
	a = (a >> 16) + (a & 0xFFFF) /* sum base 2**16 digits */
	if a < 65535 {
		return a
	}
	if a < (2 * 65535) {
		return a - 65535
	}
	return a - (2 * 65535)
}

func newCASCounter() uint16 {
	return uint16(1 + mod65535(rand.Uint32()))
}
