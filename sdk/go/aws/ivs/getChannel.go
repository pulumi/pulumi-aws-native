// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IVS::Channel
func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	var rv LookupChannelResult
	err := ctx.Invoke("aws-native:ivs:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelArgs struct {
	// Channel ARN is automatically generated on creation and assigned as the unique identifier.
	Arn string `pulumi:"arn"`
}

type LookupChannelResult struct {
	// Channel ARN is automatically generated on creation and assigned as the unique identifier.
	Arn *string `pulumi:"arn"`
	// Whether the channel is authorized.
	Authorized *bool `pulumi:"authorized"`
	// Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.
	IngestEndpoint *string `pulumi:"ingestEndpoint"`
	// Channel latency mode.
	LatencyMode *ChannelLatencyMode `pulumi:"latencyMode"`
	// Channel
	Name *string `pulumi:"name"`
	// Channel Playback URL.
	PlaybackUrl *string `pulumi:"playbackUrl"`
	// Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: “” (recording is disabled).
	RecordingConfigurationArn *string `pulumi:"recordingConfigurationArn"`
	// A list of key-value pairs that contain metadata for the asset model.
	Tags []ChannelTag `pulumi:"tags"`
	// Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
	Type *ChannelType `pulumi:"type"`
}

func LookupChannelOutput(ctx *pulumi.Context, args LookupChannelOutputArgs, opts ...pulumi.InvokeOption) LookupChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelResult, error) {
			args := v.(LookupChannelArgs)
			r, err := LookupChannel(ctx, &args, opts...)
			var s LookupChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelResultOutput)
}

type LookupChannelOutputArgs struct {
	// Channel ARN is automatically generated on creation and assigned as the unique identifier.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelArgs)(nil)).Elem()
}

type LookupChannelResultOutput struct{ *pulumi.OutputState }

func (LookupChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelResult)(nil)).Elem()
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutput() LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutputWithContext(ctx context.Context) LookupChannelResultOutput {
	return o
}

// Channel ARN is automatically generated on creation and assigned as the unique identifier.
func (o LookupChannelResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Whether the channel is authorized.
func (o LookupChannelResultOutput) Authorized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *bool { return v.Authorized }).(pulumi.BoolPtrOutput)
}

// Channel ingest endpoint, part of the definition of an ingest server, used when you set up streaming software.
func (o LookupChannelResultOutput) IngestEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.IngestEndpoint }).(pulumi.StringPtrOutput)
}

// Channel latency mode.
func (o LookupChannelResultOutput) LatencyMode() ChannelLatencyModePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelLatencyMode { return v.LatencyMode }).(ChannelLatencyModePtrOutput)
}

// Channel
func (o LookupChannelResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Channel Playback URL.
func (o LookupChannelResultOutput) PlaybackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.PlaybackUrl }).(pulumi.StringPtrOutput)
}

// Recording Configuration ARN. A value other than an empty string indicates that recording is enabled. Default: “” (recording is disabled).
func (o LookupChannelResultOutput) RecordingConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.RecordingConfigurationArn }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs that contain metadata for the asset model.
func (o LookupChannelResultOutput) Tags() ChannelTagArrayOutput {
	return o.ApplyT(func(v LookupChannelResult) []ChannelTag { return v.Tags }).(ChannelTagArrayOutput)
}

// Channel type, which determines the allowable resolution and bitrate. If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
func (o LookupChannelResultOutput) Type() ChannelTypePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelType { return v.Type }).(ChannelTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelResultOutput{})
}