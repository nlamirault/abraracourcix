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
	// "time"
)

const (
	// LEVELDB backend
	LEVELDB string = "leveldb"

	// BOLTDB backend
	BOLTDB string = "boltdb"

	// REDIS backend
	REDIS string = "redis"

	// MEMDB backend
	MEMDB string = "memdb"
)

var (
	// ErrNotSupported is thrown when the backend k/v store is not supported by libkv
	ErrNotSupported = errors.New("Backend storage not supported.")

	// ErrNotImplemented is thrown when a method is not implemented by the current backend
	ErrNotImplemented = errors.New("Call not implemented in current backend")
)

// Config represents storage configuration
type Config struct {
	Data string
	Port string
}

// Storage represents the Abraracourcix backend storage
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
func InitStorage(backend string, config *Config) (Storage, error) {
	switch backend {
	case MEMDB:
		return NewMemDB(config.Data)
	case LEVELDB:
		return NewLevelDB(config.Data)
	case BOLTDB:
		return NewBoltDB(config.Data)
	case REDIS:
		return NewRedis(config.Port)
	default:
		return nil, fmt.Errorf("%s %s", ErrNotSupported.Error(), "")
	}

}

// type StoreURL struct {
// 	Key          string
// 	ShortUrl     string
// 	LongUrl      string
// 	CreationDate int64
// }

// // NewStoreURL creates a StoreURL instance
// func NewStoreURL(key, shorturl, longurl string) *StoreURL {
// 	url := new(StoreURL)
// 	url.CreationDate = time.Now().UnixNano()
// 	url.Key = key
// 	url.LongUrl = longurl
// 	url.ShortUrl = shorturl
// 	return url
// }
