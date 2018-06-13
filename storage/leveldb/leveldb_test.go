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

package leveldb

import (
	"testing"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/storage/storagetest"
)

func getLevelDBConfiguration() (*config.Configuration, error) {
	td, err := storagetest.TempDirectory()
	if err != nil {
		return nil, err
	}
	return &config.Configuration{
		Storage: &config.StorageConfiguration{
			Name: "leveldb",
			LevelDB: &config.LevelDBConfiguration{
				Path: td,
			},
		},
	}, nil
}

func TestLevelDBStorage(t *testing.T) {
	conf, err := getLevelDBConfiguration()
	if err != nil {
		t.Fatalf("Can't create configuration")
	}
	db, err := newLevelDBStorage(conf)
	if err != nil {
		t.Fatalf("Can't create LevelDB test database.")
	}
	storagetest.ValidateBackend(t, db)
}
