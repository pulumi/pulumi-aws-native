// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build python || all
// +build python all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestSimplePython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple-py"),
		})

	integration.ProgramTest(t, &test)
}

func TestDefaultTagsPython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "default-tags-py"),
			Config: map[string]string{
				"aws-native:defaultTags": `{
					"defaultTag": "defaultTagValue"
					}`,
			},
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				assert.Equal(t, []interface{}{
					map[string]interface{}{
						"key":   "localTag",
						"value": "localTagValue",
					},
					map[string]interface{}{
						"key":   "defaultTag",
						"value": "defaultTagValue",
					},
				}, stackInfo.Outputs["logGroupTags"])

				assert.Equal(t, map[string]interface{}{
					"defaultTag": "defaultTagValue",
					"localTag":   "localTagValue",
				}, stackInfo.Outputs["policyTags"])
			},
		})

	integration.ProgramTest(t, &test)
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	basePy := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePy
}
