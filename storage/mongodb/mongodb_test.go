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

package mongodb

import (
	"testing"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/io"
	"github.com/nlamirault/abraracourcix/storage/storagetest"
)

// These tests require a MongoDB server running on "127.0.0.1:27017" (the default)
const redisTestServer = "localhost:27017"

func getMongoDBConfiguration() (*config.Configuration, error) {
	dbName, err := io.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &config.Configuration{
		Storage: &config.StorageConfiguration{
			Name: "mongodb",
			MongoDB: &config.MongoDBConfiguration{
				Address:    "127.0.0.1:27017",
				Database:   dbName,
				Collection: "myurls",
			},
		},
	}, nil
}

func TestMongoDBStorage(t *testing.T) {
	conf, err := getMongoDBConfiguration()
	if err != nil {
		t.Fatalf("Can't create configuration")
	}
	db, err := newMongoDBStorage(conf)
	if err != nil {
		t.Fatalf("Can't create MongoDB storage engine.")
	}
	storagetest.ValidateBackend(t, db)
}
