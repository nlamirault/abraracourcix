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

package vault

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	"github.com/nlamirault/abraracourcix/auth"
	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/tracing"
	"github.com/nlamirault/abraracourcix/transport"
)

const (
	mimetype = "application/json; charset=utf-8"
	cuid     = "ft02468"
	apiKey   = "cdcsdcs54-545d-48ed-85fed-5145465cscds"
	user     = "Foo Bar"
)

func createVaultSystem(url string) (auth.Authentication, error) {
	tracing.NewTestTracer("vault")
	return newVaultSystem(&config.Configuration{
		Auth: &config.AuthConfiguration{
			Name: "vault",
			Vault: &config.VaultConfiguration{
				Address:  url,
				Roleid:   "1111",
				Secretid: "2222",
			},
		},
	})
}

func mockHandler() http.Handler {
	handler := http.NewServeMux()
	handler.Handle("/v1/auth/approle/login", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", mimetype)
		w.WriteHeader(200)
		fmt.Fprintln(w, `{"request_id":"7b54b6b3-d349-b202-fcb5-94c18c384034","lease_id":"","renewable":false,"lease_duration":0,"data":null,"wrap_info":null,"warnings":null,"auth":{"client_token":"2ba6658f-9dca-6b65-5c81-df63fb0f4971","accessor":"a4b19994-054a-65f9-9e64-a9d9d14469be","policies":["default","nimbus.admin"],"metadata":{},"lease_duration":1200,"renewable":true}}`)
	}))
	handler.Handle("/v1/secret/nimbus/ft02468", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", mimetype)
		w.WriteHeader(200)
		fmt.Fprintln(w, `{"request_id":"8530d70b-2caf-359e-6e74-ed29e83304ee","lease_id":"","renewable":false,"lease_duration":28800,"data":{"apikey":"cdcsdcs54-545d-48ed-85fed-5145465cscds","cuid":"ft02468","name":"Foo Bar","roles":["user","admin"]},"wrap_info":null,"warnings":null,"auth":null}`)
	}))
	return handler
}

func setupHTTP() (*httptest.Server, *http.Client) {
	server := httptest.NewServer(mockHandler())
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(server.URL)
			},
		},
	}

	return server, client
}

func Test_AppRoleCredentials(t *testing.T) {
	server, _ := setupHTTP()
	defer server.Close()

	vs, err := createVaultSystem(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}

	ctx := context.Background()
	span := tracing.GetParentSpan(ctx, "vault")
	result, err := vs.Credentials(ctx, span, cuid, apiKey)
	if err != nil {
		t.Fatalf("Error with credentials: %s", err.Error())
	}
	if result != apiKey {
		t.Fatalf("Invalid credentials: %s %s", result, apiKey)
	}
}

func Test_AppRoleAuthentication(t *testing.T) {
	server, _ := setupHTTP()
	defer server.Close()

	vs, err := createVaultSystem(server.URL)
	if err != nil {
		t.Fatalf(err.Error())
	}
	headers := map[string]string{
		transport.Authorization: fmt.Sprintf("%s %s", vs.Key(), apiKey),
		transport.UserID:        cuid,
	}
	md := metadata.New(headers)
	ctx := metadata.NewIncomingContext(context.Background(), md)
	span := tracing.GetParentSpan(ctx, "vault")
	if _, ok := metadata.FromIncomingContext(ctx); !ok {
		t.Fatalf("Invalid metadata")
	}
	headers, err = vs.Authenticate(ctx, span, apiKey)
	if err != nil {
		t.Fatalf("Error with authentication: %s", err.Error())
	}
	if headers[transport.Username] != user {
		t.Fatalf("Invalid headers: %s", headers)
	}

}
