// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securitylake

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubscriberAccessTypesItem string

const (
	SubscriberAccessTypesItemLakeformation = SubscriberAccessTypesItem("LAKEFORMATION")
	SubscriberAccessTypesItemS3            = SubscriberAccessTypesItem("S3")
)

func (SubscriberAccessTypesItem) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriberAccessTypesItem)(nil)).Elem()
}

func (e SubscriberAccessTypesItem) ToSubscriberAccessTypesItemOutput() SubscriberAccessTypesItemOutput {
	return pulumi.ToOutput(e).(SubscriberAccessTypesItemOutput)
}

func (e SubscriberAccessTypesItem) ToSubscriberAccessTypesItemOutputWithContext(ctx context.Context) SubscriberAccessTypesItemOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SubscriberAccessTypesItemOutput)
}

func (e SubscriberAccessTypesItem) ToSubscriberAccessTypesItemPtrOutput() SubscriberAccessTypesItemPtrOutput {
	return e.ToSubscriberAccessTypesItemPtrOutputWithContext(context.Background())
}

func (e SubscriberAccessTypesItem) ToSubscriberAccessTypesItemPtrOutputWithContext(ctx context.Context) SubscriberAccessTypesItemPtrOutput {
	return SubscriberAccessTypesItem(e).ToSubscriberAccessTypesItemOutputWithContext(ctx).ToSubscriberAccessTypesItemPtrOutputWithContext(ctx)
}

func (e SubscriberAccessTypesItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriberAccessTypesItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriberAccessTypesItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SubscriberAccessTypesItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SubscriberAccessTypesItemOutput struct{ *pulumi.OutputState }

func (SubscriberAccessTypesItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriberAccessTypesItem)(nil)).Elem()
}

func (o SubscriberAccessTypesItemOutput) ToSubscriberAccessTypesItemOutput() SubscriberAccessTypesItemOutput {
	return o
}

func (o SubscriberAccessTypesItemOutput) ToSubscriberAccessTypesItemOutputWithContext(ctx context.Context) SubscriberAccessTypesItemOutput {
	return o
}

func (o SubscriberAccessTypesItemOutput) ToSubscriberAccessTypesItemPtrOutput() SubscriberAccessTypesItemPtrOutput {
	return o.ToSubscriberAccessTypesItemPtrOutputWithContext(context.Background())
}

func (o SubscriberAccessTypesItemOutput) ToSubscriberAccessTypesItemPtrOutputWithContext(ctx context.Context) SubscriberAccessTypesItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriberAccessTypesItem) *SubscriberAccessTypesItem {
		return &v
	}).(SubscriberAccessTypesItemPtrOutput)
}

func (o SubscriberAccessTypesItemOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SubscriberAccessTypesItemOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriberAccessTypesItem) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SubscriberAccessTypesItemOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriberAccessTypesItemOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriberAccessTypesItem) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SubscriberAccessTypesItemPtrOutput struct{ *pulumi.OutputState }

func (SubscriberAccessTypesItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriberAccessTypesItem)(nil)).Elem()
}

func (o SubscriberAccessTypesItemPtrOutput) ToSubscriberAccessTypesItemPtrOutput() SubscriberAccessTypesItemPtrOutput {
	return o
}

func (o SubscriberAccessTypesItemPtrOutput) ToSubscriberAccessTypesItemPtrOutputWithContext(ctx context.Context) SubscriberAccessTypesItemPtrOutput {
	return o
}

func (o SubscriberAccessTypesItemPtrOutput) Elem() SubscriberAccessTypesItemOutput {
	return o.ApplyT(func(v *SubscriberAccessTypesItem) SubscriberAccessTypesItem {
		if v != nil {
			return *v
		}
		var ret SubscriberAccessTypesItem
		return ret
	}).(SubscriberAccessTypesItemOutput)
}

func (o SubscriberAccessTypesItemPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriberAccessTypesItemPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SubscriberAccessTypesItem) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SubscriberAccessTypesItemInput is an input type that accepts values of the SubscriberAccessTypesItem enum
// A concrete instance of `SubscriberAccessTypesItemInput` can be one of the following:
//
//	SubscriberAccessTypesItemLakeformation
//	SubscriberAccessTypesItemS3
type SubscriberAccessTypesItemInput interface {
	pulumi.Input

	ToSubscriberAccessTypesItemOutput() SubscriberAccessTypesItemOutput
	ToSubscriberAccessTypesItemOutputWithContext(context.Context) SubscriberAccessTypesItemOutput
}

