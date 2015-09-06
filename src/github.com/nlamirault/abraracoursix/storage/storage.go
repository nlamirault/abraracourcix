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
	"errors"
	"fmt"
)

// Backend represents a storage Backend
type Backend string

const (
	// LEVELDB backend
	LEVELDB Backend = "leveldb"

	// BOLTDB backend
	BOLTDB Backend = "boltdb"

	// MEMDB backend
	MEMDB Backend = "memdb"
)

var (
	// ErrNotSupported is thrown when the backend k/v store is not supported by libkv
	ErrNotSupported = errors.New("Backend storage not supported yet, please choose one of")

	// ErrNotImplemented is thrown when a method is not implemented by the current backend
	ErrNotImplemented = errors.New("Call not implemented in current backend")
)

// Storage represents the Abraracoursix backend storage
// Each storage should support every call listed
// here.
type Storage interface {

	// Put a value at the specified key
	Put(key []byte, value []byte) error

	// Get a value given its key
	Get(key []byte) ([]byte, error)

	// Delete the value at the specified key
	Delete(key []byte) error

	// Verify if a Key exists in the store
	//Exists(key string) (bool, error)

	// Close the store connection
	Close()

	// Print backend informations
	Print()
}

// InitStorage creates an instance of storage
func InitStorage(backend Backend, path string) (Storage, error) {
	switch backend {
	case MEMDB:
		return NewMemDB(path)
	case LEVELDB:
		return NewLevelDB(path)
	case BOLTDB:
		return NewBoltDB(path)
	default:
		return nil, fmt.Errorf("%s %s",
			ErrNotSupported.Error(), " unsupported backend")
	}

}
