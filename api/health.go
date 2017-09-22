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
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"

	"github.com/nlamirault/abraracourcix/auth"
	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/messaging"
	"github.com/nlamirault/abraracourcix/pb/health"
	"github.com/nlamirault/abraracourcix/pkg/rbac"
	"github.com/nlamirault/abraracourcix/tracing"
	"github.com/nlamirault/abraracourcix/transport"
)

type HealthService struct {
	Authentication auth.Authentication
	HealthUser     string
	HealthKey      string
	URI            string
	Services       []string
}

func NewHealthService(conf *config.Configuration, uri string, services []string) (*HealthService, error) {
	glog.V(2).Info("Create the health service")
	rbac.AddRoles("health", "HealthService", map[string][]string{
		"Status": []string{rbac.AdminRole},
	})
	authentication, err := auth.New(conf)
	if err != nil {
		return nil, err
	}
	return &HealthService{
		// Conf:     conf,
		Authentication: authentication,
		HealthKey:      conf.Auth.Vault.HealthKey,
		HealthUser:     conf.Auth.Vault.HealthUser,
		URI:            uri,
		Services:       services,
	}, nil
}

func (service *HealthService) Status(ctx context.Context, req *health.StatusRequest) (*health.StatusResponse, error) {
	glog.V(1).Info("Check Health services")

	span := tracing.GetParentSpan(ctx, messaging.HealthEvent)
	defer span.Finish()

	conn, err := grpc.Dial(service.URI, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := healthpb.NewHealthClient(conn)
	token, err := service.Authentication.Credentials(ctx, span, service.HealthUser, service.HealthKey)
	if err != nil {
		return nil, err
	}
	glog.V(2).Infof("Auth token: %s", token)
	md := metadata.New(map[string]string{
		transport.Authorization: auth.GetAuthenticationHeader(service.Authentication, token),
		transport.UserID:        service.HealthUser,
	})
	newCtx := metadata.NewContext(ctx, md)

	servicesStatus := []*health.ServiceStatus{}
	for _, service := range service.Services {
		glog.V(2).Infof("Check health service: %s", service)
		resp, err := client.Check(newCtx, &healthpb.HealthCheckRequest{
			Service: service,
		})
		if err != nil {
			servicesStatus = append(servicesStatus, &health.ServiceStatus{
				Name:   service,
				Status: "KO",
				Text:   err.Error(),
			})
		} else {
			servicesStatus = append(servicesStatus, &health.ServiceStatus{
				Name:   service,
				Status: "OK",
				Text:   fmt.Sprintf("%s", resp.Status),
			})
		}
	}

	resp := &health.StatusResponse{}
	resp.Services = servicesStatus

	glog.V(0).Infof("Health response: %s", resp)
	span.LogFields(log.Object("response", resp))
	return resp, nil
}
