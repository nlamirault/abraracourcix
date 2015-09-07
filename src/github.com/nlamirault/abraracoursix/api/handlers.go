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

package api

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/nlamirault/abraracoursix/api/v1"
	"github.com/nlamirault/abraracoursix/io"
	"github.com/nlamirault/abraracoursix/storage"
)

// var (
// 	webservice *WebService
// )

type URL struct {
	URL string `json:"url" binding:"required"`
}

// WebService represents the Restful API
type WebService struct {
	Store storage.Storage
}

// NewWebService creates a new WebService instance
func NewWebService(store storage.Storage) *WebService {
	return &WebService{Store: store}
}

// Help send a message in JSON
func (ws *WebService) Help(c *gin.Context) {
	c.String(http.StatusOK,
		"Welcome to Abraracoursix, a simple URL Shortener")
}

// DisplayAPIVersion sends the API version in JSON format
func (ws *WebService) DisplayAPIVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": "1"})
}

// URLShow send the url store using the key
func (ws *WebService) URLShow(c *gin.Context) {
	key := c.Param("url")
	log.Info("Retrieve URL using key: ", key)
	data, err := ws.Store.Get([]byte(key))
	if err != nil {
		str := fmt.Sprintf("Error retrieving URL with key %s", key)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": str})
		return
	}
	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"Unknown key": key})
		return
	}
	url := string(data)
	log.Info("Find URL : ", url)
	c.JSON(http.StatusOK, gin.H{"URL": url})
}

// URLCreate store a long URL using a key
func (ws *WebService) URLCreate(c *gin.Context) {
	// url := c.PostForm("url")
	var url URL
	c.Bind(&url)
	log.Infof("URL to store: %v", url)
	key := io.GenerateKey()
	err := ws.Store.Put([]byte(key), []byte(url.URL))
	if err != nil {
		str := fmt.Sprintf("Can't store URL %s", url.URL)
		c.JSON(http.StatusNotFound, gin.H{"Error": str})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ShortURL": key})
	// } else {
	// 	str := fmt.Sprintf("Invalid URL %s", url.URL)
	// 	c.JSON(http.StatusBadRequest, gin.H{"Error": str})
	// 	return
	// }
}

// GetWebService return a new gin.Engine
func GetWebService(store storage.Storage) *gin.Engine {
	ws := v1.NewWebService(store)
	r := gin.Default()
	r.GET("/", ws.Help)
	r.GET("/api/version", ws.DisplayAPIVersion)
	v1 := r.Group("api/v1")
	v1.GET("/get/:url", ws.URLShow)
	v1.POST("/create", ws.URLCreate)
	return r
}
