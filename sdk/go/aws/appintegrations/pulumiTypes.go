// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appintegrations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataIntegrationScheduleConfig struct {
	// The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.
	FirstExecutionFrom string `pulumi:"firstExecutionFrom"`
	// The name of the object to pull from the data source.
	Object string `pulumi:"object"`
	// How often the data should be pulled from data source.
	ScheduleExpression string `pulumi:"scheduleExpression"`
}

// DataIntegrationScheduleConfigInput is an input type that accepts DataIntegrationScheduleConfigArgs and DataIntegrationScheduleConfigOutput values.
// You can construct a concrete instance of `DataIntegrationScheduleConfigInput` via:
//
//	DataIntegrationScheduleConfigArgs{...}
type DataIntegrationScheduleConfigInput interface {
	pulumi.Input

	ToDataIntegrationScheduleConfigOutput() DataIntegrationScheduleConfigOutput
	ToDataIntegrationScheduleConfigOutputWithContext(context.Context) DataIntegrationScheduleConfigOutput
}

type DataIntegrationScheduleConfigArgs struct {
	// The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.
	FirstExecutionFrom pulumi.StringInput `pulumi:"firstExecutionFrom"`
	// The name of the object to pull from the data source.
	Object pulumi.StringInput `pulumi:"object"`
	// How often the data should be pulled from data source.
	ScheduleExpression pulumi.StringInput `pulumi:"scheduleExpression"`
}

func (DataIntegrationScheduleConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataIntegrationScheduleConfig)(nil)).Elem()
}

func (i DataIntegrationScheduleConfigArgs) ToDataIntegrationScheduleConfigOutput() DataIntegrationScheduleConfigOutput {
	return i.ToDataIntegrationScheduleConfigOutputWithContext(context.Background())
}

func (i DataIntegrationScheduleConfigArgs) ToDataIntegrationScheduleConfigOutputWithContext(ctx context.Context) DataIntegrationScheduleConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataIntegrationScheduleConfigOutput)
}

type DataIntegrationScheduleConfigOutput struct{ *pulumi.OutputState }

func (DataIntegrationScheduleConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataIntegrationScheduleConfig)(nil)).Elem()
}

func (o DataIntegrationScheduleConfigOutput) ToDataIntegrationScheduleConfigOutput() DataIntegrationScheduleConfigOutput {
	return o
}

func (o DataIntegrationScheduleConfigOutput) ToDataIntegrationScheduleConfigOutputWithContext(ctx context.Context) DataIntegrationScheduleConfigOutput {
	return o
}

// The start date for objects to import in the first flow run. Epoch or ISO timestamp format is supported.
func (o DataIntegrationScheduleConfigOutput) FirstExecutionFrom() pulumi.StringOutput {
	return o.ApplyT(func(v DataIntegrationScheduleConfig) string { return v.FirstExecutionFrom }).(pulumi.StringOutput)
}

// The name of the object to pull from the data source.
func (o DataIntegrationScheduleConfigOutput) Object() pulumi.StringOutput {
	return o.ApplyT(func(v DataIntegrationScheduleConfig) string { return v.Object }).(pulumi.StringOutput)
}

// How often the data should be pulled from data source.
func (o DataIntegrationScheduleConfigOutput) ScheduleExpression() pulumi.StringOutput {
	return o.ApplyT(func(v DataIntegrationScheduleConfig) string { return v.ScheduleExpression }).(pulumi.StringOutput)
}

// A label for tagging DataIntegration resources
type DataIntegrationTag struct {
	// A key to identify the tag.
	Key string `pulumi:"key"`
	// Corresponding tag value for the key.
	Value string `pulumi:"value"`
}

// DataIntegrationTagInput is an input type that accepts DataIntegrationTagArgs and DataIntegrationTagOutput values.
// You can construct a concrete instance of `DataIntegrationTagInput` via:
//
//	DataIntegrationTagArgs{...}
type DataIntegrationTagInput interface {
	pulumi.Input

	ToDataIntegrationTagOutput() DataIntegrationTagOutput
	ToDataIntegrationTagOutputWithContext(context.Context) DataIntegrationTagOutput
}

