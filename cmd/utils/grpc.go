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

package utils

import (
	"bytes"
	"errors"
	"net"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/nlamirault/abraracourcix/auth"
	_ "github.com/nlamirault/abraracourcix/auth/basic"
	// _ "github.com/nlamirault/abraracourcix/auth/vault"
	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/tracing"
	"github.com/nlamirault/abraracourcix/transport"
)

var (
	ErrUsernameNotFound    = errors.New("Username not found")
	ErrApiKeyNotFound      = errors.New("API key not found")
	ErrGrpcAddressNotFound = errors.New("gRPC address not found")
)

// GRPCClient define a client using gRPC protocol
type GRPCClient struct {
	ServerAddress  string
	Username       string
	Password       string
	Authentication auth.Authentication
}

// NewGRPCClient creates a new gRPC client
func NewGRPCClient(cmd *cobra.Command) (*GRPCClient, error) {
	setupFromEnvironmentVariables()
	if len(Username) == 0 {
		return nil, ErrUsernameNotFound
	}
	if len(Password) == 0 {
		return nil, ErrApiKeyNotFound
	}
	if len(ServerAddress) == 0 {
		return nil, ErrGrpcAddressNotFound
	}
	conf := &config.Configuration{
		Auth: &config.AuthConfiguration{
			// Name: "vault",
			// Vault: &config.VaultConfiguration{
			// 	Address: "https://vault.io",
			// },
			Name: "BasicAuth",
		},
	}
	authentication, err := auth.New(conf)
	if err != nil {
		return nil, err
	}
	glog.V(2).Infof("gRPC client created: %s %s", ServerAddress, Username)
	return &GRPCClient{
		ServerAddress:  ServerAddress,
		Username:       Username,
		Password:       Password,
		Authentication: authentication,
	}, nil
}

func (client *GRPCClient) GetConn() (*grpc.ClientConn, error) {
	return grpc.Dial(
		client.ServerAddress,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor),
	)
}

func (client *GRPCClient) GetContext(cliName string) (context.Context, error) {
	ctx := context.Background()
	span := tracing.GetParentSpan(ctx, cliName)
	token, err := client.Authentication.Credentials(ctx, span, client.Username, client.Password)
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		transport.Authorization: auth.GetAuthenticationHeader(client.Authentication, token),
	}
	if host, err := os.Hostname(); err != nil {
		headers[transport.UserHostname] = host
	}
	addrs, _ := net.InterfaceAddrs()
	var buffer bytes.Buffer
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				buffer.WriteString(ipnet.IP.String() + " ")
			}
		}
	}
	headers[transport.UserIP] = buffer.String()
	headers[transport.UserID] = client.Username
	md := metadata.New(headers)
	glog.V(2).Infof("Transport metadata: %s", md)
	return metadata.NewContext(ctx, md), nil
}
