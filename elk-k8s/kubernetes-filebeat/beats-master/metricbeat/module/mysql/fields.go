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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mysql

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mysql", asset.ModuleFieldsPri, AssetMysql); err != nil {
		panic(err)
	}
}

// AssetMysql returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/mysql.
func AssetMysql() string {
	return "eJzkXF1z2zazvs+v2MlN0xnb0+tcCRIAR-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"
}