// A label for tagging DataIntegration resources
type DataIntegrationTagArgs struct {
	// A key to identify the tag.
	Key pulumi.StringInput `pulumi:"key"`
	// Corresponding tag value for the key.
	Value pulumi.StringInput `pulumi:"value"`
}

func (DataIntegrationTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataIntegrationTag)(nil)).Elem()
}

func (i DataIntegrationTagArgs) ToDataIntegrationTagOutput() DataIntegrationTagOutput {
	return i.ToDataIntegrationTagOutputWithContext(context.Background())
}

func (i DataIntegrationTagArgs) ToDataIntegrationTagOutputWithContext(ctx context.Context) DataIntegrationTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataIntegrationTagOutput)
}

// DataIntegrationTagArrayInput is an input type that accepts DataIntegrationTagArray and DataIntegrationTagArrayOutput values.
// You can construct a concrete instance of `DataIntegrationTagArrayInput` via:
//
//	DataIntegrationTagArray{ DataIntegrationTagArgs{...} }
type DataIntegrationTagArrayInput interface {
	pulumi.Input

	ToDataIntegrationTagArrayOutput() DataIntegrationTagArrayOutput
	ToDataIntegrationTagArrayOutputWithContext(context.Context) DataIntegrationTagArrayOutput
}

type DataIntegrationTagArray []DataIntegrationTagInput

func (DataIntegrationTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataIntegrationTag)(nil)).Elem()
}

func (i DataIntegrationTagArray) ToDataIntegrationTagArrayOutput() DataIntegrationTagArrayOutput {
	return i.ToDataIntegrationTagArrayOutputWithContext(context.Background())
}

func (i DataIntegrationTagArray) ToDataIntegrationTagArrayOutputWithContext(ctx context.Context) DataIntegrationTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataIntegrationTagArrayOutput)
}

// A label for tagging DataIntegration resources
type DataIntegrationTagOutput struct{ *pulumi.OutputState }

func (DataIntegrationTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataIntegrationTag)(nil)).Elem()
}

func (o DataIntegrationTagOutput) ToDataIntegrationTagOutput() DataIntegrationTagOutput {
	return o
}

func (o DataIntegrationTagOutput) ToDataIntegrationTagOutputWithContext(ctx context.Context) DataIntegrationTagOutput {
	return o
}

// A key to identify the tag.
func (o DataIntegrationTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DataIntegrationTag) string { return v.Key }).(pulumi.StringOutput)
}

// Corresponding tag value for the key.
func (o DataIntegrationTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v DataIntegrationTag) string { return v.Value }).(pulumi.StringOutput)
}

type DataIntegrationTagArrayOutput struct{ *pulumi.OutputState }

func (DataIntegrationTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataIntegrationTag)(nil)).Elem()
}

func (o DataIntegrationTagArrayOutput) ToDataIntegrationTagArrayOutput() DataIntegrationTagArrayOutput {
	return o
}

func (o DataIntegrationTagArrayOutput) ToDataIntegrationTagArrayOutputWithContext(ctx context.Context) DataIntegrationTagArrayOutput {
	return o
}

func (o DataIntegrationTagArrayOutput) Index(i pulumi.IntInput) DataIntegrationTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataIntegrationTag {
		return vs[0].([]DataIntegrationTag)[vs[1].(int)]
	}).(DataIntegrationTagOutput)
}

type EventIntegrationAssociation struct {
	// The metadata associated with the client.
	ClientAssociationMetadata []EventIntegrationMetadata `pulumi:"clientAssociationMetadata"`
	// The identifier for the client that is associated with the event integration.
	ClientId *string `pulumi:"clientId"`
	// The name of the Eventbridge rule.
	EventBridgeRuleName *string `pulumi:"eventBridgeRuleName"`
	// The Amazon Resource Name (ARN) for the event integration association.
	EventIntegrationAssociationArn *string `pulumi:"eventIntegrationAssociationArn"`
	// The identifier for the event integration association.
	EventIntegrationAssociationId *string `pulumi:"eventIntegrationAssociationId"`
}

