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

package tracing

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/blog/content/context/userip"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/transport"
)

const (
	// ServiceName used to setup the tracer
	ServiceName string = "abraracourcix"
)

type TracerFunc func(conf *config.Configuration) (opentracing.Tracer, error)

var registeredTracers = map[string](TracerFunc){}

func RegisterTracer(name string, f TracerFunc) {
	registeredTracers[name] = f
}

func New(conf *config.Configuration) (opentracing.Tracer, error) {
	glog.V(1).Infof("Opentracing setup: %s", conf.Tracing)
	f, ok := registeredTracers[conf.Tracing.Name]
	if !ok {
		return nil, fmt.Errorf("Unsupported tracer: %s", conf.Tracing.Name)
	}
	tracer, err := f(conf)
	if err != nil {
		return nil, err
	}

	// explicitly set our tracer to be the default tracer.
	opentracing.SetGlobalTracer(tracer)
	return tracer, nil
}

func GetParentSpan(ctx context.Context, operationName string) opentracing.Span {
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan == nil {
		glog.V(2).Infof("Create parent span for service")
		parentSpan, ctx = opentracing.StartSpanFromContext(ctx, operationName)
	}
	if userIP, ok := userip.FromContext(ctx); ok {
		parentSpan.SetTag("user.remote_ip", userIP.String())
	}
	if userID := ctx.Value(transport.UserID); userID != nil {
		parentSpan.SetTag("user.id", string([]byte(userID.(string))))
	}
	if md, ok := metadata.FromContext(ctx); ok {
		if hostname, ok := md[transport.UserHostname]; ok {
			parentSpan.SetTag("user.hostname", hostname)
		}
		if ip, ok := md[transport.UserIP]; ok {
			parentSpan.SetTag("user.ip", ip)
		}
	}
	return parentSpan
}
func GetChildSpan(span opentracing.Span, operationName string) opentracing.Span {
	return opentracing.StartSpan(operationName, opentracing.ChildOf(span.Context()))
}

func SetSpanTags(span opentracing.Span, provider string, url string, method string) {
	span.SetTag(string(ext.Component), provider)
	span.SetTag(string(ext.HTTPUrl), url)
	span.SetTag(string(ext.HTTPMethod), method)
}

func SetResponse(span opentracing.Span, response string) {
	span.SetTag("response", response)
	span.LogFields(log.String("response", response))
}

func GrpcError(span opentracing.Span, err error) error {
	glog.V(0).Infof(err.Error())
	span.LogFields(log.Error(err))
	return status.Errorf(codes.Internal, err.Error())
}

func Errorf(span opentracing.Span, err error) error {
	span.LogFields(log.Error(err))
	return err
}
