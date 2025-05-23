// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourcegroups

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies the type of resource query that determines this group's membership. There are two valid query types:
//
// - `TAG_FILTERS_1_0` indicates that the group is a tag-based group. To complete the group membership, you must include the `TagFilters` property to specify the tag filters to use in the query.
// - `CLOUDFORMATION_STACK_1_0` , the default, indicates that the group is a CloudFormation stack-based group. Group membership is based on the CloudFormation stack. You must specify the `StackIdentifier` property in the query to define which stack to associate the group with, or leave it empty to default to the stack where the group is defined.
type GroupResourceQueryType string

const (
	GroupResourceQueryTypeTagFilters10          = GroupResourceQueryType("TAG_FILTERS_1_0")
	GroupResourceQueryTypeCloudformationStack10 = GroupResourceQueryType("CLOUDFORMATION_STACK_1_0")
)

func (GroupResourceQueryType) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupResourceQueryType)(nil)).Elem()
}

func (e GroupResourceQueryType) ToGroupResourceQueryTypeOutput() GroupResourceQueryTypeOutput {
	return pulumi.ToOutput(e).(GroupResourceQueryTypeOutput)
}

func (e GroupResourceQueryType) ToGroupResourceQueryTypeOutputWithContext(ctx context.Context) GroupResourceQueryTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GroupResourceQueryTypeOutput)
}

func (e GroupResourceQueryType) ToGroupResourceQueryTypePtrOutput() GroupResourceQueryTypePtrOutput {
	return e.ToGroupResourceQueryTypePtrOutputWithContext(context.Background())
}

func (e GroupResourceQueryType) ToGroupResourceQueryTypePtrOutputWithContext(ctx context.Context) GroupResourceQueryTypePtrOutput {
	return GroupResourceQueryType(e).ToGroupResourceQueryTypeOutputWithContext(ctx).ToGroupResourceQueryTypePtrOutputWithContext(ctx)
}

func (e GroupResourceQueryType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GroupResourceQueryType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GroupResourceQueryType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GroupResourceQueryType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GroupResourceQueryTypeOutput struct{ *pulumi.OutputState }

func (GroupResourceQueryTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupResourceQueryType)(nil)).Elem()
}

func (o GroupResourceQueryTypeOutput) ToGroupResourceQueryTypeOutput() GroupResourceQueryTypeOutput {
	return o
}

func (o GroupResourceQueryTypeOutput) ToGroupResourceQueryTypeOutputWithContext(ctx context.Context) GroupResourceQueryTypeOutput {
	return o
}

func (o GroupResourceQueryTypeOutput) ToGroupResourceQueryTypePtrOutput() GroupResourceQueryTypePtrOutput {
	return o.ToGroupResourceQueryTypePtrOutputWithContext(context.Background())
}

func (o GroupResourceQueryTypeOutput) ToGroupResourceQueryTypePtrOutputWithContext(ctx context.Context) GroupResourceQueryTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupResourceQueryType) *GroupResourceQueryType {
		return &v
	}).(GroupResourceQueryTypePtrOutput)
}

func (o GroupResourceQueryTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GroupResourceQueryTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GroupResourceQueryType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GroupResourceQueryTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GroupResourceQueryTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GroupResourceQueryType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GroupResourceQueryTypePtrOutput struct{ *pulumi.OutputState }

func (GroupResourceQueryTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupResourceQueryType)(nil)).Elem()
}

func (o GroupResourceQueryTypePtrOutput) ToGroupResourceQueryTypePtrOutput() GroupResourceQueryTypePtrOutput {
	return o
}

func (o GroupResourceQueryTypePtrOutput) ToGroupResourceQueryTypePtrOutputWithContext(ctx context.Context) GroupResourceQueryTypePtrOutput {
	return o
}

func (o GroupResourceQueryTypePtrOutput) Elem() GroupResourceQueryTypeOutput {
	return o.ApplyT(func(v *GroupResourceQueryType) GroupResourceQueryType {
		if v != nil {
			return *v
		}
		var ret GroupResourceQueryType
		return ret
	}).(GroupResourceQueryTypeOutput)
}

