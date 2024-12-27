// Copyright 2016-2021, Pulumi Corporation.

package main

func excludedTypes() []string {
	return []string{
		// TODO[pulumi/pulumi-aws-native#1946] AnycastIpList is not supported yet
		"AWS::CloudFront::AnycastIpList",
	}
}
