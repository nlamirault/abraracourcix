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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nlamirault/abraracourcix/api"
	"github.com/nlamirault/abraracourcix/io"
	"github.com/nlamirault/abraracourcix/logging"
	"github.com/nlamirault/abraracourcix/storage"
)

var (
	port       string
	debug      bool
	version    bool
	backend    string
	backendURL string
	dataDir    string
	username   string
	password   string
)

func init() {
	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&port, "port", "8080", "port to use")
	flag.StringVar(&backend, "backend", "boltdb", "Storage backend")
	flag.StringVar(&dataDir, "data", "", "Data directory")
	flag.StringVar(&backendURL, "backend-url", "", "URL for backends")
	flag.StringVar(&username, "username", "", "Username authentication")
	flag.StringVar(&password, "password", "", "Password authentication")
	flag.Parse()
}

func getDefaultDataDir() string {
	return fmt.Sprintf("%s/.config/abraracourcix", io.UserHomeDir())
}

func getStorage() (storage.Storage, error) {
	if len(dataDir) == 0 {
		dataDir = getDefaultDataDir()
	}
	err := os.MkdirAll(dataDir, 0744)
	if err != nil {
		log.Printf("[ERROR] [abraracourcix] Unable to create data directory %v",
			err)
	}
	return storage.InitStorage(backend, &storage.Config{
		Data:       fmt.Sprintf("%s/%s", dataDir, backend),
		BackendURL: backendURL,
	})
}

func main() {
	if debug {
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	store, err := getStorage()
	if err != nil {
		log.Printf("[ERROR] [abraracourcix] %v", err)
		return
	}
	var auth *api.Authentication
	log.Printf("%s %s", username, password)
	if len(username) > 0 && len(password) > 0 {
		auth = &api.Authentication{
			Username: username,
			Password: password,
		}
	}
	e := api.GetWebService(store, auth)
	if debug {
		e.Debug()
	}
	log.Printf("[INFO] [abraracourcix] Launch Abraracourcix on %s using %s backend",
		port, backend)
	e.Run(fmt.Sprintf(":%s", port))
}
