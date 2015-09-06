// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

const bucketName = "abraracoursix"

type BoltDB struct {
	*bolt.DB
	Path string
}

// NewBoltDB opens a new BoltDB connection to the specified path and bucket
func NewBoltDB(path string) (*BoltDB, error) {
	log.Debug("Init BoltDB storage : ", path)
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("Can't create BoltDB bucket: %s", err)
		}
		return nil
	})
	return &BoltDB{DB: db, Path: path}, nil
}

func (db *BoltDB) Get(key []byte) ([]byte, error) {
	log.Debugf("[boltdb] Get : %v", string(key))
	return nil, ErrNotImplemented
}

func (db *BoltDB) Put(key []byte, value []byte) error {
	log.Debugf("[boltdb] Put : %v %v", string(key), string(value))
	return ErrNotImplemented
}

func (db *BoltDB) Delete(key []byte) error {
	log.Debugf("[boltdb] Delete : %v", string(key))
	return ErrNotImplemented
}

func (db *BoltDB) Close() {
	log.Debug("[boltdb] Close")
}

func (db *BoltDB) Print() {
	log.Debug("[boltdb] BoltDB storage backend : ")
}
