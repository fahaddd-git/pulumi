// Copyright 2016-2019, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/pulumi/pulumi/pkg/secrets/service"

	"github.com/pulumi/pulumi/pkg/backend/httpstate"
	"github.com/pulumi/pulumi/sdk/pulumi/secrets"
)

func newServiceSecretsManager(s httpstate.Stack) (secrets.Manager, error) {
	client := s.Backend().(httpstate.Backend).Client()
	id := s.StackIdentifier()

	return service.NewServiceSecretsManager(client, id)
}
