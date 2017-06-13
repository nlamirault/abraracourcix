// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/BurntSushi/toml"
)

// Configuration holds configuration for Enigma.
type Configuration struct {
	API     *APIConfiguration
	Tracing *TracingConfiguration
	Storage *StorageConfiguration
	Auth    *AuthConfiguration
}

// New returns a Configuration with default values
func New() *Configuration {
	return &Configuration{
		API:     &APIConfiguration{},
		Tracing: &TracingConfiguration{},
		Storage: &StorageConfiguration{},
		Auth:    &AuthConfiguration{},
	}
}

// LoadFileConfig returns a Configuration from reading the specified file (a toml file).
func LoadFileConfig(file string) (*Configuration, error) {
	configuration := New()
	if _, err := toml.DecodeFile(file, configuration); err != nil {
		return nil, err
	}
	return configuration, nil
}

// APIConfiguration defines the configuration for the gRPC and REST api
type APIConfiguration struct {
	GrpcPort int
	RestPort int
}

type ZipkinConfiguration struct {
	Host string
	Port int
}

type AppdashConfiguration struct {
	Host string
	Port int
}

type JaegerConfiguration struct {
	Host string
	Port int
}

// TracingConfiguration defines the OpenTracing usage
type TracingConfiguration struct {
	Name    string
	Zipkin  *ZipkinConfiguration
	Appdash *AppdashConfiguration
	Jaeger  *JaegerConfiguration
}

type StorageConfiguration struct {
	Name    string
	BoltDB  *BoltDBConfiguration
	LevelDB *LevelDBConfiguration
}

// BoltDBConfiguration defines the configuration for BoltDB storage backend
type BoltDBConfiguration struct {
	Bucket string
	File   string
}

type LevelDBConfiguration struct {
	Path string
}

type AuthConfiguration struct {
	Name  string
	Vault *VaultConfiguration
}

type VaultConfiguration struct {
	Address    string
	Roleid     string
	Secretid   string
	HealthUser string `toml:"healthuser"`
	HealthKey  string `toml:"healthkey"`
}
