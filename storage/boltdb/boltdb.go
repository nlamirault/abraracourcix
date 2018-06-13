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

package boltdb

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/golang/glog"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/storage"
)

const (
	label      = "boltdb"
	bucketName = "abraracourcix"
)

type boltDB struct {
	db     *bolt.DB
	bucket string
	path   string
}

func init() {
	storage.RegisterStorage(label, newBoltdbStorage)
}

func newBoltdbStorage(conf *config.Configuration) (storage.Storage, error) {
	glog.V(1).Infof("Create storage using BoltDB : %s", conf.Storage)
	// db, err := bolt.Open(conf.Storage.BoltDB.File, 0600, nil)
	// if err != nil {
	// 	return nil, err
	// }
	return &boltDB{
		// db:     db,
		bucket: conf.Storage.BoltDB.Bucket,
		path:   conf.Storage.BoltDB.File,
	}, nil
}

func (boltDB *boltDB) Name() string {
	return label
}

func (boltDB *boltDB) Init() error {
	glog.V(1).Info("Initialize")
	db, err := bolt.Open(boltDB.path, 0600, nil)
	if err != nil {
		return err
	}
	boltDB.db = db
	return boltDB.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(boltDB.bucket))
		if err != nil {
			return fmt.Errorf("Can't create BoltDB bucket: %s", err)
		}
		return nil
	})
}

func (boltDB *boltDB) List() ([][]byte, error) {
	glog.V(1).Info("List all URLs")
	urls := [][]byte{}
	err := boltDB.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(boltDB.bucket))
		b.ForEach(func(k, v []byte) error {
			glog.V(3).Infof("Entry : %s %s", string(k), string(v))
			urls = append(urls, k)
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	glog.V(1).Infof("URLs: %s", urls)
	return urls, nil
}

func (boltDB *boltDB) Get(key []byte) ([]byte, error) {
	glog.V(1).Infof("Search entry with key : %v", string(key))
	var value []byte
	err := boltDB.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(boltDB.bucket))
		value = b.Get(key)
		if value != nil {
			glog.V(2).Infof("Find : %s", string(value))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (boltDB *boltDB) Put(key []byte, value []byte) error {
	glog.V(1).Infof("Put : %v %v", string(key), string(value))
	return boltDB.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(boltDB.bucket))
		return b.Put(key, value)
	})
}

func (boltDB *boltDB) Delete(key []byte) error {
	glog.V(1).Infof("Delete : %v", string(key))
	return boltDB.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(boltDB.bucket))
		return b.Delete(key)
	})
}

func (boltDB *boltDB) Close() error {
	glog.V(1).Infof("Close")
	return boltDB.db.Close()
}
