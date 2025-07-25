// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::Lambda::EventSourceMapping“ resource creates a mapping between an event source and an LAMlong function. LAM reads items from the event source and triggers the function.
//
//	For details about each event source type, see the following topics. In particular, each of the topics describes the required and optional parameters for the specific event source.
//	 +   [Configuring a Dynamo DB stream as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-dynamodb-eventsourcemapping)
//	 +   [Configuring a Kinesis stream as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-eventsourcemapping)
//	 +   [Configuring an SQS queue as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-eventsource)
//	 +   [Configuring an MQ broker as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-mq.html#services-mq-eventsourcemapping)
//	 +   [Configuring MSK as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html)
//	 +   [Configuring Self-Managed Apache Kafka as an event source](https://docs.aws.amazon.com/lambda/latest/dg/kafka-smaa.html)
//	 +   [Configuring Amazon DocumentDB as an event source](https://docs.aws.amazon.com/lambda/latest/dg/with-documentdb.html)
func LookupEventSourceMapping(ctx *pulumi.Context, args *LookupEventSourceMappingArgs, opts ...pulumi.InvokeOption) (*LookupEventSourceMappingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEventSourceMappingResult
	err := ctx.Invoke("aws-native:lambda:getEventSourceMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventSourceMappingArgs struct {
	// The event source mapping's ID.
	Id string `pulumi:"id"`
}

type LookupEventSourceMappingResult struct {
	// Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.
	AmazonManagedKafkaEventSourceConfig *EventSourceMappingAmazonManagedKafkaEventSourceConfig `pulumi:"amazonManagedKafkaEventSourceConfig"`
	// The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function. Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).
	//   +  *Amazon Kinesis* – Default 100. Max 10,000.
	//   +  *Amazon DynamoDB Streams* – Default 100. Max 10,000.
	//   +  *Amazon Simple Queue Service* – Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.
	//   +  *Amazon Managed Streaming for Apache Kafka* – Default 100. Max 10,000.
	//   +  *Self-managed Apache Kafka* – Default 100. Max 10,000.
	//   +  *Amazon MQ (ActiveMQ and RabbitMQ)* – Default 100. Max 10,000.
	//   +  *DocumentDB* – Default 100. Max 10,000.
	BatchSize *int `pulumi:"batchSize"`
	// (Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry. The default value is false.
	//   When using ``BisectBatchOnFunctionError``, check the ``BatchSize`` parameter in the ``OnFailure`` destination message's metadata. The ``BatchSize`` could be greater than 1 since LAM consolidates failed messages metadata when writing to the ``OnFailure`` destination.
	BisectBatchOnFunctionError *bool `pulumi:"bisectBatchOnFunctionError"`
	// (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka event sources only) A configuration object that specifies the destination of an event after Lambda processes it.
	DestinationConfig *EventSourceMappingDestinationConfig `pulumi:"destinationConfig"`
	// Specific configuration settings for a DocumentDB event source.
	DocumentDbEventSourceConfig *EventSourceMappingDocumentDbEventSourceConfig `pulumi:"documentDbEventSourceConfig"`
	// When true, the event source mapping is active. When false, Lambda pauses polling and invocation.
	//  Default: True
	Enabled *bool `pulumi:"enabled"`
	// The Amazon Resource Name (ARN) of the event source mapping.
	EventSourceMappingArn *string `pulumi:"eventSourceMappingArn"`
	// An object that defines the filter criteria that determine whether Lambda should process an event. For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html).
	FilterCriteria *EventSourceMappingFilterCriteria `pulumi:"filterCriteria"`
	// The name or ARN of the Lambda function.
	//   **Name formats**
	//  +  *Function name* – ``MyFunction``.
	//   +  *Function ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:MyFunction``.
	//   +  *Version or Alias ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD``.
	//   +  *Partial ARN* – ``123456789012:function:MyFunction``.
	//
	//  The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.
	FunctionName *string `pulumi:"functionName"`
	// (Kinesis, DynamoDB Streams, and SQS) A list of current response type enums applied to the event source mapping.
	//  Valid Values: ``ReportBatchItemFailures``
	FunctionResponseTypes []EventSourceMappingFunctionResponseTypesItem `pulumi:"functionResponseTypes"`
	// The event source mapping's ID.
	Id *string `pulumi:"id"`
	// The ARN of the KMSlong (KMS) customer managed key that Lambda uses to encrypt your function's [filter criteria](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-basics).
	KmsKeyArn *string `pulumi:"kmsKeyArn"`
	// The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.
	//  *Default (, , event sources)*: 0
	//  *Default (, Kafka, , event sources)*: 500 ms
	//  *Related setting:* For SQS event sources, when you set ``BatchSize`` to a value greater than 10, you must set ``MaximumBatchingWindowInSeconds`` to at least 1.
	MaximumBatchingWindowInSeconds *int `pulumi:"maximumBatchingWindowInSeconds"`
	// (Kinesis and DynamoDB Streams only) Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.
	//   The minimum valid value for maximum record age is 60s. Although values less than 60 and greater than -1 fall within the parameter's absolute range, they are not allowed
	MaximumRecordAgeInSeconds *int `pulumi:"maximumRecordAgeInSeconds"`
	// (Kinesis and DynamoDB Streams only) Discard records after the specified number of retries. The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.
	MaximumRetryAttempts *int `pulumi:"maximumRetryAttempts"`
	// The metrics configuration for your event source. For more information, see [Event source mapping metrics](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics-types.html#event-source-mapping-metrics).
	MetricsConfig *EventSourceMappingMetricsConfig `pulumi:"metricsConfig"`
	// (Kinesis and DynamoDB Streams only) The number of batches to process concurrently from each shard. The default value is 1.
	ParallelizationFactor *int `pulumi:"parallelizationFactor"`
	// (Amazon MSK and self-managed Apache Kafka only) The provisioned mode configuration for the event source. For more information, see [provisioned mode](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-provisioned-mode).
	ProvisionedPollerConfig *EventSourceMappingProvisionedPollerConfig `pulumi:"provisionedPollerConfig"`
	// (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
	Queues []string `pulumi:"queues"`
	// (Amazon SQS only) The scaling configuration for the event source. For more information, see [Configuring maximum concurrency for Amazon SQS event sources](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency).
	ScalingConfig *EventSourceMappingScalingConfig `pulumi:"scalingConfig"`
	// Specific configuration settings for a self-managed Apache Kafka event source.
	SelfManagedKafkaEventSourceConfig *EventSourceMappingSelfManagedKafkaEventSourceConfig `pulumi:"selfManagedKafkaEventSourceConfig"`
	// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
	SourceAccessConfigurations []EventSourceMappingSourceAccessConfiguration `pulumi:"sourceAccessConfigurations"`
	// A list of tags to add to the event source mapping.
	//   You must have the ``lambda:TagResource``, ``lambda:UntagResource``, and ``lambda:ListTags`` permissions for your [principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) to manage the CFN stack. If you don't have these permissions, there might be unexpected behavior with stack-level tags propagating to the resource during resource creation and update.
	Tags []aws.Tag `pulumi:"tags"`
	// The name of the Kafka topic.
	Topics []string `pulumi:"topics"`
	// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources. A value of 0 seconds indicates no tumbling window.
	TumblingWindowInSeconds *int `pulumi:"tumblingWindowInSeconds"`
}

func LookupEventSourceMappingOutput(ctx *pulumi.Context, args LookupEventSourceMappingOutputArgs, opts ...pulumi.InvokeOption) LookupEventSourceMappingResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEventSourceMappingResultOutput, error) {
			args := v.(LookupEventSourceMappingArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:lambda:getEventSourceMapping", args, LookupEventSourceMappingResultOutput{}, options).(LookupEventSourceMappingResultOutput), nil
		}).(LookupEventSourceMappingResultOutput)
}

