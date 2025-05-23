// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for a Device Farm Project
type Project struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the project. See [Amazon resource names](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference guide* .
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Sets the execution timeout value (in minutes) for a project. All test runs in this project use the specified execution timeout value unless overridden when scheduling a run.
	DefaultJobTimeoutMinutes pulumi.IntPtrOutput `pulumi:"defaultJobTimeoutMinutes"`
	// The project's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The tags to add to the resource. A tag is an array of key-value pairs. Tag keys can have a maximum character length of 128 characters. Tag values can have a maximum length of 256 characters.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The VPC security groups and subnets that are attached to a project.
	VpcConfig ProjectVpcConfigPtrOutput `pulumi:"vpcConfig"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("aws-native:devicefarm:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("aws-native:devicefarm:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
}

type ProjectState struct {
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Sets the execution timeout value (in minutes) for a project. All test runs in this project use the specified execution timeout value unless overridden when scheduling a run.
	DefaultJobTimeoutMinutes *int `pulumi:"defaultJobTimeoutMinutes"`
	// The project's name.
	Name *string `pulumi:"name"`
	// The tags to add to the resource. A tag is an array of key-value pairs. Tag keys can have a maximum character length of 128 characters. Tag values can have a maximum length of 256 characters.
	Tags []aws.Tag `pulumi:"tags"`
	// The VPC security groups and subnets that are attached to a project.
	VpcConfig *ProjectVpcConfig `pulumi:"vpcConfig"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Sets the execution timeout value (in minutes) for a project. All test runs in this project use the specified execution timeout value unless overridden when scheduling a run.
	DefaultJobTimeoutMinutes pulumi.IntPtrInput
	// The project's name.
	Name pulumi.StringPtrInput
	// The tags to add to the resource. A tag is an array of key-value pairs. Tag keys can have a maximum character length of 128 characters. Tag values can have a maximum length of 256 characters.
	Tags aws.TagArrayInput
	// The VPC security groups and subnets that are attached to a project.
	VpcConfig ProjectVpcConfigPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// The Amazon Resource Name (ARN) of the project. See [Amazon resource names](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference guide* .
func (o ProjectOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Sets the execution timeout value (in minutes) for a project. All test runs in this project use the specified execution timeout value unless overridden when scheduling a run.
func (o ProjectOutput) DefaultJobTimeoutMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.IntPtrOutput { return v.DefaultJobTimeoutMinutes }).(pulumi.IntPtrOutput)
}

// The project's name.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The tags to add to the resource. A tag is an array of key-value pairs. Tag keys can have a maximum character length of 128 characters. Tag values can have a maximum length of 256 characters.
func (o ProjectOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Project) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The VPC security groups and subnets that are attached to a project.
func (o ProjectOutput) VpcConfig() ProjectVpcConfigPtrOutput {
	return o.ApplyT(func(v *Project) ProjectVpcConfigPtrOutput { return v.VpcConfig }).(ProjectVpcConfigPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterOutputType(ProjectOutput{})
}
