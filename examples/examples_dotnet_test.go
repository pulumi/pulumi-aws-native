// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build dotnet || all
// +build dotnet all

package examples

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/require"
)

func TestSimpleCs(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "simple-cs"),
			SkipRefresh: true,
		})

	integration.ProgramTest(t, &test)
}

func TestRoleCs(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "role-cs"),
			SkipRefresh: true,
		})

	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	if os.Getenv("PULUMI_LOCAL_NUGET") == "" {
		localNugetDir, err := filepath.Abs("../nuget")
		if err != nil {
			t.Fatalf("Failed to get absolute path to nuget directory, ensure you run `make build_dotnet install_dotnet_sdk` first: %v", err)
		}
		os.Setenv("PULUMI_LOCAL_NUGET", localNugetDir)
		sourceName := "pulumi-aws-native"
		output, err := exec.Command("dotnet", "nuget", "list", "source").CombinedOutput()
		require.NoError(t, err, "failed to list nuget sources")
		if !strings.Contains(string(output), sourceName) {
			err := exec.Command("dotnet", "nuget", "add", "source", localNugetDir, "-n", sourceName).Run()
			require.NoError(t, err, "failed to add nuget source")
		}
	}
	base := getBaseOptions(t)
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.AwsNative",
		},
	})

	return baseCsharp
}
