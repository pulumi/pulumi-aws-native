// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resolve parameter value conflicts
type AddonResolveConflicts string

const (
	AddonResolveConflictsNone      = AddonResolveConflicts("NONE")
	AddonResolveConflictsOverwrite = AddonResolveConflicts("OVERWRITE")
	AddonResolveConflictsPreserve  = AddonResolveConflicts("PRESERVE")
)

func (AddonResolveConflicts) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonResolveConflicts)(nil)).Elem()
}

func (e AddonResolveConflicts) ToAddonResolveConflictsOutput() AddonResolveConflictsOutput {
	return pulumi.ToOutput(e).(AddonResolveConflictsOutput)
}

func (e AddonResolveConflicts) ToAddonResolveConflictsOutputWithContext(ctx context.Context) AddonResolveConflictsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AddonResolveConflictsOutput)
}

func (e AddonResolveConflicts) ToAddonResolveConflictsPtrOutput() AddonResolveConflictsPtrOutput {
	return e.ToAddonResolveConflictsPtrOutputWithContext(context.Background())
}

func (e AddonResolveConflicts) ToAddonResolveConflictsPtrOutputWithContext(ctx context.Context) AddonResolveConflictsPtrOutput {
	return AddonResolveConflicts(e).ToAddonResolveConflictsOutputWithContext(ctx).ToAddonResolveConflictsPtrOutputWithContext(ctx)
}

func (e AddonResolveConflicts) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddonResolveConflicts) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AddonResolveConflicts) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AddonResolveConflicts) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AddonResolveConflictsOutput struct{ *pulumi.OutputState }

func (AddonResolveConflictsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonResolveConflicts)(nil)).Elem()
}

func (o AddonResolveConflictsOutput) ToAddonResolveConflictsOutput() AddonResolveConflictsOutput {
	return o
}

func (o AddonResolveConflictsOutput) ToAddonResolveConflictsOutputWithContext(ctx context.Context) AddonResolveConflictsOutput {
	return o
}

func (o AddonResolveConflictsOutput) ToAddonResolveConflictsPtrOutput() AddonResolveConflictsPtrOutput {
	return o.ToAddonResolveConflictsPtrOutputWithContext(context.Background())
}

func (o AddonResolveConflictsOutput) ToAddonResolveConflictsPtrOutputWithContext(ctx context.Context) AddonResolveConflictsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddonResolveConflicts) *AddonResolveConflicts {
		return &v
	}).(AddonResolveConflictsPtrOutput)
}

func (o AddonResolveConflictsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AddonResolveConflictsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddonResolveConflicts) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AddonResolveConflictsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddonResolveConflictsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AddonResolveConflicts) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AddonResolveConflictsPtrOutput struct{ *pulumi.OutputState }

func (AddonResolveConflictsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddonResolveConflicts)(nil)).Elem()
}

func (o AddonResolveConflictsPtrOutput) ToAddonResolveConflictsPtrOutput() AddonResolveConflictsPtrOutput {
	return o
}

func (o AddonResolveConflictsPtrOutput) ToAddonResolveConflictsPtrOutputWithContext(ctx context.Context) AddonResolveConflictsPtrOutput {
	return o
}

func (o AddonResolveConflictsPtrOutput) Elem() AddonResolveConflictsOutput {
	return o.ApplyT(func(v *AddonResolveConflicts) AddonResolveConflicts {
		if v != nil {
			return *v
		}
		var ret AddonResolveConflicts
		return ret
	}).(AddonResolveConflictsOutput)
}

func (o AddonResolveConflictsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AddonResolveConflictsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AddonResolveConflicts) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AddonResolveConflictsInput is an input type that accepts AddonResolveConflictsArgs and AddonResolveConflictsOutput values.
// You can construct a concrete instance of `AddonResolveConflictsInput` via:
//
//	AddonResolveConflictsArgs{...}
type AddonResolveConflictsInput interface {
	pulumi.Input

	ToAddonResolveConflictsOutput() AddonResolveConflictsOutput
	ToAddonResolveConflictsOutputWithContext(context.Context) AddonResolveConflictsOutput
}

