// Copyright (C) 2015-2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package leveldb

import (
	"github.com/golang/glog"
	goleveldb "github.com/syndtr/goleveldb/leveldb"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/storage"
)

const (
	label = "leveldb"
)

// LevelDB represents a storage using the BoltDB database
type levelDB struct {
	db   *goleveldb.DB
	path string
}

func init() {
	storage.RegisterStorage(label, newLevelDBStorage)
}

func newLevelDBStorage(conf *config.Configuration) (storage.Storage, error) {
	glog.V(1).Infof("Create storage using LevelDB : %s", conf.Storage)
	// db, err := goleveldb.OpenFile(conf.Storage.LevelDB.Path, nil)
	// if err != nil {
	// 	return nil, err
	// }
	return &levelDB{
		// db:   db,
		path: conf.Storage.LevelDB.Path,
	}, nil
}

func (levelDB *levelDB) Name() string {
	return label
}

func (levelDB *levelDB) Init() error {
	db, err := goleveldb.OpenFile(levelDB.path, nil)
	if err != nil {
		return err
	}
	levelDB.db = db
	return nil
}

func (levelDB *levelDB) List() ([][]byte, error) {
	glog.V(1).Info("List all URLs")
	iter := levelDB.db.NewIterator(nil, nil)
	content := [][]byte{}
	for iter.Next() {
		key := iter.Key()
		content = append(content, key)
	}
	return content, nil
}

func (levelDB *levelDB) Get(key []byte) ([]byte, error) {
	glog.V(1).Infof("Search entry with key : %v", string(key))
	value, err := levelDB.db.Get(key, nil)
	if err != nil {
		if err == goleveldb.ErrNotFound {
			return nil, nil
		}
		return nil, err

	}
	glog.V(2).Infof("Find: %s", value)
	return value, nil
}

func (levelDB *levelDB) Put(key []byte, value []byte) error {
	glog.V(1).Infof("Put : %v %v", string(key), string(value))
	return levelDB.db.Put(key, value, nil)
}

func (levelDB *levelDB) Delete(key []byte) error {
	glog.V(1).Infof("Put : %v", string(key))
	return levelDB.db.Delete(key, nil)
}

func (levelDB *levelDB) Close() error {
	glog.V(1).Info("Close")
	return levelDB.db.Close()
}
