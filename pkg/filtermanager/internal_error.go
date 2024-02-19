// Copyright The HTNN Authors.
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

package filtermanager

import "mosn.io/htnn/pkg/filtermanager/api"

type internalErrorFilter struct {
	api.PassThroughFilter
}

func (f *internalErrorFilter) DecodeHeaders(headers api.RequestHeaderMap, endStream bool) api.ResultAction {
	return &api.LocalResponse{
		Code: 500,
	}
}

func InternalErrorFactory(interface{}, api.FilterCallbackHandler) api.Filter {
	return &internalErrorFilter{}
}