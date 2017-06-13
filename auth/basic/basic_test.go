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

package basic

import (
	"testing"

	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	"github.com/nlamirault/abraracourcix/auth"
	"github.com/nlamirault/abraracourcix/transport"
)

func createBasicAuthSystem() *basicAuthSystem {
	return &basicAuthSystem{}
}

func Test_BasicAuthWithValidUsernamePassword(t *testing.T) {
	sys := createBasicAuthSystem()
	ctx := context.Background()
	span := opentracing.SpanFromContext(ctx)
	headers, err := sys.Authenticate(context.Background(), span, "YWRtaW46bmltZGE=")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if headers[transport.Username] != auth.Username {
		t.Fatalf("Invalid headers: %s", headers)
	}
}

func Test_BasicAuthWithInvalidUsernameOrPassword(t *testing.T) {
	sys := createBasicAuthSystem()
	ctx := context.Background()
	span := opentracing.SpanFromContext(ctx)
	_, err := sys.Authenticate(ctx, span, "Zm9vOmJhcg==")
	if err == nil {
		t.Fatalf("No error with invalid username/password.")
	}
	if err.Error() != "Unauthorized" {
		t.Fatalf("Invalid error: %s", err.Error())
	}
}

func Test_BasicAuthWithInvalidCredentials(t *testing.T) {
	sys := createBasicAuthSystem()
	ctx := context.Background()
	span := opentracing.SpanFromContext(ctx)
	_, err := sys.Authenticate(ctx, span, "Zm9v")
	if err == nil {
		t.Fatalf("No error with invalid credentials.")
	}
	if err.Error() != "Not Authorized" {
		t.Fatalf("Invalid error: %s", err.Error())
	}
}

func Test_BasicAuthWithInvalidBase64(t *testing.T) {
	sys := createBasicAuthSystem()
	ctx := context.Background()
	span := opentracing.SpanFromContext(ctx)
	_, err := sys.Authenticate(ctx, span, "csdmlcsdcsd")
	if err == nil {
		t.Fatalf("No error with invalid base64.")
	}
}
