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
	//"fmt"
	"log"
	"net/http"

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

// APIErrorResponse reprensents an error in JSON
type APIErrorResponse struct {
	Error string `json:"error"`
}

// NewWebService creates a new WebService instance
func NewWebService(store storage.Storage) *WebService {
	log.Printf("[DEBUG] [abraracourcix] Creates webservice with backend : %v",
		store)
	return &WebService{Store: store}
}

func (ws *WebService) storeURL(key []byte, url *storage.URL) error {
	data, err := storage.EncodeURL(url)
	if err != nil {
		return storage.ErrURLNotEncoded
	}
	err = ws.Store.Put(key, data)
	if err != nil {
		return storage.ErrEntityNotSaved
	}
	return nil
}

func (ws *WebService) retrieveURL(key []byte) (*storage.URL, error) {
	data, err := ws.Store.Get(key)
	if err != nil {
		return nil, storage.ErrEntityNotStore
	}
	if len(data) == 0 {
		return nil, nil
	}
	url, err := storage.DecodeURL(data)
	if err != nil {
		return nil, storage.ErrURLNotDecoded
	}
	return url, nil
}

func (ws *WebService) storeAnalytics(key []byte, stat *storage.Analytics) error {
	data, err := storage.EncodeAnalytics(stat)
	if err != nil {
		return storage.ErrAnalyticsNotEncoded
	}
	err = ws.Store.Put(key, data)
	if err != nil {
		return storage.ErrEntityNotSaved
	}
	return nil
}

func (ws *WebService) retrieveAnalytics(key []byte) (*storage.Analytics, error) {
	data, err := ws.Store.Get(key)
	if err != nil {
		return nil, storage.ErrEntityNotStore
	}
	stat, err := storage.DecodeAnalytics(data)
	if err != nil {
		return nil, storage.ErrAnalyticsNotDecoded
	}
	return stat, nil
}

func (ws *WebService) manageAnalytics(url *storage.URL, request *http.Request, longURL bool, shortURL bool) {
	log.Printf("[INFO] [abraracourcix] Analytics for URL : %v %s %s",
		url, request.UserAgent(), request.Referer())
	key := storage.GetAnalyticsKey(url.Key)
	stat, err := ws.retrieveAnalytics([]byte(key))
	if err != nil {
		log.Printf("[WARN] [abraracourcix] Can't decode Analytics %v", err)
		return
	}
	log.Printf("[INFO] [abraracourcix] Analytics find : %v", stat)
	ua := request.UserAgent()
	if stat.UserAgents != nil {
		stat.UserAgents[ua] = stat.UserAgents[ua] + 1
	} else {
		stat.UserAgents = make(map[string]int64)
		stat.UserAgents[ua] = 1
	}
	if longURL {
		stat.LongURLClicks = stat.LongURLClicks + 1
	}
	if shortURL {
		stat.ShortURLClicks = stat.ShortURLClicks + 1
	}
	err = ws.storeAnalytics([]byte(key), stat)
	if err != nil {
		log.Printf("[WARN] [abraracourcix] Can't store analytics URL %s %v",
			url, stat)
	}
	log.Printf("[INFO] [abraracourcix] Analytics updated : %v", stat)
}

func (ws *WebService) createAnalytics(url *storage.URL) {
	log.Printf("[INFO] [abraracourcix] Analytics for URL : %v", url)
	stat := storage.NewAnalytics()
	key := storage.GetAnalyticsKey(url.Key)
	err := ws.storeAnalytics([]byte(key), stat)
	if err != nil {
		log.Printf("[WARN] [abraracourcix] Can't store analytics URL %s %v",
			url, stat)
	}
	log.Printf("[INFO] [abraracourcix] Analytics added : %s -> %v", key, stat)
}
