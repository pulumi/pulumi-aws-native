// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackagev2

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
	case "aws-native:mediapackagev2:Channel":
		r = &Channel{}
	case "aws-native:mediapackagev2:ChannelGroup":
		r = &ChannelGroup{}
	case "aws-native:mediapackagev2:ChannelPolicy":
		r = &ChannelPolicy{}
	case "aws-native:mediapackagev2:OriginEndpoint":
		r = &OriginEndpoint{}
	case "aws-native:mediapackagev2:OriginEndpointPolicy":
		r = &OriginEndpointPolicy{}
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
		"mediapackagev2",
		&module{version},
	)
}
