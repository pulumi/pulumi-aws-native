// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws-native:rds:DBCluster":
		r = &DBCluster{}
	case "aws-native:rds:DBClusterParameterGroup":
		r = &DBClusterParameterGroup{}
	case "aws-native:rds:DBInstance":
		r = &DBInstance{}
	case "aws-native:rds:DBParameterGroup":
		r = &DBParameterGroup{}
	case "aws-native:rds:DBProxy":
		r = &DBProxy{}
	case "aws-native:rds:DBProxyEndpoint":
		r = &DBProxyEndpoint{}
	case "aws-native:rds:DBProxyTargetGroup":
		r = &DBProxyTargetGroup{}
	case "aws-native:rds:DBSecurityGroup":
		r = &DBSecurityGroup{}
	case "aws-native:rds:DBSecurityGroupIngress":
		r = &DBSecurityGroupIngress{}
	case "aws-native:rds:DBSubnetGroup":
		r = &DBSubnetGroup{}
	case "aws-native:rds:EventSubscription":
		r = &EventSubscription{}
	case "aws-native:rds:GlobalCluster":
		r = &GlobalCluster{}
	case "aws-native:rds:OptionGroup":
		r = &OptionGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws-native",
		"rds",
		&module{version},
	)
}