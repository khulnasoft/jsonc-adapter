// Copyright 2015 Matthew Holt and The Kengine Authors
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

package jsoncadapter

import (
	"encoding/json"

	"github.com/khulnasoft/kengine/v2/kengineconfig"
	"github.com/muhammadmuzzammil1998/jsonc"
)

func init() {
	kengineconfig.RegisterAdapter("jsonc", Adapter{})
}

// Adapter adapts JSON-C to Kengine JSON.
type Adapter struct{}

// Adapt converts the JSON-C config in body to Kengine JSON.
func (a Adapter) Adapt(body []byte, options map[string]interface{}) (result []byte, warnings []kengineconfig.Warning, err error) {
	result = jsonc.ToJSON(body)

	// any errors in the JSON will be
	// reported during config load, but
	// we can at least warn here that
	// it is not valid JSON
	if !json.Valid(result) {
		warnings = append(warnings, kengineconfig.Warning{
			Message: "Resulting JSON is invalid.",
		})
	}

	return
}

// Interface guard
var _ kengineconfig.Adapter = (*Adapter)(nil)
