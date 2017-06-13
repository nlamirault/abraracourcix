// Copyright (C) 2015-2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/storage"
)

const (
	label = "redis"
)

type redisDB struct {
	//Conn      redis.Conn
	keyprefix string
	pool      *redis.Pool
}

func init() {
	storage.RegisterStorage(label, newRedisStorage)
}

func newRedisStorage(conf *config.Configuration) (storage.Storage, error) {
	glog.V(1).Infof("Create storage using Redis : %s", conf.Storage)
	pool := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf(":%s", conf.Storage.Redis.Address))
		},
	}
	return &redisDB{
		pool:      pool,
		keyprefix: conf.Storage.Redis.Keyprefix,
	}, nil
}

func (redisDB *redisDB) Name() string {
	return label
}

func (redisDB *redisDB) Init() error {
	glog.V(1).Info("Initialize")
	return nil
}

// Print backend informations
func (redisDB *redisDB) List() ([][]byte, error) {
	glog.V(1).Infof("List all URLs")
	urls := [][]byte{}
	keys, err := redis.Strings(redisDB.pool.Get().Do("KEYS", redisDB.keyprefix))
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		urls = append(urls, []byte(key))
	}
	return urls, nil
}

// Get a value given its key
func (redisDB *redisDB) Get(key []byte) ([]byte, error) {
	glog.V(1).Infof("Search entry with key : %v", string(key))
	val, err := redisDB.pool.Get().Do("HGET", redisDB.keyprefix, string(key))
	if err != nil {
		return nil, err
	}
	data, err := redis.String(val, nil)
	if err != nil {
		return nil, err
	}
	glog.V(2).Infof("Find : %s", data)
	return []byte(data), err
}

// Put a value at the specified key
func (redisDB *redisDB) Put(key []byte, value []byte) error {
	glog.V(1).Infof("Put : %v %v", string(key), string(value))
	_, err := redisDB.pool.Get().Do("HSET", redisDB.keyprefix, string(key), value)
	return err
}

// Delete the value at the specified key
func (redisDB *redisDB) Delete(key []byte) error {
	glog.V(1).Infof("Delete : %v", string(key))
	_, err := redisDB.pool.Get().Do("HDEL", string(key))
	return err
}

// Close the store connection
func (redisDB *redisDB) Close() error {
	glog.V(1).Infof("Close")
	if redisDB.pool != nil {
		return redisDB.pool.Close()
	}
	return nil
}

// Print backend informations
func (redisDB *redisDB) Print() error {
	glog.V(1).Infof("Storage backend: %s", label)
	keys, err := redis.Strings(redisDB.pool.Get().Do("KEYS", "*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		fmt.Printf("%s", key)
	}
	return nil
}
