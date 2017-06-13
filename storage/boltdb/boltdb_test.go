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

	// "github.com/boltdb/bolt"

	"github.com/nlamirault/abraracourcix/config"
)

// tempfile returns a temporary file path.
func tempfile() (string, error) {
	f, _ := ioutil.TempFile("", "boltdb-")
	err := f.Close()
	if err != nil {
		return "", err
	}
	err = os.Remove(f.Name())
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}

func getBoltDBConfiguration() (*config.Configuration, error) {
	tf, err := tempfile()
	if err != nil {
		return nil, err
	}
	return &config.Configuration{
		Storage: &config.StorageConfiguration{
			Name: "boltdb",
			BoltDB: &config.BoltDBConfiguration{
				Bucket: "UT",
				File:   tf,
			},
		},
	}, nil
}

// Ensure that gets a non-existent key returns nil.
func TestBoltDB_Get_NonExistent(t *testing.T) {
	conf, err := getBoltDBConfiguration()
	if err != nil {
		t.Fatalf("Can't create configuration")
	}
	db, err := newBoltdbStorage(conf)
	if err != nil {
		t.Fatalf("Can't create BoltDB test database.")
	}
	if err := db.Init(); err != nil {
		t.Fatalf("Can't initialize BoltDB storage.")
	}
	// defer db.Close()
	defer func() {
		err := db.Close()
		if err != nil {
			t.Fatalf("Can't close BoltDB test database.")
		}
	}()

	value, err := db.Get([]byte("foo"))
	if err != nil {
		t.Fatalf("Can't retrieve BoltDB key.")
	}
	// fmt.Println("Value: ", string(value))
	if value != nil {
		t.Fatalf("Error retrieve invalid key.")
	}
}

// Ensure that that gets an existent key returns value.
func TestBoltDB_Get_Existent(t *testing.T) {
	// tf, err := tempfile()
	// if err != nil {
	// 	t.Fatalf("Can't create temporary file: %v", err)
	// }
	conf, err := getBoltDBConfiguration()
	if err != nil {
		t.Fatalf("Can't create configuration")
	}
	db, err := newBoltdbStorage(conf)
	if err != nil {
		t.Fatalf("Can't create BoltDB test database.")
	}
	if err := db.Init(); err != nil {
		t.Fatalf("Can't initialize BoltDB storage.")
	}
	//defer db.Close()
	defer func() {
		err := db.Close()
		if err != nil {
			t.Fatalf("Can't close BoltDB test database.")
		}
	}()

	err = db.Put([]byte("foo"), []byte("bar"))
	if err != nil {
		t.Fatalf("Can't put values into BoltDB database: %v", err)
	}
	value, err := db.Get([]byte("foo"))
	if err != nil {
		t.Fatalf("Can't retrieve BoltDB key.")
	}
	// fmt.Println("Value: ", string(value))
	if string(value) != "bar" {
		t.Fatalf("Error retrieve invalid value.")
	}
}
