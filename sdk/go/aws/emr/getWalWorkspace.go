// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::EMR::WALWorkspace Type
func LookupWalWorkspace(ctx *pulumi.Context, args *LookupWalWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWalWorkspaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWalWorkspaceResult
	err := ctx.Invoke("aws-native:emr:getWalWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWalWorkspaceArgs struct {
	// The name of the emrwal container
	WalWorkspaceName string `pulumi:"walWorkspaceName"`
}

type LookupWalWorkspaceResult struct {
	// An array of key-value pairs to apply to this resource.
	Tags []WalWorkspaceTag `pulumi:"tags"`
}

func LookupWalWorkspaceOutput(ctx *pulumi.Context, args LookupWalWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupWalWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWalWorkspaceResult, error) {
			args := v.(LookupWalWorkspaceArgs)
			r, err := LookupWalWorkspace(ctx, &args, opts...)
			var s LookupWalWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWalWorkspaceResultOutput)
}

type LookupWalWorkspaceOutputArgs struct {
	// The name of the emrwal container
	WalWorkspaceName pulumi.StringInput `pulumi:"walWorkspaceName"`
}

func (LookupWalWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWalWorkspaceArgs)(nil)).Elem()
}

type LookupWalWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWalWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWalWorkspaceResult)(nil)).Elem()
}

func (o LookupWalWorkspaceResultOutput) ToLookupWalWorkspaceResultOutput() LookupWalWorkspaceResultOutput {
	return o
}

func (o LookupWalWorkspaceResultOutput) ToLookupWalWorkspaceResultOutputWithContext(ctx context.Context) LookupWalWorkspaceResultOutput {
	return o
}

func (o LookupWalWorkspaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWalWorkspaceResult] {
	return pulumix.Output[LookupWalWorkspaceResult]{
		OutputState: o.OutputState,
	}
}

// An array of key-value pairs to apply to this resource.
func (o LookupWalWorkspaceResultOutput) Tags() WalWorkspaceTagArrayOutput {
	return o.ApplyT(func(v LookupWalWorkspaceResult) []WalWorkspaceTag { return v.Tags }).(WalWorkspaceTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWalWorkspaceResultOutput{})
}