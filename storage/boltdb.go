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
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const bucketName = "abraracourcix"

// BoltDB represents a storage using the BoltDB database
type BoltDB struct {
	*bolt.DB
	BucketName string
	Path       string
}

// NewBoltDB opens a new BoltDB connection to the specified path and bucket
func NewBoltDB(path string) (*BoltDB, error) {
	log.Printf("[DEBUG] [abraracourcix] Init BoltDB storage : %v", path)
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("Can't create BoltDB bucket: %s", err)
		}
		return nil
	})
	return &BoltDB{DB: db, Path: path}, nil
}

// Get a value given its key
func (db *BoltDB) Get(key []byte) ([]byte, error) {
	log.Printf("[DEBUG] [abraracourcix] Search entry with key : %v", string(key))
	var value []byte
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		b.ForEach(func(k, v []byte) error {
			// log.Printf("[BoltDB] Entry : %s %s", string(k), string(v))
			if string(k) == string(key) {
				log.Printf("[INFO] [abraracourcix] Find : %s",
					string(v))
				value = v
			}
			return nil
		})
		return nil
	})
	return value, nil
}

// Put a value at the specified key
func (db *BoltDB) Put(key []byte, value []byte) error {
	log.Printf("[DEBUG] [abraracourcix] Put : %v %v", string(key), string(value))
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		b.Put(key, value)
		return nil
	})
	return nil
}

// Delete the value at the specified key
func (db *BoltDB) Delete(key []byte) error {
	log.Printf("[DEBUG] [abraracourcix] Delete : %v", string(key))
	return ErrNotImplemented
}

// Close the store connection
func (db *BoltDB) Close() {
	log.Printf("[DEBUG] [abraracourcix] Close")
}

// Print backend informations
func (db *BoltDB) Print() {
	log.Printf("[DEBUG] [abraracourcix] storage backend: %s", BOLTDB)
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		b.ForEach(func(key, value []byte) error {
			log.Println(string(key), string(value))
			return nil
		})
		return nil
	})
}
