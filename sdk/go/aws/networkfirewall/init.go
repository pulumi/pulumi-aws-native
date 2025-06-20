// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

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
	case "aws-native:networkfirewall:Firewall":
		r = &Firewall{}
	case "aws-native:networkfirewall:FirewallPolicy":
		r = &FirewallPolicy{}
	case "aws-native:networkfirewall:LoggingConfiguration":
		r = &LoggingConfiguration{}
	case "aws-native:networkfirewall:RuleGroup":
		r = &RuleGroup{}
	case "aws-native:networkfirewall:TlsInspectionConfiguration":
		r = &TlsInspectionConfiguration{}
	case "aws-native:networkfirewall:VpcEndpointAssociation":
		r = &VpcEndpointAssociation{}
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
		"networkfirewall",
		&module{version},
	)
}