var addonResolveConflictsPtrType = reflect.TypeOf((**AddonResolveConflicts)(nil)).Elem()

type AddonResolveConflictsPtrInput interface {
	pulumi.Input

	ToAddonResolveConflictsPtrOutput() AddonResolveConflictsPtrOutput
	ToAddonResolveConflictsPtrOutputWithContext(context.Context) AddonResolveConflictsPtrOutput
}

type addonResolveConflictsPtr string

func AddonResolveConflictsPtr(v string) AddonResolveConflictsPtrInput {
	return (*addonResolveConflictsPtr)(&v)
}

func (*addonResolveConflictsPtr) ElementType() reflect.Type {
	return addonResolveConflictsPtrType
}

func (in *addonResolveConflictsPtr) ToAddonResolveConflictsPtrOutput() AddonResolveConflictsPtrOutput {
	return pulumi.ToOutput(in).(AddonResolveConflictsPtrOutput)
}

func (in *addonResolveConflictsPtr) ToAddonResolveConflictsPtrOutputWithContext(ctx context.Context) AddonResolveConflictsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AddonResolveConflictsPtrOutput)
}

// Ipv4 or Ipv6. You can only specify ipv6 for 1.21 and later clusters that use version 1.10.1 or later of the Amazon VPC CNI add-on
type ClusterKubernetesNetworkConfigIpFamily string

const (
	ClusterKubernetesNetworkConfigIpFamilyIpv4 = ClusterKubernetesNetworkConfigIpFamily("ipv4")
	ClusterKubernetesNetworkConfigIpFamilyIpv6 = ClusterKubernetesNetworkConfigIpFamily("ipv6")
)

func (ClusterKubernetesNetworkConfigIpFamily) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKubernetesNetworkConfigIpFamily)(nil)).Elem()
}

func (e ClusterKubernetesNetworkConfigIpFamily) ToClusterKubernetesNetworkConfigIpFamilyOutput() ClusterKubernetesNetworkConfigIpFamilyOutput {
	return pulumi.ToOutput(e).(ClusterKubernetesNetworkConfigIpFamilyOutput)
}

func (e ClusterKubernetesNetworkConfigIpFamily) ToClusterKubernetesNetworkConfigIpFamilyOutputWithContext(ctx context.Context) ClusterKubernetesNetworkConfigIpFamilyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterKubernetesNetworkConfigIpFamilyOutput)
}

func (e ClusterKubernetesNetworkConfigIpFamily) ToClusterKubernetesNetworkConfigIpFamilyPtrOutput() ClusterKubernetesNetworkConfigIpFamilyPtrOutput {
	return e.ToClusterKubernetesNetworkConfigIpFamilyPtrOutputWithContext(context.Background())
}

func (e ClusterKubernetesNetworkConfigIpFamily) ToClusterKubernetesNetworkConfigIpFamilyPtrOutputWithContext(ctx context.Context) ClusterKubernetesNetworkConfigIpFamilyPtrOutput {
	return ClusterKubernetesNetworkConfigIpFamily(e).ToClusterKubernetesNetworkConfigIpFamilyOutputWithContext(ctx).ToClusterKubernetesNetworkConfigIpFamilyPtrOutputWithContext(ctx)
}

func (e ClusterKubernetesNetworkConfigIpFamily) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterKubernetesNetworkConfigIpFamily) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterKubernetesNetworkConfigIpFamily) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterKubernetesNetworkConfigIpFamily) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterKubernetesNetworkConfigIpFamilyOutput struct{ *pulumi.OutputState }

func (ClusterKubernetesNetworkConfigIpFamilyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKubernetesNetworkConfigIpFamily)(nil)).Elem()
}

func (o ClusterKubernetesNetworkConfigIpFamilyOutput) ToClusterKubernetesNetworkConfigIpFamilyOutput() ClusterKubernetesNetworkConfigIpFamilyOutput {
	return o
}

