// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cleanroomsml

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::CleanRoomsML::TrainingDataset Resource Type
type TrainingDataset struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput      `pulumi:"description"`
	Name        pulumi.StringOutput         `pulumi:"name"`
	RoleArn     pulumi.StringOutput         `pulumi:"roleArn"`
	Status      TrainingDatasetStatusOutput `pulumi:"status"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms-ml training dataset.
	Tags               aws.TagArrayOutput                `pulumi:"tags"`
	TrainingData       TrainingDatasetDatasetArrayOutput `pulumi:"trainingData"`
	TrainingDatasetArn pulumi.StringOutput               `pulumi:"trainingDatasetArn"`
}

// NewTrainingDataset registers a new resource with the given unique name, arguments, and options.
func NewTrainingDataset(ctx *pulumi.Context,
	name string, args *TrainingDatasetArgs, opts ...pulumi.ResourceOption) (*TrainingDataset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.TrainingData == nil {
		return nil, errors.New("invalid value for required argument 'TrainingData'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"description",
		"name",
		"roleArn",
		"trainingData[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TrainingDataset
	err := ctx.RegisterResource("aws-native:cleanroomsml:TrainingDataset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrainingDataset gets an existing TrainingDataset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrainingDataset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrainingDatasetState, opts ...pulumi.ResourceOption) (*TrainingDataset, error) {
	var resource TrainingDataset
	err := ctx.ReadResource("aws-native:cleanroomsml:TrainingDataset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrainingDataset resources.
type trainingDatasetState struct {
}

type TrainingDatasetState struct {
}

func (TrainingDatasetState) ElementType() reflect.Type {
	return reflect.TypeOf((*trainingDatasetState)(nil)).Elem()
}

type trainingDatasetArgs struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	RoleArn     string  `pulumi:"roleArn"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms-ml training dataset.
	Tags         []aws.Tag                `pulumi:"tags"`
	TrainingData []TrainingDatasetDataset `pulumi:"trainingData"`
}

// The set of arguments for constructing a TrainingDataset resource.
type TrainingDatasetArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	RoleArn     pulumi.StringInput
	// An arbitrary set of tags (key-value pairs) for this cleanrooms-ml training dataset.
	Tags         aws.TagArrayInput
	TrainingData TrainingDatasetDatasetArrayInput
}

func (TrainingDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trainingDatasetArgs)(nil)).Elem()
}

type TrainingDatasetInput interface {
	pulumi.Input

	ToTrainingDatasetOutput() TrainingDatasetOutput
	ToTrainingDatasetOutputWithContext(ctx context.Context) TrainingDatasetOutput
}

func (*TrainingDataset) ElementType() reflect.Type {
	return reflect.TypeOf((**TrainingDataset)(nil)).Elem()
}

func (i *TrainingDataset) ToTrainingDatasetOutput() TrainingDatasetOutput {
	return i.ToTrainingDatasetOutputWithContext(context.Background())
}

func (i *TrainingDataset) ToTrainingDatasetOutputWithContext(ctx context.Context) TrainingDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrainingDatasetOutput)
}

type TrainingDatasetOutput struct{ *pulumi.OutputState }

func (TrainingDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrainingDataset)(nil)).Elem()
}

func (o TrainingDatasetOutput) ToTrainingDatasetOutput() TrainingDatasetOutput {
	return o
}

func (o TrainingDatasetOutput) ToTrainingDatasetOutputWithContext(ctx context.Context) TrainingDatasetOutput {
	return o
}

func (o TrainingDatasetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrainingDataset) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TrainingDatasetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrainingDataset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TrainingDatasetOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *TrainingDataset) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o TrainingDatasetOutput) Status() TrainingDatasetStatusOutput {
	return o.ApplyT(func(v *TrainingDataset) TrainingDatasetStatusOutput { return v.Status }).(TrainingDatasetStatusOutput)
}

// An arbitrary set of tags (key-value pairs) for this cleanrooms-ml training dataset.
func (o TrainingDatasetOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *TrainingDataset) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func (o TrainingDatasetOutput) TrainingData() TrainingDatasetDatasetArrayOutput {
	return o.ApplyT(func(v *TrainingDataset) TrainingDatasetDatasetArrayOutput { return v.TrainingData }).(TrainingDatasetDatasetArrayOutput)
}

func (o TrainingDatasetOutput) TrainingDatasetArn() pulumi.StringOutput {
	return o.ApplyT(func(v *TrainingDataset) pulumi.StringOutput { return v.TrainingDatasetArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrainingDatasetInput)(nil)).Elem(), &TrainingDataset{})
	pulumi.RegisterOutputType(TrainingDatasetOutput{})
}