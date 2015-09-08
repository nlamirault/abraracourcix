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

package v1

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/nlamirault/abraracoursix/io"
)

type URL struct {
	URL string `json:"url" binding:"required"`
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
	var url URL
	c.BindJSON(&url)
	log.Infof("URL to store: %v", url)
	if len(url.URL) > 0 {
		key := io.GenerateKey()
		err := ws.Store.Put([]byte(key), []byte(url.URL))
		if err != nil {
			str := fmt.Sprintf("Can't store URL %s", url.URL)
			c.JSON(http.StatusNotFound, gin.H{"Error": str})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ShortURL": key})
	} else {
		str := fmt.Sprintf("Invalid URL : [%s]", url.URL)
		c.JSON(http.StatusBadRequest, gin.H{"Error": str})
		return
	}
}
