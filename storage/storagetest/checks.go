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

package storagetest

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/nlamirault/abraracourcix/io"
	"github.com/nlamirault/abraracourcix/storage"
)

func ValidateBackend(t *testing.T, db storage.Storage) {
	checkGetUnknownEntry(t, db)
	key, err := io.GenerateKey()
	if err != nil {
		t.Fatalf(err.Error())
	}
	value, err := io.GenerateKey()
	if err != nil {
		t.Fatalf(err.Error())
	}
	checkManageEntry(t, db, []byte(key), []byte(value))
	checkListEntries(t, db)
}

func checkGetUnknownEntry(t *testing.T, db storage.Storage) {
	if err := db.Init(); err != nil {
		t.Fatalf("Can't initialize storage: %v", err)
	}
	key, err := io.GenerateKey()
	if err != nil {
		t.Fatalf(err.Error())
	}
	value, err := db.Get([]byte(key))
	if err != nil {
		t.Fatalf("Can't retrieve key.")
	}
	fmt.Println("Value: ", string(value))
	if value != nil {
		t.Fatalf("Error retrieve invalid key %s", string(value))
	}
	if err := db.Close(); err != nil {
		t.Fatalf("Can't close test database.")
	}
}

func checkManageEntry(t *testing.T, db storage.Storage, key []byte, value []byte) {
	if err := db.Init(); err != nil {
		t.Fatalf("Can't initialize storage: %v", err)
	}
	if err := db.Put(key, value); err != nil {
		t.Fatalf("Can't store key: %v", err)
	}
	newValue, err := db.Get(key)
	if err != nil {
		t.Fatalf("Can't retrieve entry: %v", err)
	}
	if string(newValue) != string(value) {
		t.Fatalf("Error retrieve invalid value.")
	}

	if err := db.Delete(key); err != nil {
		t.Fatalf("Can't delete entry: %v", err)
	}

	val, err := db.Get(key)
	if err != nil {
		t.Fatalf("Can't retrieve deleted entry: %v", err)
	}
	if val != nil {
		t.Fatalf("Can retrieve a deleted entry: %s", string(val))
	}

	if err := db.Close(); err != nil {
		t.Fatalf("Can't close test database.")
	}
}

func checkListEntries(t *testing.T, db storage.Storage) {
	if err := db.Init(); err != nil {
		t.Fatalf("Can't initialize storage: %v", err)
	}
	nb := 3
	for i := 1; i <= nb; i++ {
		key, err := io.GenerateKey()
		if err != nil {
			t.Fatalf(err.Error())
		}
		value, err := io.GenerateKey()
		if err != nil {
			t.Fatalf(err.Error())
		}
		if err := db.Put([]byte(key), []byte(value)); err != nil {
			t.Fatalf("Can't store key: %v", err)
		}
	}
	keys, err := db.List()
	if err != nil {
		t.Fatalf("Can't list entries: %v", err)
	}
	if len(keys) <= 0 {
		t.Fatalf("Invalid number of keys")
	}
}

func TempDirectory() (string, error) {
	d, err := ioutil.TempDir("", "storage-ut-")
	if err != nil {
		return "", err
	}
	if err := os.Remove(d); err != nil {
		return "", err
	}
	return d, nil
}

func TempFile() (string, error) {
	f, err := ioutil.TempFile("", "storage-ut-")
	if err != nil {
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}
	if err := os.Remove(f.Name()); err != nil {
		return "", err
	}
	return f.Name(), nil
}
