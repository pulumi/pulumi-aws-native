// Copyright 2016-2021, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
)

func TestSimpleGo(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple-go"),
		})

	integration.ProgramTest(t, &test)
}

func TestWriteOnlyGo(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			ExpectRefreshChanges: false,
			Dir:                  filepath.Join(getCwd(t), "write-only-go"),
			Config: map[string]string{
				"aws-native:defaultTags": `{
					"defaultTag": "defaultTagValue"
					}`,
			},
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				// We should see both the default tag and the local tag in the log group tags.
				// If the default tag is missing, it's because there was a missing output from the provider and the SDK filled in the value with the original input.
				tags := stack.Outputs["tags"].(map[string]interface{})
				assert.Equal(t, map[string]interface{}{
					"defaultTag": "defaultTagValue",
					"localTag":   "localTagValue",
				}, tags)
			},
		})

	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)

	sdkPath, err := filepath.Abs("../sdk")
	if err != nil {
		t.Fatal(err)
	}

	baseGo := base.With(integration.ProgramTestOptions{
		Verbose: true,
		Dependencies: []string{
			fmt.Sprintf("github.com/pulumi/pulumi-aws-native/sdk=%s", sdkPath),
		},
	})

	return baseGo
}
