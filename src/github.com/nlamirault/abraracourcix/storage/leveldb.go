// Copyright (C) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

// LevelDB represents a storage using the BoltDB database
type LevelDB struct {
	*leveldb.DB
	Path string
}

// NewLevelDB opens a new LevelDB connection to the specified path
func NewLevelDB(path string) (*LevelDB, error) {
	log.Printf("[DEBUG] [abraracourcix] Init LevelDB storage : %s", path)
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}
	return &LevelDB{DB: db, Path: path}, nil
}

// Get a value given its key
func (db *LevelDB) Get(key []byte) ([]byte, error) {
	log.Printf("[DEBUG] [abraracourcix] Get : %v", string(key))
	value, err := db.DB.Get(key, nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err

	}
	log.Printf("[DEBUG] [abraracourcix] Find value : %s", value)
	return value, nil
}

// Put a value at the specified key
func (db *LevelDB) Put(key []byte, value []byte) error {
	log.Printf("[DEBUG] [abraracourcix] Put : %v %v",
		string(key), string(value))
	err := db.DB.Put(key, value, nil)
	return err
}

// Delete the value at the specified key
func (db *LevelDB) Delete(key []byte) error {
	log.Printf("[DEBUG] [abraracourcix] Delete : %v", string(key))
	return db.DB.Delete(key, nil)
}

// Close the store connection
func (db *LevelDB) Close() {
	log.Printf("[DEBUG] [abraracourcix] Close")
	db.DB.Close()
}

// Print backend informations
func (db *LevelDB) Print() {
	log.Printf("[DEBUG] [abraracourcix] Print database")
	iter := db.DB.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		log.Printf("[DEBUG] [abraracourcix] [%X]:\t[%X]\n", key, value)
	}
}
