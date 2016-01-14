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
	//"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/nlamirault/abraracourcix/storage"
)

// URLStats send the url analytics using the key
func (ws *WebService) URLStats(c *echo.Context) error {
	url := c.Param("url")
	log.Printf("[INFO] [abraracourcix] Retrieve URL analytics using key: %v",
		url)
	stat, err := ws.retrieveAnalytics([]byte(storage.GetAnalyticsKey(url)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			&APIErrorResponse{Error: err.Error()})
	}
	log.Printf("[INFO] [abraracourcix] Find Analytics : %v", stat)
	return c.JSON(http.StatusOK, stat)
}
