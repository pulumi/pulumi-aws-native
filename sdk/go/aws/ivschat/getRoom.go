// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivschat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::IVSChat::Room.
func LookupRoom(ctx *pulumi.Context, args *LookupRoomArgs, opts ...pulumi.InvokeOption) (*LookupRoomResult, error) {
	var rv LookupRoomResult
	err := ctx.Invoke("aws-native:ivschat:getRoom", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoomArgs struct {
	// Room ARN is automatically generated on creation and assigned as the unique identifier.
	Arn string `pulumi:"arn"`
}

type LookupRoomResult struct {
	// Room ARN is automatically generated on creation and assigned as the unique identifier.
	Arn *string `pulumi:"arn"`
	// The system-generated ID of the room.
	Id *string `pulumi:"id"`
	// Array of logging configuration identifiers attached to the room.
	LoggingConfigurationIdentifiers []string `pulumi:"loggingConfigurationIdentifiers"`
	// The maximum number of characters in a single message.
	MaximumMessageLength *int `pulumi:"maximumMessageLength"`
	// The maximum number of messages per second that can be sent to the room.
	MaximumMessageRatePerSecond *int                      `pulumi:"maximumMessageRatePerSecond"`
	MessageReviewHandler        *RoomMessageReviewHandler `pulumi:"messageReviewHandler"`
	// The name of the room. The value does not need to be unique.
	Name *string `pulumi:"name"`
	// An array of key-value pairs to apply to this resource.
	Tags []RoomTag `pulumi:"tags"`
}

func LookupRoomOutput(ctx *pulumi.Context, args LookupRoomOutputArgs, opts ...pulumi.InvokeOption) LookupRoomResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRoomResult, error) {
			args := v.(LookupRoomArgs)
			r, err := LookupRoom(ctx, &args, opts...)
			var s LookupRoomResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRoomResultOutput)
}

type LookupRoomOutputArgs struct {
	// Room ARN is automatically generated on creation and assigned as the unique identifier.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupRoomOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoomArgs)(nil)).Elem()
}

type LookupRoomResultOutput struct{ *pulumi.OutputState }

func (LookupRoomResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRoomResult)(nil)).Elem()
}

func (o LookupRoomResultOutput) ToLookupRoomResultOutput() LookupRoomResultOutput {
	return o
}

func (o LookupRoomResultOutput) ToLookupRoomResultOutputWithContext(ctx context.Context) LookupRoomResultOutput {
	return o
}

// Room ARN is automatically generated on creation and assigned as the unique identifier.
func (o LookupRoomResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoomResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The system-generated ID of the room.
func (o LookupRoomResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoomResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Array of logging configuration identifiers attached to the room.
func (o LookupRoomResultOutput) LoggingConfigurationIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRoomResult) []string { return v.LoggingConfigurationIdentifiers }).(pulumi.StringArrayOutput)
}

// The maximum number of characters in a single message.
func (o LookupRoomResultOutput) MaximumMessageLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRoomResult) *int { return v.MaximumMessageLength }).(pulumi.IntPtrOutput)
}

// The maximum number of messages per second that can be sent to the room.
func (o LookupRoomResultOutput) MaximumMessageRatePerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRoomResult) *int { return v.MaximumMessageRatePerSecond }).(pulumi.IntPtrOutput)
}

func (o LookupRoomResultOutput) MessageReviewHandler() RoomMessageReviewHandlerPtrOutput {
	return o.ApplyT(func(v LookupRoomResult) *RoomMessageReviewHandler { return v.MessageReviewHandler }).(RoomMessageReviewHandlerPtrOutput)
}

// The name of the room. The value does not need to be unique.
func (o LookupRoomResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRoomResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupRoomResultOutput) Tags() RoomTagArrayOutput {
	return o.ApplyT(func(v LookupRoomResult) []RoomTag { return v.Tags }).(RoomTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRoomResultOutput{})
}