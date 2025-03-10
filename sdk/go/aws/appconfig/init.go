// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
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
	case "aws-native:appconfig:Application":
		r = &Application{}
	case "aws-native:appconfig:ConfigurationProfile":
		r = &ConfigurationProfile{}
	case "aws-native:appconfig:Deployment":
		r = &Deployment{}
	case "aws-native:appconfig:DeploymentStrategy":
		r = &DeploymentStrategy{}
	case "aws-native:appconfig:Environment":
		r = &Environment{}
	case "aws-native:appconfig:Extension":
		r = &Extension{}
	case "aws-native:appconfig:ExtensionAssociation":
		r = &ExtensionAssociation{}
	case "aws-native:appconfig:HostedConfigurationVersion":
		r = &HostedConfigurationVersion{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"aws-native",
		"appconfig",
		&module{version},
	)
}
