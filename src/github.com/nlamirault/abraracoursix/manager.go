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
	// "github.com/syndtr/goleveldb/leveldb"
)

type UrlMapping struct {
	ShortURL string `json:shorturl`
	LongURL  string `json:longurl`
}

type APIResponse struct {
	Message string `json:message`
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
	data, err := database.Get([]byte(key))
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
	err := database.Put([]byte(key), []byte(url))
	if err != nil {
		str := fmt.Sprintf("Can't store URL %s", url)
		c.JSON(http.StatusNotFound, gin.H{"Error": str})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ShortURL": key})
}
