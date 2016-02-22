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
	// "fmt"
	"testing"
	// "github.com/boltdb/bolt"
)

// Ensure that gets a non-existent key returns nil.
func TestMemDB(t *testing.T) {
	db, err := NewMemDB("/tmp")
	if err != nil {
		t.Fatalf("Can't create MemDB test database.")
	}
	err = db.Put([]byte("foo"), []byte("bar"))
	if err != nil {
		t.Fatalf("Can't store MemDB values: %v.", err)
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
