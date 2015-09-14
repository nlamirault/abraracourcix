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

	"github.com/nlamirault/abraracourcix/storage"
)

// WebService represents the Restful API
type WebService struct {
	Store storage.Storage
}

// APIVersion represents version of the REST API
type APIVersion struct {
	Version string `json:"version"`
}

// NewWebService creates a new WebService instance
func NewWebService(store storage.Storage) *WebService {
	log.Printf("[DEBUG] [abraracourcix] Creates webservice with backend : %v",
		store)
	return &WebService{Store: store}
}

// Help send a message in JSON
func (ws *WebService) Help(c *echo.Context) error {
	return c.String(http.StatusOK,
		"Welcome to Abraracourcix, a simple URL Shortener\n")
}

// DisplayAPIVersion sends the API version in JSON format
func (ws *WebService) DisplayAPIVersion(c *echo.Context) error {
	return c.JSON(http.StatusOK, &APIVersion{Version: "1"})
}

// Redirect retrieve longUrl from storage and send a HTTP Redirect
func (ws *WebService) Redirect(c *echo.Context) error {
	key := c.Param("url")
	log.Printf("[INFO] [abraracourcix] Retrieve URL using key: %v", key)
	data, err := ws.Store.Get([]byte(key))
	if err != nil {
		str := &APIErrorResponse{
			Error: fmt.Sprintf("Error retrieving URL with key %s", key),
		}
		return c.JSON(http.StatusInternalServerError, str)
	}
	if data == nil {
		str := &APIErrorResponse{
			Error: fmt.Sprintf("Unknown key %s", key),
		}
		return c.JSON(http.StatusNotFound, str)
	}
	//url := &URL{URL: string(data), Key: key}
	url, err := storage.DecodeURL(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	log.Printf("[INFO] [abraracourcix] Redirect to URL : %#v", url)
	return c.Redirect(http.StatusMovedPermanently, url.LongURL)
}
