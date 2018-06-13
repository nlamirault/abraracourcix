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
	"github.com/golang/glog"
)

// gLogger is implementation of the Logger interface that delegates to `glog` package
var gLogger = &glogLogger{}

type glogLogger struct{}

func (l *glogLogger) Error(msg string) {
	glog.Error(msg)
}

// Infof logs a message at info priority
func (l *glogLogger) Infof(msg string, args ...interface{}) {
	glog.Infof(msg, args...)
}
