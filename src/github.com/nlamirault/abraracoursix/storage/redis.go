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
	//"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/garyburd/redigo/redis"
)

const keyprefix = "abraracoursix"

// Redis represents a storage using the Redis database
type Redis struct {
	Conn      redis.Conn
	Keyprefix string
}

// NewRedis instantiates a new Redis database client
func NewRedis(address string) (*Redis, error) {
	conn, err := redis.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	return &Redis{Conn: conn, Keyprefix: keyprefix}, nil
}

// Get a value given its key
func (db *Redis) Get(key []byte) ([]byte, error) {
	log.Debugf("[%s] Delete : %v", REDIS, string(key))
	val, err := db.Conn.Do("HGET", db.Keyprefix, string(key))
	return val.([]byte), err
}

// Put a value at the specified key
func (db *Redis) Put(key []byte, value []byte) error {
	log.Debugf("[%s] Delete : %v", REDIS, string(key))
	_, err := db.Conn.Do("HSET", db.Keyprefix, string(key), value)
	return err
}

// Delete the value at the specified key
func (db *Redis) Delete(key []byte) error {
	log.Debugf("[%s] Delete : %v", REDIS, string(key))
	_, err := db.Conn.Do("HDEL", string(key))
	return err
}

// Close the store connection
func (db *Redis) Close() {
	log.Debugf("[%s] Close", REDIS)
	db.Conn.Close()
}

// Print backend informations
func (db *Redis) Print() {
	log.Debugf("[%s] Print", REDIS)
}
