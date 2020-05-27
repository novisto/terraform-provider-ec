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

package deploymentstate

import (
	"fmt"

	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deploymentsize"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// MemoryToState parses a megabyte int notation to a gigabyte notation.
func MemoryToState(mem int32) string {
	if mem%1024 > 1 && mem%512 == 0 {
		return fmt.Sprintf("%0.1fg", float32(mem)/1024)
	}
	return fmt.Sprintf("%dg", mem/1024)
}

// ParseTopologySize parses a flattened topology into its model.
func ParseTopologySize(topology map[string]interface{}) (models.TopologySize, error) {
	if mem, ok := topology["memory_per_node"]; ok {
		val, err := deploymentsize.Parse(mem.(string))
		if err != nil {
			return models.TopologySize{}, err
		}

		return models.TopologySize{
			// TODO: For now the resource is assumed to be "memory". This can
			// and will change in the future, we need to accommodate for this case.
			Value: ec.Int32(val), Resource: ec.String("memory"),
		}, nil
	}

	return models.TopologySize{}, nil
}
