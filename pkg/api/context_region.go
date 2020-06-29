// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package api

import "context"

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key string

// regionKey is the key for region values in Contexts. It is
// unexported; clients use user.NewContext and api.WithRegion
// instead of using this key directly.
var regionKey key

// WithRegion creates a new context with a region value. This needs to be used
// when calling any auto-generated platform APIs directly.
// client..Stack.GetVersionStacks(stack.NewGetVersionStacksParams().
//		WithContext(api.WithRegion(context.Background(), "us-east-1")), nil)
func WithRegion(ctx context.Context, region string) context.Context {
	return context.WithValue(ctx, regionKey, region)
}

// GetContextRegion returns the region from a specified context, if any.
func GetContextRegion(ctx context.Context) (string, bool) {
	if ctx == nil {
		return "", false
	}

	// This notation is requiredd in case the context doesn't contain the key
	// so it doesn't panic.
	region, _ := ctx.Value(regionKey).(string)
	return region, region != ""
}
