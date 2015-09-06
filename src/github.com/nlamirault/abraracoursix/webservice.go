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
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/nlamirault/abraracoursix/storage"
)

var (
	webservice *WebService
)

// WebService represents the Restful API
type WebService struct {
	Store storage.Storage
}

func setupWebService(store storage.Storage, port string) {
	store.Print()
	webservice := &WebService{Store: store}
	log.Info("Get db key")
	data, err := webservice.Store.Get([]byte("foo"))
	if err != nil {
		log.Info("Unknown URL with key")
	}
	log.Info("Data: ", data)
	webservice.Store.Put([]byte("foo"), []byte("bar"))
	log.Info("Start web service")
	router := gin.Default()
	router.GET("/", help)
	router.GET("/api/version", displayAPIVersion)
	v1 := router.Group("api/v1")
	v1.GET("/get/:url", urlShow)
	v1.POST("/create/:url", urlCreate)
	router.Run(fmt.Sprintf(":%s", port))
}

func help(c *gin.Context) {
	c.String(http.StatusOK,
		"Welcome to Abraracoursix, a simple URL Shortener")
}

func displayAPIVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "1"})
}

func urlShow(c *gin.Context) {
	key := c.Param("url")
	log.Info("Retrieve URL using key: ", key)
	data, err := webservice.Store.Get([]byte(key))
	if err != nil {
		str := fmt.Sprintf("Unknown URL with key %s", key)
		c.JSON(http.StatusNotFound, gin.H{"Error": str})
		return
	}
	c.JSON(http.StatusOK, gin.H{"URL": data})
}

func urlCreate(c *gin.Context) {
	url := c.Param("url")
	key := "aaa"
	err := webservice.Store.Put([]byte(key), []byte(url))
	if err != nil {
		str := fmt.Sprintf("Can't store URL %s", url)
		c.JSON(http.StatusNotFound, gin.H{"Error": str})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ShortURL": key})
}