func (o ClusterKubernetesNetworkConfigIpFamilyOutput) ToClusterKubernetesNetworkConfigIpFamilyOutputWithContext(ctx context.Context) ClusterKubernetesNetworkConfigIpFamilyOutput {
	return o
}

func (o ClusterKubernetesNetworkConfigIpFamilyOutput) ToClusterKubernetesNetworkConfigIpFamilyPtrOutput() ClusterKubernetesNetworkConfigIpFamilyPtrOutput {
	return o.ToClusterKubernetesNetworkConfigIpFamilyPtrOutputWithContext(context.Background())
}

func (o ClusterKubernetesNetworkConfigIpFamilyOutput) ToClusterKubernetesNetworkConfigIpFamilyPtrOutputWithContext(ctx context.Context) ClusterKubernetesNetworkConfigIpFamilyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterKubernetesNetworkConfigIpFamily) *ClusterKubernetesNetworkConfigIpFamily {
		return &v
	}).(ClusterKubernetesNetworkConfigIpFamilyPtrOutput)
}

func (o ClusterKubernetesNetworkConfigIpFamilyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterKubernetesNetworkConfigIpFamilyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterKubernetesNetworkConfigIpFamily) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterKubernetesNetworkConfigIpFamilyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterKubernetesNetworkConfigIpFamilyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterKubernetesNetworkConfigIpFamily) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterKubernetesNetworkConfigIpFamilyPtrOutput struct{ *pulumi.OutputState }

func (ClusterKubernetesNetworkConfigIpFamilyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterKubernetesNetworkConfigIpFamily)(nil)).Elem()
}

func (o ClusterKubernetesNetworkConfigIpFamilyPtrOutput) ToClusterKubernetesNetworkConfigIpFamilyPtrOutput() ClusterKubernetesNetworkConfigIpFamilyPtrOutput {
	return o
}

func (o ClusterKubernetesNetworkConfigIpFamilyPtrOutput) ToClusterKubernetesNetworkConfigIpFamilyPtrOutputWithContext(ctx context.Context) ClusterKubernetesNetworkConfigIpFamilyPtrOutput {
	return o
}

func (o ClusterKubernetesNetworkConfigIpFamilyPtrOutput) Elem() ClusterKubernetesNetworkConfigIpFamilyOutput {
	return o.ApplyT(func(v *ClusterKubernetesNetworkConfigIpFamily) ClusterKubernetesNetworkConfigIpFamily {
		if v != nil {
			return *v
		}
		var ret ClusterKubernetesNetworkConfigIpFamily
		return ret
	}).(ClusterKubernetesNetworkConfigIpFamilyOutput)
}

func (o ClusterKubernetesNetworkConfigIpFamilyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterKubernetesNetworkConfigIpFamilyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterKubernetesNetworkConfigIpFamily) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ClusterKubernetesNetworkConfigIpFamilyInput is an input type that accepts ClusterKubernetesNetworkConfigIpFamilyArgs and ClusterKubernetesNetworkConfigIpFamilyOutput values.
// You can construct a concrete instance of `ClusterKubernetesNetworkConfigIpFamilyInput` via:
//
//	ClusterKubernetesNetworkConfigIpFamilyArgs{...}
type ClusterKubernetesNetworkConfigIpFamilyInput interface {
	pulumi.Input

	ToClusterKubernetesNetworkConfigIpFamilyOutput() ClusterKubernetesNetworkConfigIpFamilyOutput
	ToClusterKubernetesNetworkConfigIpFamilyOutputWithContext(context.Context) ClusterKubernetesNetworkConfigIpFamilyOutput
}

var clusterKubernetesNetworkConfigIpFamilyPtrType = reflect.TypeOf((**ClusterKubernetesNetworkConfigIpFamily)(nil)).Elem()

