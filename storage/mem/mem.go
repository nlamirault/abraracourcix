// Copyright (C) 2015, 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memdb

import (
	"github.com/golang/glog"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/storage"
)

const (
	label = "memdb"
)

type memDB struct {
	db map[string][]byte
}

func newMemDBStorage(conf *config.Configuration) (storage.Storage, error) {
	return &memDB{}, nil
}

func (memDB *memDB) Name() string {
	return label
}

func (memDB *memDB) Init() error {
	glog.V(1).Infof("Initialize")
	memDB.db = make(map[string][]byte)
	return nil
}

func (memDB *memDB) List() ([][]byte, error) {
	glog.V(1).Infof("Initialize all keys")
	keys := make([][]byte, 0, len(memDB.db))
	for key := range memDB.db {
		keys = append(keys, []byte(key))
	}
	return keys, nil
}

func (memDB *memDB) Get(key []byte) ([]byte, error) {
	glog.V(1).Infof("Search entry with key : %v", string(key))
	return memDB.db[string(key)], nil
}

func (memDB *memDB) Put(key []byte, value []byte) (err error) {
	glog.V(1).Infof("Put : %v %v", string(key), string(value))
	memDB.db[string(key)] = value
	return nil
}

func (memDB *memDB) Delete(key []byte) (err error) {
	glog.V(1).Infof("Delete entry with key : %v", string(key))
	delete(memDB.db, string(key))
	return nil
}

func (memDB *memDB) Close() error {
	glog.V(1).Infof("Close")
	memDB.db = nil
	return nil
}
