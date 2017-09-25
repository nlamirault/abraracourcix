# Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM alpine:latest
MAINTAINER Nicolas Lamirault <nicolas.lamirault@gmail.com>

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN	apk add --no-cache \
	ca-certificates

COPY . /go/src/github.com/nlamirault/abraracourcix

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		go \
		git \
		gcc \
		libc-dev \
		libgcc \
	&& cd /go/src/github.com/nlamirault/abraracourcix \
	&& go build -o /usr/bin/abraracourcixd github.com/nlamirault/abraracourcix/cmd/abraracourcixd \
    && go build -o /usr/bin/abraracourcixctl github.com/nlamirault/abraracourcix/cmd/abraracourcixctl \
    && go build -o /usr/bin/abraracourcixadm github.com/nlamirault/abraracourcix/cmd/abraracourcixadm \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."

VOLUME ["/etc/abraracourcix"]
