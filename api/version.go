// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"encoding/json"
	"net/http"

	"github.com/golang/glog"

	"github.com/nlamirault/abraracourcix/version"
)

type VersionResponse struct {
	Version string `json:"version"`
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	glog.V(2).Infof("Retrieve version")
	response := VersionResponse{
		Version: version.Version,
	}
	json.NewEncoder(w).Encode(response)
	return
}
