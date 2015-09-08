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

	log "github.com/Sirupsen/logrus"

	"github.com/nlamirault/abraracoursix/api"
	"github.com/nlamirault/abraracoursix/io"
	"github.com/nlamirault/abraracoursix/storage"
)

var (
	port    string
	debug   bool
	version bool
	backend string
)

func init() {
	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&port, "port", "8080", "port to use")
	flag.StringVar(&backend, "backend", "boltdb", "Storage backend")
	flag.Parse()
}

func getConfigDir() string {
	return fmt.Sprintf("%s/.config/abraracoursix", io.UserHomeDir())
}

func main() {
	if debug {
		log.SetLevel(log.DebugLevel)
	}
	if version {
		fmt.Println("Abraracoursix v", Version)
		return
	}
	store, err := storage.InitStorage(backend, //"leveldb",
		fmt.Sprintf("%s/%s", getConfigDir(), backend))
	if err != nil {
		log.Fatalln("Database is not load, err - ", err)
		return
	}
	e := api.GetWebService(store)
	if debug {
		e.Debug()
	}
	log.Infof("Launch Abraracoursix on %s using %s backend",
		port, backend)
	e.Run(fmt.Sprintf(":%s", port))
}
