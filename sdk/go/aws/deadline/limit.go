// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deadline

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Deadline::Limit Resource Type
type Limit struct {
	pulumi.CustomResourceState

	// The value that you specify as the `name` in the `amounts` field of the `hostRequirements` in a step of a job template to declare the limit requirement.
	AmountRequirementName pulumi.StringOutput `pulumi:"amountRequirementName"`
	// The number of resources from the limit that are being used by jobs. The result is delayed and may not be the count at the time that you called the operation.
	CurrentCount pulumi.IntOutput `pulumi:"currentCount"`
	// A description of the limit. A clear description helps you identify the purpose of the limit.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the limit used in lists to identify the limit.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The unique identifier of the farm that contains the limit.
	FarmId pulumi.StringOutput `pulumi:"farmId"`
	// The unique identifier of the limit.
	LimitId pulumi.StringOutput `pulumi:"limitId"`
	// The maximum number of resources constrained by this limit. When all of the resources are in use, steps that require the limit won't be scheduled until the resource is available.
	//
	// The `maxValue` must not be 0. If the value is -1, there is no restriction on the number of resources that can be acquired for this limit.
	MaxCount pulumi.IntOutput `pulumi:"maxCount"`
}

// NewLimit registers a new resource with the given unique name, arguments, and options.
func NewLimit(ctx *pulumi.Context,
	name string, args *LimitArgs, opts ...pulumi.ResourceOption) (*Limit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AmountRequirementName == nil {
		return nil, errors.New("invalid value for required argument 'AmountRequirementName'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.FarmId == nil {
		return nil, errors.New("invalid value for required argument 'FarmId'")
	}
	if args.MaxCount == nil {
		return nil, errors.New("invalid value for required argument 'MaxCount'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"amountRequirementName",
		"farmId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Limit
	err := ctx.RegisterResource("aws-native:deadline:Limit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLimit gets an existing Limit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLimit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LimitState, opts ...pulumi.ResourceOption) (*Limit, error) {
	var resource Limit
	err := ctx.ReadResource("aws-native:deadline:Limit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Limit resources.
type limitState struct {
}

type LimitState struct {
}

func (LimitState) ElementType() reflect.Type {
	return reflect.TypeOf((*limitState)(nil)).Elem()
}

type limitArgs struct {
	// The value that you specify as the `name` in the `amounts` field of the `hostRequirements` in a step of a job template to declare the limit requirement.
	AmountRequirementName string `pulumi:"amountRequirementName"`
	// A description of the limit. A clear description helps you identify the purpose of the limit.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	Description *string `pulumi:"description"`
	// The name of the limit used in lists to identify the limit.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	DisplayName string `pulumi:"displayName"`
	// The unique identifier of the farm that contains the limit.
	FarmId string `pulumi:"farmId"`
	// The maximum number of resources constrained by this limit. When all of the resources are in use, steps that require the limit won't be scheduled until the resource is available.
	//
	// The `maxValue` must not be 0. If the value is -1, there is no restriction on the number of resources that can be acquired for this limit.
	MaxCount int `pulumi:"maxCount"`
}

// The set of arguments for constructing a Limit resource.
type LimitArgs struct {
	// The value that you specify as the `name` in the `amounts` field of the `hostRequirements` in a step of a job template to declare the limit requirement.
	AmountRequirementName pulumi.StringInput
	// A description of the limit. A clear description helps you identify the purpose of the limit.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	Description pulumi.StringPtrInput
	// The name of the limit used in lists to identify the limit.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	DisplayName pulumi.StringInput
	// The unique identifier of the farm that contains the limit.
	FarmId pulumi.StringInput
	// The maximum number of resources constrained by this limit. When all of the resources are in use, steps that require the limit won't be scheduled until the resource is available.
	//
	// The `maxValue` must not be 0. If the value is -1, there is no restriction on the number of resources that can be acquired for this limit.
	MaxCount pulumi.IntInput
}

func (LimitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*limitArgs)(nil)).Elem()
}

type LimitInput interface {
	pulumi.Input

	ToLimitOutput() LimitOutput
	ToLimitOutputWithContext(ctx context.Context) LimitOutput
}

func (*Limit) ElementType() reflect.Type {
	return reflect.TypeOf((**Limit)(nil)).Elem()
}

func (i *Limit) ToLimitOutput() LimitOutput {
	return i.ToLimitOutputWithContext(context.Background())
}

func (i *Limit) ToLimitOutputWithContext(ctx context.Context) LimitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LimitOutput)
}

type LimitOutput struct{ *pulumi.OutputState }

func (LimitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Limit)(nil)).Elem()
}

func (o LimitOutput) ToLimitOutput() LimitOutput {
	return o
}

func (o LimitOutput) ToLimitOutputWithContext(ctx context.Context) LimitOutput {
	return o
}

// The value that you specify as the `name` in the `amounts` field of the `hostRequirements` in a step of a job template to declare the limit requirement.
func (o LimitOutput) AmountRequirementName() pulumi.StringOutput {
	return o.ApplyT(func(v *Limit) pulumi.StringOutput { return v.AmountRequirementName }).(pulumi.StringOutput)
}

// The number of resources from the limit that are being used by jobs. The result is delayed and may not be the count at the time that you called the operation.
func (o LimitOutput) CurrentCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Limit) pulumi.IntOutput { return v.CurrentCount }).(pulumi.IntOutput)
}

// A description of the limit. A clear description helps you identify the purpose of the limit.
//
// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
func (o LimitOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Limit) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the limit used in lists to identify the limit.
//
// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
func (o LimitOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Limit) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The unique identifier of the farm that contains the limit.
func (o LimitOutput) FarmId() pulumi.StringOutput {
	return o.ApplyT(func(v *Limit) pulumi.StringOutput { return v.FarmId }).(pulumi.StringOutput)
}

// The unique identifier of the limit.
func (o LimitOutput) LimitId() pulumi.StringOutput {
	return o.ApplyT(func(v *Limit) pulumi.StringOutput { return v.LimitId }).(pulumi.StringOutput)
}

// The maximum number of resources constrained by this limit. When all of the resources are in use, steps that require the limit won't be scheduled until the resource is available.
//
// The `maxValue` must not be 0. If the value is -1, there is no restriction on the number of resources that can be acquired for this limit.
func (o LimitOutput) MaxCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Limit) pulumi.IntOutput { return v.MaxCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LimitInput)(nil)).Elem(), &Limit{})
	pulumi.RegisterOutputType(LimitOutput{})
}
