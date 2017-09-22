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
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	ghealth "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/nlamirault/abraracourcix/api"
	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/middleware"
	"github.com/nlamirault/abraracourcix/pb/health"
	"github.com/nlamirault/abraracourcix/pb/info"
	"github.com/nlamirault/abraracourcix/pb/v2beta"
	"github.com/nlamirault/abraracourcix/storage"
)

func registerServer(backend storage.Storage, serverAuth *serverAuthentication, tracer opentracing.Tracer, conf *config.Configuration, grpcAddr string) (*grpc.Server, error) {
	glog.V(1).Info("Create the gRPC server")
	tagsOpts := []grpc_ctxtags.Option{
		grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.TagBasedRequestFieldExtractor("log_fields")),
	}
	server := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_ctxtags.StreamServerInterceptor(tagsOpts...),
				grpc_opentracing.StreamServerInterceptor(),
				grpc_prometheus.StreamServerInterceptor)),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				middleware.ServerLoggingInterceptor(true),
				grpc_ctxtags.UnaryServerInterceptor(tagsOpts...),
				grpc_prometheus.UnaryServerInterceptor,
				grpc_opentracing.UnaryServerInterceptor(),
				grpc_auth.UnaryServerInterceptor(serverAuth.authenticate))),
	)

	v2beta.RegisterUrlServiceServer(server, api.NewUrlService(backend))

	info.RegisterInfoServiceServer(server, api.NewInfoService(conf))
	healthService, err := api.NewHealthService(conf, grpcAddr, api.Services)
	if err != nil {
		return nil, err
	}
	health.RegisterHealthServiceServer(server, healthService)

	healthServer := ghealth.NewServer()
	healthpb.RegisterHealthServer(server, healthServer)
	for _, service := range api.Services {
		healthServer.SetServingStatus(service, healthpb.HealthCheckResponse_SERVING)
	}

	grpc_prometheus.Register(server)

	return server, nil
}

func registerGateway(ctx context.Context, addr string) (*runtime.ServeMux, error) {
	glog.V(1).Info("Create the REST gateway")
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	gwmux := runtime.NewServeMux()
	if err := v2beta.RegisterUrlServiceHandlerFromEndpoint(ctx, gwmux, addr, opts); err != nil {
		return nil, err
	}
	return gwmux, nil
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