type EventIntegrationAssociationOutput struct{ *pulumi.OutputState }

func (EventIntegrationAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventIntegrationAssociation)(nil)).Elem()
}

func (o EventIntegrationAssociationOutput) ToEventIntegrationAssociationOutput() EventIntegrationAssociationOutput {
	return o
}

func (o EventIntegrationAssociationOutput) ToEventIntegrationAssociationOutputWithContext(ctx context.Context) EventIntegrationAssociationOutput {
	return o
}

// The metadata associated with the client.
func (o EventIntegrationAssociationOutput) ClientAssociationMetadata() EventIntegrationMetadataArrayOutput {
	return o.ApplyT(func(v EventIntegrationAssociation) []EventIntegrationMetadata { return v.ClientAssociationMetadata }).(EventIntegrationMetadataArrayOutput)
}

// The identifier for the client that is associated with the event integration.
func (o EventIntegrationAssociationOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventIntegrationAssociation) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The name of the Eventbridge rule.
func (o EventIntegrationAssociationOutput) EventBridgeRuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventIntegrationAssociation) *string { return v.EventBridgeRuleName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) for the event integration association.
func (o EventIntegrationAssociationOutput) EventIntegrationAssociationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventIntegrationAssociation) *string { return v.EventIntegrationAssociationArn }).(pulumi.StringPtrOutput)
}

// The identifier for the event integration association.
func (o EventIntegrationAssociationOutput) EventIntegrationAssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventIntegrationAssociation) *string { return v.EventIntegrationAssociationId }).(pulumi.StringPtrOutput)
}

type EventIntegrationAssociationArrayOutput struct{ *pulumi.OutputState }

func (EventIntegrationAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventIntegrationAssociation)(nil)).Elem()
}

func (o EventIntegrationAssociationArrayOutput) ToEventIntegrationAssociationArrayOutput() EventIntegrationAssociationArrayOutput {
	return o
}

func (o EventIntegrationAssociationArrayOutput) ToEventIntegrationAssociationArrayOutputWithContext(ctx context.Context) EventIntegrationAssociationArrayOutput {
	return o
}

func (o EventIntegrationAssociationArrayOutput) Index(i pulumi.IntInput) EventIntegrationAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventIntegrationAssociation {
		return vs[0].([]EventIntegrationAssociation)[vs[1].(int)]
	}).(EventIntegrationAssociationOutput)
}

type EventIntegrationEventFilter struct {
	// The source of the events.
	Source string `pulumi:"source"`
}

// EventIntegrationEventFilterInput is an input type that accepts EventIntegrationEventFilterArgs and EventIntegrationEventFilterOutput values.
// You can construct a concrete instance of `EventIntegrationEventFilterInput` via:
//
//	EventIntegrationEventFilterArgs{...}
type EventIntegrationEventFilterInput interface {
	pulumi.Input

	ToEventIntegrationEventFilterOutput() EventIntegrationEventFilterOutput
	ToEventIntegrationEventFilterOutputWithContext(context.Context) EventIntegrationEventFilterOutput
}

type EventIntegrationEventFilterArgs struct {
	// The source of the events.
	Source pulumi.StringInput `pulumi:"source"`
}

func (EventIntegrationEventFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventIntegrationEventFilter)(nil)).Elem()
}

func (i EventIntegrationEventFilterArgs) ToEventIntegrationEventFilterOutput() EventIntegrationEventFilterOutput {
	return i.ToEventIntegrationEventFilterOutputWithContext(context.Background())
}

func (i EventIntegrationEventFilterArgs) ToEventIntegrationEventFilterOutputWithContext(ctx context.Context) EventIntegrationEventFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventIntegrationEventFilterOutput)
}

