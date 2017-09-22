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

	"github.com/nlamirault/abraracourcix/io"
	"github.com/nlamirault/abraracourcix/messaging"
	"github.com/nlamirault/abraracourcix/pb/v2beta"
	"github.com/nlamirault/abraracourcix/storage"
	"github.com/nlamirault/abraracourcix/tracing"
)

const (
	UrlServiceName = "UrlService"
)

type UrlService struct {
	Storage storage.Storage
}

func NewUrlService(storage storage.Storage) *UrlService {
	glog.V(2).Infof("Create the Url service using %v", storage)
	service := &UrlService{
		Storage: storage,
	}
	service.Register()
	return service
}

func (service *UrlService) Register() {
	Services = append(Services, UrlServiceName)
}

func (ls *UrlService) List(ctx context.Context, request *v2beta.GetUrlsRequest) (*v2beta.GetUrlsResponse, error) {
	glog.V(1).Info("[url] List all urls")
	span := tracing.GetParentSpan(ctx, messaging.ListUrlsEvent)
	defer span.Finish()

	urlKeys, err := ls.Storage.List()
	if err != nil {
		return nil, tracing.GrpcError(span, err)
	}
	span.LogFields(log.Object("storage response", urlKeys))
	keys := []string{}
	for _, key := range urlKeys {
		keys = append(keys, string(key))
	}

	return &v2beta.GetUrlsResponse{
		Keys: keys,
	}, nil
}

func (ls *UrlService) Create(ctx context.Context, request *v2beta.CreateUrlRequest) (*v2beta.CreateUrlResponse, error) {
	glog.V(1).Info("[url] Create a new url")
	span := tracing.GetParentSpan(ctx, messaging.CreateUrlEvent)
	defer span.Finish()

	key, err := io.GenerateKey()
	if err != nil {
		return nil, tracing.GrpcError(span, err)
	}

	url := &storage.URL{
		LongURL:      request.Link,
		Key:          key,
		CreationDate: io.GetCreationDate(),
	}

	content, err := defaultMarshaler.Marshal(url)
	if err != nil {
		return nil, tracing.GrpcError(span, err)
	}
	if err := ls.Storage.Put([]byte(key), content); err != nil {
		return nil, tracing.GrpcError(span, err)
	}
	return &v2beta.CreateUrlResponse{
		Url: &v2beta.Url{
			Key:      url.Key,
			Link:     url.LongURL,
			Creation: fmt.Sprintf("%s", url.CreationDate),
		},
	}, nil
}

func (ls *UrlService) Get(ctx context.Context, request *v2beta.GetUrlRequest) (*v2beta.GetUrlResponse, error) {
	glog.V(1).Info("[url] Retrieve a url")
	span := tracing.GetParentSpan(ctx, messaging.GetUrlEvent)
	defer span.Finish()

	content, err := ls.Storage.Get([]byte(request.Key))
	if err != nil {
		return nil, tracing.GrpcError(span, err)
	}
	span.LogFields(log.Object("storage response", string(content)))

	var url *storage.URL
	if err := defaultMarshaler.Unmarshal(content, &url); err != nil {
		return nil, tracing.GrpcError(span, err)
	}

	return &v2beta.GetUrlResponse{
		Url: &v2beta.Url{
			Key:      url.Key,
			Link:     url.LongURL,
			Creation: fmt.Sprintf("%s", url.CreationDate),
		},
	}, nil
}
