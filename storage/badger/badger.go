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

package badger

import (
	"os"

	dbadger "github.com/dgraph-io/badger/badger"
	"github.com/golang/glog"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/storage"
)

const (
	label = "badger"
	mode  = 0755
)

type badger struct {
	kv   *dbadger.KV
	path string
}

func init() {
	storage.RegisterStorage(label, newBadgerStorage)
}

func newBadgerStorage(conf *config.Configuration) (storage.Storage, error) {
	glog.V(1).Infof("Create storage using Badger : %s", conf.Storage)
	return &badger{
		// kv: kv,
		path: conf.Storage.Badger.Path,
	}, nil

}

func (badgerDB *badger) Name() string {
	return label
}

func (badgerDB *badger) Init() error {
	glog.V(1).Info("Initialize")
	opt := dbadger.DefaultOptions
	opt.Dir = badgerDB.path
	opt.ValueDir = badgerDB.path
	if err := ensureDirectoriesExists([]string{
		opt.Dir,
		opt.ValueDir,
	}); err != nil {
		return err
	}
	kv, err := dbadger.NewKV(&opt)
	if err != nil {
		return err
	}
	badgerDB.kv = kv
	return nil
}

func (badgerDB *badger) List() ([][]byte, error) {
	glog.V(1).Infof("List all URLs")
	itrOpt := dbadger.IteratorOptions{
		PrefetchSize: 1000,
		FetchValues:  true,
		Reverse:      false,
	}
	urls := [][]byte{}
	itr := badgerDB.kv.NewIterator(itrOpt)
	for itr.Rewind(); itr.Valid(); itr.Next() {
		item := itr.Item()
		urls = append(urls, item.Key())
	}
	return urls, nil
}

func (badgerDB *badger) Get(key []byte) ([]byte, error) {
	glog.V(1).Infof("Search entry with key : %v", string(key))
	var item dbadger.KVItem
	if err := badgerDB.kv.Get(key, &item); err != nil {
		return nil, err
	}
	if item.Value() == nil {
		return nil, nil
	}
	return item.Value(), nil
}

func (badgerDB *badger) Put(key []byte, value []byte) error {
	glog.V(1).Infof("Put : %v %v", string(key), string(value))
	return badgerDB.kv.Set(key, value)
}

func (badgerDB *badger) Delete(key []byte) error {
	glog.V(1).Infof("Delete : %v", string(key))
	return badgerDB.kv.Delete(key)
}

func (badgerDB *badger) Close() error {
	glog.V(1).Infof("Close")
	return badgerDB.kv.Close()
}

func ensureDirectoriesExists(directories []string) error {
	for _, dir := range directories {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return os.Mkdir(dir, mode)
		}
	}
	return nil
}
