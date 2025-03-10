// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::DMS::InstanceProfile.
func LookupInstanceProfile(ctx *pulumi.Context, args *LookupInstanceProfileArgs, opts ...pulumi.InvokeOption) (*LookupInstanceProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInstanceProfileResult
	err := ctx.Invoke("aws-native:dms:getInstanceProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceProfileArgs struct {
	// The property describes an ARN of the instance profile.
	InstanceProfileArn string `pulumi:"instanceProfileArn"`
}

type LookupInstanceProfileResult struct {
	// The property describes an availability zone of the instance profile.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The optional description of the instance profile.
	Description *string `pulumi:"description"`
	// The property describes an ARN of the instance profile.
	InstanceProfileArn *string `pulumi:"instanceProfileArn"`
	// The property describes a creating time of the instance profile.
	InstanceProfileCreationTime *string `pulumi:"instanceProfileCreationTime"`
	// The property describes a name for the instance profile.
	InstanceProfileName *string `pulumi:"instanceProfileName"`
	// The property describes kms key arn for the instance profile.
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The property describes a network type for the instance profile.
	NetworkType *InstanceProfileNetworkType `pulumi:"networkType"`
	// The property describes the publicly accessible of the instance profile
	PubliclyAccessible *bool `pulumi:"publiclyAccessible"`
	// The property describes a subnet group identifier for the instance profile.
	SubnetGroupIdentifier *string `pulumi:"subnetGroupIdentifier"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
	// The property describes vps security groups for the instance profile.
	VpcSecurityGroups []string `pulumi:"vpcSecurityGroups"`
}

func LookupInstanceProfileOutput(ctx *pulumi.Context, args LookupInstanceProfileOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupInstanceProfileResultOutput, error) {
			args := v.(LookupInstanceProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:dms:getInstanceProfile", args, LookupInstanceProfileResultOutput{}, options).(LookupInstanceProfileResultOutput), nil
		}).(LookupInstanceProfileResultOutput)
}

type LookupInstanceProfileOutputArgs struct {
	// The property describes an ARN of the instance profile.
	InstanceProfileArn pulumi.StringInput `pulumi:"instanceProfileArn"`
}

func (LookupInstanceProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceProfileArgs)(nil)).Elem()
}

type LookupInstanceProfileResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceProfileResult)(nil)).Elem()
}

func (o LookupInstanceProfileResultOutput) ToLookupInstanceProfileResultOutput() LookupInstanceProfileResultOutput {
	return o
}

func (o LookupInstanceProfileResultOutput) ToLookupInstanceProfileResultOutputWithContext(ctx context.Context) LookupInstanceProfileResultOutput {
	return o
}

// The property describes an availability zone of the instance profile.
func (o LookupInstanceProfileResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The optional description of the instance profile.
func (o LookupInstanceProfileResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The property describes an ARN of the instance profile.
func (o LookupInstanceProfileResultOutput) InstanceProfileArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.InstanceProfileArn }).(pulumi.StringPtrOutput)
}

// The property describes a creating time of the instance profile.
func (o LookupInstanceProfileResultOutput) InstanceProfileCreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.InstanceProfileCreationTime }).(pulumi.StringPtrOutput)
}

// The property describes a name for the instance profile.
func (o LookupInstanceProfileResultOutput) InstanceProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.InstanceProfileName }).(pulumi.StringPtrOutput)
}

// The property describes kms key arn for the instance profile.
func (o LookupInstanceProfileResultOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.KmsKeyArn }).(pulumi.StringPtrOutput)
}

// The property describes a network type for the instance profile.
func (o LookupInstanceProfileResultOutput) NetworkType() InstanceProfileNetworkTypePtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *InstanceProfileNetworkType { return v.NetworkType }).(InstanceProfileNetworkTypePtrOutput)
}

// The property describes the publicly accessible of the instance profile
func (o LookupInstanceProfileResultOutput) PubliclyAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *bool { return v.PubliclyAccessible }).(pulumi.BoolPtrOutput)
}

// The property describes a subnet group identifier for the instance profile.
func (o LookupInstanceProfileResultOutput) SubnetGroupIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.SubnetGroupIdentifier }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupInstanceProfileResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The property describes vps security groups for the instance profile.
func (o LookupInstanceProfileResultOutput) VpcSecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) []string { return v.VpcSecurityGroups }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceProfileResultOutput{})
}
