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

package vault

import (
	"bytes"
	"crypto/tls"
	// "encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/vault/api"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/nlamirault/abraracourcix/auth"
	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/tracing"
	"github.com/nlamirault/abraracourcix/transport"
)

const (
	label = "vault"

	key = "Vault"

	namespace = "/secret/nimbus"

	authretries       = 10
	retrydelayseconds = 3

	InvalidAPIKey         = "API key invalid"
	NoAuthenticationName  = "Authentication name is not set"
	NoAuthenticationRoles = "Authentication roles are not set"
)

type vaultSystem struct {
	Client   *api.Client
	Token    string
	RoleID   string
	SecretID string
}

func init() {
	auth.RegisterAuthentication(label, newVaultSystem)
}

func newVaultSystem(conf *config.Configuration) (auth.Authentication, error) {
	glog.V(1).Infof("Configure Vault using: %s", conf.Auth.Vault)

	if len(conf.Auth.Vault.Address) == 0 {
		return nil, fmt.Errorf("Invalid Vault host: %s", conf.Auth.Vault)
	}
	cfg := &api.Config{
		Address:    conf.Auth.Vault.Address,
		HttpClient: cleanhttp.DefaultClient(),
		MaxRetries: 3,
	}

	cfg.HttpClient.Timeout = time.Second * 60
	httpTransport := cfg.HttpClient.Transport.(*http.Transport)
	httpTransport.TLSHandshakeTimeout = 10 * time.Second
	httpTransport.TLSClientConfig = &tls.Config{
		// MinVersion: tls.VersionTLS12,
		InsecureSkipVerify: true,
	}
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &vaultSystem{
		Client:   client,
		RoleID:   conf.Auth.Vault.Roleid,
		SecretID: conf.Auth.Vault.Secretid,
	}, nil
}

func (vs vaultSystem) Name() string {
	return label
}

func (vs vaultSystem) Key() string {
	return key
}

func (vs vaultSystem) Credentials(ctx context.Context, parentSpan opentracing.Span, cuid string, apiKey string) (string, error) {
	glog.V(1).Infof("Set credentials: %s %s", cuid, apiKey)
	// auth := cuid + ":" + apiKey
	// token := base64.StdEncoding.EncodeToString([]byte(auth))
	// return token, nil
	return apiKey, nil
}

func (vs vaultSystem) Authenticate(ctx context.Context, parentSpan opentracing.Span, apiKey string) (map[string]string, error) {
	glog.V(1).Infof("Check Vault token: %s", apiKey)

	span := tracing.GetChildSpan(parentSpan, "vault_authenticate")
	defer span.Finish()
	span.SetTag("provider", "vault")

	// b, err := base64.StdEncoding.DecodeString(token)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Unauthenticated, "Can't check Vault authentication: %s", err)
	// }
	// pair := strings.SplitN(string(b), ":", 2)
	// if len(pair) != 2 {
	// 	return nil, status.Errorf(codes.Unauthenticated, "Invalid gRPC metadata token")
	// }
	// cuid := pair[0]
	// apiKey := pair[1]

	if err := vs.login(); err != nil {
		return nil, err
	}

	cuid, err := transport.ExtractMetadata(ctx, transport.UserID)
	if err != nil {
		return nil, err
	}
	span.LogFields(log.String("cuid", cuid))
	entry := fmt.Sprintf("%s/%s", namespace, cuid)
	glog.V(2).Infof("Check authentication for %s", entry)
	secret, err := vs.Client.Logical().Read(entry)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}
	if secret == nil || secret.Data == nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid entry: %s", entry))
	}
	glog.V(2).Infof("Vault data: %#v", secret.Data)
	apiKeyAuth, err := getEntry(secret, "apikey")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprintf("Invalid authentication with this secret. %s %s", cuid, secret.Data))
	}
	if apiKeyAuth != apiKey {
		return nil, status.Errorf(codes.Unauthenticated, InvalidAPIKey)
	}
	name, err := getEntry(secret, "name")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, NoAuthenticationName)
	}
	roles, err := getEntries(secret, "roles")
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, NoAuthenticationRoles)
	}
	glog.V(2).Infof("User correctly authenticated %s %s", cuid, roles)
	headers := map[string]string{}
	headers[transport.Username] = name
	headers[transport.UserRoles] = roles
	headers[transport.UserID] = cuid
	span.LogFields(log.String("name", name))
	span.LogFields(log.Object("roles", roles))
	parentSpan.SetTag("user", name)
	parentSpan.SetTag("roles", roles)
	return headers, nil
}

func (vs vaultSystem) login() error {
	glog.V(2).Infof("Login into Vault")
	if len(vs.RoleID) == 0 ||
		len(vs.SecretID) == 0 {
		return status.Errorf(codes.Internal, "Invalid Vault configuration")
	}

	loginRequest := map[string]interface{}{
		"role_id":   vs.RoleID,
		"secret_id": vs.SecretID,
	}
	glog.V(2).Infof("Retrieve token")
	secretLogin, err := vs.Client.Logical().Write("auth/approle/login", loginRequest)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	vs.Client.SetToken(secretLogin.Auth.ClientToken)
	return nil
}

func getEntry(secret *api.Secret, key string) (string, error) {
	value, ok := secret.Data[key].(string)
	if !ok {
		return "", fmt.Errorf("Unknown entry")
	}
	return value, nil
}

func getEntries(secret *api.Secret, key string) (string, error) {
	var buffer bytes.Buffer
	values, ok := secret.Data[key].([]interface{})
	if !ok {
		return "", fmt.Errorf("Unknown entry")
	}
	for _, role := range values {
		glog.V(2).Infof("Vault role: %s", role)
		buffer.WriteString(fmt.Sprintf("%s ", role))
	}
	return strings.TrimSpace(buffer.String()), nil
}
