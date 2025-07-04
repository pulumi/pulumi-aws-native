// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dsql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

// A map of key and value pairs to use to tag your cluster.
type ClusterTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.
	Value string `pulumi:"value"`
}

// The encryption configuration details for the cluster.
type EncryptionDetailsProperties struct {
	// The status of encryption for the cluster.
	EncryptionStatus *string `pulumi:"encryptionStatus"`
	// The type of encryption that protects data in the cluster.
	EncryptionType *string `pulumi:"encryptionType"`
	// The Amazon Resource Name (ARN) of the KMS key that encrypts data in the cluster.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
}

// The encryption configuration details for the cluster.
type EncryptionDetailsPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionDetailsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionDetailsProperties)(nil)).Elem()
}

func (o EncryptionDetailsPropertiesOutput) ToEncryptionDetailsPropertiesOutput() EncryptionDetailsPropertiesOutput {
	return o
}

func (o EncryptionDetailsPropertiesOutput) ToEncryptionDetailsPropertiesOutputWithContext(ctx context.Context) EncryptionDetailsPropertiesOutput {
	return o
}

// The status of encryption for the cluster.
func (o EncryptionDetailsPropertiesOutput) EncryptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionDetailsProperties) *string { return v.EncryptionStatus }).(pulumi.StringPtrOutput)
}

// The type of encryption that protects data in the cluster.
func (o EncryptionDetailsPropertiesOutput) EncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionDetailsProperties) *string { return v.EncryptionType }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the KMS key that encrypts data in the cluster.
func (o EncryptionDetailsPropertiesOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionDetailsProperties) *string { return v.KmsKeyArn }).(pulumi.StringPtrOutput)
}

type EncryptionDetailsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionDetailsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionDetailsProperties)(nil)).Elem()
}

func (o EncryptionDetailsPropertiesPtrOutput) ToEncryptionDetailsPropertiesPtrOutput() EncryptionDetailsPropertiesPtrOutput {
	return o
}

func (o EncryptionDetailsPropertiesPtrOutput) ToEncryptionDetailsPropertiesPtrOutputWithContext(ctx context.Context) EncryptionDetailsPropertiesPtrOutput {
	return o
}

func (o EncryptionDetailsPropertiesPtrOutput) Elem() EncryptionDetailsPropertiesOutput {
	return o.ApplyT(func(v *EncryptionDetailsProperties) EncryptionDetailsProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionDetailsProperties
		return ret
	}).(EncryptionDetailsPropertiesOutput)
}

// The status of encryption for the cluster.
func (o EncryptionDetailsPropertiesPtrOutput) EncryptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionDetailsProperties) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionStatus
	}).(pulumi.StringPtrOutput)
}

// The type of encryption that protects data in the cluster.
func (o EncryptionDetailsPropertiesPtrOutput) EncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionDetailsProperties) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionType
	}).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the KMS key that encrypts data in the cluster.
func (o EncryptionDetailsPropertiesPtrOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionDetailsProperties) *string {
		if v == nil {
			return nil
		}
		return v.KmsKeyArn
	}).(pulumi.StringPtrOutput)
}

// The Multi-region properties associated to this cluster.
type MultiRegionPropertiesProperties struct {
	// The set of peered clusters that form the multi-Region cluster configuration. Each peered cluster represents a database instance in a different Region.
	Clusters []string `pulumi:"clusters"`
	// The witness region in a multi-region cluster.
	WitnessRegion *string `pulumi:"witnessRegion"`
}

// MultiRegionPropertiesPropertiesInput is an input type that accepts MultiRegionPropertiesPropertiesArgs and MultiRegionPropertiesPropertiesOutput values.
// You can construct a concrete instance of `MultiRegionPropertiesPropertiesInput` via:
//
//	MultiRegionPropertiesPropertiesArgs{...}
type MultiRegionPropertiesPropertiesInput interface {
	pulumi.Input

	ToMultiRegionPropertiesPropertiesOutput() MultiRegionPropertiesPropertiesOutput
	ToMultiRegionPropertiesPropertiesOutputWithContext(context.Context) MultiRegionPropertiesPropertiesOutput
}

// The Multi-region properties associated to this cluster.
type MultiRegionPropertiesPropertiesArgs struct {
	// The set of peered clusters that form the multi-Region cluster configuration. Each peered cluster represents a database instance in a different Region.
	Clusters pulumi.StringArrayInput `pulumi:"clusters"`
	// The witness region in a multi-region cluster.
	WitnessRegion pulumi.StringPtrInput `pulumi:"witnessRegion"`
}

func (MultiRegionPropertiesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiRegionPropertiesProperties)(nil)).Elem()
}

func (i MultiRegionPropertiesPropertiesArgs) ToMultiRegionPropertiesPropertiesOutput() MultiRegionPropertiesPropertiesOutput {
	return i.ToMultiRegionPropertiesPropertiesOutputWithContext(context.Background())
}

func (i MultiRegionPropertiesPropertiesArgs) ToMultiRegionPropertiesPropertiesOutputWithContext(ctx context.Context) MultiRegionPropertiesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiRegionPropertiesPropertiesOutput)
}

func (i MultiRegionPropertiesPropertiesArgs) ToMultiRegionPropertiesPropertiesPtrOutput() MultiRegionPropertiesPropertiesPtrOutput {
	return i.ToMultiRegionPropertiesPropertiesPtrOutputWithContext(context.Background())
}

func (i MultiRegionPropertiesPropertiesArgs) ToMultiRegionPropertiesPropertiesPtrOutputWithContext(ctx context.Context) MultiRegionPropertiesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiRegionPropertiesPropertiesOutput).ToMultiRegionPropertiesPropertiesPtrOutputWithContext(ctx)
}