type LookupEventSourceMappingOutputArgs struct {
	// The event source mapping's ID.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEventSourceMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSourceMappingArgs)(nil)).Elem()
}

type LookupEventSourceMappingResultOutput struct{ *pulumi.OutputState }

func (LookupEventSourceMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSourceMappingResult)(nil)).Elem()
}

func (o LookupEventSourceMappingResultOutput) ToLookupEventSourceMappingResultOutput() LookupEventSourceMappingResultOutput {
	return o
}

func (o LookupEventSourceMappingResultOutput) ToLookupEventSourceMappingResultOutputWithContext(ctx context.Context) LookupEventSourceMappingResultOutput {
	return o
}

// Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.
func (o LookupEventSourceMappingResultOutput) AmazonManagedKafkaEventSourceConfig() EventSourceMappingAmazonManagedKafkaEventSourceConfigPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingAmazonManagedKafkaEventSourceConfig {
		return v.AmazonManagedKafkaEventSourceConfig
	}).(EventSourceMappingAmazonManagedKafkaEventSourceConfigPtrOutput)
}

// The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function. Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).
//   - *Amazon Kinesis* – Default 100. Max 10,000.
//   - *Amazon DynamoDB Streams* – Default 100. Max 10,000.
//   - *Amazon Simple Queue Service* – Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.
//   - *Amazon Managed Streaming for Apache Kafka* – Default 100. Max 10,000.
//   - *Self-managed Apache Kafka* – Default 100. Max 10,000.
//   - *Amazon MQ (ActiveMQ and RabbitMQ)* – Default 100. Max 10,000.
//   - *DocumentDB* – Default 100. Max 10,000.
func (o LookupEventSourceMappingResultOutput) BatchSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.BatchSize }).(pulumi.IntPtrOutput)
}

