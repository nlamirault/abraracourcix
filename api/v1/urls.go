// Copyright (C) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/nlamirault/abraracourcix/io"
	"github.com/nlamirault/abraracourcix/storage"
)

// URLShow send the url store using the key
func (ws *WebService) URLShow(c *echo.Context) error {
	key := c.Param("url")
	log.Printf("[INFO] [abraracourcix] Retrieve URL using key: %v", key)
	url, err := ws.retrieveURL([]byte(key))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	if url == nil {
		return c.JSON(http.StatusNotFound,
			&APIErrorResponse{
				Error: fmt.Sprintf("Unknown key %s", key),
			})
	}
	ws.manageAnalytics(url, c.Request(), false, true)
	log.Printf("[INFO] [abraracourcix] Find URL : %v", url)
	return c.JSON(http.StatusOK, url)
}

// URLCreate store a long URL using a key
func (ws *WebService) URLCreate(c *echo.Context) error {
	var url storage.URL
	err := c.Bind(&url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	log.Printf("[INFO] [abraracourcix] URL to store: %v", url)
	if len(url.LongURL) > 0 {
		url.Key, err = io.GenerateKey()
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				&APIErrorResponse{Error: err.Error()})
		}
		url.CreationDate = io.GetCreationDate()
		err := ws.storeURL([]byte(url.Key), &url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				&APIErrorResponse{Error: err.Error()})
		}
		log.Printf("[INFO] [abraracourcix] URL stored : %s -> %v",
			url.Key, url)
		ws.createAnalytics(&url)
		return c.JSON(http.StatusOK, url)
	}
	str := &APIErrorResponse{
		Error: fmt.Sprintf("Invalid URL"),
	}
	return c.JSON(http.StatusBadRequest, str)
}
