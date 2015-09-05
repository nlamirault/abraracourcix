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
	"github.com/gin-gonic/gin"

	"github.com/nlamirault/abraracoursix/storage"
)

var (
	port     string
	debug    bool
	version  bool
	database *storage.LevelDB
)

func init() {
	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&version, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&port, "port", "8080", "port to use")
	flag.Parse()
}

func main() {
	// set log level
	if debug {
		log.SetLevel(log.DebugLevel)
	}

	if version {
		fmt.Println("Abraracoursix v", Version)
		return
	}
	router := gin.Default()
	router.GET("/", help)
	router.GET("/api/version", displayAPIVersion)
	v1 := router.Group("api/v1")
	v1.GET("/get/:url", urlShow)
	v1.POST("/create/:url", urlCreate)
	database, err := storage.NewDatabase("/home/nlamirault/.config/abraracoursix/db")
	if err != nil {
		log.Fatalln("Database is not load, err - ", err)
		return
	}
	database.Print()
	log.Info("Get db key")
	data, err := database.Get([]byte("foo"))
	if err != nil {
		log.Info("Unknown URL with key")
	}
	log.Info("Data: ", data)
	database.Put([]byte("foo"), []byte("bar"))
	log.Info("Start web service")
	router.Run(":8080")
}
