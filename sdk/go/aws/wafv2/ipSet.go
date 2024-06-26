// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains a list of IP addresses. This can be either IPV4 or IPV6. The list will be mutually
type IpSet struct {
	pulumi.CustomResourceState

	// List of IPAddresses.
	Addresses pulumi.StringArrayOutput `pulumi:"addresses"`
	// The Amazon Resource Name (ARN) of the IP set.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the IP set.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// A description of the IP set that helps with identification.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The version of the IP addresses, either `IPV4` or `IPV6` .
	IpAddressVersion IpSetIpAddressVersionOutput `pulumi:"ipAddressVersion"`
	// The name of the IP set. You cannot change the name of an `IPSet` after you create it.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application. A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope IpSetScopeOutput `pulumi:"scope"`
	// Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewIpSet registers a new resource with the given unique name, arguments, and options.
func NewIpSet(ctx *pulumi.Context,
	name string, args *IpSetArgs, opts ...pulumi.ResourceOption) (*IpSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Addresses == nil {
		return nil, errors.New("invalid value for required argument 'Addresses'")
	}
	if args.IpAddressVersion == nil {
		return nil, errors.New("invalid value for required argument 'IpAddressVersion'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"scope",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpSet
	err := ctx.RegisterResource("aws-native:wafv2:IpSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpSet gets an existing IpSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpSetState, opts ...pulumi.ResourceOption) (*IpSet, error) {
	var resource IpSet
	err := ctx.ReadResource("aws-native:wafv2:IpSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpSet resources.
type ipSetState struct {
}

type IpSetState struct {
}

func (IpSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetState)(nil)).Elem()
}

type ipSetArgs struct {
	// List of IPAddresses.
	Addresses []string `pulumi:"addresses"`
	// A description of the IP set that helps with identification.
	Description *string `pulumi:"description"`
	// The version of the IP addresses, either `IPV4` or `IPV6` .
	IpAddressVersion IpSetIpAddressVersion `pulumi:"ipAddressVersion"`
	// The name of the IP set. You cannot change the name of an `IPSet` after you create it.
	Name *string `pulumi:"name"`
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application. A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope IpSetScope `pulumi:"scope"`
	// Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a IpSet resource.
type IpSetArgs struct {
	// List of IPAddresses.
	Addresses pulumi.StringArrayInput
	// A description of the IP set that helps with identification.
	Description pulumi.StringPtrInput
	// The version of the IP addresses, either `IPV4` or `IPV6` .
	IpAddressVersion IpSetIpAddressVersionInput
	// The name of the IP set. You cannot change the name of an `IPSet` after you create it.
	Name pulumi.StringPtrInput
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application. A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope IpSetScopeInput
	// Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
	Tags aws.TagArrayInput
}

func (IpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipSetArgs)(nil)).Elem()
}

type IpSetInput interface {
	pulumi.Input

	ToIpSetOutput() IpSetOutput
	ToIpSetOutputWithContext(ctx context.Context) IpSetOutput
}

func (*IpSet) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSet)(nil)).Elem()
}

func (i *IpSet) ToIpSetOutput() IpSetOutput {
	return i.ToIpSetOutputWithContext(context.Background())
}

func (i *IpSet) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSetOutput)
}

type IpSetOutput struct{ *pulumi.OutputState }

func (IpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpSet)(nil)).Elem()
}

func (o IpSetOutput) ToIpSetOutput() IpSetOutput {
	return o
}

func (o IpSetOutput) ToIpSetOutputWithContext(ctx context.Context) IpSetOutput {
	return o
}

// List of IPAddresses.
func (o IpSetOutput) Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringArrayOutput { return v.Addresses }).(pulumi.StringArrayOutput)
}

// The Amazon Resource Name (ARN) of the IP set.
func (o IpSetOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the IP set.
func (o IpSetOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// A description of the IP set that helps with identification.
func (o IpSetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The version of the IP addresses, either `IPV4` or `IPV6` .
func (o IpSetOutput) IpAddressVersion() IpSetIpAddressVersionOutput {
	return o.ApplyT(func(v *IpSet) IpSetIpAddressVersionOutput { return v.IpAddressVersion }).(IpSetIpAddressVersionOutput)
}

// The name of the IP set. You cannot change the name of an `IPSet` after you create it.
func (o IpSetOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpSet) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Specifies whether this is for an Amazon CloudFront distribution or for a regional application. A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and `REGIONAL` .
//
// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
func (o IpSetOutput) Scope() IpSetScopeOutput {
	return o.ApplyT(func(v *IpSet) IpSetScopeOutput { return v.Scope }).(IpSetScopeOutput)
}

// Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
//
// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
func (o IpSetOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *IpSet) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpSetInput)(nil)).Elem(), &IpSet{})
	pulumi.RegisterOutputType(IpSetOutput{})
}
