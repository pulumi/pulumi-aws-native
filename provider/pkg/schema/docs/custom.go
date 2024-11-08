// Copyright 2016-2024, Pulumi Corporation.

package docs

import (
	_ "embed"
)

// TODO[pulumi/pulumi-cdk#109] Add examples for the other languages.
//go:embed content/cfn-custom-resource.md
var CfnCustomResourceEmulatorDocs string
