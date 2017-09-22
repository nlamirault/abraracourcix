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

package abraracourcixd

import (
	"github.com/golang/glog"
	"github.com/mwitkow/go-grpc-middleware/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"github.com/nlamirault/abraracourcix/auth"
	_ "github.com/nlamirault/abraracourcix/auth/basic"
	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/tracing"
)

type serverAuthentication struct {
	Authentication auth.Authentication
}

func newServerAuthentication(conf *config.Configuration) (*serverAuthentication, error) {
	authentication, err := auth.New(conf)
	if err != nil {
		return nil, err
	}
	return &serverAuthentication{
		Authentication: authentication,
	}, nil
}

func (sa *serverAuthentication) authenticate(ctx context.Context) (context.Context, error) {
	glog.V(2).Infof("Check authentication using %s", sa.Authentication.Key())

	span := tracing.GetParentSpan(ctx, "authenticate")
	defer span.Finish()

	token, err := grpc_auth.AuthFromMD(ctx, sa.Authentication.Key())
	if err != nil {
		return nil, err
	}
	headers, err := sa.Authentication.Authenticate(ctx, span, token)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, err.Error())
	}
	glog.V(2).Infof("Authentication add headers: %s", headers)
	return metadata.NewIncomingContext(ctx, metadata.New(headers)), nil
}
