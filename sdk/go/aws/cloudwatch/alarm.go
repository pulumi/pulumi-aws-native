// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::CloudWatch::Alarm“ type specifies an alarm and associates it with the specified metric or metric math expression.
//
//	When this operation creates an alarm, the alarm state is immediately set to ``INSUFFICIENT_DATA``. The alarm is then evaluated and its state is set appropriately. Any actions associated with the new state are then executed.
//	When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.
type Alarm struct {
	pulumi.CustomResourceState

	// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
	ActionsEnabled pulumi.BoolPtrOutput `pulumi:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *API Reference*.
	AlarmActions pulumi.StringArrayOutput `pulumi:"alarmActions"`
	// The description of the alarm.
	AlarmDescription pulumi.StringPtrOutput `pulumi:"alarmDescription"`
	// The name of the alarm. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the alarm name.
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	AlarmName pulumi.StringPtrOutput `pulumi:"alarmName"`
	// The ARN of the CloudWatch alarm, such as `arn:aws:cloudwatch:us-west-2:123456789012:alarm:myCloudWatchAlarm-CPUAlarm-UXMMZK36R55Z` .
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand.
	ComparisonOperator pulumi.StringOutput `pulumi:"comparisonOperator"`
	// The number of datapoints that must be breaching to trigger the alarm. This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for ``EvaluationPeriods`` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
	//  If you omit this parameter, CW uses the same value here that you set for ``EvaluationPeriods``, and the alarm goes to alarm state if that many consecutive periods are breaching.
	DatapointsToAlarm pulumi.IntPtrOutput `pulumi:"datapointsToAlarm"`
	// The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify ``Dimensions``. Instead, you use ``Metrics``.
	Dimensions AlarmDimensionArrayOutput `pulumi:"dimensions"`
	// Used only for alarms based on percentiles. If ``ignore``, the alarm state does not change during periods with too few data points to be statistically significant. If ``evaluate`` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.
	EvaluateLowSampleCountPercentile pulumi.StringPtrOutput `pulumi:"evaluateLowSampleCountPercentile"`
	// The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and ``DatapointsToAlarm`` is the M.
	//  For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
	EvaluationPeriods pulumi.IntOutput `pulumi:"evaluationPeriods"`
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	//  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
	//  For an alarm based on a math expression, you can't specify ``ExtendedStatistic``. Instead, you use ``Metrics``.
	ExtendedStatistic pulumi.StringPtrOutput `pulumi:"extendedStatistic"`
	// The actions to execute when this alarm transitions to the ``INSUFFICIENT_DATA`` state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions pulumi.StringArrayOutput `pulumi:"insufficientDataActions"`
	// The name of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you use ``Metrics`` instead and you can't specify ``MetricName``.
	MetricName pulumi.StringPtrOutput `pulumi:"metricName"`
	// An array that enables you to create an alarm based on the result of a metric math expression. Each item in the array either retrieves a metric or performs a math expression.
	//  If you specify the ``Metrics`` parameter, you cannot specify ``MetricName``, ``Dimensions``, ``Period``, ``Namespace``, ``Statistic``, ``ExtendedStatistic``, or ``Unit``.
	Metrics AlarmMetricDataQueryArrayOutput `pulumi:"metrics"`
	// The namespace of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify ``Namespace`` and you use ``Metrics`` instead.
	//  For a list of namespaces for metrics from AWS services, see [Services That Publish Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The actions to execute when this alarm transitions to the ``OK`` state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	OkActions pulumi.StringArrayOutput `pulumi:"okActions"`
	// The period, in seconds, over which the statistic is applied. This is required for an alarm based on a metric. Valid values are 10, 30, 60, and any multiple of 60.
	//  For an alarm based on a math expression, you can't specify ``Period``, and instead you use the ``Metrics`` parameter.
	//   *Minimum:* 10
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use ``ExtendedStatistic``.
	//  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
	//  For an alarm based on a math expression, you can't specify ``Statistic``. Instead, you use ``Metrics``.
	Statistic pulumi.StringPtrOutput `pulumi:"statistic"`
	// A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the ``cloudwatch:TagResource`` permission.
	//  Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// The value to compare with the specified statistic.
	Threshold pulumi.Float64PtrOutput `pulumi:"threshold"`
	// In an alarm based on an anomaly detection model, this is the ID of the ``ANOMALY_DETECTION_BAND`` function used as the threshold for the alarm.
	ThresholdMetricId pulumi.StringPtrOutput `pulumi:"thresholdMetricId"`
	// Sets how this alarm is to handle missing data points. Valid values are ``breaching``, ``notBreaching``, ``ignore``, and ``missing``. For more information, see [Configuring How Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon User Guide*.
	//  If you omit this parameter, the default behavior of ``missing`` is used.
	TreatMissingData pulumi.StringPtrOutput `pulumi:"treatMissingData"`
	// The unit of the metric associated with the alarm. Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a ``Metrics`` array.
	//   You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
	Unit pulumi.StringPtrOutput `pulumi:"unit"`
}

// NewAlarm registers a new resource with the given unique name, arguments, and options.
func NewAlarm(ctx *pulumi.Context,
	name string, args *AlarmArgs, opts ...pulumi.ResourceOption) (*Alarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComparisonOperator == nil {
		return nil, errors.New("invalid value for required argument 'ComparisonOperator'")
	}
	if args.EvaluationPeriods == nil {
		return nil, errors.New("invalid value for required argument 'EvaluationPeriods'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"alarmName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Alarm
	err := ctx.RegisterResource("aws-native:cloudwatch:Alarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlarm gets an existing Alarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlarmState, opts ...pulumi.ResourceOption) (*Alarm, error) {
	var resource Alarm
	err := ctx.ReadResource("aws-native:cloudwatch:Alarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alarm resources.
type alarmState struct {
}

type AlarmState struct {
}

func (AlarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmState)(nil)).Elem()
}

type alarmArgs struct {
	// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
	ActionsEnabled *bool `pulumi:"actionsEnabled"`
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *API Reference*.
	AlarmActions []string `pulumi:"alarmActions"`
	// The description of the alarm.
	AlarmDescription *string `pulumi:"alarmDescription"`
	// The name of the alarm. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the alarm name.
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	AlarmName *string `pulumi:"alarmName"`
	// The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand.
	ComparisonOperator string `pulumi:"comparisonOperator"`
	// The number of datapoints that must be breaching to trigger the alarm. This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for ``EvaluationPeriods`` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
	//  If you omit this parameter, CW uses the same value here that you set for ``EvaluationPeriods``, and the alarm goes to alarm state if that many consecutive periods are breaching.
	DatapointsToAlarm *int `pulumi:"datapointsToAlarm"`
	// The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify ``Dimensions``. Instead, you use ``Metrics``.
	Dimensions []AlarmDimension `pulumi:"dimensions"`
	// Used only for alarms based on percentiles. If ``ignore``, the alarm state does not change during periods with too few data points to be statistically significant. If ``evaluate`` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.
	EvaluateLowSampleCountPercentile *string `pulumi:"evaluateLowSampleCountPercentile"`
	// The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and ``DatapointsToAlarm`` is the M.
	//  For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
	EvaluationPeriods int `pulumi:"evaluationPeriods"`
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	//  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
	//  For an alarm based on a math expression, you can't specify ``ExtendedStatistic``. Instead, you use ``Metrics``.
	ExtendedStatistic *string `pulumi:"extendedStatistic"`
	// The actions to execute when this alarm transitions to the ``INSUFFICIENT_DATA`` state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions []string `pulumi:"insufficientDataActions"`
	// The name of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you use ``Metrics`` instead and you can't specify ``MetricName``.
	MetricName *string `pulumi:"metricName"`
	// An array that enables you to create an alarm based on the result of a metric math expression. Each item in the array either retrieves a metric or performs a math expression.
	//  If you specify the ``Metrics`` parameter, you cannot specify ``MetricName``, ``Dimensions``, ``Period``, ``Namespace``, ``Statistic``, ``ExtendedStatistic``, or ``Unit``.
	Metrics []AlarmMetricDataQuery `pulumi:"metrics"`
	// The namespace of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify ``Namespace`` and you use ``Metrics`` instead.
	//  For a list of namespaces for metrics from AWS services, see [Services That Publish Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
	Namespace *string `pulumi:"namespace"`
	// The actions to execute when this alarm transitions to the ``OK`` state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	OkActions []string `pulumi:"okActions"`
	// The period, in seconds, over which the statistic is applied. This is required for an alarm based on a metric. Valid values are 10, 30, 60, and any multiple of 60.
	//  For an alarm based on a math expression, you can't specify ``Period``, and instead you use the ``Metrics`` parameter.
	//   *Minimum:* 10
	Period *int `pulumi:"period"`
	// The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use ``ExtendedStatistic``.
	//  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
	//  For an alarm based on a math expression, you can't specify ``Statistic``. Instead, you use ``Metrics``.
	Statistic *string `pulumi:"statistic"`
	// A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the ``cloudwatch:TagResource`` permission.
	//  Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	Tags []aws.Tag `pulumi:"tags"`
	// The value to compare with the specified statistic.
	Threshold *float64 `pulumi:"threshold"`
	// In an alarm based on an anomaly detection model, this is the ID of the ``ANOMALY_DETECTION_BAND`` function used as the threshold for the alarm.
	ThresholdMetricId *string `pulumi:"thresholdMetricId"`
	// Sets how this alarm is to handle missing data points. Valid values are ``breaching``, ``notBreaching``, ``ignore``, and ``missing``. For more information, see [Configuring How Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon User Guide*.
	//  If you omit this parameter, the default behavior of ``missing`` is used.
	TreatMissingData *string `pulumi:"treatMissingData"`
	// The unit of the metric associated with the alarm. Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a ``Metrics`` array.
	//   You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
	Unit *string `pulumi:"unit"`
}

// The set of arguments for constructing a Alarm resource.
type AlarmArgs struct {
	// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
	ActionsEnabled pulumi.BoolPtrInput
	// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *API Reference*.
	AlarmActions pulumi.StringArrayInput
	// The description of the alarm.
	AlarmDescription pulumi.StringPtrInput
	// The name of the alarm. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the alarm name.
	//   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	AlarmName pulumi.StringPtrInput
	// The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand.
	ComparisonOperator pulumi.StringInput
	// The number of datapoints that must be breaching to trigger the alarm. This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for ``EvaluationPeriods`` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
	//  If you omit this parameter, CW uses the same value here that you set for ``EvaluationPeriods``, and the alarm goes to alarm state if that many consecutive periods are breaching.
	DatapointsToAlarm pulumi.IntPtrInput
	// The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify ``Dimensions``. Instead, you use ``Metrics``.
	Dimensions AlarmDimensionArrayInput
	// Used only for alarms based on percentiles. If ``ignore``, the alarm state does not change during periods with too few data points to be statistically significant. If ``evaluate`` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.
	EvaluateLowSampleCountPercentile pulumi.StringPtrInput
	// The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and ``DatapointsToAlarm`` is the M.
	//  For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
	EvaluationPeriods pulumi.IntInput
	// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
	//  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
	//  For an alarm based on a math expression, you can't specify ``ExtendedStatistic``. Instead, you use ``Metrics``.
	ExtendedStatistic pulumi.StringPtrInput
	// The actions to execute when this alarm transitions to the ``INSUFFICIENT_DATA`` state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	InsufficientDataActions pulumi.StringArrayInput
	// The name of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you use ``Metrics`` instead and you can't specify ``MetricName``.
	MetricName pulumi.StringPtrInput
	// An array that enables you to create an alarm based on the result of a metric math expression. Each item in the array either retrieves a metric or performs a math expression.
	//  If you specify the ``Metrics`` parameter, you cannot specify ``MetricName``, ``Dimensions``, ``Period``, ``Namespace``, ``Statistic``, ``ExtendedStatistic``, or ``Unit``.
	Metrics AlarmMetricDataQueryArrayInput
	// The namespace of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify ``Namespace`` and you use ``Metrics`` instead.
	//  For a list of namespaces for metrics from AWS services, see [Services That Publish Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
	Namespace pulumi.StringPtrInput
	// The actions to execute when this alarm transitions to the ``OK`` state from any other state. Each action is specified as an Amazon Resource Name (ARN).
	OkActions pulumi.StringArrayInput
	// The period, in seconds, over which the statistic is applied. This is required for an alarm based on a metric. Valid values are 10, 30, 60, and any multiple of 60.
	//  For an alarm based on a math expression, you can't specify ``Period``, and instead you use the ``Metrics`` parameter.
	//   *Minimum:* 10
	Period pulumi.IntPtrInput
	// The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use ``ExtendedStatistic``.
	//  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
	//  For an alarm based on a math expression, you can't specify ``Statistic``. Instead, you use ``Metrics``.
	Statistic pulumi.StringPtrInput
	// A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the ``cloudwatch:TagResource`` permission.
	//  Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	Tags aws.TagArrayInput
	// The value to compare with the specified statistic.
	Threshold pulumi.Float64PtrInput
	// In an alarm based on an anomaly detection model, this is the ID of the ``ANOMALY_DETECTION_BAND`` function used as the threshold for the alarm.
	ThresholdMetricId pulumi.StringPtrInput
	// Sets how this alarm is to handle missing data points. Valid values are ``breaching``, ``notBreaching``, ``ignore``, and ``missing``. For more information, see [Configuring How Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon User Guide*.
	//  If you omit this parameter, the default behavior of ``missing`` is used.
	TreatMissingData pulumi.StringPtrInput
	// The unit of the metric associated with the alarm. Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a ``Metrics`` array.
	//   You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
	Unit pulumi.StringPtrInput
}

func (AlarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alarmArgs)(nil)).Elem()
}

type AlarmInput interface {
	pulumi.Input

	ToAlarmOutput() AlarmOutput
	ToAlarmOutputWithContext(ctx context.Context) AlarmOutput
}

func (*Alarm) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (i *Alarm) ToAlarmOutput() AlarmOutput {
	return i.ToAlarmOutputWithContext(context.Background())
}

func (i *Alarm) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlarmOutput)
}

type AlarmOutput struct{ *pulumi.OutputState }

func (AlarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alarm)(nil)).Elem()
}

func (o AlarmOutput) ToAlarmOutput() AlarmOutput {
	return o
}

func (o AlarmOutput) ToAlarmOutputWithContext(ctx context.Context) AlarmOutput {
	return o
}

// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
func (o AlarmOutput) ActionsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.BoolPtrOutput { return v.ActionsEnabled }).(pulumi.BoolPtrOutput)
}

// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *API Reference*.
func (o AlarmOutput) AlarmActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringArrayOutput { return v.AlarmActions }).(pulumi.StringArrayOutput)
}

// The description of the alarm.
func (o AlarmOutput) AlarmDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.AlarmDescription }).(pulumi.StringPtrOutput)
}

// The name of the alarm. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the alarm name.
//
//	If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
func (o AlarmOutput) AlarmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.AlarmName }).(pulumi.StringPtrOutput)
}

// The ARN of the CloudWatch alarm, such as `arn:aws:cloudwatch:us-west-2:123456789012:alarm:myCloudWatchAlarm-CPUAlarm-UXMMZK36R55Z` .
func (o AlarmOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand.
func (o AlarmOutput) ComparisonOperator() pulumi.StringOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringOutput { return v.ComparisonOperator }).(pulumi.StringOutput)
}

// The number of datapoints that must be breaching to trigger the alarm. This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for “EvaluationPeriods“ is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
//
//	If you omit this parameter, CW uses the same value here that you set for ``EvaluationPeriods``, and the alarm goes to alarm state if that many consecutive periods are breaching.
func (o AlarmOutput) DatapointsToAlarm() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntPtrOutput { return v.DatapointsToAlarm }).(pulumi.IntPtrOutput)
}

// The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify “Dimensions“. Instead, you use “Metrics“.
func (o AlarmOutput) Dimensions() AlarmDimensionArrayOutput {
	return o.ApplyT(func(v *Alarm) AlarmDimensionArrayOutput { return v.Dimensions }).(AlarmDimensionArrayOutput)
}

// Used only for alarms based on percentiles. If “ignore“, the alarm state does not change during periods with too few data points to be statistically significant. If “evaluate“ or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.
func (o AlarmOutput) EvaluateLowSampleCountPercentile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.EvaluateLowSampleCountPercentile }).(pulumi.StringPtrOutput)
}

// The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and “DatapointsToAlarm“ is the M.
//
//	For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
func (o AlarmOutput) EvaluationPeriods() pulumi.IntOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntOutput { return v.EvaluationPeriods }).(pulumi.IntOutput)
}

// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
//
//	For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
//	For an alarm based on a math expression, you can't specify ``ExtendedStatistic``. Instead, you use ``Metrics``.
func (o AlarmOutput) ExtendedStatistic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.ExtendedStatistic }).(pulumi.StringPtrOutput)
}

// The actions to execute when this alarm transitions to the “INSUFFICIENT_DATA“ state from any other state. Each action is specified as an Amazon Resource Name (ARN).
func (o AlarmOutput) InsufficientDataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringArrayOutput { return v.InsufficientDataActions }).(pulumi.StringArrayOutput)
}

// The name of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you use “Metrics“ instead and you can't specify “MetricName“.
func (o AlarmOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.MetricName }).(pulumi.StringPtrOutput)
}

// An array that enables you to create an alarm based on the result of a metric math expression. Each item in the array either retrieves a metric or performs a math expression.
//
//	If you specify the ``Metrics`` parameter, you cannot specify ``MetricName``, ``Dimensions``, ``Period``, ``Namespace``, ``Statistic``, ``ExtendedStatistic``, or ``Unit``.
func (o AlarmOutput) Metrics() AlarmMetricDataQueryArrayOutput {
	return o.ApplyT(func(v *Alarm) AlarmMetricDataQueryArrayOutput { return v.Metrics }).(AlarmMetricDataQueryArrayOutput)
}

// The namespace of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify “Namespace“ and you use “Metrics“ instead.
//
//	For a list of namespaces for metrics from AWS services, see [Services That Publish Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
func (o AlarmOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The actions to execute when this alarm transitions to the “OK“ state from any other state. Each action is specified as an Amazon Resource Name (ARN).
func (o AlarmOutput) OkActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringArrayOutput { return v.OkActions }).(pulumi.StringArrayOutput)
}

// The period, in seconds, over which the statistic is applied. This is required for an alarm based on a metric. Valid values are 10, 30, 60, and any multiple of 60.
//
//	For an alarm based on a math expression, you can't specify ``Period``, and instead you use the ``Metrics`` parameter.
//	 *Minimum:* 10
func (o AlarmOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use “ExtendedStatistic“.
//
//	For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
//	For an alarm based on a math expression, you can't specify ``Statistic``. Instead, you use ``Metrics``.
func (o AlarmOutput) Statistic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.Statistic }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the “cloudwatch:TagResource“ permission.
//
//	Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
func (o AlarmOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Alarm) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

// The value to compare with the specified statistic.
func (o AlarmOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.Float64PtrOutput { return v.Threshold }).(pulumi.Float64PtrOutput)
}

// In an alarm based on an anomaly detection model, this is the ID of the “ANOMALY_DETECTION_BAND“ function used as the threshold for the alarm.
func (o AlarmOutput) ThresholdMetricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.ThresholdMetricId }).(pulumi.StringPtrOutput)
}

// Sets how this alarm is to handle missing data points. Valid values are “breaching“, “notBreaching“, “ignore“, and “missing“. For more information, see [Configuring How Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon User Guide*.
//
//	If you omit this parameter, the default behavior of ``missing`` is used.
func (o AlarmOutput) TreatMissingData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.TreatMissingData }).(pulumi.StringPtrOutput)
}

// The unit of the metric associated with the alarm. Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a “Metrics“ array.
//
//	You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
func (o AlarmOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alarm) pulumi.StringPtrOutput { return v.Unit }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlarmInput)(nil)).Elem(), &Alarm{})
	pulumi.RegisterOutputType(AlarmOutput{})
}
