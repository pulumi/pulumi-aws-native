// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package medialive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::MediaLive::CloudWatchAlarmTemplateGroup Resource Type
type CloudWatchAlarmTemplateGroup struct {
	pulumi.CustomResourceState

	// A cloudwatch alarm template group's ARN (Amazon Resource Name)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-`
	AwsId pulumi.StringOutput `pulumi:"awsId"`
	// The date and time of resource creation.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A resource's optional description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Identifier  pulumi.StringOutput    `pulumi:"identifier"`
	// The date and time of latest resource modification.
	ModifiedAt pulumi.StringOutput `pulumi:"modifiedAt"`
	// A resource's name. Names must be unique within the scope of a resource type in a specific region.
	Name pulumi.StringOutput    `pulumi:"name"`
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewCloudWatchAlarmTemplateGroup registers a new resource with the given unique name, arguments, and options.
func NewCloudWatchAlarmTemplateGroup(ctx *pulumi.Context,
	name string, args *CloudWatchAlarmTemplateGroupArgs, opts ...pulumi.ResourceOption) (*CloudWatchAlarmTemplateGroup, error) {
	if args == nil {
		args = &CloudWatchAlarmTemplateGroupArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"tags.*",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudWatchAlarmTemplateGroup
	err := ctx.RegisterResource("aws-native:medialive:CloudWatchAlarmTemplateGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudWatchAlarmTemplateGroup gets an existing CloudWatchAlarmTemplateGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudWatchAlarmTemplateGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudWatchAlarmTemplateGroupState, opts ...pulumi.ResourceOption) (*CloudWatchAlarmTemplateGroup, error) {
	var resource CloudWatchAlarmTemplateGroup
	err := ctx.ReadResource("aws-native:medialive:CloudWatchAlarmTemplateGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudWatchAlarmTemplateGroup resources.
type cloudWatchAlarmTemplateGroupState struct {
}

type CloudWatchAlarmTemplateGroupState struct {
}

func (CloudWatchAlarmTemplateGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudWatchAlarmTemplateGroupState)(nil)).Elem()
}

type cloudWatchAlarmTemplateGroupArgs struct {
	// A resource's optional description.
	Description *string `pulumi:"description"`
	// A resource's name. Names must be unique within the scope of a resource type in a specific region.
	Name *string           `pulumi:"name"`
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a CloudWatchAlarmTemplateGroup resource.
type CloudWatchAlarmTemplateGroupArgs struct {
	// A resource's optional description.
	Description pulumi.StringPtrInput
	// A resource's name. Names must be unique within the scope of a resource type in a specific region.
	Name pulumi.StringPtrInput
	Tags pulumi.StringMapInput
}

func (CloudWatchAlarmTemplateGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudWatchAlarmTemplateGroupArgs)(nil)).Elem()
}

type CloudWatchAlarmTemplateGroupInput interface {
	pulumi.Input

	ToCloudWatchAlarmTemplateGroupOutput() CloudWatchAlarmTemplateGroupOutput
	ToCloudWatchAlarmTemplateGroupOutputWithContext(ctx context.Context) CloudWatchAlarmTemplateGroupOutput
}

func (*CloudWatchAlarmTemplateGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudWatchAlarmTemplateGroup)(nil)).Elem()
}

func (i *CloudWatchAlarmTemplateGroup) ToCloudWatchAlarmTemplateGroupOutput() CloudWatchAlarmTemplateGroupOutput {
	return i.ToCloudWatchAlarmTemplateGroupOutputWithContext(context.Background())
}

func (i *CloudWatchAlarmTemplateGroup) ToCloudWatchAlarmTemplateGroupOutputWithContext(ctx context.Context) CloudWatchAlarmTemplateGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudWatchAlarmTemplateGroupOutput)
}

type CloudWatchAlarmTemplateGroupOutput struct{ *pulumi.OutputState }

func (CloudWatchAlarmTemplateGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudWatchAlarmTemplateGroup)(nil)).Elem()
}

func (o CloudWatchAlarmTemplateGroupOutput) ToCloudWatchAlarmTemplateGroupOutput() CloudWatchAlarmTemplateGroupOutput {
	return o
}

func (o CloudWatchAlarmTemplateGroupOutput) ToCloudWatchAlarmTemplateGroupOutputWithContext(ctx context.Context) CloudWatchAlarmTemplateGroupOutput {
	return o
}

// A cloudwatch alarm template group's ARN (Amazon Resource Name)
func (o CloudWatchAlarmTemplateGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudWatchAlarmTemplateGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-`
func (o CloudWatchAlarmTemplateGroupOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudWatchAlarmTemplateGroup) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

// The date and time of resource creation.
func (o CloudWatchAlarmTemplateGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudWatchAlarmTemplateGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// A resource's optional description.
func (o CloudWatchAlarmTemplateGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudWatchAlarmTemplateGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CloudWatchAlarmTemplateGroupOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudWatchAlarmTemplateGroup) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// The date and time of latest resource modification.
func (o CloudWatchAlarmTemplateGroupOutput) ModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudWatchAlarmTemplateGroup) pulumi.StringOutput { return v.ModifiedAt }).(pulumi.StringOutput)
}

// A resource's name. Names must be unique within the scope of a resource type in a specific region.
func (o CloudWatchAlarmTemplateGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudWatchAlarmTemplateGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudWatchAlarmTemplateGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloudWatchAlarmTemplateGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudWatchAlarmTemplateGroupInput)(nil)).Elem(), &CloudWatchAlarmTemplateGroup{})
	pulumi.RegisterOutputType(CloudWatchAlarmTemplateGroupOutput{})
}