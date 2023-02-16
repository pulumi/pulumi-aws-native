// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies a metric filter that describes how CloudWatch Logs extracts information from logs and transforms it into Amazon CloudWatch metrics.
type MetricFilter struct {
	pulumi.CustomResourceState

	// A name for the metric filter.
	FilterName pulumi.StringPtrOutput `pulumi:"filterName"`
	// Pattern that Logs follows to interpret each entry in a log.
	FilterPattern pulumi.StringOutput `pulumi:"filterPattern"`
	// Existing log group that you want to associate with this filter.
	LogGroupName pulumi.StringOutput `pulumi:"logGroupName"`
	// A collection of information that defines how metric data gets emitted.
	MetricTransformations MetricFilterMetricTransformationArrayOutput `pulumi:"metricTransformations"`
}

// NewMetricFilter registers a new resource with the given unique name, arguments, and options.
func NewMetricFilter(ctx *pulumi.Context,
	name string, args *MetricFilterArgs, opts ...pulumi.ResourceOption) (*MetricFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FilterPattern == nil {
		return nil, errors.New("invalid value for required argument 'FilterPattern'")
	}
	if args.LogGroupName == nil {
		return nil, errors.New("invalid value for required argument 'LogGroupName'")
	}
	if args.MetricTransformations == nil {
		return nil, errors.New("invalid value for required argument 'MetricTransformations'")
	}
	var resource MetricFilter
	err := ctx.RegisterResource("aws-native:logs:MetricFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricFilter gets an existing MetricFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricFilterState, opts ...pulumi.ResourceOption) (*MetricFilter, error) {
	var resource MetricFilter
	err := ctx.ReadResource("aws-native:logs:MetricFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricFilter resources.
type metricFilterState struct {
}

type MetricFilterState struct {
}

func (MetricFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricFilterState)(nil)).Elem()
}

type metricFilterArgs struct {
	// A name for the metric filter.
	FilterName *string `pulumi:"filterName"`
	// Pattern that Logs follows to interpret each entry in a log.
	FilterPattern string `pulumi:"filterPattern"`
	// Existing log group that you want to associate with this filter.
	LogGroupName string `pulumi:"logGroupName"`
	// A collection of information that defines how metric data gets emitted.
	MetricTransformations []MetricFilterMetricTransformation `pulumi:"metricTransformations"`
}

// The set of arguments for constructing a MetricFilter resource.
type MetricFilterArgs struct {
	// A name for the metric filter.
	FilterName pulumi.StringPtrInput
	// Pattern that Logs follows to interpret each entry in a log.
	FilterPattern pulumi.StringInput
	// Existing log group that you want to associate with this filter.
	LogGroupName pulumi.StringInput
	// A collection of information that defines how metric data gets emitted.
	MetricTransformations MetricFilterMetricTransformationArrayInput
}

func (MetricFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricFilterArgs)(nil)).Elem()
}

type MetricFilterInput interface {
	pulumi.Input

	ToMetricFilterOutput() MetricFilterOutput
	ToMetricFilterOutputWithContext(ctx context.Context) MetricFilterOutput
}

func (*MetricFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricFilter)(nil)).Elem()
}

func (i *MetricFilter) ToMetricFilterOutput() MetricFilterOutput {
	return i.ToMetricFilterOutputWithContext(context.Background())
}

func (i *MetricFilter) ToMetricFilterOutputWithContext(ctx context.Context) MetricFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricFilterOutput)
}

type MetricFilterOutput struct{ *pulumi.OutputState }

func (MetricFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricFilter)(nil)).Elem()
}

func (o MetricFilterOutput) ToMetricFilterOutput() MetricFilterOutput {
	return o
}

func (o MetricFilterOutput) ToMetricFilterOutputWithContext(ctx context.Context) MetricFilterOutput {
	return o
}

// A name for the metric filter.
func (o MetricFilterOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetricFilter) pulumi.StringPtrOutput { return v.FilterName }).(pulumi.StringPtrOutput)
}

// Pattern that Logs follows to interpret each entry in a log.
func (o MetricFilterOutput) FilterPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricFilter) pulumi.StringOutput { return v.FilterPattern }).(pulumi.StringOutput)
}

// Existing log group that you want to associate with this filter.
func (o MetricFilterOutput) LogGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricFilter) pulumi.StringOutput { return v.LogGroupName }).(pulumi.StringOutput)
}

// A collection of information that defines how metric data gets emitted.
func (o MetricFilterOutput) MetricTransformations() MetricFilterMetricTransformationArrayOutput {
	return o.ApplyT(func(v *MetricFilter) MetricFilterMetricTransformationArrayOutput { return v.MetricTransformations }).(MetricFilterMetricTransformationArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetricFilterInput)(nil)).Elem(), &MetricFilter{})
	pulumi.RegisterOutputType(MetricFilterOutput{})
}