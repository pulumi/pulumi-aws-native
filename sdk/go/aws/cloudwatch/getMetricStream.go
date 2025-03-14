// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for Metric Stream
func LookupMetricStream(ctx *pulumi.Context, args *LookupMetricStreamArgs, opts ...pulumi.InvokeOption) (*LookupMetricStreamResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMetricStreamResult
	err := ctx.Invoke("aws-native:cloudwatch:getMetricStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetricStreamArgs struct {
	// Name of the metric stream.
	Name string `pulumi:"name"`
}

type LookupMetricStreamResult struct {
	// Amazon Resource Name of the metric stream.
	Arn *string `pulumi:"arn"`
	// The date of creation of the metric stream.
	CreationDate *string `pulumi:"creationDate"`
	// Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
	ExcludeFilters []MetricStreamFilter `pulumi:"excludeFilters"`
	// The ARN of the Kinesis Firehose where to stream the data.
	FirehoseArn *string `pulumi:"firehoseArn"`
	// Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
	IncludeFilters []MetricStreamFilter `pulumi:"includeFilters"`
	// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false.
	IncludeLinkedAccountsMetrics *bool `pulumi:"includeLinkedAccountsMetrics"`
	// The date of the last update of the metric stream.
	LastUpdateDate *string `pulumi:"lastUpdateDate"`
	// The output format of the data streamed to the Kinesis Firehose.
	OutputFormat *string `pulumi:"outputFormat"`
	// The ARN of the role that provides access to the Kinesis Firehose.
	RoleArn *string `pulumi:"roleArn"`
	// Displays the state of the Metric Stream.
	State *string `pulumi:"state"`
	// By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.
	StatisticsConfigurations []MetricStreamStatisticsConfiguration `pulumi:"statisticsConfigurations"`
	// A set of tags to assign to the delivery stream.
	Tags []aws.Tag `pulumi:"tags"`
}

func LookupMetricStreamOutput(ctx *pulumi.Context, args LookupMetricStreamOutputArgs, opts ...pulumi.InvokeOption) LookupMetricStreamResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupMetricStreamResultOutput, error) {
			args := v.(LookupMetricStreamArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("aws-native:cloudwatch:getMetricStream", args, LookupMetricStreamResultOutput{}, options).(LookupMetricStreamResultOutput), nil
		}).(LookupMetricStreamResultOutput)
}

type LookupMetricStreamOutputArgs struct {
	// Name of the metric stream.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupMetricStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricStreamArgs)(nil)).Elem()
}

type LookupMetricStreamResultOutput struct{ *pulumi.OutputState }

func (LookupMetricStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricStreamResult)(nil)).Elem()
}

func (o LookupMetricStreamResultOutput) ToLookupMetricStreamResultOutput() LookupMetricStreamResultOutput {
	return o
}

func (o LookupMetricStreamResultOutput) ToLookupMetricStreamResultOutputWithContext(ctx context.Context) LookupMetricStreamResultOutput {
	return o
}

// Amazon Resource Name of the metric stream.
func (o LookupMetricStreamResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The date of creation of the metric stream.
func (o LookupMetricStreamResultOutput) CreationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) *string { return v.CreationDate }).(pulumi.StringPtrOutput)
}

// Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
func (o LookupMetricStreamResultOutput) ExcludeFilters() MetricStreamFilterArrayOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) []MetricStreamFilter { return v.ExcludeFilters }).(MetricStreamFilterArrayOutput)
}

// The ARN of the Kinesis Firehose where to stream the data.
func (o LookupMetricStreamResultOutput) FirehoseArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) *string { return v.FirehoseArn }).(pulumi.StringPtrOutput)
}

// Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
func (o LookupMetricStreamResultOutput) IncludeFilters() MetricStreamFilterArrayOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) []MetricStreamFilter { return v.IncludeFilters }).(MetricStreamFilterArrayOutput)
}

// If you are creating a metric stream in a monitoring account, specify true to include metrics from source accounts that are linked to this monitoring account, in the metric stream. The default is false.
func (o LookupMetricStreamResultOutput) IncludeLinkedAccountsMetrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) *bool { return v.IncludeLinkedAccountsMetrics }).(pulumi.BoolPtrOutput)
}

// The date of the last update of the metric stream.
func (o LookupMetricStreamResultOutput) LastUpdateDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) *string { return v.LastUpdateDate }).(pulumi.StringPtrOutput)
}

// The output format of the data streamed to the Kinesis Firehose.
func (o LookupMetricStreamResultOutput) OutputFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) *string { return v.OutputFormat }).(pulumi.StringPtrOutput)
}

// The ARN of the role that provides access to the Kinesis Firehose.
func (o LookupMetricStreamResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// Displays the state of the Metric Stream.
func (o LookupMetricStreamResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.
func (o LookupMetricStreamResultOutput) StatisticsConfigurations() MetricStreamStatisticsConfigurationArrayOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) []MetricStreamStatisticsConfiguration {
		return v.StatisticsConfigurations
	}).(MetricStreamStatisticsConfigurationArrayOutput)
}

// A set of tags to assign to the delivery stream.
func (o LookupMetricStreamResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupMetricStreamResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetricStreamResultOutput{})
}
