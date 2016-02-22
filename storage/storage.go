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
	"errors"
	"fmt"
	//"log"
	"time"
)

const (
	// LEVELDB backend
	LEVELDB string = "leveldb"

	// BOLTDB backend
	BOLTDB string = "boltdb"

	// REDIS backend
	REDIS string = "redis"

	// MONGODB backend
	MONGODB string = "mongodb"

	// MEMDB backend
	MEMDB string = "memdb"
)

var (
	// ErrNotSupported is thrown when the backend store is not supported
	ErrNotSupported = errors.New("Backend storage not supported.")

	// ErrNotImplemented is thrown when a method is not implemented by the current backend
	ErrNotImplemented = errors.New("Call not implemented in current backend")

	// ErrEntityNotSaved is thrown when an entity can't be save into the backend
	ErrEntityNotSaved = errors.New("Can't save data")

	// ErrEntityNotStore is thrown when an entity isn't store into the backend
	ErrEntityNotStore = errors.New("Not store data")
)

// Config represents storage configuration
type Config struct {
	Data       string
	BackendURL string
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
	Close() error

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
		return NewRedis(config.BackendURL)
	case MONGODB:
		return NewMongo(config.BackendURL)
	default:
		return nil, fmt.Errorf("%s %s", ErrNotSupported.Error(), "")
	}

}

// URL represents an URL into storage backend
type URL struct {
	// Key is the short URL that expands to the long URL you provided
	Key string `json:"key"`
	// LongURL is the long URL to which it expands.
	LongURL string `json:"longUrl"`
	// CreationDate is the time at which this short URL was created
	CreationDate time.Time `json:"creation_date"`
}

// StringCount represents a label and a count
// type StringCount struct {
// 	// Count: Number of clicks for this top entry
// 	Count int64 `json:"count,omitempty,string"`

// 	// Id: Label assigned to this top entry
// 	ID string `json:"id,omitempty"`
// }

// Analytics contains click statistics
type Analytics struct {
	// LongUrlClicks: Number of clicks on all short URLs pointing to
	// this long URL.
	LongURLClicks int64 `json:"longUrlClicks,omitempty,string"`

	// ShortUrlClicks: Number of clicks on this short URL.
	ShortURLClicks int64 `json:"shortUrlClicks,omitempty,string"`

	// // Platforms Top platforms or OSes, e.g. "Linux, Windows, ..."
	// Platforms []*StringCount `json:"platforms,omitempty"`

	// // Browsers: Top browsers, e.g. "Chrome"; Only present if this data
	// // is available.
	// Browsers []*StringCount `json:"browsers,omitempty"`

	// UserAgent represents the user agent requester
	//UserAgents []*StringCount   `json:"user_agents,omitempty"`
	UserAgents map[string]int64 `json:"user_agents,omitempty"`

	// // Countries: Top countries (expressed as country codes), e.g. "US" or "FR"
	// Countries []*StringCount `json:"countries,omitempty"`
}

// NewAnalytics creates an Analytics instance
func NewAnalytics() *Analytics {
	return &Analytics{
		LongURLClicks:  1,
		ShortURLClicks: 0,
		UserAgents:     make(map[string]int64),
	}
}

// GetAnalyticsKey returns database key for Analytics
func GetAnalyticsKey(key string) string {
	return fmt.Sprintf("stat_%s", key)
}
