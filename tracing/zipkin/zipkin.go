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

package zipkin

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go-opentracing"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/tracing"
)

const (
	zipkinTracerLabel = "zipkin"

	// Set to true for RPC style spans (Zipkin V1) vs Node style (OpenTracing)
	sameSpan = false
)

func init() {
	tracing.RegisterTracer(zipkinTracerLabel, newTracer)
}

func newTracer(conf *config.Configuration) (opentracing.Tracer, error) {
	glog.V(1).Infof("Create OpenTracing tracer using Zipkin: %s", conf.Tracing.Zipkin)
	if conf.Tracing.Zipkin == nil {
		return nil, fmt.Errorf("No configuration for Zipkin Opentracing: %s", conf)
	}
	collector, err := zipkintracer.NewHTTPCollector(
		fmt.Sprintf("http://%s:%d/api/v1/spans", conf.Tracing.Zipkin.Host, conf.Tracing.Zipkin.Port))
	if err != nil {
		return nil, err
	}
	tracer, err := zipkintracer.NewTracer(
		zipkintracer.NewRecorder(collector, true, fmt.Sprintf("%s:0", conf.Tracing.Zipkin.Host), tracing.ServiceName),
		zipkintracer.ClientServerSameSpan(false))
	if err != nil {
		return nil, err
	}
	return tracer, nil
}
