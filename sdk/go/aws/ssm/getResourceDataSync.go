// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSM::ResourceDataSync
func LookupResourceDataSync(ctx *pulumi.Context, args *LookupResourceDataSyncArgs, opts ...pulumi.InvokeOption) (*LookupResourceDataSyncResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceDataSyncResult
	err := ctx.Invoke("aws-native:ssm:getResourceDataSync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceDataSyncArgs struct {
	// A name for the resource data sync.
	SyncName string `pulumi:"syncName"`
}

type LookupResourceDataSyncResult struct {
	// Information about the source where the data was synchronized.
	SyncSource *ResourceDataSyncSyncSource `pulumi:"syncSource"`
}

func LookupResourceDataSyncOutput(ctx *pulumi.Context, args LookupResourceDataSyncOutputArgs, opts ...pulumi.InvokeOption) LookupResourceDataSyncResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupResourceDataSyncResultOutput, error) {
			args := v.(LookupResourceDataSyncArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ssm:getResourceDataSync", args, LookupResourceDataSyncResultOutput{}, options).(LookupResourceDataSyncResultOutput), nil
		}).(LookupResourceDataSyncResultOutput)
}

type LookupResourceDataSyncOutputArgs struct {
	// A name for the resource data sync.
	SyncName pulumi.StringInput `pulumi:"syncName"`
}

func (LookupResourceDataSyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceDataSyncArgs)(nil)).Elem()
}

type LookupResourceDataSyncResultOutput struct{ *pulumi.OutputState }

func (LookupResourceDataSyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceDataSyncResult)(nil)).Elem()
}

func (o LookupResourceDataSyncResultOutput) ToLookupResourceDataSyncResultOutput() LookupResourceDataSyncResultOutput {
	return o
}

func (o LookupResourceDataSyncResultOutput) ToLookupResourceDataSyncResultOutputWithContext(ctx context.Context) LookupResourceDataSyncResultOutput {
	return o
}

// Information about the source where the data was synchronized.
func (o LookupResourceDataSyncResultOutput) SyncSource() ResourceDataSyncSyncSourcePtrOutput {
	return o.ApplyT(func(v LookupResourceDataSyncResult) *ResourceDataSyncSyncSource { return v.SyncSource }).(ResourceDataSyncSyncSourcePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceDataSyncResultOutput{})
}
