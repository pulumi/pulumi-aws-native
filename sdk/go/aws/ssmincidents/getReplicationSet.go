// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssmincidents

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::SSMIncidents::ReplicationSet
func LookupReplicationSet(ctx *pulumi.Context, args *LookupReplicationSetArgs, opts ...pulumi.InvokeOption) (*LookupReplicationSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicationSetResult
	err := ctx.Invoke("aws-native:ssmincidents:getReplicationSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicationSetArgs struct {
	// The ARN of the ReplicationSet.
	Arn string `pulumi:"arn"`
}

type LookupReplicationSetResult struct {
	// The ARN of the ReplicationSet.
	Arn *string `pulumi:"arn"`
	// Determines if the replication set deletion protection is enabled or not. If deletion protection is enabled, you can't delete the last Region in the replication set.
	DeletionProtected *bool `pulumi:"deletionProtected"`
	// The ReplicationSet configuration.
	Regions []ReplicationSetReplicationRegion `pulumi:"regions"`
	// The tags to apply to the replication set.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupReplicationSetOutput(ctx *pulumi.Context, args LookupReplicationSetOutputArgs, opts ...pulumi.InvokeOption) LookupReplicationSetResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupReplicationSetResultOutput, error) {
			args := v.(LookupReplicationSetArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ssmincidents:getReplicationSet", args, LookupReplicationSetResultOutput{}, options).(LookupReplicationSetResultOutput), nil
		}).(LookupReplicationSetResultOutput)
}

type LookupReplicationSetOutputArgs struct {
	// The ARN of the ReplicationSet.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupReplicationSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationSetArgs)(nil)).Elem()
}

type LookupReplicationSetResultOutput struct{ *pulumi.OutputState }

func (LookupReplicationSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicationSetResult)(nil)).Elem()
}

func (o LookupReplicationSetResultOutput) ToLookupReplicationSetResultOutput() LookupReplicationSetResultOutput {
	return o
}

func (o LookupReplicationSetResultOutput) ToLookupReplicationSetResultOutputWithContext(ctx context.Context) LookupReplicationSetResultOutput {
	return o
}

// The ARN of the ReplicationSet.
func (o LookupReplicationSetResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicationSetResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Determines if the replication set deletion protection is enabled or not. If deletion protection is enabled, you can't delete the last Region in the replication set.
func (o LookupReplicationSetResultOutput) DeletionProtected() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupReplicationSetResult) *bool { return v.DeletionProtected }).(pulumi.BoolPtrOutput)
}

// The ReplicationSet configuration.
func (o LookupReplicationSetResultOutput) Regions() ReplicationSetReplicationRegionArrayOutput {
	return o.ApplyT(func(v LookupReplicationSetResult) []ReplicationSetReplicationRegion { return v.Regions }).(ReplicationSetReplicationRegionArrayOutput)
}

// The tags to apply to the replication set.
func (o LookupReplicationSetResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupReplicationSetResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicationSetResultOutput{})
}
