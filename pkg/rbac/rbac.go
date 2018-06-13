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

package rbac

import (
	"fmt"

	"github.com/golang/glog"
)

const (
	AdminRole = "admin"

	UserRole = "user"

	GrpcHealthKey = "/grpc.health.v1.Health/Check"

	HealthRole = "health"
)

var userRights = map[string][]string{}

func Roles() map[string][]string {
	return userRights
}

func AddRoles(api string, service string, roles map[string][]string) {
	for name, role := range roles {
		userRights[fmt.Sprintf("/%s.%s/%s", api, service, name)] = role
	}
}

func HasRights(key string, roles []string) error {
	glog.V(2).Infof("Check Roles: %s %s %s", key, roles, userRights[key])
	if key == GrpcHealthKey && len(roles) == 1 && HealthRole == roles[0] {
		return nil
	}
	for _, value := range roles {
		for _, val := range userRights[key] {
			if val == value {
				return nil
			}
		}
	}
	return fmt.Errorf("Invalid user rights: %s", roles)
}
