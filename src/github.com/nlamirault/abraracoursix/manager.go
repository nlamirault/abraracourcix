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
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/golang/leveldb"
	// ldb "github.com/golang/leveldb/db"
)

type UrlMapping struct {
	ShortURL string `json:shorturl`
	LongURL  string `json:longurl`
}

type APIResponse struct {
	Message string `json:message`
}

type Shortner struct {
	// db *leveldb.DB
}

func NewShortner() *Shortner {
	return &Shortner{
	//myconnection: NewDBConnection(),
	}
}

func help(c *gin.Context) {
	c.String(http.StatusOK,
		"Welcome to Abraracoursix, a simple URL Shortener")
}

func displayAPIVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "1"})
}

func urlShow(c *gin.Context) {
	url := c.Param("url")
	c.JSON(http.StatusOK,
		gin.H{"URL": url})
}

func urlCreate(c *gin.Context) {
	url := c.Param("url")
	c.JSON(http.StatusOK,
		gin.H{"URL": url})
}
