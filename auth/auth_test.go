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

package auth

import (
	// "fmt"
	"testing"

	"github.com/nlamirault/abraracourcix/config"
)

func Test_InvalidAuthenticationSystem(t *testing.T) {
	_, err := New(&config.Configuration{})
	if err == nil {
		t.Fatalf("Auth error: %s", err)
	}
}

func Test_UnknownAuthenticationSystem(t *testing.T) {
	_, err := New(&config.Configuration{
		Tracing: &config.TracingConfiguration{
			Name: "OpenBar",
		},
	})
	if err == nil {
		t.Fatal("Configure auth failed.")
	}
}
