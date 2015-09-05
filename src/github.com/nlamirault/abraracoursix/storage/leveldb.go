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
	// "fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDB struct {
	*leveldb.DB
}

func NewDatabase(path string) (*LevelDB, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDB{DB: db}, nil
}

func (db *LevelDB) Get(key []byte) (value []byte, err error) {
	return db.DB.Get(key, nil)
}

func (db *LevelDB) Put(key, value []byte) (err error) {
	err = db.DB.Put(key, value, nil)
	return err
}

func (db *LevelDB) Delete(key []byte) (err error) {
	return db.DB.Delete(key, nil)
}

func (db *LevelDB) Close() {
	db.DB.Close()
}

func (db *LevelDB) Print() {
	log.Info("Print database")
	iter := db.DB.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		log.Debug("[%X]:\t[%X]\n", key, value)
	}
}
