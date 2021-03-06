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

APP = abraracourcix

VERSION=$(shell \
        grep "const Version" version/version.go \
        |awk -F'=' '{print $$2}' \
        |sed -e "s/[^0-9.]//g" \
	|sed -e "s/ //g")

SHELL = /bin/bash

DIR = $(shell pwd)

DOCKER = docker

GO = go

GOX = gox -os="linux darwin windows freebsd openbsd netbsd"
GOX_ARGS = "-output={{.Dir}}-$(VERSION)_{{.OS}}_{{.Arch}}"

BINTRAY_URI = https://api.bintray.com
BINTRAY_USERNAME = nlamirault
BINTRAY_ORG = nlamirault
BINTRAY_REPOSITORY= oss

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

MAKE_COLOR=\033[33;01m%-20s\033[0m

MAIN = github.com/nlamirault/abraracourcix
SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
PKGS = $(shell glide novendor)
EXE = $(shell ls abraracourcixd-*_* abraracourcixctl-*_* abraracourcixadm-*_*)

PACKAGE=$(APP)-$(VERSION)
ARCHIVE=$(PACKAGE).tar

GOX_ARGS = "-output={{.Dir}}-$(VERSION)_{{.OS}}_{{.Arch}}"

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(MAKE_COLOR) : %s\n", $$1, $$2}'

clean: ## Cleanup
	@echo -e "$(OK_COLOR)[$(APP)] Cleanup$(NO_COLOR)"
	@rm -fr abraracourcixadm abraracourcixctl abraracourcixd *.tar.gz

.PHONY: tools
tools: ## Install tools
	@echo -e "$(OK_COLOR)[$(APP)] Install requirements$(NO_COLOR)"
	@go get -u github.com/golang/glog
	@go get -u github.com/kardianos/govendor
	@go get -u github.com/Masterminds/rmvcsdir
	@go get -u github.com/golang/lint/golint
	@go get -u github.com/kisielk/errcheck
	@go get -u github.com/mitchellh/gox
	@go get -u gopkg.in/mikedanese/gazel.v17/gazel
	@wget https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip

.PHONY: proto
proto: ## Install protocol buffer tools
	@go get -u github.com/golang/protobuf/protoc-gen-go
	@go get -u github.com/golang/protobuf/proto
	@go install ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/
	@go install ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

init: tools proto ## Install requirements

.PHONY: deps
deps: ## Install dependencies
	@echo -e "$(OK_COLOR)[$(APP)] Update dependencies$(NO_COLOR)"
	@govendor update

.PHONY: pb
pb: ## Generate Protobuf
	@go generate pb/api.go

.PHONY: swagger
swagger: ## Generate Swagger
	go-bindata-assetfs -pkg swagger third_party/swagger-ui/... && mv bindata_assetfs.go pkg/ui/swagger/

.PHONY: changelog
changelog:
	@$(GO) generate -x ./pkg/static/

.PHONY: doc
doc: ## Generate documentation
	@echo -e "$(OK_COLOR)[$(APP)] Documentation $(NO_COLOR)"
	docker run -it -v `pwd`/doc:/documents/ rochdev/alpine-asciidoctor asciidoctor index.adoc -a stylesheet=dt-oab.css -a stylesdir=/documents/stylesheets -d book -a toc2

.PHONY: webdoc
webdoc: doc
	@$(GO) generate -x ./pkg/webdoc/

.PHONY: build
build: ## Make binary
	@echo -e "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	@$(GO) build -o abraracourcixd github.com/nlamirault/abraracourcix/cmd/abraracourcixd
	@$(GO) build -o abraracourcixctl github.com/nlamirault/abraracourcix/cmd/abraracourcixctl
	@$(GO) build -o abraracourcixadm github.com/nlamirault/abraracourcix/cmd/abraracourcixadm

.PHONY: test
test: ## Launch unit tests
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@govendor test +local

.PHONY: lint
lint: ## Launch golint
	@$(foreach file,$(SRCS),golint $(file) || exit;)

.PHONY: vet
vet: ## Launch go vet
	@$(foreach file,$(SRCS),$(GO) vet $(file) || exit;)

.PHONY: errcheck
errcheck: ## Launch go errcheck
	@echo -e "$(OK_COLOR)[$(APP)] Go Errcheck $(NO_COLOR)"
	@$(foreach pkg,$(PKGS),errcheck $(pkg) $(glide novendor) || exit;)

.PHONY: coverage
coverage: ## Launch code coverage
	@$(foreach pkg,$(PKGS),$(GO) test -cover $(pkg) $(glide novendor) || exit;)

gox: ## Make all binaries
	@echo -e "$(OK_COLOR)[$(APP)] Create binaries $(NO_COLOR)"
	$(GOX) -output=abraracourcixctl-$(VERSION)_{{.OS}}_{{.Arch}} -osarch="linux/amd64 darwin/amd64 windows/amd64" github.com/nlamirault/abraracourcix/cmd/abraracourcixctl
	$(GOX) -output=abraracourcixadm-$(VERSION)_{{.OS}}_{{.Arch}} -osarch="linux/amd64 darwin/amd64 windows/amd64" github.com/nlamirault/abraracourcix/cmd/abraracourcixadm
	$(GOX) -output=abraracourcixd-$(VERSION)_{{.OS}}_{{.Arch}} -osarch="linux/amd64 darwin/amd64 windows/amd64" github.com/nlamirault/abraracourcix/cmd/abraracourcixd

.PHONY: binaries
binaries: ## Upload all binaries
	@echo -e "$(OK_COLOR)[$(APP)] Upload binaries to Bintray $(NO_COLOR)"
	for i in $(EXE); do \
		curl -T $$i \
			-u$(BINTRAY_USERNAME):$(BINTRAY_APIKEY) \
			"$(BINTRAY_URI)/content/$(BINTRAY_ORG)/$(BINTRAY_REPOSITORY)/$(APP)/${VERSION}/$$i;publish=1"; \
        done

# for goprojectile
.PHONY: gopath
gopath:
	@echo `pwd`:`pwd`/vendor