var subscriberAccessTypesItemPtrType = reflect.TypeOf((**SubscriberAccessTypesItem)(nil)).Elem()

type SubscriberAccessTypesItemPtrInput interface {
	pulumi.Input

	ToSubscriberAccessTypesItemPtrOutput() SubscriberAccessTypesItemPtrOutput
	ToSubscriberAccessTypesItemPtrOutputWithContext(context.Context) SubscriberAccessTypesItemPtrOutput
}

type subscriberAccessTypesItemPtr string

func SubscriberAccessTypesItemPtr(v string) SubscriberAccessTypesItemPtrInput {
	return (*subscriberAccessTypesItemPtr)(&v)
}

func (*subscriberAccessTypesItemPtr) ElementType() reflect.Type {
	return subscriberAccessTypesItemPtrType
}

func (in *subscriberAccessTypesItemPtr) ToSubscriberAccessTypesItemPtrOutput() SubscriberAccessTypesItemPtrOutput {
	return pulumi.ToOutput(in).(SubscriberAccessTypesItemPtrOutput)
}

func (in *subscriberAccessTypesItemPtr) ToSubscriberAccessTypesItemPtrOutputWithContext(ctx context.Context) SubscriberAccessTypesItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SubscriberAccessTypesItemPtrOutput)
}

// SubscriberAccessTypesItemArrayInput is an input type that accepts SubscriberAccessTypesItemArray and SubscriberAccessTypesItemArrayOutput values.
// You can construct a concrete instance of `SubscriberAccessTypesItemArrayInput` via:
//
//	SubscriberAccessTypesItemArray{ SubscriberAccessTypesItemArgs{...} }
type SubscriberAccessTypesItemArrayInput interface {
	pulumi.Input

	ToSubscriberAccessTypesItemArrayOutput() SubscriberAccessTypesItemArrayOutput
	ToSubscriberAccessTypesItemArrayOutputWithContext(context.Context) SubscriberAccessTypesItemArrayOutput
}

type SubscriberAccessTypesItemArray []SubscriberAccessTypesItem

func (SubscriberAccessTypesItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriberAccessTypesItem)(nil)).Elem()
}

func (i SubscriberAccessTypesItemArray) ToSubscriberAccessTypesItemArrayOutput() SubscriberAccessTypesItemArrayOutput {
	return i.ToSubscriberAccessTypesItemArrayOutputWithContext(context.Background())
}

func (i SubscriberAccessTypesItemArray) ToSubscriberAccessTypesItemArrayOutputWithContext(ctx context.Context) SubscriberAccessTypesItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriberAccessTypesItemArrayOutput)
}

type SubscriberAccessTypesItemArrayOutput struct{ *pulumi.OutputState }

func (SubscriberAccessTypesItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriberAccessTypesItem)(nil)).Elem()
}

func (o SubscriberAccessTypesItemArrayOutput) ToSubscriberAccessTypesItemArrayOutput() SubscriberAccessTypesItemArrayOutput {
	return o
}

func (o SubscriberAccessTypesItemArrayOutput) ToSubscriberAccessTypesItemArrayOutputWithContext(ctx context.Context) SubscriberAccessTypesItemArrayOutput {
	return o
}

func (o SubscriberAccessTypesItemArrayOutput) Index(i pulumi.IntInput) SubscriberAccessTypesItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriberAccessTypesItem {
		return vs[0].([]SubscriberAccessTypesItem)[vs[1].(int)]
	}).(SubscriberAccessTypesItemOutput)
}

// The HTTPS method used for the notification subscription.
type SubscriberNotificationHttpsNotificationConfigurationHttpMethod string

const (
	SubscriberNotificationHttpsNotificationConfigurationHttpMethodPost = SubscriberNotificationHttpsNotificationConfigurationHttpMethod("POST")
	SubscriberNotificationHttpsNotificationConfigurationHttpMethodPut  = SubscriberNotificationHttpsNotificationConfigurationHttpMethod("PUT")
)

func (SubscriberNotificationHttpsNotificationConfigurationHttpMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriberNotificationHttpsNotificationConfigurationHttpMethod)(nil)).Elem()
}

func (e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput() SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput {
	return pulumi.ToOutput(e).(SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput)
}

func (e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodOutputWithContext(ctx context.Context) SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput)
}

