// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appstream

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppStream::AppBlockBuilder.
type AppBlockBuilder struct {
	pulumi.CustomResourceState

	// The access endpoints of the app block builder.
	AccessEndpoints AppBlockBuilderAccessEndpointArrayOutput `pulumi:"accessEndpoints"`
	// The ARN of the app block.
	//
	// *Maximum* : `1`
	AppBlockArns pulumi.StringArrayOutput `pulumi:"appBlockArns"`
	// The ARN of the app block builder.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The time when the app block builder was created.
	CreatedTime pulumi.StringOutput `pulumi:"createdTime"`
	// The description of the app block builder.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the app block builder.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Indicates whether default internet access is enabled for the app block builder.
	EnableDefaultInternetAccess pulumi.BoolPtrOutput `pulumi:"enableDefaultInternetAccess"`
	// The ARN of the IAM role that is applied to the app block builder.
	IamRoleArn pulumi.StringPtrOutput `pulumi:"iamRoleArn"`
	// The instance type of the app block builder.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
	// The name of the app block builder.
	Name pulumi.StringOutput `pulumi:"name"`
	// The platform of the app block builder.
	//
	// *Allowed values* : `WINDOWS_SERVER_2019`
	Platform pulumi.StringOutput `pulumi:"platform"`
	// The tags of the app block builder.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The VPC configuration for the app block builder.
	VpcConfig AppBlockBuilderVpcConfigOutput `pulumi:"vpcConfig"`
}

