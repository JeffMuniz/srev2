// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "eJzMWV9z27gRf8+n2MlL7RlJCRIAR-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"
}
