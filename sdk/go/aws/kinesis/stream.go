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
type Stream struct {
	pulumi.CustomResourceState

	// The Amazon resource name (ARN) of the Kinesis stream
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The final list of shard-level metrics
	DesiredShardLevelMetrics StreamEnhancedMetricArrayOutput `pulumi:"desiredShardLevelMetrics"`
	// The name of the Kinesis stream.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The number of hours for the data records that are stored in shards to remain accessible.
	RetentionPeriodHours pulumi.IntPtrOutput `pulumi:"retentionPeriodHours"`
	// The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
	ShardCount pulumi.IntPtrOutput `pulumi:"shardCount"`
	// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
	StreamEncryption StreamEncryptionPtrOutput `pulumi:"streamEncryption"`
	// The mode in which the stream is running.
	StreamModeDetails StreamModeDetailsPtrOutput `pulumi:"streamModeDetails"`
	// An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewStream registers a new resource with the given unique name, arguments, and options.
func NewStream(ctx *pulumi.Context,
	name string, args *StreamArgs, opts ...pulumi.ResourceOption) (*Stream, error) {
	if args == nil {
		args = &StreamArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Stream
	err := ctx.RegisterResource("aws-native:kinesis:Stream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStream gets an existing Stream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamState, opts ...pulumi.ResourceOption) (*Stream, error) {
	var resource Stream
	err := ctx.ReadResource("aws-native:kinesis:Stream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stream resources.
type streamState struct {
}

type StreamState struct {
}

func (StreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamState)(nil)).Elem()
}

type streamArgs struct {
	// The final list of shard-level metrics
	DesiredShardLevelMetrics []StreamEnhancedMetric `pulumi:"desiredShardLevelMetrics"`
	// The name of the Kinesis stream.
	Name *string `pulumi:"name"`
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

// The set of arguments for constructing a Stream resource.
type StreamArgs struct {
	// The final list of shard-level metrics
	DesiredShardLevelMetrics StreamEnhancedMetricArrayInput
	// The name of the Kinesis stream.
	Name pulumi.StringPtrInput
	// The number of hours for the data records that are stored in shards to remain accessible.
	RetentionPeriodHours pulumi.IntPtrInput
	// The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
	ShardCount pulumi.IntPtrInput
	// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
	StreamEncryption StreamEncryptionPtrInput
	// The mode in which the stream is running.
	StreamModeDetails StreamModeDetailsPtrInput
	// An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
	Tags aws.TagArrayInput
}

func (StreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamArgs)(nil)).Elem()
}

type StreamInput interface {
	pulumi.Input

	ToStreamOutput() StreamOutput
	ToStreamOutputWithContext(ctx context.Context) StreamOutput
}

func (*Stream) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (i *Stream) ToStreamOutput() StreamOutput {
	return i.ToStreamOutputWithContext(context.Background())
}

func (i *Stream) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamOutput)
}

type StreamOutput struct{ *pulumi.OutputState }

func (StreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stream)(nil)).Elem()
}

func (o StreamOutput) ToStreamOutput() StreamOutput {
	return o
}

func (o StreamOutput) ToStreamOutputWithContext(ctx context.Context) StreamOutput {
	return o
}

// The Amazon resource name (ARN) of the Kinesis stream
func (o StreamOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The final list of shard-level metrics
func (o StreamOutput) DesiredShardLevelMetrics() StreamEnhancedMetricArrayOutput {
	return o.ApplyT(func(v *Stream) StreamEnhancedMetricArrayOutput { return v.DesiredShardLevelMetrics }).(StreamEnhancedMetricArrayOutput)
}

// The name of the Kinesis stream.
func (o StreamOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The number of hours for the data records that are stored in shards to remain accessible.
func (o StreamOutput) RetentionPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.RetentionPeriodHours }).(pulumi.IntPtrOutput)
}

// The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
func (o StreamOutput) ShardCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Stream) pulumi.IntPtrOutput { return v.ShardCount }).(pulumi.IntPtrOutput)
}

// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
func (o StreamOutput) StreamEncryption() StreamEncryptionPtrOutput {
	return o.ApplyT(func(v *Stream) StreamEncryptionPtrOutput { return v.StreamEncryption }).(StreamEncryptionPtrOutput)
}

// The mode in which the stream is running.
func (o StreamOutput) StreamModeDetails() StreamModeDetailsPtrOutput {
	return o.ApplyT(func(v *Stream) StreamModeDetailsPtrOutput { return v.StreamModeDetails }).(StreamModeDetailsPtrOutput)
}

// An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.
func (o StreamOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Stream) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StreamInput)(nil)).Elem(), &Stream{})
	pulumi.RegisterOutputType(StreamOutput{})
}
