// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediatailor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::MediaTailor::Channel Resource Type
func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupChannelResult
	err := ctx.Invoke("aws-native:mediatailor:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelArgs struct {
	// The name of the channel.
	ChannelName string `pulumi:"channelName"`
}

type LookupChannelResult struct {
	// <p>The ARN of the channel.</p>
	Arn *string `pulumi:"arn"`
	// <p>The list of audiences defined in channel.</p>
	Audiences []string `pulumi:"audiences"`
	// The slate used to fill gaps between programs in the schedule. You must configure filler slate if your channel uses the `LINEAR` `PlaybackMode` . MediaTailor doesn't support filler slate for channels using the `LOOP` `PlaybackMode` .
	FillerSlate *ChannelSlateSource `pulumi:"fillerSlate"`
	// The log configuration.
	LogConfiguration *ChannelLogConfigurationForChannel `pulumi:"logConfiguration"`
	// <p>The channel's output properties.</p>
	Outputs []ChannelRequestOutputItem `pulumi:"outputs"`
	// The type of playback mode for this channel.
	//
	// `LINEAR` - Programs play back-to-back only once.
	//
	// `LOOP` - Programs play back-to-back in an endless loop. When the last program in the schedule plays, playback loops back to the first program in the schedule.
	PlaybackMode *ChannelPlaybackMode `pulumi:"playbackMode"`
	// The tags to assign to the channel.
	Tags []aws.Tag `pulumi:"tags"`
	// The configuration for time-shifted viewing.
	TimeShiftConfiguration *ChannelTimeShiftConfiguration `pulumi:"timeShiftConfiguration"`
}

func LookupChannelOutput(ctx *pulumi.Context, args LookupChannelOutputArgs, opts ...pulumi.InvokeOption) LookupChannelResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupChannelResultOutput, error) {
			args := v.(LookupChannelArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:mediatailor:getChannel", args, LookupChannelResultOutput{}, options).(LookupChannelResultOutput), nil
		}).(LookupChannelResultOutput)
}

type LookupChannelOutputArgs struct {
	// The name of the channel.
	ChannelName pulumi.StringInput `pulumi:"channelName"`
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

// <p>The ARN of the channel.</p>
func (o LookupChannelResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// <p>The list of audiences defined in channel.</p>
func (o LookupChannelResultOutput) Audiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupChannelResult) []string { return v.Audiences }).(pulumi.StringArrayOutput)
}

// The slate used to fill gaps between programs in the schedule. You must configure filler slate if your channel uses the `LINEAR` `PlaybackMode` . MediaTailor doesn't support filler slate for channels using the `LOOP` `PlaybackMode` .
func (o LookupChannelResultOutput) FillerSlate() ChannelSlateSourcePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelSlateSource { return v.FillerSlate }).(ChannelSlateSourcePtrOutput)
}

// The log configuration.
func (o LookupChannelResultOutput) LogConfiguration() ChannelLogConfigurationForChannelPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelLogConfigurationForChannel { return v.LogConfiguration }).(ChannelLogConfigurationForChannelPtrOutput)
}

// <p>The channel's output properties.</p>
func (o LookupChannelResultOutput) Outputs() ChannelRequestOutputItemArrayOutput {
	return o.ApplyT(func(v LookupChannelResult) []ChannelRequestOutputItem { return v.Outputs }).(ChannelRequestOutputItemArrayOutput)
}

// The type of playback mode for this channel.
//
// `LINEAR` - Programs play back-to-back only once.
//
// `LOOP` - Programs play back-to-back in an endless loop. When the last program in the schedule plays, playback loops back to the first program in the schedule.
func (o LookupChannelResultOutput) PlaybackMode() ChannelPlaybackModePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelPlaybackMode { return v.PlaybackMode }).(ChannelPlaybackModePtrOutput)
}

// The tags to assign to the channel.
func (o LookupChannelResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupChannelResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The configuration for time-shifted viewing.
func (o LookupChannelResultOutput) TimeShiftConfiguration() ChannelTimeShiftConfigurationPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelTimeShiftConfiguration { return v.TimeShiftConfiguration }).(ChannelTimeShiftConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelResultOutput{})
}
