// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The connection group for your distribution tenants. When you first create a distribution tenant and you don't specify a connection group, CloudFront will automatically create a default connection group for you. When you create a new distribution tenant and don't specify a connection group, the default one will be associated with your distribution tenant.
func LookupConnectionGroup(ctx *pulumi.Context, args *LookupConnectionGroupArgs, opts ...pulumi.InvokeOption) (*LookupConnectionGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConnectionGroupResult
	err := ctx.Invoke("aws-native:cloudfront:getConnectionGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionGroupArgs struct {
	// The ID of the connection group.
	Id string `pulumi:"id"`
}

type LookupConnectionGroupResult struct {
	// The ID of the Anycast static IP list.
	AnycastIpListId *string `pulumi:"anycastIpListId"`
	// The Amazon Resource Name (ARN) of the connection group.
	Arn *string `pulumi:"arn"`
	// The date and time when the connection group was created.
	CreatedTime *string `pulumi:"createdTime"`
	// The current version of the connection group.
	ETag *string `pulumi:"eTag"`
	// Whether the connection group is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the connection group.
	Id *string `pulumi:"id"`
	// IPv6 is enabled for the connection group.
	Ipv6Enabled *bool `pulumi:"ipv6Enabled"`
	// Whether the connection group is the default connection group for the distribution tenants.
	IsDefault *bool `pulumi:"isDefault"`
	// The date and time when the connection group was updated.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The routing endpoint (also known as the DNS name) that is assigned to the connection group, such as d111111abcdef8.cloudfront.net.
	RoutingEndpoint *string `pulumi:"routingEndpoint"`
	// The status of the connection group.
	Status *string `pulumi:"status"`
	// A complex type that contains zero or more ``Tag`` elements.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupConnectionGroupOutput(ctx *pulumi.Context, args LookupConnectionGroupOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupConnectionGroupResultOutput, error) {
			args := v.(LookupConnectionGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:cloudfront:getConnectionGroup", args, LookupConnectionGroupResultOutput{}, options).(LookupConnectionGroupResultOutput), nil
		}).(LookupConnectionGroupResultOutput)
}

type LookupConnectionGroupOutputArgs struct {
	// The ID of the connection group.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupConnectionGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionGroupArgs)(nil)).Elem()
}

type LookupConnectionGroupResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionGroupResult)(nil)).Elem()
}

func (o LookupConnectionGroupResultOutput) ToLookupConnectionGroupResultOutput() LookupConnectionGroupResultOutput {
	return o
}

func (o LookupConnectionGroupResultOutput) ToLookupConnectionGroupResultOutputWithContext(ctx context.Context) LookupConnectionGroupResultOutput {
	return o
}

// The ID of the Anycast static IP list.
func (o LookupConnectionGroupResultOutput) AnycastIpListId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *string { return v.AnycastIpListId }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the connection group.
func (o LookupConnectionGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The date and time when the connection group was created.
func (o LookupConnectionGroupResultOutput) CreatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *string { return v.CreatedTime }).(pulumi.StringPtrOutput)
}

// The current version of the connection group.
func (o LookupConnectionGroupResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

// Whether the connection group is enabled.
func (o LookupConnectionGroupResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The ID of the connection group.
func (o LookupConnectionGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// IPv6 is enabled for the connection group.
func (o LookupConnectionGroupResultOutput) Ipv6Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *bool { return v.Ipv6Enabled }).(pulumi.BoolPtrOutput)
}

// Whether the connection group is the default connection group for the distribution tenants.
func (o LookupConnectionGroupResultOutput) IsDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *bool { return v.IsDefault }).(pulumi.BoolPtrOutput)
}

// The date and time when the connection group was updated.
func (o LookupConnectionGroupResultOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// The routing endpoint (also known as the DNS name) that is assigned to the connection group, such as d111111abcdef8.cloudfront.net.
func (o LookupConnectionGroupResultOutput) RoutingEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *string { return v.RoutingEndpoint }).(pulumi.StringPtrOutput)
}

// The status of the connection group.
func (o LookupConnectionGroupResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// A complex type that contains zero or more “Tag“ elements.
func (o LookupConnectionGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupConnectionGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionGroupResultOutput{})
}
