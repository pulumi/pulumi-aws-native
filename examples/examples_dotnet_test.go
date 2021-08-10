
// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
// +build dotnet all

package examples

import (
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"path/filepath"
	"testing"
)

func TestSimpleCs(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "simple-cs"),
			SkipRefresh: true,
		})

	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.AwsNative",
		},
	})

	return baseCsharp
}
