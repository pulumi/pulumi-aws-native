// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema of AWS::EC2::IPAMPool Type
func LookupIpamPool(ctx *pulumi.Context, args *LookupIpamPoolArgs, opts ...pulumi.InvokeOption) (*LookupIpamPoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpamPoolResult
	err := ctx.Invoke("aws-native:ec2:getIpamPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIpamPoolArgs struct {
	// Id of the IPAM Pool.
	IpamPoolId string `pulumi:"ipamPoolId"`
}

type LookupIpamPoolResult struct {
	// The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.
	AllocationDefaultNetmaskLength *int `pulumi:"allocationDefaultNetmaskLength"`
	// The maximum allowed netmask length for allocations made from this pool.
	AllocationMaxNetmaskLength *int `pulumi:"allocationMaxNetmaskLength"`
	// The minimum allowed netmask length for allocations made from this pool.
	AllocationMinNetmaskLength *int `pulumi:"allocationMinNetmaskLength"`
	// When specified, an allocation will not be allowed unless a resource has a matching set of tags.
	AllocationResourceTags []IpamPoolTag `pulumi:"allocationResourceTags"`
	// The Amazon Resource Name (ARN) of the IPAM Pool.
	Arn *string `pulumi:"arn"`
	// Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.
	AutoImport *bool `pulumi:"autoImport"`
	// The description of the IPAM pool.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the IPAM this pool is a part of.
	IpamArn *string `pulumi:"ipamArn"`
	// Id of the IPAM Pool.
	IpamPoolId *string `pulumi:"ipamPoolId"`
	// The Amazon Resource Name (ARN) of the scope this pool is a part of.
	IpamScopeArn *string `pulumi:"ipamScopeArn"`
	// Determines whether this scope contains publicly routable space or space for a private network
	IpamScopeType *IpamPoolIpamScopeType `pulumi:"ipamScopeType"`
	// The depth of this pool in the source pool hierarchy.
	PoolDepth *int `pulumi:"poolDepth"`
	// A list of cidrs representing the address space available for allocation in this pool.
	ProvisionedCidrs []IpamPoolProvisionedCidr `pulumi:"provisionedCidrs"`
	// The state of this pool. This can be one of the following values: "create-in-progress", "create-complete", "modify-in-progress", "modify-complete", "delete-in-progress", or "delete-complete"
	State *IpamPoolStateEnum `pulumi:"state"`
	// An explanation of how the pool arrived at it current state.
	StateMessage *string `pulumi:"stateMessage"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupIpamPoolOutput(ctx *pulumi.Context, args LookupIpamPoolOutputArgs, opts ...pulumi.InvokeOption) LookupIpamPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIpamPoolResultOutput, error) {
			args := v.(LookupIpamPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:ec2:getIpamPool", args, LookupIpamPoolResultOutput{}, options).(LookupIpamPoolResultOutput), nil
		}).(LookupIpamPoolResultOutput)
}

type LookupIpamPoolOutputArgs struct {
	// Id of the IPAM Pool.
	IpamPoolId pulumi.StringInput `pulumi:"ipamPoolId"`
}

func (LookupIpamPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpamPoolArgs)(nil)).Elem()
}

type LookupIpamPoolResultOutput struct{ *pulumi.OutputState }

func (LookupIpamPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpamPoolResult)(nil)).Elem()
}

func (o LookupIpamPoolResultOutput) ToLookupIpamPoolResultOutput() LookupIpamPoolResultOutput {
	return o
}

func (o LookupIpamPoolResultOutput) ToLookupIpamPoolResultOutputWithContext(ctx context.Context) LookupIpamPoolResultOutput {
	return o
}

// The default netmask length for allocations made from this pool. This value is used when the netmask length of an allocation isn't specified.
func (o LookupIpamPoolResultOutput) AllocationDefaultNetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *int { return v.AllocationDefaultNetmaskLength }).(pulumi.IntPtrOutput)
}

// The maximum allowed netmask length for allocations made from this pool.
func (o LookupIpamPoolResultOutput) AllocationMaxNetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *int { return v.AllocationMaxNetmaskLength }).(pulumi.IntPtrOutput)
}

// The minimum allowed netmask length for allocations made from this pool.
func (o LookupIpamPoolResultOutput) AllocationMinNetmaskLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *int { return v.AllocationMinNetmaskLength }).(pulumi.IntPtrOutput)
}

// When specified, an allocation will not be allowed unless a resource has a matching set of tags.
func (o LookupIpamPoolResultOutput) AllocationResourceTags() IpamPoolTagArrayOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) []IpamPoolTag { return v.AllocationResourceTags }).(IpamPoolTagArrayOutput)
}

// The Amazon Resource Name (ARN) of the IPAM Pool.
func (o LookupIpamPoolResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Determines what to do if IPAM discovers resources that haven't been assigned an allocation. If set to true, an allocation will be made automatically.
func (o LookupIpamPoolResultOutput) AutoImport() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *bool { return v.AutoImport }).(pulumi.BoolPtrOutput)
}

// The description of the IPAM pool.
func (o LookupIpamPoolResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the IPAM this pool is a part of.
func (o LookupIpamPoolResultOutput) IpamArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *string { return v.IpamArn }).(pulumi.StringPtrOutput)
}

// Id of the IPAM Pool.
func (o LookupIpamPoolResultOutput) IpamPoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *string { return v.IpamPoolId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the scope this pool is a part of.
func (o LookupIpamPoolResultOutput) IpamScopeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *string { return v.IpamScopeArn }).(pulumi.StringPtrOutput)
}

// Determines whether this scope contains publicly routable space or space for a private network
func (o LookupIpamPoolResultOutput) IpamScopeType() IpamPoolIpamScopeTypePtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *IpamPoolIpamScopeType { return v.IpamScopeType }).(IpamPoolIpamScopeTypePtrOutput)
}

// The depth of this pool in the source pool hierarchy.
func (o LookupIpamPoolResultOutput) PoolDepth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *int { return v.PoolDepth }).(pulumi.IntPtrOutput)
}

// A list of cidrs representing the address space available for allocation in this pool.
func (o LookupIpamPoolResultOutput) ProvisionedCidrs() IpamPoolProvisionedCidrArrayOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) []IpamPoolProvisionedCidr { return v.ProvisionedCidrs }).(IpamPoolProvisionedCidrArrayOutput)
}

// The state of this pool. This can be one of the following values: "create-in-progress", "create-complete", "modify-in-progress", "modify-complete", "delete-in-progress", or "delete-complete"
func (o LookupIpamPoolResultOutput) State() IpamPoolStateEnumPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *IpamPoolStateEnum { return v.State }).(IpamPoolStateEnumPtrOutput)
}

// An explanation of how the pool arrived at it current state.
func (o LookupIpamPoolResultOutput) StateMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) *string { return v.StateMessage }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupIpamPoolResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupIpamPoolResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpamPoolResultOutput{})
}