type EventIntegrationEventFilterOutput struct{ *pulumi.OutputState }

func (EventIntegrationEventFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventIntegrationEventFilter)(nil)).Elem()
}

func (o EventIntegrationEventFilterOutput) ToEventIntegrationEventFilterOutput() EventIntegrationEventFilterOutput {
	return o
}

func (o EventIntegrationEventFilterOutput) ToEventIntegrationEventFilterOutputWithContext(ctx context.Context) EventIntegrationEventFilterOutput {
	return o
}

// The source of the events.
func (o EventIntegrationEventFilterOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v EventIntegrationEventFilter) string { return v.Source }).(pulumi.StringOutput)
}

type EventIntegrationMetadata struct {
	// A key to identify the metadata.
	Key string `pulumi:"key"`
	// Corresponding metadata value for the key.
	Value string `pulumi:"value"`
}

type EventIntegrationMetadataOutput struct{ *pulumi.OutputState }

func (EventIntegrationMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventIntegrationMetadata)(nil)).Elem()
}

func (o EventIntegrationMetadataOutput) ToEventIntegrationMetadataOutput() EventIntegrationMetadataOutput {
	return o
}

func (o EventIntegrationMetadataOutput) ToEventIntegrationMetadataOutputWithContext(ctx context.Context) EventIntegrationMetadataOutput {
	return o
}

// A key to identify the metadata.
func (o EventIntegrationMetadataOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v EventIntegrationMetadata) string { return v.Key }).(pulumi.StringOutput)
}

// Corresponding metadata value for the key.
func (o EventIntegrationMetadataOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v EventIntegrationMetadata) string { return v.Value }).(pulumi.StringOutput)
}

type EventIntegrationMetadataArrayOutput struct{ *pulumi.OutputState }

func (EventIntegrationMetadataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventIntegrationMetadata)(nil)).Elem()
}

func (o EventIntegrationMetadataArrayOutput) ToEventIntegrationMetadataArrayOutput() EventIntegrationMetadataArrayOutput {
	return o
}

func (o EventIntegrationMetadataArrayOutput) ToEventIntegrationMetadataArrayOutputWithContext(ctx context.Context) EventIntegrationMetadataArrayOutput {
	return o
}

func (o EventIntegrationMetadataArrayOutput) Index(i pulumi.IntInput) EventIntegrationMetadataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventIntegrationMetadata {
		return vs[0].([]EventIntegrationMetadata)[vs[1].(int)]
	}).(EventIntegrationMetadataOutput)
}

type EventIntegrationTag struct {
	// A key to identify the tag.
	Key string `pulumi:"key"`
	// Corresponding tag value for the key.
	Value string `pulumi:"value"`
}

// EventIntegrationTagInput is an input type that accepts EventIntegrationTagArgs and EventIntegrationTagOutput values.
// You can construct a concrete instance of `EventIntegrationTagInput` via:
//
//	EventIntegrationTagArgs{...}
type EventIntegrationTagInput interface {
	pulumi.Input

	ToEventIntegrationTagOutput() EventIntegrationTagOutput
	ToEventIntegrationTagOutputWithContext(context.Context) EventIntegrationTagOutput
}

type EventIntegrationTagArgs struct {
	// A key to identify the tag.
	Key pulumi.StringInput `pulumi:"key"`
	// Corresponding tag value for the key.
	Value pulumi.StringInput `pulumi:"value"`
}

func (EventIntegrationTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventIntegrationTag)(nil)).Elem()
}

func (i EventIntegrationTagArgs) ToEventIntegrationTagOutput() EventIntegrationTagOutput {
	return i.ToEventIntegrationTagOutputWithContext(context.Background())
}

func (i EventIntegrationTagArgs) ToEventIntegrationTagOutputWithContext(ctx context.Context) EventIntegrationTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventIntegrationTagOutput)
}

