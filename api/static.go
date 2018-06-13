// Copyright (C) 2016-2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	//"fmt"
	"net/http"

	"github.com/golang/glog"

	"github.com/nlamirault/abraracourcix/pkg/static"
)

// ServeStaticFile expose static files
func ServeStaticFile(mux *http.ServeMux) {
	glog.V(1).Infof("Create the Static file handler")

	mux.HandleFunc("/changelog", func(w http.ResponseWriter, req *http.Request) {
		if data, err := static.ChangelogMdBytes(); err == nil {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write(data)
		}
	})
}
