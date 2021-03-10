// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"path"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/pkg/errors"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

var client *cloudformation.CloudFormation

func init() {
	region := os.Getenv("AWS_REGION")
	if region != "" {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(region),
		})
		if err != nil {
			panic(errors.Errorf("could not create AWS session: %v", err))
		}
		client = cloudformation.New(sess)
	}
}

func TestVpcSingleInstanceInSubnet(t *testing.T) {
	stack := getStackName(t)
	defer deleteStack(t, stack)

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: path.Join(getCwd(t), "vpc-single-instance-in-subnet"),
		})

	integration.ProgramTest(t, &test)
}

func getStackName(t *testing.T) string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		t.Fatalf("failed to generate stack name: %v", err)
	}
	return "pulumi-cfn-test-" + hex.EncodeToString(b)
}

func deleteStack(t *testing.T, name string) {
	if t.Skipped() {
		return
	}

	_, err := client.DeleteStack(&cloudformation.DeleteStackInput{
		StackName: aws.String(name),
	})
	if err != nil {
		t.Fatalf("failed to delete stack %v: %v", name, err)
	}
}

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

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
		SkipRefresh:          true,
		Quick:                true,
	}
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	envRegion := getEnvRegion(t)
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"aws-native:region": envRegion,
		},
		Dependencies: []string{
			"@pulumi/aws-native",
		},
	})

	return baseJS
}
