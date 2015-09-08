#!/bin/bash

# Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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



APP="abraracourcix"

set -e
if [ -z "$1" ]; then
    echo -e "\033[31;01m[$APP] Pass the version number as the first arg. E.g.: $0 1.2.3 \033[0m"
    exit 1
fi
if [ -z "$GITHUB_TOKEN" ]; then
    echo -e "\033[31;01m[$APP] GITHUB_TOKEN must be set for github-release \033[0m"
    exit 1
fi

VERSION=$1
REPO="abraracourcix"
USERNAME="nlamirault"
OS_PLATFORM_ARG=(-os="darwin linux windows")

# git tag $VERSION
# git push --tags

echo -e "\033[32;01m[$APP] Build image \033[0m"
docker build -t $(REPO)/release .

echo -e "\033[32;01m[$APP] Make binaries \033[0m"
docker run --rm \
       -v `pwd`:/tmp/ \
       $(REPO)/release \
       gox "${OS_PLATFORM_ARG[@]}" "${OS_ARCH_ARG[@]}" -output="/tmp/abraracourcix_{{.OS}}-{{.Arch}}"

echo -e "\033[32;01m[$APP] Make release \033[0m"
docker run --rm -e GITHUB_TOKEN $(REPO)/release \
    github-release release \
    --user $USERNAME \
    --repo $REPO \
    --tag $VERSION \
    --name $VERSION \
    --description ""

echo -e "\033[32;01m[$APP] Upload archive \033[0m"
for BINARY in abraracourcix_*; do
    docker run --rm -e GITHUB_TOKEN -v `pwd`:/go/src/github.com/nlamirault/abraracourcix \
           $(REPO)/release github-release upload \
           --user $USERNAME \
           --repo $REPO \
           --tag $VERSION \
           --name $BINARY \
           --file $BINARY
done