// NewAppBlockBuilder registers a new resource with the given unique name, arguments, and options.
func NewAppBlockBuilder(ctx *pulumi.Context,
	name string, args *AppBlockBuilderArgs, opts ...pulumi.ResourceOption) (*AppBlockBuilder, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	if args.Platform == nil {
		return nil, errors.New("invalid value for required argument 'Platform'")
	}
	if args.VpcConfig == nil {
		return nil, errors.New("invalid value for required argument 'VpcConfig'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AppBlockBuilder
	err := ctx.RegisterResource("aws-native:appstream:AppBlockBuilder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppBlockBuilder gets an existing AppBlockBuilder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppBlockBuilder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppBlockBuilderState, opts ...pulumi.ResourceOption) (*AppBlockBuilder, error) {
	var resource AppBlockBuilder
	err := ctx.ReadResource("aws-native:appstream:AppBlockBuilder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppBlockBuilder resources.
type appBlockBuilderState struct {
}

type AppBlockBuilderState struct {
}

func (AppBlockBuilderState) ElementType() reflect.Type {
	return reflect.TypeOf((*appBlockBuilderState)(nil)).Elem()
}

type appBlockBuilderArgs struct {
	// The access endpoints of the app block builder.
	AccessEndpoints []AppBlockBuilderAccessEndpoint `pulumi:"accessEndpoints"`
	// The ARN of the app block.
	//
	// *Maximum* : `1`
	AppBlockArns []string `pulumi:"appBlockArns"`
	// The description of the app block builder.
	Description *string `pulumi:"description"`
	// The display name of the app block builder.
	DisplayName *string `pulumi:"displayName"`
	// Indicates whether default internet access is enabled for the app block builder.
	EnableDefaultInternetAccess *bool `pulumi:"enableDefaultInternetAccess"`
	// The ARN of the IAM role that is applied to the app block builder.
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// The instance type of the app block builder.
	InstanceType string `pulumi:"instanceType"`
	// The name of the app block builder.
	Name *string `pulumi:"name"`
	// The platform of the app block builder.
	//
	// *Allowed values* : `WINDOWS_SERVER_2019`
	Platform string `pulumi:"platform"`
	// The tags of the app block builder.
	Tags []aws.Tag `pulumi:"tags"`
	// The VPC configuration for the app block builder.
	VpcConfig AppBlockBuilderVpcConfig `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a AppBlockBuilder resource.
type AppBlockBuilderArgs struct {
	// The access endpoints of the app block builder.
	AccessEndpoints AppBlockBuilderAccessEndpointArrayInput
	// The ARN of the app block.
	//
	// *Maximum* : `1`
	AppBlockArns pulumi.StringArrayInput
	// The description of the app block builder.
	Description pulumi.StringPtrInput
	// The display name of the app block builder.
	DisplayName pulumi.StringPtrInput
	// Indicates whether default internet access is enabled for the app block builder.
	EnableDefaultInternetAccess pulumi.BoolPtrInput
	// The ARN of the IAM role that is applied to the app block builder.
	IamRoleArn pulumi.StringPtrInput
	// The instance type of the app block builder.
	InstanceType pulumi.StringInput
	// The name of the app block builder.
	Name pulumi.StringPtrInput
	// The platform of the app block builder.
	//
	// *Allowed values* : `WINDOWS_SERVER_2019`
	Platform pulumi.StringInput
	// The tags of the app block builder.
	Tags aws.TagArrayInput
	// The VPC configuration for the app block builder.
	VpcConfig AppBlockBuilderVpcConfigInput
}

func (AppBlockBuilderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appBlockBuilderArgs)(nil)).Elem()
}

type AppBlockBuilderInput interface {
	pulumi.Input

	ToAppBlockBuilderOutput() AppBlockBuilderOutput
	ToAppBlockBuilderOutputWithContext(ctx context.Context) AppBlockBuilderOutput
}

func (*AppBlockBuilder) ElementType() reflect.Type {
	return reflect.TypeOf((**AppBlockBuilder)(nil)).Elem()
}

func (i *AppBlockBuilder) ToAppBlockBuilderOutput() AppBlockBuilderOutput {
	return i.ToAppBlockBuilderOutputWithContext(context.Background())
}

func (i *AppBlockBuilder) ToAppBlockBuilderOutputWithContext(ctx context.Context) AppBlockBuilderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppBlockBuilderOutput)
}

type AppBlockBuilderOutput struct{ *pulumi.OutputState }

func (AppBlockBuilderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppBlockBuilder)(nil)).Elem()
}

func (o AppBlockBuilderOutput) ToAppBlockBuilderOutput() AppBlockBuilderOutput {
	return o
}

func (o AppBlockBuilderOutput) ToAppBlockBuilderOutputWithContext(ctx context.Context) AppBlockBuilderOutput {
	return o
}

// The access endpoints of the app block builder.
func (o AppBlockBuilderOutput) AccessEndpoints() AppBlockBuilderAccessEndpointArrayOutput {
	return o.ApplyT(func(v *AppBlockBuilder) AppBlockBuilderAccessEndpointArrayOutput { return v.AccessEndpoints }).(AppBlockBuilderAccessEndpointArrayOutput)
}

// The ARN of the app block.
//
// *Maximum* : `1`
func (o AppBlockBuilderOutput) AppBlockArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringArrayOutput { return v.AppBlockArns }).(pulumi.StringArrayOutput)
}

// The ARN of the app block builder.
func (o AppBlockBuilderOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The time when the app block builder was created.
func (o AppBlockBuilderOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

// The description of the app block builder.
func (o AppBlockBuilderOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the app block builder.
func (o AppBlockBuilderOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Indicates whether default internet access is enabled for the app block builder.
func (o AppBlockBuilderOutput) EnableDefaultInternetAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.BoolPtrOutput { return v.EnableDefaultInternetAccess }).(pulumi.BoolPtrOutput)
}

// The ARN of the IAM role that is applied to the app block builder.
func (o AppBlockBuilderOutput) IamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringPtrOutput { return v.IamRoleArn }).(pulumi.StringPtrOutput)
}

// The instance type of the app block builder.
func (o AppBlockBuilderOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

// The name of the app block builder.
func (o AppBlockBuilderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The platform of the app block builder.
//
// *Allowed values* : `WINDOWS_SERVER_2019`
func (o AppBlockBuilderOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v *AppBlockBuilder) pulumi.StringOutput { return v.Platform }).(pulumi.StringOutput)
}

// The tags of the app block builder.
func (o AppBlockBuilderOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *AppBlockBuilder) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The VPC configuration for the app block builder.
func (o AppBlockBuilderOutput) VpcConfig() AppBlockBuilderVpcConfigOutput {
	return o.ApplyT(func(v *AppBlockBuilder) AppBlockBuilderVpcConfigOutput { return v.VpcConfig }).(AppBlockBuilderVpcConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppBlockBuilderInput)(nil)).Elem(), &AppBlockBuilder{})
	pulumi.RegisterOutputType(AppBlockBuilderOutput{})
}
