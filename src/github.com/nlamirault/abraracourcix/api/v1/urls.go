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
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/nlamirault/abraracourcix/io"
	"github.com/nlamirault/abraracourcix/storage"
)

// APIErrorResponse reprensents an error in JSON
type APIErrorResponse struct {
	Error string `json:"error"`
}

// URLShow send the url store using the key
func (ws *WebService) URLShow(c *echo.Context) error {
	key := c.Param("url")
	log.Printf("[INFO] [abraracourcix] Retrieve URL using key: %v", key)
	data, err := ws.Store.Get([]byte(key))
	if err != nil {
		str := &APIErrorResponse{
			Error: fmt.Sprintf("Error retrieving URL with key %s : %v",
				key, err),
		}
		return c.JSON(http.StatusInternalServerError, str)
	}
	if data == nil {
		str := &APIErrorResponse{
			Error: fmt.Sprintf("Unknown key %s", key),
		}
		return c.JSON(http.StatusNotFound, str)
	}
	//url := string(data)
	//url := &URL{URL: string(data), Key: key}
	url, err := storage.DecodeURL(data)
	if err != nil {
		str := &APIErrorResponse{
			Error: fmt.Sprintf("Can't decore URL %v", err),
		}
		return c.JSON(http.StatusInternalServerError, str)
	}
	log.Printf("[INFO] [abraracourcix] Find URL : %v", url)
	return c.JSON(http.StatusOK, url)
}

// URLCreate store a long URL using a key
func (ws *WebService) URLCreate(c *echo.Context) error {
	var url storage.URL
	c.Bind(&url)
	log.Printf("[INFO] [abraracourcix] URL to store: %v", url)
	if len(url.LongURL) > 0 {
		// key := io.GenerateKey()
		url.Key = io.GenerateKey()
		url.CreationDate = io.GetCreationDate()
		// err := ws.Store.Put([]byte(key), []byte(url.LongURL))
		data, err := storage.EncodeURL(&url)
		if err != nil {
			str := &APIErrorResponse{
				Error: fmt.Sprintf("Can't encode URL %v", err),
			}
			return c.JSON(http.StatusNotFound, str)
		}
		err = ws.Store.Put([]byte(url.Key), data)
		if err != nil {
			str := &APIErrorResponse{
				Error: fmt.Sprintf("Can't store URL %s", url.LongURL),
			}
			return c.JSON(http.StatusNotFound, str)
		}
		// url.Key = key
		log.Printf("[INFO] [abraracourcix] URL stored : %v", url)
		return c.JSON(http.StatusOK, url)
	}
	str := &APIErrorResponse{
		Error: fmt.Sprintf("Invalid URL"),
	}
	return c.JSON(http.StatusBadRequest, str)
}
