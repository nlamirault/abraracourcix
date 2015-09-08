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
	"time"

	// "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"

	"github.com/nlamirault/abraracoursix/api/v1"
	"github.com/nlamirault/abraracoursix/storage"
)

// GetWebService return a new gin.Engine
func GetWebService(store storage.Storage) *gin.Engine {
	ws := v1.NewWebService(store)
	r := gin.Default()
	r.GET("/", ws.Help)
	r.GET("/api/version", ws.DisplayAPIVersion)
	v1 := r.Group("api/v1")
	v1.GET("/urls/:url", ws.URLShow)
	v1.POST("/urls", ws.URLCreate)
	return r
}
