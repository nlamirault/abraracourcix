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
	// "fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/nlamirault/abraracourcix/config"
)

// tempdir returns a temporary directory path.
func tempdir() (string, error) {
	d, _ := ioutil.TempDir("", "leveldb-")
	err := os.Remove(d)
	if err != nil {
		return "", err
	}
	return d, nil
}

func getLevelDBConfiguration() (*config.Configuration, error) {
	td, err := tempdir()
	if err != nil {
		return nil, err
	}
	return &config.Configuration{
		Storage: &config.StorageConfiguration{
			Name: "leveldb",
			LevelDB: &config.LevelDBConfiguration{
				Path: td,
			},
		},
	}, nil
}

// Ensure that gets a non-existent key returns nil.
func TestLevelDB_Get_NonExistent(t *testing.T) {
	conf, err := getLevelDBConfiguration()
	if err != nil {
		t.Fatalf("Can't create configuration")
	}
	db, err := newLevelDBStorage(conf)
	if err != nil {
		t.Fatalf("Can't create LevelDB test database.")
	}
	//defer db.Close()
	defer func() {
		err := db.Close()
		if err != nil {
			t.Fatalf("Can't close LevelDB test database.")
		}
	}()
	value, err := db.Get([]byte("foo"))
	if err != nil {
		t.Fatalf("Can't retrieve LevelDB key.")
	}
	// fmt.Println("Value: ", string(value))
	if value != nil {
		t.Fatalf("Error retrieve invalid key.")
	}
}

// Ensure that that gets an existent key returns value.
func TestLevelDBDB_Get_Existent(t *testing.T) {
	conf, err := getLevelDBConfiguration()
	if err != nil {
		t.Fatalf("Can't create configuration")
	}
	db, err := newLevelDBStorage(conf)
	if err != nil {
		t.Fatalf("Can't create LevelDB test database.")
	}
	// td, err := tempdir()
	// if err != nil {
	// 	t.Fatalf("Can't create temporary directory: %v", err)
	// }
	// db, err := NewLevelDB(td)
	// if err != nil {
	// 	t.Fatalf("Can't create LevelDB test database.")
	// }
	//defer db.Close()
	defer func() {
		err := db.Close()
		if err != nil {
			t.Fatalf("Can't close LevelDB test database.")
		}
	}()
	err = db.Put([]byte("foo"), []byte("bar"))
	if err != nil {
		t.Fatalf("Can't store LevelDB key: %v", err)
	}
	value, err := db.Get([]byte("foo"))
	if err != nil {
		t.Fatalf("Can't retrieve LevelDB key: %v", err)
	}
	// fmt.Println("Value: ", string(value))
	if string(value) != "bar" {
		t.Fatalf("Error retrieve invalid value.")
	}
}
