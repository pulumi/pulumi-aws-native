// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch
{
    public static class GetAlarm
    {
        /// <summary>
        /// The ``AWS::CloudWatch::Alarm`` type specifies an alarm and associates it with the specified metric or metric math expression.
        ///  When this operation creates an alarm, the alarm state is immediately set to ``INSUFFICIENT_DATA``. The alarm is then evaluated and its state is set appropriately. Any actions associated with the new state are then executed.
        ///  When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.
        /// </summary>
        public static Task<GetAlarmResult> InvokeAsync(GetAlarmArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlarmResult>("aws-native:cloudwatch:getAlarm", args ?? new GetAlarmArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::CloudWatch::Alarm`` type specifies an alarm and associates it with the specified metric or metric math expression.
        ///  When this operation creates an alarm, the alarm state is immediately set to ``INSUFFICIENT_DATA``. The alarm is then evaluated and its state is set appropriately. Any actions associated with the new state are then executed.
        ///  When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.
        /// </summary>
        public static Output<GetAlarmResult> Invoke(GetAlarmInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlarmResult>("aws-native:cloudwatch:getAlarm", args ?? new GetAlarmInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::CloudWatch::Alarm`` type specifies an alarm and associates it with the specified metric or metric math expression.
        ///  When this operation creates an alarm, the alarm state is immediately set to ``INSUFFICIENT_DATA``. The alarm is then evaluated and its state is set appropriately. Any actions associated with the new state are then executed.
        ///  When you update an existing alarm, its state is left unchanged, but the update completely overwrites the previous configuration of the alarm.
        /// </summary>
        public static Output<GetAlarmResult> Invoke(GetAlarmInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlarmResult>("aws-native:cloudwatch:getAlarm", args ?? new GetAlarmInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlarmArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the alarm. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the alarm name. 
        ///   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Input("alarmName", required: true)]
        public string AlarmName { get; set; } = null!;

        public GetAlarmArgs()
        {
        }
        public static new GetAlarmArgs Empty => new GetAlarmArgs();
    }

    public sealed class GetAlarmInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the alarm. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the alarm name. 
        ///   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Input("alarmName", required: true)]
        public Input<string> AlarmName { get; set; } = null!;

        public GetAlarmInvokeArgs()
        {
        }
        public static new GetAlarmInvokeArgs Empty => new GetAlarmInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlarmResult
    {
        /// <summary>
        /// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
        /// </summary>
        public readonly bool? ActionsEnabled;
        /// <summary>
        /// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN). For more information about creating alarms and the actions that you can specify, see [PutMetricAlarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_PutMetricAlarm.html) in the *API Reference*.
        /// </summary>
        public readonly ImmutableArray<string> AlarmActions;
        /// <summary>
        /// The description of the alarm.
        /// </summary>
        public readonly string? AlarmDescription;
        /// <summary>
        /// The ARN of the CloudWatch alarm, such as `arn:aws:cloudwatch:us-west-2:123456789012:alarm:myCloudWatchAlarm-CPUAlarm-UXMMZK36R55Z` .
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The arithmetic operation to use when comparing the specified statistic and threshold. The specified statistic value is used as the first operand.
        /// </summary>
        public readonly string? ComparisonOperator;
        /// <summary>
        /// The number of datapoints that must be breaching to trigger the alarm. This is used only if you are setting an "M out of N" alarm. In that case, this value is the M, and the value that you set for ``EvaluationPeriods`` is the N value. For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
        ///  If you omit this parameter, CW uses the same value here that you set for ``EvaluationPeriods``, and the alarm goes to alarm state if that many consecutive periods are breaching.
        /// </summary>
        public readonly int? DatapointsToAlarm;
        /// <summary>
        /// The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify ``Dimensions``. Instead, you use ``Metrics``.
        /// </summary>
        public readonly ImmutableArray<Outputs.AlarmDimension> Dimensions;
        /// <summary>
        /// Used only for alarms based on percentiles. If ``ignore``, the alarm state does not change during periods with too few data points to be statistically significant. If ``evaluate`` or this parameter is not used, the alarm is always evaluated and possibly changes state no matter how many data points are available.
        /// </summary>
        public readonly string? EvaluateLowSampleCountPercentile;
        /// <summary>
        /// The number of periods over which data is compared to the specified threshold. If you are setting an alarm that requires that a number of consecutive data points be breaching to trigger the alarm, this value specifies that number. If you are setting an "M out of N" alarm, this value is the N, and ``DatapointsToAlarm`` is the M.
        ///  For more information, see [Evaluating an Alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation) in the *User Guide*.
        /// </summary>
        public readonly int? EvaluationPeriods;
        /// <summary>
        /// The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
        ///  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
        ///  For an alarm based on a math expression, you can't specify ``ExtendedStatistic``. Instead, you use ``Metrics``.
        /// </summary>
        public readonly string? ExtendedStatistic;
        /// <summary>
        /// The actions to execute when this alarm transitions to the ``INSUFFICIENT_DATA`` state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        /// </summary>
        public readonly ImmutableArray<string> InsufficientDataActions;
        /// <summary>
        /// The name of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you use ``Metrics`` instead and you can't specify ``MetricName``.
        /// </summary>
        public readonly string? MetricName;
        /// <summary>
        /// An array that enables you to create an alarm based on the result of a metric math expression. Each item in the array either retrieves a metric or performs a math expression.
        ///  If you specify the ``Metrics`` parameter, you cannot specify ``MetricName``, ``Dimensions``, ``Period``, ``Namespace``, ``Statistic``, ``ExtendedStatistic``, or ``Unit``.
        /// </summary>
        public readonly ImmutableArray<Outputs.AlarmMetricDataQuery> Metrics;
        /// <summary>
        /// The namespace of the metric associated with the alarm. This is required for an alarm based on a metric. For an alarm based on a math expression, you can't specify ``Namespace`` and you use ``Metrics`` instead.
        ///  For a list of namespaces for metrics from AWS services, see [Services That Publish Metrics.](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html)
        /// </summary>
        public readonly string? Namespace;
        /// <summary>
        /// The actions to execute when this alarm transitions to the ``OK`` state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        /// </summary>
        public readonly ImmutableArray<string> OkActions;
        /// <summary>
        /// The period, in seconds, over which the statistic is applied. This is required for an alarm based on a metric. Valid values are 10, 30, 60, and any multiple of 60.
        ///  For an alarm based on a math expression, you can't specify ``Period``, and instead you use the ``Metrics`` parameter.
        ///   *Minimum:* 10
        /// </summary>
        public readonly int? Period;
        /// <summary>
        /// The statistic for the metric associated with the alarm, other than percentile. For percentile statistics, use ``ExtendedStatistic``.
        ///  For an alarm based on a metric, you must specify either ``Statistic`` or ``ExtendedStatistic`` but not both.
        ///  For an alarm based on a math expression, you can't specify ``Statistic``. Instead, you use ``Metrics``.
        /// </summary>
        public readonly string? Statistic;
        /// <summary>
        /// A list of key-value pairs to associate with the alarm. You can associate as many as 50 tags with an alarm. To be able to associate tags with the alarm when you create the alarm, you must have the ``cloudwatch:TagResource`` permission.
        ///  Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The value to compare with the specified statistic.
        /// </summary>
        public readonly double? Threshold;
        /// <summary>
        /// In an alarm based on an anomaly detection model, this is the ID of the ``ANOMALY_DETECTION_BAND`` function used as the threshold for the alarm.
        /// </summary>
        public readonly string? ThresholdMetricId;
        /// <summary>
        /// Sets how this alarm is to handle missing data points. Valid values are ``breaching``, ``notBreaching``, ``ignore``, and ``missing``. For more information, see [Configuring How Alarms Treat Missing Data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data) in the *Amazon User Guide*.
        ///  If you omit this parameter, the default behavior of ``missing`` is used.
        /// </summary>
        public readonly string? TreatMissingData;
        /// <summary>
        /// The unit of the metric associated with the alarm. Specify this only if you are creating an alarm based on a single metric. Do not specify this if you are specifying a ``Metrics`` array.
        ///   You can specify the following values: Seconds, Microseconds, Milliseconds, Bytes, Kilobytes, Megabytes, Gigabytes, Terabytes, Bits, Kilobits, Megabits, Gigabits, Terabits, Percent, Count, Bytes/Second, Kilobytes/Second, Megabytes/Second, Gigabytes/Second, Terabytes/Second, Bits/Second, Kilobits/Second, Megabits/Second, Gigabits/Second, Terabits/Second, Count/Second, or None.
        /// </summary>
        public readonly string? Unit;

        [OutputConstructor]
        private GetAlarmResult(
            bool? actionsEnabled,

            ImmutableArray<string> alarmActions,

            string? alarmDescription,

            string? arn,

            string? comparisonOperator,

            int? datapointsToAlarm,

            ImmutableArray<Outputs.AlarmDimension> dimensions,

            string? evaluateLowSampleCountPercentile,

            int? evaluationPeriods,

            string? extendedStatistic,

            ImmutableArray<string> insufficientDataActions,

            string? metricName,

            ImmutableArray<Outputs.AlarmMetricDataQuery> metrics,

            string? @namespace,

            ImmutableArray<string> okActions,

            int? period,

            string? statistic,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            double? threshold,

            string? thresholdMetricId,

            string? treatMissingData,

            string? unit)
        {
            ActionsEnabled = actionsEnabled;
            AlarmActions = alarmActions;
            AlarmDescription = alarmDescription;
            Arn = arn;
            ComparisonOperator = comparisonOperator;
            DatapointsToAlarm = datapointsToAlarm;
            Dimensions = dimensions;
            EvaluateLowSampleCountPercentile = evaluateLowSampleCountPercentile;
            EvaluationPeriods = evaluationPeriods;
            ExtendedStatistic = extendedStatistic;
            InsufficientDataActions = insufficientDataActions;
            MetricName = metricName;
            Metrics = metrics;
            Namespace = @namespace;
            OkActions = okActions;
            Period = period;
            Statistic = statistic;
            Tags = tags;
            Threshold = threshold;
            ThresholdMetricId = thresholdMetricId;
            TreatMissingData = treatMissingData;
            Unit = unit;
        }
    }
}
