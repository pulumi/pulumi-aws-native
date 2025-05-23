// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xray

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This schema provides construct and validation rules for AWS-XRay Group resource parameters.
//
// ## Example Usage
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/xray"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := xray.NewGroup(ctx, "myGroupResource", &xray.GroupArgs{
//				GroupName:        pulumi.String("MyGroup"),
//				FilterExpression: pulumi.String("duration > 10"),
//				InsightsConfiguration: &xray.GroupInsightsConfigurationArgs{
//					InsightsEnabled:      pulumi.Bool(false),
//					NotificationsEnabled: pulumi.Bool(false),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Example
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/xray"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := xray.NewGroup(ctx, "myGroupResource", &xray.GroupArgs{
//				GroupName:        pulumi.String("MyGroup"),
//				FilterExpression: pulumi.String("duration > 10"),
//				InsightsConfiguration: &xray.GroupInsightsConfigurationArgs{
//					InsightsEnabled:      pulumi.Bool(false),
//					NotificationsEnabled: pulumi.Bool(false),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Group struct {
	pulumi.CustomResourceState

	// The filter expression defining criteria by which to group traces.
	FilterExpression pulumi.StringPtrOutput `pulumi:"filterExpression"`
	// The ARN of the group that was generated on creation.
	GroupArn pulumi.StringOutput `pulumi:"groupArn"`
	// The case-sensitive name of the new group. Names must be unique.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// The structure containing configurations related to insights.
	//
	// - The InsightsEnabled boolean can be set to true to enable insights for the group or false to disable insights for the group.
	// - The NotificationsEnabled boolean can be set to true to enable insights notifications through Amazon EventBridge for the group.
	InsightsConfiguration GroupInsightsConfigurationPtrOutput `pulumi:"insightsConfiguration"`
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil {
		args = &GroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Group
	err := ctx.RegisterResource("aws-native:xray:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("aws-native:xray:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
}

type GroupState struct {
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The filter expression defining criteria by which to group traces.
	FilterExpression *string `pulumi:"filterExpression"`
	// The case-sensitive name of the new group. Names must be unique.
	GroupName *string `pulumi:"groupName"`
	// The structure containing configurations related to insights.
	//
	// - The InsightsEnabled boolean can be set to true to enable insights for the group or false to disable insights for the group.
	// - The NotificationsEnabled boolean can be set to true to enable insights notifications through Amazon EventBridge for the group.
	InsightsConfiguration *GroupInsightsConfiguration `pulumi:"insightsConfiguration"`
	// An array of key-value pairs to apply to this resource.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The filter expression defining criteria by which to group traces.
	FilterExpression pulumi.StringPtrInput
	// The case-sensitive name of the new group. Names must be unique.
	GroupName pulumi.StringPtrInput
	// The structure containing configurations related to insights.
	//
	// - The InsightsEnabled boolean can be set to true to enable insights for the group or false to disable insights for the group.
	// - The NotificationsEnabled boolean can be set to true to enable insights notifications through Amazon EventBridge for the group.
	InsightsConfiguration GroupInsightsConfigurationPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags aws.TagArrayInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}

type GroupInput interface {
	pulumi.Input

	ToGroupOutput() GroupOutput
	ToGroupOutputWithContext(ctx context.Context) GroupOutput
}

func (*Group) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (i *Group) ToGroupOutput() GroupOutput {
	return i.ToGroupOutputWithContext(context.Background())
}

func (i *Group) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupOutput)
}

type GroupOutput struct{ *pulumi.OutputState }

func (GroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Group)(nil)).Elem()
}

func (o GroupOutput) ToGroupOutput() GroupOutput {
	return o
}

func (o GroupOutput) ToGroupOutputWithContext(ctx context.Context) GroupOutput {
	return o
}

// The filter expression defining criteria by which to group traces.
func (o GroupOutput) FilterExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Group) pulumi.StringPtrOutput { return v.FilterExpression }).(pulumi.StringPtrOutput)
}

// The ARN of the group that was generated on creation.
func (o GroupOutput) GroupArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.GroupArn }).(pulumi.StringOutput)
}

// The case-sensitive name of the new group. Names must be unique.
func (o GroupOutput) GroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Group) pulumi.StringOutput { return v.GroupName }).(pulumi.StringOutput)
}

// The structure containing configurations related to insights.
//
// - The InsightsEnabled boolean can be set to true to enable insights for the group or false to disable insights for the group.
// - The NotificationsEnabled boolean can be set to true to enable insights notifications through Amazon EventBridge for the group.
func (o GroupOutput) InsightsConfiguration() GroupInsightsConfigurationPtrOutput {
	return o.ApplyT(func(v *Group) GroupInsightsConfigurationPtrOutput { return v.InsightsConfiguration }).(GroupInsightsConfigurationPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o GroupOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Group) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupInput)(nil)).Elem(), &Group{})
	pulumi.RegisterOutputType(GroupOutput{})
}
