// Copyright 2019 Yunion
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

package worker

import (
	"context"
	"fmt"

	"yunion.io/x/onecloud/pkg/appsrv"
	"yunion.io/x/onecloud/pkg/vpcagent/options"
)

var workers = map[string]NewWorkerFunc{}

type NewWorkerFunc func(opts *options.Options) IWorker

type IWorker interface {
	Start(ctx context.Context, app *appsrv.Application, prefix string)
}

func NewWorker(opts *options.Options) IWorker {
	n, ok := workers[opts.VpcProvider]
	if !ok {
		return nil
	}
	return n(opts)
}

func RegisterNewWorkerFunc(name string, n NewWorkerFunc) {
	if _, ok := workers[name]; ok {
		panic(fmt.Sprintf("worker %s already registered", name))
	}
	workers[name] = n
}
