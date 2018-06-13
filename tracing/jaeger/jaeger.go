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

package jaeger

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	// jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/nlamirault/abraracourcix/config"
	"github.com/nlamirault/abraracourcix/tracing"
)

const (
	label = "jaeger"
)

func init() {
	tracing.RegisterTracer(label, newTracer)
}

func newTracer(conf *config.Configuration) (opentracing.Tracer, error) {
	glog.V(1).Infof("Create OpenTracing tracer using Jaeger: %s %d", conf.Tracing.Jaeger.Host, conf.Tracing.Jaeger.Port)

	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%s:%d", conf.Tracing.Jaeger.Host, conf.Tracing.Jaeger.Port),
		},
	}
	// jLogger := jaegerlog.StdLogger
	jLogger := gLogger
	jMetricsFactory := metrics.NullFactory
	// Initialize tracer with a logger and a metrics factory
	tracer, _, err := cfg.New(
		tracing.ServiceName,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return nil, err
	}
	// defer closer.Close()
	return tracer, nil
}
