// Copyright (C) 2015-2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"io"
	"strings"
	//"fmt"
	"mime"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/golang/glog"

	"github.com/nlamirault/abraracourcix/pb"
	"github.com/nlamirault/abraracourcix/pkg/ui/swagger"
)

const (
	prefix = "/swagger-ui/"
)

// ServeSwagger expose files in third_party/swagger-ui/ on <host>/swagger-ui
func ServeSwagger(mux *http.ServeMux) {
	glog.V(1).Infof("Create the SwaggerUI handler")
	mime.AddExtensionType(".svg", "image/svg+xml")

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:     swagger.Asset,
		AssetDir:  swagger.AssetDir,
		AssetInfo: swagger.AssetInfo,
		Prefix:    "third_party/swagger-ui",
	})
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(pb.Swagger))
	})
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
