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

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	templateFile, err := ioutil.TempFile("", "configuration")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(templateFile.Name())
	data := []byte(`# configuration file

[api]
grpcPort = 8080
restPort = 9090

[storage]
name = "boltdb"

[storage.boltdb]
file = "/tmp/ut.db"
bucket = "foobar"

[tracing]
name = "zipkin"

[tracing.zipkin]
host = "10.2.4.6"
port = 9441

[tracing.appdash]
host = "10.1.3.5"
port = 8080

[tracing.jaeger]
host = "10.2.2.2"
port = 8888
`)
	err = ioutil.WriteFile(templateFile.Name(), data, 0700)
	if err != nil {
		t.Fatal(err)
	}
	configuration, err := LoadFileConfig(templateFile.Name())
	if err != nil {
		t.Fatalf("Error with configuration: %v", err)
	}
	fmt.Printf("Configuration : %#v\n", configuration)
	if configuration.Storage.Name != "boltdb" {
		t.Fatalf("Configuration backend failed")
	}

	// Storage
	if configuration.Storage.BoltDB.Bucket != "foobar" ||
		configuration.Storage.BoltDB.File != "/tmp/ut.db" {
		t.Fatalf("Configuration BoltDB failed")
	}

	// API
	if configuration.API.GrpcPort != 8080 ||
		configuration.API.RestPort != 9090 {
		t.Fatalf("Configuration API failed")
	}

	// Tracing
	if configuration.Tracing.Name != "zipkin" {
		t.Fatalf("Configuration OpenTracing tracer failed")
	}
	if configuration.Tracing.Zipkin.Host != "10.2.4.6" ||
		configuration.Tracing.Zipkin.Port != 9441 {
		t.Fatalf("Configuration OpenTracing Zipkin failed")
	}
	if configuration.Tracing.Appdash.Host != "10.1.3.5" ||
		configuration.Tracing.Appdash.Port != 8080 {
		t.Fatalf("Configuration OpenTracing Appdash failed")
	}
	if configuration.Tracing.Jaeger.Host != "10.2.2.2" ||
		configuration.Tracing.Jaeger.Port != 8888 {
		t.Fatalf("Configuration OpenTracing Jaeger failed")
	}
}
