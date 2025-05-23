// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Kinesis::Stream
func LookupStream(ctx *pulumi.Context, args *LookupStreamArgs, opts ...pulumi.InvokeOption) (*LookupStreamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStreamResult
	err := ctx.Invoke("aws-native:kinesis:getStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStreamArgs struct {
	// The name of the Kinesis stream.
	Name string `pulumi:"name"`
}

type LookupStreamResult struct {
	// The Amazon resource name (ARN) of the Kinesis stream
	Arn *string `pulumi:"arn"`
	// The final list of shard-level metrics
	DesiredShardLevelMetrics []StreamEnhancedMetric `pulumi:"desiredShardLevelMetrics"`
	// The number of hours for the data records that are stored in shards to remain accessible.
	RetentionPeriodHours *int `pulumi:"retentionPeriodHours"`
	// The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
	ShardCount *int `pulumi:"shardCount"`
	// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
	StreamEncryption *StreamEncryption `pulumi:"streamEncryption"`
	// The mode in which the stream is running.
	StreamModeDetails *StreamModeDetails `pulumi:"streamModeDetails"`
	// An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupStreamOutput(ctx *pulumi.Context, args LookupStreamOutputArgs, opts ...pulumi.InvokeOption) LookupStreamResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStreamResultOutput, error) {
			args := v.(LookupStreamArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:kinesis:getStream", args, LookupStreamResultOutput{}, options).(LookupStreamResultOutput), nil
		}).(LookupStreamResultOutput)
}

type LookupStreamOutputArgs struct {
	// The name of the Kinesis stream.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamArgs)(nil)).Elem()
}

type LookupStreamResultOutput struct{ *pulumi.OutputState }

func (LookupStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamResult)(nil)).Elem()
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutput() LookupStreamResultOutput {
	return o
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutputWithContext(ctx context.Context) LookupStreamResultOutput {
	return o
}

// The Amazon resource name (ARN) of the Kinesis stream
func (o LookupStreamResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The final list of shard-level metrics
func (o LookupStreamResultOutput) DesiredShardLevelMetrics() StreamEnhancedMetricArrayOutput {
	return o.ApplyT(func(v LookupStreamResult) []StreamEnhancedMetric { return v.DesiredShardLevelMetrics }).(StreamEnhancedMetricArrayOutput)
}

// The number of hours for the data records that are stored in shards to remain accessible.
func (o LookupStreamResultOutput) RetentionPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *int { return v.RetentionPeriodHours }).(pulumi.IntPtrOutput)
}

// The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
func (o LookupStreamResultOutput) ShardCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *int { return v.ShardCount }).(pulumi.IntPtrOutput)
}

// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
func (o LookupStreamResultOutput) StreamEncryption() StreamEncryptionPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *StreamEncryption { return v.StreamEncryption }).(StreamEncryptionPtrOutput)
}

// The mode in which the stream is running.
func (o LookupStreamResultOutput) StreamModeDetails() StreamModeDetailsPtrOutput {
	return o.ApplyT(func(v LookupStreamResult) *StreamModeDetails { return v.StreamModeDetails }).(StreamModeDetailsPtrOutput)
}

// An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
func (o LookupStreamResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupStreamResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamResultOutput{})
}
