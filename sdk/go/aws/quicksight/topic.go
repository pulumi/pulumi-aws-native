// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Topic Resource Type.
type Topic struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the topic.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the AWS account that you want to create a topic in.
	AwsAccountId pulumi.StringPtrOutput `pulumi:"awsAccountId"`
	// Configuration options for a `Topic` .
	ConfigOptions TopicConfigOptionsPtrOutput `pulumi:"configOptions"`
	// The data sets that the topic is associated with.
	DataSets TopicDatasetMetadataArrayOutput `pulumi:"dataSets"`
	// The description of the topic.
	Description pulumi.StringPtrOutput   `pulumi:"description"`
	FolderArns  pulumi.StringArrayOutput `pulumi:"folderArns"`
	// The name of the topic.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The ID for the topic. This ID is unique per AWS Region for each AWS account.
	TopicId pulumi.StringPtrOutput `pulumi:"topicId"`
	// The user experience version of the topic.
	UserExperienceVersion TopicUserExperienceVersionPtrOutput `pulumi:"userExperienceVersion"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		args = &TopicArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"awsAccountId",
		"folderArns[*]",
		"topicId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Topic
	err := ctx.RegisterResource("aws-native:quicksight:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("aws-native:quicksight:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
}

type TopicState struct {
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// The ID of the AWS account that you want to create a topic in.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// Configuration options for a `Topic` .
	ConfigOptions *TopicConfigOptions `pulumi:"configOptions"`
	// The data sets that the topic is associated with.
	DataSets []TopicDatasetMetadata `pulumi:"dataSets"`
	// The description of the topic.
	Description *string  `pulumi:"description"`
	FolderArns  []string `pulumi:"folderArns"`
	// The name of the topic.
	Name *string `pulumi:"name"`
	// The ID for the topic. This ID is unique per AWS Region for each AWS account.
	TopicId *string `pulumi:"topicId"`
	// The user experience version of the topic.
	UserExperienceVersion *TopicUserExperienceVersion `pulumi:"userExperienceVersion"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// The ID of the AWS account that you want to create a topic in.
	AwsAccountId pulumi.StringPtrInput
	// Configuration options for a `Topic` .
	ConfigOptions TopicConfigOptionsPtrInput
	// The data sets that the topic is associated with.
	DataSets TopicDatasetMetadataArrayInput
	// The description of the topic.
	Description pulumi.StringPtrInput
	FolderArns  pulumi.StringArrayInput
	// The name of the topic.
	Name pulumi.StringPtrInput
	// The ID for the topic. This ID is unique per AWS Region for each AWS account.
	TopicId pulumi.StringPtrInput
	// The user experience version of the topic.
	UserExperienceVersion TopicUserExperienceVersionPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

// The Amazon Resource Name (ARN) of the topic.
func (o TopicOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The ID of the AWS account that you want to create a topic in.
func (o TopicOutput) AwsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.AwsAccountId }).(pulumi.StringPtrOutput)
}

// Configuration options for a `Topic` .
func (o TopicOutput) ConfigOptions() TopicConfigOptionsPtrOutput {
	return o.ApplyT(func(v *Topic) TopicConfigOptionsPtrOutput { return v.ConfigOptions }).(TopicConfigOptionsPtrOutput)
}

// The data sets that the topic is associated with.
func (o TopicOutput) DataSets() TopicDatasetMetadataArrayOutput {
	return o.ApplyT(func(v *Topic) TopicDatasetMetadataArrayOutput { return v.DataSets }).(TopicDatasetMetadataArrayOutput)
}

// The description of the topic.
func (o TopicOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TopicOutput) FolderArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringArrayOutput { return v.FolderArns }).(pulumi.StringArrayOutput)
}

// The name of the topic.
func (o TopicOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The ID for the topic. This ID is unique per AWS Region for each AWS account.
func (o TopicOutput) TopicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.TopicId }).(pulumi.StringPtrOutput)
}

// The user experience version of the topic.
func (o TopicOutput) UserExperienceVersion() TopicUserExperienceVersionPtrOutput {
	return o.ApplyT(func(v *Topic) TopicUserExperienceVersionPtrOutput { return v.UserExperienceVersion }).(TopicUserExperienceVersionPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicInput)(nil)).Elem(), &Topic{})
	pulumi.RegisterOutputType(TopicOutput{})
}