type ClusterKubernetesNetworkConfigIpFamilyPtrInput interface {
	pulumi.Input

	ToClusterKubernetesNetworkConfigIpFamilyPtrOutput() ClusterKubernetesNetworkConfigIpFamilyPtrOutput
	ToClusterKubernetesNetworkConfigIpFamilyPtrOutputWithContext(context.Context) ClusterKubernetesNetworkConfigIpFamilyPtrOutput
}

type clusterKubernetesNetworkConfigIpFamilyPtr string

func ClusterKubernetesNetworkConfigIpFamilyPtr(v string) ClusterKubernetesNetworkConfigIpFamilyPtrInput {
	return (*clusterKubernetesNetworkConfigIpFamilyPtr)(&v)
}

func (*clusterKubernetesNetworkConfigIpFamilyPtr) ElementType() reflect.Type {
	return clusterKubernetesNetworkConfigIpFamilyPtrType
}

func (in *clusterKubernetesNetworkConfigIpFamilyPtr) ToClusterKubernetesNetworkConfigIpFamilyPtrOutput() ClusterKubernetesNetworkConfigIpFamilyPtrOutput {
	return pulumi.ToOutput(in).(ClusterKubernetesNetworkConfigIpFamilyPtrOutput)
}

func (in *clusterKubernetesNetworkConfigIpFamilyPtr) ToClusterKubernetesNetworkConfigIpFamilyPtrOutputWithContext(ctx context.Context) ClusterKubernetesNetworkConfigIpFamilyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterKubernetesNetworkConfigIpFamilyPtrOutput)
}

// The type of the identity provider configuration.
type IdentityProviderConfigType string

const (
	IdentityProviderConfigTypeOidc = IdentityProviderConfigType("oidc")
)

func (IdentityProviderConfigType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderConfigType)(nil)).Elem()
}

func (e IdentityProviderConfigType) ToIdentityProviderConfigTypeOutput() IdentityProviderConfigTypeOutput {
	return pulumi.ToOutput(e).(IdentityProviderConfigTypeOutput)
}

func (e IdentityProviderConfigType) ToIdentityProviderConfigTypeOutputWithContext(ctx context.Context) IdentityProviderConfigTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityProviderConfigTypeOutput)
}

func (e IdentityProviderConfigType) ToIdentityProviderConfigTypePtrOutput() IdentityProviderConfigTypePtrOutput {
	return e.ToIdentityProviderConfigTypePtrOutputWithContext(context.Background())
}

func (e IdentityProviderConfigType) ToIdentityProviderConfigTypePtrOutputWithContext(ctx context.Context) IdentityProviderConfigTypePtrOutput {
	return IdentityProviderConfigType(e).ToIdentityProviderConfigTypeOutputWithContext(ctx).ToIdentityProviderConfigTypePtrOutputWithContext(ctx)
}

func (e IdentityProviderConfigType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityProviderConfigType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityProviderConfigType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityProviderConfigType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityProviderConfigTypeOutput struct{ *pulumi.OutputState }

func (IdentityProviderConfigTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderConfigType)(nil)).Elem()
}

func (o IdentityProviderConfigTypeOutput) ToIdentityProviderConfigTypeOutput() IdentityProviderConfigTypeOutput {
	return o
}

func (o IdentityProviderConfigTypeOutput) ToIdentityProviderConfigTypeOutputWithContext(ctx context.Context) IdentityProviderConfigTypeOutput {
	return o
}

func (o IdentityProviderConfigTypeOutput) ToIdentityProviderConfigTypePtrOutput() IdentityProviderConfigTypePtrOutput {
	return o.ToIdentityProviderConfigTypePtrOutputWithContext(context.Background())
}

func (o IdentityProviderConfigTypeOutput) ToIdentityProviderConfigTypePtrOutputWithContext(ctx context.Context) IdentityProviderConfigTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProviderConfigType) *IdentityProviderConfigType {
		return &v
	}).(IdentityProviderConfigTypePtrOutput)
}