// (Kinesis and DynamoDB Streams only) If the function returns an error, split the batch in two and retry. The default value is false.
//
//	When using ``BisectBatchOnFunctionError``, check the ``BatchSize`` parameter in the ``OnFailure`` destination message's metadata. The ``BatchSize`` could be greater than 1 since LAM consolidates failed messages metadata when writing to the ``OnFailure`` destination.
func (o LookupEventSourceMappingResultOutput) BisectBatchOnFunctionError() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *bool { return v.BisectBatchOnFunctionError }).(pulumi.BoolPtrOutput)
}

// (Kinesis, DynamoDB Streams, Amazon MSK, and self-managed Apache Kafka event sources only) A configuration object that specifies the destination of an event after Lambda processes it.
func (o LookupEventSourceMappingResultOutput) DestinationConfig() EventSourceMappingDestinationConfigPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingDestinationConfig {
		return v.DestinationConfig
	}).(EventSourceMappingDestinationConfigPtrOutput)
}

// Specific configuration settings for a DocumentDB event source.
func (o LookupEventSourceMappingResultOutput) DocumentDbEventSourceConfig() EventSourceMappingDocumentDbEventSourceConfigPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingDocumentDbEventSourceConfig {
		return v.DocumentDbEventSourceConfig
	}).(EventSourceMappingDocumentDbEventSourceConfigPtrOutput)
}

// When true, the event source mapping is active. When false, Lambda pauses polling and invocation.
//
//	Default: True
func (o LookupEventSourceMappingResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The Amazon Resource Name (ARN) of the event source mapping.
func (o LookupEventSourceMappingResultOutput) EventSourceMappingArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *string { return v.EventSourceMappingArn }).(pulumi.StringPtrOutput)
}

// An object that defines the filter criteria that determine whether Lambda should process an event. For more information, see [Lambda event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html).
func (o LookupEventSourceMappingResultOutput) FilterCriteria() EventSourceMappingFilterCriteriaPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingFilterCriteria { return v.FilterCriteria }).(EventSourceMappingFilterCriteriaPtrOutput)
}

// The name or ARN of the Lambda function.
//
//	 **Name formats**
//	+  *Function name* – ``MyFunction``.
//	 +  *Function ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:MyFunction``.
//	 +  *Version or Alias ARN* – ``arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD``.
//	 +  *Partial ARN* – ``123456789012:function:MyFunction``.
//
//	The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.
func (o LookupEventSourceMappingResultOutput) FunctionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *string { return v.FunctionName }).(pulumi.StringPtrOutput)
}

// (Kinesis, DynamoDB Streams, and SQS) A list of current response type enums applied to the event source mapping.
//
//	Valid Values: ``ReportBatchItemFailures``
func (o LookupEventSourceMappingResultOutput) FunctionResponseTypes() EventSourceMappingFunctionResponseTypesItemArrayOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) []EventSourceMappingFunctionResponseTypesItem {
		return v.FunctionResponseTypes
	}).(EventSourceMappingFunctionResponseTypesItemArrayOutput)
}

// The event source mapping's ID.
func (o LookupEventSourceMappingResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The ARN of the KMSlong (KMS) customer managed key that Lambda uses to encrypt your function's [filter criteria](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-basics).
func (o LookupEventSourceMappingResultOutput) KmsKeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *string { return v.KmsKeyArn }).(pulumi.StringPtrOutput)
}