// EventIntegrationTagArrayInput is an input type that accepts EventIntegrationTagArray and EventIntegrationTagArrayOutput values.
// You can construct a concrete instance of `EventIntegrationTagArrayInput` via:
//
//	EventIntegrationTagArray{ EventIntegrationTagArgs{...} }
type EventIntegrationTagArrayInput interface {
	pulumi.Input

	ToEventIntegrationTagArrayOutput() EventIntegrationTagArrayOutput
	ToEventIntegrationTagArrayOutputWithContext(context.Context) EventIntegrationTagArrayOutput
}

type EventIntegrationTagArray []EventIntegrationTagInput

func (EventIntegrationTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventIntegrationTag)(nil)).Elem()
}

func (i EventIntegrationTagArray) ToEventIntegrationTagArrayOutput() EventIntegrationTagArrayOutput {
	return i.ToEventIntegrationTagArrayOutputWithContext(context.Background())
}

func (i EventIntegrationTagArray) ToEventIntegrationTagArrayOutputWithContext(ctx context.Context) EventIntegrationTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventIntegrationTagArrayOutput)
}

type EventIntegrationTagOutput struct{ *pulumi.OutputState }

func (EventIntegrationTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventIntegrationTag)(nil)).Elem()
}

func (o EventIntegrationTagOutput) ToEventIntegrationTagOutput() EventIntegrationTagOutput {
	return o
}

func (o EventIntegrationTagOutput) ToEventIntegrationTagOutputWithContext(ctx context.Context) EventIntegrationTagOutput {
	return o
}

// A key to identify the tag.
func (o EventIntegrationTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v EventIntegrationTag) string { return v.Key }).(pulumi.StringOutput)
}

// Corresponding tag value for the key.
func (o EventIntegrationTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v EventIntegrationTag) string { return v.Value }).(pulumi.StringOutput)
}

type EventIntegrationTagArrayOutput struct{ *pulumi.OutputState }

func (EventIntegrationTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventIntegrationTag)(nil)).Elem()
}

func (o EventIntegrationTagArrayOutput) ToEventIntegrationTagArrayOutput() EventIntegrationTagArrayOutput {
	return o
}

func (o EventIntegrationTagArrayOutput) ToEventIntegrationTagArrayOutputWithContext(ctx context.Context) EventIntegrationTagArrayOutput {
	return o
}

func (o EventIntegrationTagArrayOutput) Index(i pulumi.IntInput) EventIntegrationTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventIntegrationTag {
		return vs[0].([]EventIntegrationTag)[vs[1].(int)]
	}).(EventIntegrationTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataIntegrationScheduleConfigInput)(nil)).Elem(), DataIntegrationScheduleConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataIntegrationTagInput)(nil)).Elem(), DataIntegrationTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataIntegrationTagArrayInput)(nil)).Elem(), DataIntegrationTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventIntegrationEventFilterInput)(nil)).Elem(), EventIntegrationEventFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventIntegrationTagInput)(nil)).Elem(), EventIntegrationTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventIntegrationTagArrayInput)(nil)).Elem(), EventIntegrationTagArray{})
	pulumi.RegisterOutputType(DataIntegrationScheduleConfigOutput{})
	pulumi.RegisterOutputType(DataIntegrationTagOutput{})
	pulumi.RegisterOutputType(DataIntegrationTagArrayOutput{})
	pulumi.RegisterOutputType(EventIntegrationAssociationOutput{})
	pulumi.RegisterOutputType(EventIntegrationAssociationArrayOutput{})
	pulumi.RegisterOutputType(EventIntegrationEventFilterOutput{})
	pulumi.RegisterOutputType(EventIntegrationMetadataOutput{})
	pulumi.RegisterOutputType(EventIntegrationMetadataArrayOutput{})
	pulumi.RegisterOutputType(EventIntegrationTagOutput{})
	pulumi.RegisterOutputType(EventIntegrationTagArrayOutput{})
}