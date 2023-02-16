// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Location::PlaceIndex Resource Type
type PlaceIndex struct {
	pulumi.CustomResourceState

	Arn                     pulumi.StringOutput                        `pulumi:"arn"`
	CreateTime              pulumi.StringOutput                        `pulumi:"createTime"`
	DataSource              pulumi.StringOutput                        `pulumi:"dataSource"`
	DataSourceConfiguration PlaceIndexDataSourceConfigurationPtrOutput `pulumi:"dataSourceConfiguration"`
	Description             pulumi.StringPtrOutput                     `pulumi:"description"`
	IndexArn                pulumi.StringOutput                        `pulumi:"indexArn"`
	IndexName               pulumi.StringOutput                        `pulumi:"indexName"`
	PricingPlan             PlaceIndexPricingPlanPtrOutput             `pulumi:"pricingPlan"`
	UpdateTime              pulumi.StringOutput                        `pulumi:"updateTime"`
}

// NewPlaceIndex registers a new resource with the given unique name, arguments, and options.
func NewPlaceIndex(ctx *pulumi.Context,
	name string, args *PlaceIndexArgs, opts ...pulumi.ResourceOption) (*PlaceIndex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataSource == nil {
		return nil, errors.New("invalid value for required argument 'DataSource'")
	}
	if args.IndexName == nil {
		return nil, errors.New("invalid value for required argument 'IndexName'")
	}
	var resource PlaceIndex
	err := ctx.RegisterResource("aws-native:location:PlaceIndex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlaceIndex gets an existing PlaceIndex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlaceIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlaceIndexState, opts ...pulumi.ResourceOption) (*PlaceIndex, error) {
	var resource PlaceIndex
	err := ctx.ReadResource("aws-native:location:PlaceIndex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlaceIndex resources.
type placeIndexState struct {
}

type PlaceIndexState struct {
}

func (PlaceIndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*placeIndexState)(nil)).Elem()
}

type placeIndexArgs struct {
	DataSource              string                             `pulumi:"dataSource"`
	DataSourceConfiguration *PlaceIndexDataSourceConfiguration `pulumi:"dataSourceConfiguration"`
	Description             *string                            `pulumi:"description"`
	IndexName               string                             `pulumi:"indexName"`
	PricingPlan             *PlaceIndexPricingPlan             `pulumi:"pricingPlan"`
}

// The set of arguments for constructing a PlaceIndex resource.
type PlaceIndexArgs struct {
	DataSource              pulumi.StringInput
	DataSourceConfiguration PlaceIndexDataSourceConfigurationPtrInput
	Description             pulumi.StringPtrInput
	IndexName               pulumi.StringInput
	PricingPlan             PlaceIndexPricingPlanPtrInput
}

func (PlaceIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*placeIndexArgs)(nil)).Elem()
}

type PlaceIndexInput interface {
	pulumi.Input

	ToPlaceIndexOutput() PlaceIndexOutput
	ToPlaceIndexOutputWithContext(ctx context.Context) PlaceIndexOutput
}

func (*PlaceIndex) ElementType() reflect.Type {
	return reflect.TypeOf((**PlaceIndex)(nil)).Elem()
}

func (i *PlaceIndex) ToPlaceIndexOutput() PlaceIndexOutput {
	return i.ToPlaceIndexOutputWithContext(context.Background())
}

func (i *PlaceIndex) ToPlaceIndexOutputWithContext(ctx context.Context) PlaceIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlaceIndexOutput)
}

type PlaceIndexOutput struct{ *pulumi.OutputState }

func (PlaceIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlaceIndex)(nil)).Elem()
}

func (o PlaceIndexOutput) ToPlaceIndexOutput() PlaceIndexOutput {
	return o
}

func (o PlaceIndexOutput) ToPlaceIndexOutputWithContext(ctx context.Context) PlaceIndexOutput {
	return o
}

func (o PlaceIndexOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaceIndex) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o PlaceIndexOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaceIndex) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o PlaceIndexOutput) DataSource() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaceIndex) pulumi.StringOutput { return v.DataSource }).(pulumi.StringOutput)
}

func (o PlaceIndexOutput) DataSourceConfiguration() PlaceIndexDataSourceConfigurationPtrOutput {
	return o.ApplyT(func(v *PlaceIndex) PlaceIndexDataSourceConfigurationPtrOutput { return v.DataSourceConfiguration }).(PlaceIndexDataSourceConfigurationPtrOutput)
}

func (o PlaceIndexOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlaceIndex) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PlaceIndexOutput) IndexArn() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaceIndex) pulumi.StringOutput { return v.IndexArn }).(pulumi.StringOutput)
}

func (o PlaceIndexOutput) IndexName() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaceIndex) pulumi.StringOutput { return v.IndexName }).(pulumi.StringOutput)
}

func (o PlaceIndexOutput) PricingPlan() PlaceIndexPricingPlanPtrOutput {
	return o.ApplyT(func(v *PlaceIndex) PlaceIndexPricingPlanPtrOutput { return v.PricingPlan }).(PlaceIndexPricingPlanPtrOutput)
}

func (o PlaceIndexOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaceIndex) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlaceIndexInput)(nil)).Elem(), &PlaceIndex{})
	pulumi.RegisterOutputType(PlaceIndexOutput{})
}