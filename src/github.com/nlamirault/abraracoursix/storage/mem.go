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
)

type MemDB struct {
	db map[string][]byte
}

func NewMemDB(path string) (*MemDB, error) {
	database := &MemDB{db: make(map[string][]byte)}
	return database, nil
}

func (db *MemDB) Get(key []byte) (value []byte, err error) {
	return db.db[string(key)], nil
}

func (db *MemDB) Put(key []byte, value []byte) (err error) {
	db.db[string(key)] = value
	return nil
}

func (db *MemDB) Delete(key []byte) (err error) {
	delete(db.db, string(key))
	return nil
}

func (db *MemDB) Close() {
	db = nil
}

func (db *MemDB) Print() {
	for key, value := range db.db {
		fmt.Printf("[%X]:\t[%X]\n", []byte(key), value)
	}
}
