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

package storage

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
	db, err := bolt.Open(conf.Storage.BoltDB.File, 0600, nil)
	if err != nil {
		return nil, err
	}
	return &boltDB{
		db:     db,
		bucket: conf.Storage.BoltDB.Bucket,
		path:   conf.Storage.BoltDB.File,
	}, nil
}

func (boltDB *boltDB) Name() string {
	return label
}

func (boltDB *boltDB) Init() error {
	return boltDB.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(boltDB.bucket))
		if err != nil {
			return fmt.Errorf("Can't create BoltDB bucket: %s", err)
		}
		return nil
	})
}

func (boltDB *boltDB) List() ([][]byte, error) {
	glog.V(1).Infof("List all URLs")
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
	err := boltDB.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(boltDB.bucket))
		err := b.ForEach(func(k, v []byte) error {
			glog.V(3).Infof("Entry : %s %s", string(k), string(v))
			if string(k) == string(key) {
				glog.V(2).Infof("Find : %s", string(v))
				value = v
			}
			return nil
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (boltDB *boltDB) Put(key []byte, value []byte) error {
	glog.V(1).Infof("Put : %v %v", string(key), string(value))
	err := boltDB.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(boltDB.bucket))
		err := b.Put(key, value)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (boltDB *boltDB) Delete(key []byte) error {
	glog.V(1).Infof("Delete : %v", string(key))
	return storage.ErrNotImplemented
}

func (boltDB *boltDB) Close() error {
	glog.V(1).Infof("Close")
	return nil
}

func (boltDB *boltDB) Print() error {
	glog.V(1).Infof("Storage backend: %s", label)
	return boltDB.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(boltDB.bucket))
		err := b.ForEach(func(key, value []byte) error {
			fmt.Printf("%s %s", string(key), string(value))
			return nil
		})
		return err
	})
}