func (e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput() SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput {
	return e.ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutputWithContext(context.Background())
}

func (e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutputWithContext(ctx context.Context) SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput {
	return SubscriberNotificationHttpsNotificationConfigurationHttpMethod(e).ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodOutputWithContext(ctx).ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutputWithContext(ctx)
}

func (e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput struct{ *pulumi.OutputState }

func (SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriberNotificationHttpsNotificationConfigurationHttpMethod)(nil)).Elem()
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput() SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput {
	return o
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodOutputWithContext(ctx context.Context) SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput {
	return o
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput() SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput {
	return o.ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutputWithContext(context.Background())
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutputWithContext(ctx context.Context) SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriberNotificationHttpsNotificationConfigurationHttpMethod) *SubscriberNotificationHttpsNotificationConfigurationHttpMethod {
		return &v
	}).(SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput)
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriberNotificationHttpsNotificationConfigurationHttpMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput struct{ *pulumi.OutputState }

func (SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriberNotificationHttpsNotificationConfigurationHttpMethod)(nil)).Elem()
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput() SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput {
	return o
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutputWithContext(ctx context.Context) SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput {
	return o
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput) Elem() SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput {
	return o.ApplyT(func(v *SubscriberNotificationHttpsNotificationConfigurationHttpMethod) SubscriberNotificationHttpsNotificationConfigurationHttpMethod {
		if v != nil {
			return *v
		}
		var ret SubscriberNotificationHttpsNotificationConfigurationHttpMethod
		return ret
	}).(SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput)
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SubscriberNotificationHttpsNotificationConfigurationHttpMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SubscriberNotificationHttpsNotificationConfigurationHttpMethodInput is an input type that accepts values of the SubscriberNotificationHttpsNotificationConfigurationHttpMethod enum
// A concrete instance of `SubscriberNotificationHttpsNotificationConfigurationHttpMethodInput` can be one of the following:
//
//	SubscriberNotificationHttpsNotificationConfigurationHttpMethodPost
//	SubscriberNotificationHttpsNotificationConfigurationHttpMethodPut
type SubscriberNotificationHttpsNotificationConfigurationHttpMethodInput interface {
	pulumi.Input

	ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput() SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput
	ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodOutputWithContext(context.Context) SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput
}

var subscriberNotificationHttpsNotificationConfigurationHttpMethodPtrType = reflect.TypeOf((**SubscriberNotificationHttpsNotificationConfigurationHttpMethod)(nil)).Elem()

type SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrInput interface {
	pulumi.Input

	ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput() SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput
	ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutputWithContext(context.Context) SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput
}

type subscriberNotificationHttpsNotificationConfigurationHttpMethodPtr string

func SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtr(v string) SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrInput {
	return (*subscriberNotificationHttpsNotificationConfigurationHttpMethodPtr)(&v)
}

func (*subscriberNotificationHttpsNotificationConfigurationHttpMethodPtr) ElementType() reflect.Type {
	return subscriberNotificationHttpsNotificationConfigurationHttpMethodPtrType
}

func (in *subscriberNotificationHttpsNotificationConfigurationHttpMethodPtr) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput() SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput {
	return pulumi.ToOutput(in).(SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput)
}

func (in *subscriberNotificationHttpsNotificationConfigurationHttpMethodPtr) ToSubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutputWithContext(ctx context.Context) SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriberAccessTypesItemInput)(nil)).Elem(), SubscriberAccessTypesItem("LAKEFORMATION"))
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriberAccessTypesItemPtrInput)(nil)).Elem(), SubscriberAccessTypesItem("LAKEFORMATION"))
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriberAccessTypesItemArrayInput)(nil)).Elem(), SubscriberAccessTypesItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriberNotificationHttpsNotificationConfigurationHttpMethodInput)(nil)).Elem(), SubscriberNotificationHttpsNotificationConfigurationHttpMethod("POST"))
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrInput)(nil)).Elem(), SubscriberNotificationHttpsNotificationConfigurationHttpMethod("POST"))
	pulumi.RegisterOutputType(SubscriberAccessTypesItemOutput{})
	pulumi.RegisterOutputType(SubscriberAccessTypesItemPtrOutput{})
	pulumi.RegisterOutputType(SubscriberAccessTypesItemArrayOutput{})
	pulumi.RegisterOutputType(SubscriberNotificationHttpsNotificationConfigurationHttpMethodOutput{})
	pulumi.RegisterOutputType(SubscriberNotificationHttpsNotificationConfigurationHttpMethodPtrOutput{})
}