// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package customerprofiles

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Event Stream resource of Amazon Connect Customer Profiles
func LookupEventStream(ctx *pulumi.Context, args *LookupEventStreamArgs, opts ...pulumi.InvokeOption) (*LookupEventStreamResult, error) {
	var rv LookupEventStreamResult
	err := ctx.Invoke("aws-native:customerprofiles:getEventStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventStreamArgs struct {
	// The unique name of the domain.
	DomainName string `pulumi:"domainName"`
	// The name of the event stream.
	EventStreamName string `pulumi:"eventStreamName"`
}

type LookupEventStreamResult struct {
	// The timestamp of when the export was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Details regarding the Kinesis stream.
	DestinationDetails *DestinationDetailsProperties `pulumi:"destinationDetails"`
	// A unique identifier for the event stream.
	EventStreamArn *string `pulumi:"eventStreamArn"`
	// The operational state of destination stream for export.
	State *EventStreamStateEnum `pulumi:"state"`
	// The tags used to organize, track, or control access for this resource.
	Tags []EventStreamTag `pulumi:"tags"`
}

func LookupEventStreamOutput(ctx *pulumi.Context, args LookupEventStreamOutputArgs, opts ...pulumi.InvokeOption) LookupEventStreamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventStreamResult, error) {
			args := v.(LookupEventStreamArgs)
			r, err := LookupEventStream(ctx, &args, opts...)
			var s LookupEventStreamResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEventStreamResultOutput)
}

type LookupEventStreamOutputArgs struct {
	// The unique name of the domain.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The name of the event stream.
	EventStreamName pulumi.StringInput `pulumi:"eventStreamName"`
}

func (LookupEventStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventStreamArgs)(nil)).Elem()
}

type LookupEventStreamResultOutput struct{ *pulumi.OutputState }

func (LookupEventStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventStreamResult)(nil)).Elem()
}

func (o LookupEventStreamResultOutput) ToLookupEventStreamResultOutput() LookupEventStreamResultOutput {
	return o
}

func (o LookupEventStreamResultOutput) ToLookupEventStreamResultOutputWithContext(ctx context.Context) LookupEventStreamResultOutput {
	return o
}

// The timestamp of when the export was created.
func (o LookupEventStreamResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventStreamResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// Details regarding the Kinesis stream.
func (o LookupEventStreamResultOutput) DestinationDetails() DestinationDetailsPropertiesPtrOutput {
	return o.ApplyT(func(v LookupEventStreamResult) *DestinationDetailsProperties { return v.DestinationDetails }).(DestinationDetailsPropertiesPtrOutput)
}

// A unique identifier for the event stream.
func (o LookupEventStreamResultOutput) EventStreamArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventStreamResult) *string { return v.EventStreamArn }).(pulumi.StringPtrOutput)
}

// The operational state of destination stream for export.
func (o LookupEventStreamResultOutput) State() EventStreamStateEnumPtrOutput {
	return o.ApplyT(func(v LookupEventStreamResult) *EventStreamStateEnum { return v.State }).(EventStreamStateEnumPtrOutput)
}

// The tags used to organize, track, or control access for this resource.
func (o LookupEventStreamResultOutput) Tags() EventStreamTagArrayOutput {
	return o.ApplyT(func(v LookupEventStreamResult) []EventStreamTag { return v.Tags }).(EventStreamTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventStreamResultOutput{})
}