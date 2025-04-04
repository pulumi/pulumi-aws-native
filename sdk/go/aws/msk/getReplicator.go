// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::MSK::Replicator
func LookupReplicator(ctx *pulumi.Context, args *LookupReplicatorArgs, opts ...pulumi.InvokeOption) (*LookupReplicatorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReplicatorResult
	err := ctx.Invoke("aws-native:msk:getReplicator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReplicatorArgs struct {
	// Amazon Resource Name for the created replicator.
	ReplicatorArn string `pulumi:"replicatorArn"`
}

type LookupReplicatorResult struct {
	// The current version of the MSK replicator.
	CurrentVersion *string `pulumi:"currentVersion"`
	// A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
	ReplicationInfoList []ReplicatorReplicationInfo `pulumi:"replicationInfoList"`
	// Amazon Resource Name for the created replicator.
	ReplicatorArn *string `pulumi:"replicatorArn"`
	// A collection of tags associated with a resource
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupReplicatorOutput(ctx *pulumi.Context, args LookupReplicatorOutputArgs, opts ...pulumi.InvokeOption) LookupReplicatorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupReplicatorResultOutput, error) {
			args := v.(LookupReplicatorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:msk:getReplicator", args, LookupReplicatorResultOutput{}, options).(LookupReplicatorResultOutput), nil
		}).(LookupReplicatorResultOutput)
}

type LookupReplicatorOutputArgs struct {
	// Amazon Resource Name for the created replicator.
	ReplicatorArn pulumi.StringInput `pulumi:"replicatorArn"`
}

func (LookupReplicatorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicatorArgs)(nil)).Elem()
}

type LookupReplicatorResultOutput struct{ *pulumi.OutputState }

func (LookupReplicatorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReplicatorResult)(nil)).Elem()
}

func (o LookupReplicatorResultOutput) ToLookupReplicatorResultOutput() LookupReplicatorResultOutput {
	return o
}

func (o LookupReplicatorResultOutput) ToLookupReplicatorResultOutputWithContext(ctx context.Context) LookupReplicatorResultOutput {
	return o
}

// The current version of the MSK replicator.
func (o LookupReplicatorResultOutput) CurrentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicatorResult) *string { return v.CurrentVersion }).(pulumi.StringPtrOutput)
}

// A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
func (o LookupReplicatorResultOutput) ReplicationInfoList() ReplicatorReplicationInfoArrayOutput {
	return o.ApplyT(func(v LookupReplicatorResult) []ReplicatorReplicationInfo { return v.ReplicationInfoList }).(ReplicatorReplicationInfoArrayOutput)
}

// Amazon Resource Name for the created replicator.
func (o LookupReplicatorResultOutput) ReplicatorArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReplicatorResult) *string { return v.ReplicatorArn }).(pulumi.StringPtrOutput)
}

// A collection of tags associated with a resource
func (o LookupReplicatorResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupReplicatorResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReplicatorResultOutput{})
}
