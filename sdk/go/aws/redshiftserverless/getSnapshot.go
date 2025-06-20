// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redshiftserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RedshiftServerless::Snapshot Resource Type.
func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSnapshotResult
	err := ctx.Invoke("aws-native:redshiftserverless:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotArgs struct {
	// The name of the snapshot.
	SnapshotName string `pulumi:"snapshotName"`
}

type LookupSnapshotResult struct {
	// The owner account of the snapshot.
	OwnerAccount *string `pulumi:"ownerAccount"`
	// The retention period of the snapshot.
	RetentionPeriod *int `pulumi:"retentionPeriod"`
	// Definition for snapshot resource
	Snapshot *SnapshotType `pulumi:"snapshot"`
}

func LookupSnapshotOutput(ctx *pulumi.Context, args LookupSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSnapshotResultOutput, error) {
			args := v.(LookupSnapshotArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:redshiftserverless:getSnapshot", args, LookupSnapshotResultOutput{}, options).(LookupSnapshotResultOutput), nil
		}).(LookupSnapshotResultOutput)
}

type LookupSnapshotOutputArgs struct {
	// The name of the snapshot.
	SnapshotName pulumi.StringInput `pulumi:"snapshotName"`
}

func (LookupSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotArgs)(nil)).Elem()
}

type LookupSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotResult)(nil)).Elem()
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutput() LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutputWithContext(ctx context.Context) LookupSnapshotResultOutput {
	return o
}

// The owner account of the snapshot.
func (o LookupSnapshotResultOutput) OwnerAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.OwnerAccount }).(pulumi.StringPtrOutput)
}

// The retention period of the snapshot.
func (o LookupSnapshotResultOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *int { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

// Definition for snapshot resource
func (o LookupSnapshotResultOutput) Snapshot() SnapshotTypePtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *SnapshotType { return v.Snapshot }).(SnapshotTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}
