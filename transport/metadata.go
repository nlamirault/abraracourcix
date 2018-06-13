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

package transport

import (
	"fmt"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	// UserID define the key to setup the ID of the hostname
	UserID = "user.id"

	// UserIP define the key to setup the user IP address into the context header
	UserIP = "user.ip"

	// UserRoles define roles associated to the user
	UserRoles = "user.roles"

	// Username is the common user name
	Username = "user.name"

	// UserHostname define the key to setup the user hostname into the context header
	UserHostname = "user.hostname"

	// Authorization define the key into the context header
	Authorization = "authorization"
)

func GetFromMetadata(md metadata.MD, key string) (string, error) {
	glog.V(2).Infof("Metadata %s get %s", md, key)
	slice, ok := md[key]
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "No value specified in metadata")
	}
	if len(slice) != 1 {
		return "", status.Errorf(codes.Unauthenticated, fmt.Sprintf("Request contains invalid apiKey: %s", slice))
	}
	return slice[0], nil
}

func ExtractMetadata(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	glog.V(2).Infof("Metadata from context: %s", md)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "gRPC context lacks metadata")
	}
	return GetFromMetadata(md, key)
}