func (o GroupResourceQueryTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GroupResourceQueryTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GroupResourceQueryType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// GroupResourceQueryTypeInput is an input type that accepts values of the GroupResourceQueryType enum
// A concrete instance of `GroupResourceQueryTypeInput` can be one of the following:
//
//	GroupResourceQueryTypeTagFilters10
//	GroupResourceQueryTypeCloudformationStack10
type GroupResourceQueryTypeInput interface {
	pulumi.Input

	ToGroupResourceQueryTypeOutput() GroupResourceQueryTypeOutput
	ToGroupResourceQueryTypeOutputWithContext(context.Context) GroupResourceQueryTypeOutput
}

var groupResourceQueryTypePtrType = reflect.TypeOf((**GroupResourceQueryType)(nil)).Elem()

type GroupResourceQueryTypePtrInput interface {
	pulumi.Input

	ToGroupResourceQueryTypePtrOutput() GroupResourceQueryTypePtrOutput
	ToGroupResourceQueryTypePtrOutputWithContext(context.Context) GroupResourceQueryTypePtrOutput
}

type groupResourceQueryTypePtr string

func GroupResourceQueryTypePtr(v string) GroupResourceQueryTypePtrInput {
	return (*groupResourceQueryTypePtr)(&v)
}

func (*groupResourceQueryTypePtr) ElementType() reflect.Type {
	return groupResourceQueryTypePtrType
}

func (in *groupResourceQueryTypePtr) ToGroupResourceQueryTypePtrOutput() GroupResourceQueryTypePtrOutput {
	return pulumi.ToOutput(in).(GroupResourceQueryTypePtrOutput)
}

func (in *groupResourceQueryTypePtr) ToGroupResourceQueryTypePtrOutputWithContext(ctx context.Context) GroupResourceQueryTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GroupResourceQueryTypePtrOutput)
}

// The status of the TagSyncTask
type TagSyncTaskStatus string

const (
	TagSyncTaskStatusActive = TagSyncTaskStatus("ACTIVE")
	TagSyncTaskStatusError  = TagSyncTaskStatus("ERROR")
)

type TagSyncTaskStatusOutput struct{ *pulumi.OutputState }

func (TagSyncTaskStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagSyncTaskStatus)(nil)).Elem()
}

func (o TagSyncTaskStatusOutput) ToTagSyncTaskStatusOutput() TagSyncTaskStatusOutput {
	return o
}

func (o TagSyncTaskStatusOutput) ToTagSyncTaskStatusOutputWithContext(ctx context.Context) TagSyncTaskStatusOutput {
	return o
}

func (o TagSyncTaskStatusOutput) ToTagSyncTaskStatusPtrOutput() TagSyncTaskStatusPtrOutput {
	return o.ToTagSyncTaskStatusPtrOutputWithContext(context.Background())
}

func (o TagSyncTaskStatusOutput) ToTagSyncTaskStatusPtrOutputWithContext(ctx context.Context) TagSyncTaskStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagSyncTaskStatus) *TagSyncTaskStatus {
		return &v
	}).(TagSyncTaskStatusPtrOutput)
}

func (o TagSyncTaskStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TagSyncTaskStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TagSyncTaskStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TagSyncTaskStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TagSyncTaskStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TagSyncTaskStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TagSyncTaskStatusPtrOutput struct{ *pulumi.OutputState }

func (TagSyncTaskStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagSyncTaskStatus)(nil)).Elem()
}

func (o TagSyncTaskStatusPtrOutput) ToTagSyncTaskStatusPtrOutput() TagSyncTaskStatusPtrOutput {
	return o
}

func (o TagSyncTaskStatusPtrOutput) ToTagSyncTaskStatusPtrOutputWithContext(ctx context.Context) TagSyncTaskStatusPtrOutput {
	return o
}

func (o TagSyncTaskStatusPtrOutput) Elem() TagSyncTaskStatusOutput {
	return o.ApplyT(func(v *TagSyncTaskStatus) TagSyncTaskStatus {
		if v != nil {
			return *v
		}
		var ret TagSyncTaskStatus
		return ret
	}).(TagSyncTaskStatusOutput)
}

func (o TagSyncTaskStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TagSyncTaskStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TagSyncTaskStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupResourceQueryTypeInput)(nil)).Elem(), GroupResourceQueryType("TAG_FILTERS_1_0"))
	pulumi.RegisterInputType(reflect.TypeOf((*GroupResourceQueryTypePtrInput)(nil)).Elem(), GroupResourceQueryType("TAG_FILTERS_1_0"))
	pulumi.RegisterOutputType(GroupResourceQueryTypeOutput{})
	pulumi.RegisterOutputType(GroupResourceQueryTypePtrOutput{})
	pulumi.RegisterOutputType(TagSyncTaskStatusOutput{})
	pulumi.RegisterOutputType(TagSyncTaskStatusPtrOutput{})
}
