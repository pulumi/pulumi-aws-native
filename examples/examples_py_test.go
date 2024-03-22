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

func TestUntypedPython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "untyped-py", "step1"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join(getCwd(t), "untyped-py", "step2"),
					Additive: true,
					ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
						assert.Equal(t, "updatedValue", stackInfo.Outputs["outputs"].(map[string]interface{})["Value"])
						assert.Equal(t,
							map[string]interface{}{"defaultTag": "defaultTagValue", "localTag": "localTagValue"},
							stackInfo.Outputs["outputs"].(map[string]interface{})["Tags"])
					},
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestDefaultTagsPython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "default-tags-py"),
			Config: map[string]string{
				"aws-native:defaultTags": `{
						"tags": {
							"defaultTag": "defaultTagValue"
						}
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

func TestVpcPython(t *testing.T) {
	// This test is not stable for several reasons so we can't run it in CI.
	// Deletion of the aws-native:ec2:IpamPoolCidr regularly fails with "GeneralServiceException: Error occurred during operation 'The CIDR has one or more allocations.'" even after the VPC is deleted.
	// Sometimes Update of the aws-native:ec2:IpamPool fails with "NotStabilized: IpamPoolCidrFailureReason(Message=The CIDR has one or more allocations."
	// The refresh of the aws-native:ec2:IpamPool fails due to a diff in the provisionedCidrs property.
	// It's also not ideal that we have to create the IPAM manually to avoid the test failing if run in parallel.
	t.SkipNow()
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "vpc-py", "step1"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join(getCwd(t), "vpc-py", "step2"),
					Additive: true,
				},
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
