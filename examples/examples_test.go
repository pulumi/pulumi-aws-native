// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getEnvRegion(t *testing.T) string {
	envRegion := os.Getenv("AWS_REGION")
	if envRegion == "" {
		t.Skipf("Skipping test due to missing AWS_REGION environment variable")
	}

	return envRegion
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	envRegion := getEnvRegion(t)
	binPath, err := filepath.Abs("../bin")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Using binPath %s\n", binPath)
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"aws-native:region": envRegion,
		},
		ExpectRefreshChanges: true,
		SkipRefresh:          true,
		Quick:                true,
		LocalProviders: []integration.LocalDependency{
			{
				Package: "aws-native",
				Path:    binPath,
			},
		},
	}
}