func (o IdentityProviderConfigTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityProviderConfigTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityProviderConfigType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityProviderConfigTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityProviderConfigTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityProviderConfigType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityProviderConfigTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityProviderConfigTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderConfigType)(nil)).Elem()
}

func (o IdentityProviderConfigTypePtrOutput) ToIdentityProviderConfigTypePtrOutput() IdentityProviderConfigTypePtrOutput {
	return o
}

func (o IdentityProviderConfigTypePtrOutput) ToIdentityProviderConfigTypePtrOutputWithContext(ctx context.Context) IdentityProviderConfigTypePtrOutput {
	return o
}

func (o IdentityProviderConfigTypePtrOutput) Elem() IdentityProviderConfigTypeOutput {
	return o.ApplyT(func(v *IdentityProviderConfigType) IdentityProviderConfigType {
		if v != nil {
			return *v
		}
		var ret IdentityProviderConfigType
		return ret
	}).(IdentityProviderConfigTypeOutput)
}

func (o IdentityProviderConfigTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityProviderConfigTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityProviderConfigType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// IdentityProviderConfigTypeInput is an input type that accepts IdentityProviderConfigTypeArgs and IdentityProviderConfigTypeOutput values.
// You can construct a concrete instance of `IdentityProviderConfigTypeInput` via:
//
//	IdentityProviderConfigTypeArgs{...}
type IdentityProviderConfigTypeInput interface {
	pulumi.Input

	ToIdentityProviderConfigTypeOutput() IdentityProviderConfigTypeOutput
	ToIdentityProviderConfigTypeOutputWithContext(context.Context) IdentityProviderConfigTypeOutput
}

var identityProviderConfigTypePtrType = reflect.TypeOf((**IdentityProviderConfigType)(nil)).Elem()

type IdentityProviderConfigTypePtrInput interface {
	pulumi.Input

	ToIdentityProviderConfigTypePtrOutput() IdentityProviderConfigTypePtrOutput
	ToIdentityProviderConfigTypePtrOutputWithContext(context.Context) IdentityProviderConfigTypePtrOutput
}

type identityProviderConfigTypePtr string

func IdentityProviderConfigTypePtr(v string) IdentityProviderConfigTypePtrInput {
	return (*identityProviderConfigTypePtr)(&v)
}

func (*identityProviderConfigTypePtr) ElementType() reflect.Type {
	return identityProviderConfigTypePtrType
}

func (in *identityProviderConfigTypePtr) ToIdentityProviderConfigTypePtrOutput() IdentityProviderConfigTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityProviderConfigTypePtrOutput)
}

func (in *identityProviderConfigTypePtr) ToIdentityProviderConfigTypePtrOutputWithContext(ctx context.Context) IdentityProviderConfigTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityProviderConfigTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddonResolveConflictsInput)(nil)).Elem(), AddonResolveConflicts("NONE"))
	pulumi.RegisterInputType(reflect.TypeOf((*AddonResolveConflictsPtrInput)(nil)).Elem(), AddonResolveConflicts("NONE"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterKubernetesNetworkConfigIpFamilyInput)(nil)).Elem(), ClusterKubernetesNetworkConfigIpFamily("ipv4"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterKubernetesNetworkConfigIpFamilyPtrInput)(nil)).Elem(), ClusterKubernetesNetworkConfigIpFamily("ipv4"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderConfigTypeInput)(nil)).Elem(), IdentityProviderConfigType("oidc"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderConfigTypePtrInput)(nil)).Elem(), IdentityProviderConfigType("oidc"))
	pulumi.RegisterOutputType(AddonResolveConflictsOutput{})
	pulumi.RegisterOutputType(AddonResolveConflictsPtrOutput{})
	pulumi.RegisterOutputType(ClusterKubernetesNetworkConfigIpFamilyOutput{})
	pulumi.RegisterOutputType(ClusterKubernetesNetworkConfigIpFamilyPtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderConfigTypeOutput{})
	pulumi.RegisterOutputType(IdentityProviderConfigTypePtrOutput{})
}