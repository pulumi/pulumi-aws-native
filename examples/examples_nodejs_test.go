// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestSimpleTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple-ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestGetTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "get-ts"),
		})

	integration.ProgramTest(t, &test)
}

func TestUpdate(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "update", "step1"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join(getCwd(t), "update", "step2"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestNamingConventions(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "aws-native-naming-conventions"),
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				// Check that the name of the created queue is correct.
				fifoName := stackInfo.Outputs["fifoName"].(string)
				if !(fifoName[len(fifoName)-5:] == ".fifo") {
					t.Errorf("Expected fifo queue name to end with '.fifo', got '%s'", fifoName)
				}
				standardName := stackInfo.Outputs["standardName"].(string)
				if standardName[len(standardName)-5:] == ".fifo" {
					t.Errorf("Expected standard queue name to not end with '.fifo', got '%s'", standardName)
				}
			},
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/aws-native",
		},
	})

	return baseJS
}
