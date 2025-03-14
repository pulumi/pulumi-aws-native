// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Macie AllowList resource schema
type AllowList struct {
	pulumi.CustomResourceState

	// AllowList ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// AllowList ID.
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// AllowList criteria.
	Criteria AllowListCriteriaOutput `pulumi:"criteria"`
	// Description of AllowList.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of AllowList.
	Name pulumi.StringOutput `pulumi:"name"`
	// AllowList status.
	Status AllowListStatusOutput `pulumi:"status"`
	// A collection of tags associated with a resource
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewAllowList registers a new resource with the given unique name, arguments, and options.
func NewAllowList(ctx *pulumi.Context,
	name string, args *AllowListArgs, opts ...pulumi.ResourceOption) (*AllowList, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AllowList
	err := ctx.RegisterResource("aws-native:macie:AllowList", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAllowList gets an existing AllowList resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAllowList(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AllowListState, opts ...pulumi.ResourceOption) (*AllowList, error) {
	var resource AllowList
	err := ctx.ReadResource("aws-native:macie:AllowList", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AllowList resources.
type allowListState struct {
}

type AllowListState struct {
}

func (AllowListState) ElementType() reflect.Type {
	return reflect.TypeOf((*allowListState)(nil)).Elem()
}

type allowListArgs struct {
	// AllowList criteria.
	Criteria AllowListCriteria `pulumi:"criteria"`
	// Description of AllowList.
	Description *string `pulumi:"description"`
	// Name of AllowList.
	Name *string `pulumi:"name"`
	// A collection of tags associated with a resource
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a AllowList resource.
type AllowListArgs struct {
	// AllowList criteria.
	Criteria AllowListCriteriaInput
	// Description of AllowList.
	Description pulumi.StringPtrInput
	// Name of AllowList.
	Name pulumi.StringPtrInput
	// A collection of tags associated with a resource
	Tags aws.TagArrayInput
}

func (AllowListArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*allowListArgs)(nil)).Elem()
}

type AllowListInput interface {
	pulumi.Input

	ToAllowListOutput() AllowListOutput
	ToAllowListOutputWithContext(ctx context.Context) AllowListOutput
}

func (*AllowList) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowList)(nil)).Elem()
}

func (i *AllowList) ToAllowListOutput() AllowListOutput {
	return i.ToAllowListOutputWithContext(context.Background())
}

func (i *AllowList) ToAllowListOutputWithContext(ctx context.Context) AllowListOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowListOutput)
}

type AllowListOutput struct{ *pulumi.OutputState }

func (AllowListOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowList)(nil)).Elem()
}

func (o AllowListOutput) ToAllowListOutput() AllowListOutput {
	return o
}

func (o AllowListOutput) ToAllowListOutputWithContext(ctx context.Context) AllowListOutput {
	return o
}

// AllowList ARN.
func (o AllowListOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *AllowList) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// AllowList ID.
func (o AllowListOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *AllowList) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// AllowList criteria.
func (o AllowListOutput) Criteria() AllowListCriteriaOutput {
	return o.ApplyT(func(v *AllowList) AllowListCriteriaOutput { return v.Criteria }).(AllowListCriteriaOutput)
}

// Description of AllowList.
func (o AllowListOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AllowList) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of AllowList.
func (o AllowListOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AllowList) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// AllowList status.
func (o AllowListOutput) Status() AllowListStatusOutput {
	return o.ApplyT(func(v *AllowList) AllowListStatusOutput { return v.Status }).(AllowListStatusOutput)
}

// A collection of tags associated with a resource
func (o AllowListOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *AllowList) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AllowListInput)(nil)).Elem(), &AllowList{})
	pulumi.RegisterOutputType(AllowListOutput{})
}