// The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.
//
//	*Default (, , event sources)*: 0
//	*Default (, Kafka, , event sources)*: 500 ms
//	*Related setting:* For SQS event sources, when you set ``BatchSize`` to a value greater than 10, you must set ``MaximumBatchingWindowInSeconds`` to at least 1.
func (o LookupEventSourceMappingResultOutput) MaximumBatchingWindowInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.MaximumBatchingWindowInSeconds }).(pulumi.IntPtrOutput)
}

// (Kinesis and DynamoDB Streams only) Discard records older than the specified age. The default value is -1, which sets the maximum age to infinite. When the value is set to infinite, Lambda never discards old records.
//
//	The minimum valid value for maximum record age is 60s. Although values less than 60 and greater than -1 fall within the parameter's absolute range, they are not allowed
func (o LookupEventSourceMappingResultOutput) MaximumRecordAgeInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.MaximumRecordAgeInSeconds }).(pulumi.IntPtrOutput)
}

// (Kinesis and DynamoDB Streams only) Discard records after the specified number of retries. The default value is -1, which sets the maximum number of retries to infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records until the record expires in the event source.
func (o LookupEventSourceMappingResultOutput) MaximumRetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.MaximumRetryAttempts }).(pulumi.IntPtrOutput)
}

// The metrics configuration for your event source. For more information, see [Event source mapping metrics](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-metrics-types.html#event-source-mapping-metrics).
func (o LookupEventSourceMappingResultOutput) MetricsConfig() EventSourceMappingMetricsConfigPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingMetricsConfig { return v.MetricsConfig }).(EventSourceMappingMetricsConfigPtrOutput)
}

// (Kinesis and DynamoDB Streams only) The number of batches to process concurrently from each shard. The default value is 1.
func (o LookupEventSourceMappingResultOutput) ParallelizationFactor() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.ParallelizationFactor }).(pulumi.IntPtrOutput)
}

// (Amazon MSK and self-managed Apache Kafka only) The provisioned mode configuration for the event source. For more information, see [provisioned mode](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventsourcemapping.html#invocation-eventsourcemapping-provisioned-mode).
func (o LookupEventSourceMappingResultOutput) ProvisionedPollerConfig() EventSourceMappingProvisionedPollerConfigPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingProvisionedPollerConfig {
		return v.ProvisionedPollerConfig
	}).(EventSourceMappingProvisionedPollerConfigPtrOutput)
}

// (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
func (o LookupEventSourceMappingResultOutput) Queues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) []string { return v.Queues }).(pulumi.StringArrayOutput)
}

// (Amazon SQS only) The scaling configuration for the event source. For more information, see [Configuring maximum concurrency for Amazon SQS event sources](https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency).
func (o LookupEventSourceMappingResultOutput) ScalingConfig() EventSourceMappingScalingConfigPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingScalingConfig { return v.ScalingConfig }).(EventSourceMappingScalingConfigPtrOutput)
}

// Specific configuration settings for a self-managed Apache Kafka event source.
func (o LookupEventSourceMappingResultOutput) SelfManagedKafkaEventSourceConfig() EventSourceMappingSelfManagedKafkaEventSourceConfigPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *EventSourceMappingSelfManagedKafkaEventSourceConfig {
		return v.SelfManagedKafkaEventSourceConfig
	}).(EventSourceMappingSelfManagedKafkaEventSourceConfigPtrOutput)
}

// An array of the authentication protocol, VPC components, or virtual host to secure and define your event source.
func (o LookupEventSourceMappingResultOutput) SourceAccessConfigurations() EventSourceMappingSourceAccessConfigurationArrayOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) []EventSourceMappingSourceAccessConfiguration {
		return v.SourceAccessConfigurations
	}).(EventSourceMappingSourceAccessConfigurationArrayOutput)
}

// A list of tags to add to the event source mapping.
//
//	You must have the ``lambda:TagResource``, ``lambda:UntagResource``, and ``lambda:ListTags`` permissions for your [principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) to manage the CFN stack. If you don't have these permissions, there might be unexpected behavior with stack-level tags propagating to the resource during resource creation and update.
func (o LookupEventSourceMappingResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The name of the Kafka topic.
func (o LookupEventSourceMappingResultOutput) Topics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) []string { return v.Topics }).(pulumi.StringArrayOutput)
}

// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing window for DynamoDB and Kinesis Streams event sources. A value of 0 seconds indicates no tumbling window.
func (o LookupEventSourceMappingResultOutput) TumblingWindowInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEventSourceMappingResult) *int { return v.TumblingWindowInSeconds }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventSourceMappingResultOutput{})
}
