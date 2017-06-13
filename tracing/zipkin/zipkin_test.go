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

package zipkin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/opentracing/opentracing-go"

	"github.com/nlamirault/abraracourcix/config"
)

type FakeServer struct {
	*httptest.Server
}

func handler(server *FakeServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}

func NewFakeServer() *FakeServer {
	h := &FakeServer{}
	h.Server = httptest.NewServer(handler(h))
	return h
}

func Test_ZipkinWithoutConfiguration(t *testing.T) {
	if _, err := newTracer(&config.Configuration{
		Tracing: &config.TracingConfiguration{
			Name: "zipkin",
		},
	}); err == nil {
		t.Fatalf("Can create tracer without configuration: %s", err)
	}
}

func Test_ZipkinTracer(t *testing.T) {
	server := NewFakeServer()
	defer server.Server.Close()

	tracer, err := newTracer(&config.Configuration{
		Tracing: &config.TracingConfiguration{
			Name: "zipkin",
			Zipkin: &config.ZipkinConfiguration{
				Host: server.Listener.Addr().String(),
			},
		},
	})
	if err != nil {
		t.Fatalf("Can't create tracer: %s", err)
	}

	span := tracer.StartSpan("root_span")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", server.Server.URL, nil)
	if err := tracer.Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header)); err != nil {
		t.Fatalf("InjectSpan returns an error: %s", err)
	}
	client.Do(req)
	span.Finish()
}
