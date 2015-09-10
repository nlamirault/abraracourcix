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
	// "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nlamirault/abraracourcix/api/v1"
	"github.com/nlamirault/abraracourcix/storage"
)

// GetWebService return a new gin.Engine
func GetWebService(store storage.Storage) *echo.Echo {
	ws := v1.NewWebService(store)
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.Get("/", ws.Help)
	e.Get("/:url", ws.Redirect)
	e.Get("/api/version", ws.DisplayAPIVersion)
	v1 := e.Group("/api/v1")
	v1.Get("/urls/:url", ws.URLShow)
	v1.Post("/urls", ws.URLCreate)
	return e
}
