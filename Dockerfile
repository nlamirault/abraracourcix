# Copyright (C) 2015-2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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

FROM alpine:3.18.3

LABEL maintainer="Nicolas Lamirault <nicolas.lamirault@gmail.com>" \
      summary="A simple URL Shortener" \
      description="A simple URL Shortener" \
      name="nlamirault/abraracourcix" \
      url="https://github.com/nlamirault/abraracourcix" \
      io.k8s.description="A simple URL Shortener" \
      io.k8s.display-name="abraracourcix"

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
        bash \
        build-base \
	&& cd /go/src/github.com/nlamirault/abraracourcix \
    && make build \
    && cp abraracourcixd /usr/bin \
    && cp abraracourcixadm /usr/bin/ \
    && cp abraracourcixctl /usr/bin \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."

VOLUME ["/etc/abraracourcix"]
