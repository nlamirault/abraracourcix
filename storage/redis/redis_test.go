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

	"testing"

	"github.com/nlamirault/abraracourcix/config"
)

// These tests require redis server running on localhost:6379 (the default)
const redisTestServer = "localhost:6379"

func getRedisConfiguration() (*config.Configuration, error) {
	return &config.Configuration{
		Storage: &config.StorageConfiguration{
			Name: "redis",
			Redis: &config.RedisConfiguration{
				Address:   "6379",
				Keyprefix: "foobar",
			},
		},
	}, nil
}

// Ensure that gets a non-existent key returns nil.
func TestRedisDB_Get_InvalidURI(t *testing.T) {
	conf, err := getRedisConfiguration()
	if err != nil {
		t.Fatalf("Can't create configuration")
	}
	db, err := newRedisStorage(conf)
	if err != nil {
		t.Fatalf("Can't create RedisDB test database.")
	}
	if err := db.Init(); err != nil {
		t.Fatalf("Can't initialize RedisDB storage.")
	}
}
