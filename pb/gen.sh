#!/usr/bin/env bash

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

# generate the gRPC code

function generate_grpcgw {
    pushd $1
    rm -rf *.pb.go
    protoc -I/usr/local/include \
           -I. -I${GOPATH}/src \
           -I../../vendor/github.com/googleapis/googleapis \
           --go_out=plugins=grpc:. *.proto

    rm -rf *.pb.gw.go
    protoc -I /usr/local/include -I . \
           -I ${GOPATH}/src \
           -I../../vendor/github.com/googleapis/googleapis \
           --grpc-gateway_out=logtostderr=true:. *.proto

    rm -rf ../swagger/*.swagger.json
    protoc -I /usr/local/include -I . \
           -I ${GOPATH}/src \
           -I../../vendor/github.com/googleapis/googleapis \
           --swagger_out=logtostderr=true:. *.proto
    popd
}

function generate_grpc {
    pushd $1
    rm -rf *.pb.go
    protoc -I/usr/local/include \
           -I. -I${GOPATH}/src \
           -I../../vendor/github.com/googleapis/googleapis \
           --go_out=plugins=grpc:. *.proto
    popd
}


function generate_swagger {
    find . -name "*.json" | xargs -I '{}' mv '{}' swagger/
    rm -f swagger/api.swagger.json
    ls swagger
    go run swagger/swagger.go swagger > swagger/api.swagger.json
}

generate_grpcgw v2beta
generate_grpc health
generate_grpc info

generate_swagger