// MultiRegionPropertiesPropertiesPtrInput is an input type that accepts MultiRegionPropertiesPropertiesArgs, MultiRegionPropertiesPropertiesPtr and MultiRegionPropertiesPropertiesPtrOutput values.
// You can construct a concrete instance of `MultiRegionPropertiesPropertiesPtrInput` via:
//
//	        MultiRegionPropertiesPropertiesArgs{...}
//
//	or:
//
//	        nil
type MultiRegionPropertiesPropertiesPtrInput interface {
	pulumi.Input

	ToMultiRegionPropertiesPropertiesPtrOutput() MultiRegionPropertiesPropertiesPtrOutput
	ToMultiRegionPropertiesPropertiesPtrOutputWithContext(context.Context) MultiRegionPropertiesPropertiesPtrOutput
}

type multiRegionPropertiesPropertiesPtrType MultiRegionPropertiesPropertiesArgs

func MultiRegionPropertiesPropertiesPtr(v *MultiRegionPropertiesPropertiesArgs) MultiRegionPropertiesPropertiesPtrInput {
	return (*multiRegionPropertiesPropertiesPtrType)(v)
}

func (*multiRegionPropertiesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MultiRegionPropertiesProperties)(nil)).Elem()
}

func (i *multiRegionPropertiesPropertiesPtrType) ToMultiRegionPropertiesPropertiesPtrOutput() MultiRegionPropertiesPropertiesPtrOutput {
	return i.ToMultiRegionPropertiesPropertiesPtrOutputWithContext(context.Background())
}

func (i *multiRegionPropertiesPropertiesPtrType) ToMultiRegionPropertiesPropertiesPtrOutputWithContext(ctx context.Context) MultiRegionPropertiesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiRegionPropertiesPropertiesPtrOutput)
}

// The Multi-region properties associated to this cluster.
type MultiRegionPropertiesPropertiesOutput struct{ *pulumi.OutputState }

func (MultiRegionPropertiesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiRegionPropertiesProperties)(nil)).Elem()
}

func (o MultiRegionPropertiesPropertiesOutput) ToMultiRegionPropertiesPropertiesOutput() MultiRegionPropertiesPropertiesOutput {
	return o
}

func (o MultiRegionPropertiesPropertiesOutput) ToMultiRegionPropertiesPropertiesOutputWithContext(ctx context.Context) MultiRegionPropertiesPropertiesOutput {
	return o
}

func (o MultiRegionPropertiesPropertiesOutput) ToMultiRegionPropertiesPropertiesPtrOutput() MultiRegionPropertiesPropertiesPtrOutput {
	return o.ToMultiRegionPropertiesPropertiesPtrOutputWithContext(context.Background())
}

func (o MultiRegionPropertiesPropertiesOutput) ToMultiRegionPropertiesPropertiesPtrOutputWithContext(ctx context.Context) MultiRegionPropertiesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MultiRegionPropertiesProperties) *MultiRegionPropertiesProperties {
		return &v
	}).(MultiRegionPropertiesPropertiesPtrOutput)
}

// The set of peered clusters that form the multi-Region cluster configuration. Each peered cluster represents a database instance in a different Region.
func (o MultiRegionPropertiesPropertiesOutput) Clusters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MultiRegionPropertiesProperties) []string { return v.Clusters }).(pulumi.StringArrayOutput)
}

// The witness region in a multi-region cluster.
func (o MultiRegionPropertiesPropertiesOutput) WitnessRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MultiRegionPropertiesProperties) *string { return v.WitnessRegion }).(pulumi.StringPtrOutput)
}

type MultiRegionPropertiesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MultiRegionPropertiesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MultiRegionPropertiesProperties)(nil)).Elem()
}

func (o MultiRegionPropertiesPropertiesPtrOutput) ToMultiRegionPropertiesPropertiesPtrOutput() MultiRegionPropertiesPropertiesPtrOutput {
	return o
}

func (o MultiRegionPropertiesPropertiesPtrOutput) ToMultiRegionPropertiesPropertiesPtrOutputWithContext(ctx context.Context) MultiRegionPropertiesPropertiesPtrOutput {
	return o
}

func (o MultiRegionPropertiesPropertiesPtrOutput) Elem() MultiRegionPropertiesPropertiesOutput {
	return o.ApplyT(func(v *MultiRegionPropertiesProperties) MultiRegionPropertiesProperties {
		if v != nil {
			return *v
		}
		var ret MultiRegionPropertiesProperties
		return ret
	}).(MultiRegionPropertiesPropertiesOutput)
}

// The set of peered clusters that form the multi-Region cluster configuration. Each peered cluster represents a database instance in a different Region.
func (o MultiRegionPropertiesPropertiesPtrOutput) Clusters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MultiRegionPropertiesProperties) []string {
		if v == nil {
			return nil
		}
		return v.Clusters
	}).(pulumi.StringArrayOutput)
}

// The witness region in a multi-region cluster.
func (o MultiRegionPropertiesPropertiesPtrOutput) WitnessRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultiRegionPropertiesProperties) *string {
		if v == nil {
			return nil
		}
		return v.WitnessRegion
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MultiRegionPropertiesPropertiesInput)(nil)).Elem(), MultiRegionPropertiesPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MultiRegionPropertiesPropertiesPtrInput)(nil)).Elem(), MultiRegionPropertiesPropertiesArgs{})
	pulumi.RegisterOutputType(EncryptionDetailsPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MultiRegionPropertiesPropertiesOutput{})
	pulumi.RegisterOutputType(MultiRegionPropertiesPropertiesPtrOutput{})
}
