// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspaces

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WorkSpaces::WorkspacesPool
func LookupWorkspacesPool(ctx *pulumi.Context, args *LookupWorkspacesPoolArgs, opts ...pulumi.InvokeOption) (*LookupWorkspacesPoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspacesPoolResult
	err := ctx.Invoke("aws-native:workspaces:getWorkspacesPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspacesPoolArgs struct {
	// The identifier of the pool.
	PoolId string `pulumi:"poolId"`
}

type LookupWorkspacesPoolResult struct {
	// The persistent application settings for users of the pool.
	ApplicationSettings *WorkspacesPoolApplicationSettings `pulumi:"applicationSettings"`
	// The identifier of the bundle used by the pool.
	BundleId *string `pulumi:"bundleId"`
	// Describes the user capacity for the pool.
	Capacity *WorkspacesPoolCapacity `pulumi:"capacity"`
	// The time the pool was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The description of the pool.
	Description *string `pulumi:"description"`
	// The identifier of the directory used by the pool.
	DirectoryId *string `pulumi:"directoryId"`
	// The Amazon Resource Name (ARN) for the pool.
	PoolArn *string `pulumi:"poolArn"`
	// The identifier of the pool.
	PoolId *string `pulumi:"poolId"`
	// The amount of time that a pool session remains active after users disconnect. If they try to reconnect to the pool session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new pool instance.
	TimeoutSettings *WorkspacesPoolTimeoutSettings `pulumi:"timeoutSettings"`
}

func LookupWorkspacesPoolOutput(ctx *pulumi.Context, args LookupWorkspacesPoolOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspacesPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspacesPoolResult, error) {
			args := v.(LookupWorkspacesPoolArgs)
			r, err := LookupWorkspacesPool(ctx, &args, opts...)
			var s LookupWorkspacesPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspacesPoolResultOutput)
}

type LookupWorkspacesPoolOutputArgs struct {
	// The identifier of the pool.
	PoolId pulumi.StringInput `pulumi:"poolId"`
}

func (LookupWorkspacesPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspacesPoolArgs)(nil)).Elem()
}

type LookupWorkspacesPoolResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspacesPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspacesPoolResult)(nil)).Elem()
}

func (o LookupWorkspacesPoolResultOutput) ToLookupWorkspacesPoolResultOutput() LookupWorkspacesPoolResultOutput {
	return o
}

func (o LookupWorkspacesPoolResultOutput) ToLookupWorkspacesPoolResultOutputWithContext(ctx context.Context) LookupWorkspacesPoolResultOutput {
	return o
}

// The persistent application settings for users of the pool.
func (o LookupWorkspacesPoolResultOutput) ApplicationSettings() WorkspacesPoolApplicationSettingsPtrOutput {
	return o.ApplyT(func(v LookupWorkspacesPoolResult) *WorkspacesPoolApplicationSettings { return v.ApplicationSettings }).(WorkspacesPoolApplicationSettingsPtrOutput)
}

// The identifier of the bundle used by the pool.
func (o LookupWorkspacesPoolResultOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspacesPoolResult) *string { return v.BundleId }).(pulumi.StringPtrOutput)
}

// Describes the user capacity for the pool.
func (o LookupWorkspacesPoolResultOutput) Capacity() WorkspacesPoolCapacityPtrOutput {
	return o.ApplyT(func(v LookupWorkspacesPoolResult) *WorkspacesPoolCapacity { return v.Capacity }).(WorkspacesPoolCapacityPtrOutput)
}

// The time the pool was created.
func (o LookupWorkspacesPoolResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspacesPoolResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The description of the pool.
func (o LookupWorkspacesPoolResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspacesPoolResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The identifier of the directory used by the pool.
func (o LookupWorkspacesPoolResultOutput) DirectoryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspacesPoolResult) *string { return v.DirectoryId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) for the pool.
func (o LookupWorkspacesPoolResultOutput) PoolArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspacesPoolResult) *string { return v.PoolArn }).(pulumi.StringPtrOutput)
}

// The identifier of the pool.
func (o LookupWorkspacesPoolResultOutput) PoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspacesPoolResult) *string { return v.PoolId }).(pulumi.StringPtrOutput)
}

// The amount of time that a pool session remains active after users disconnect. If they try to reconnect to the pool session after a disconnection or network interruption within this time interval, they are connected to their previous session. Otherwise, they are connected to a new session with a new pool instance.
func (o LookupWorkspacesPoolResultOutput) TimeoutSettings() WorkspacesPoolTimeoutSettingsPtrOutput {
	return o.ApplyT(func(v LookupWorkspacesPoolResult) *WorkspacesPoolTimeoutSettings { return v.TimeoutSettings }).(WorkspacesPoolTimeoutSettingsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspacesPoolResultOutput{})
}