// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudWatch
{
    public static class GetCompositeAlarm
    {
        /// <summary>
        /// The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression
        /// </summary>
        public static Task<GetCompositeAlarmResult> InvokeAsync(GetCompositeAlarmArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCompositeAlarmResult>("aws-native:cloudwatch:getCompositeAlarm", args ?? new GetCompositeAlarmArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression
        /// </summary>
        public static Output<GetCompositeAlarmResult> Invoke(GetCompositeAlarmInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCompositeAlarmResult>("aws-native:cloudwatch:getCompositeAlarm", args ?? new GetCompositeAlarmInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCompositeAlarmArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Composite Alarm
        /// </summary>
        [Input("alarmName", required: true)]
        public string AlarmName { get; set; } = null!;

        public GetCompositeAlarmArgs()
        {
        }
        public static new GetCompositeAlarmArgs Empty => new GetCompositeAlarmArgs();
    }

    public sealed class GetCompositeAlarmInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Composite Alarm
        /// </summary>
        [Input("alarmName", required: true)]
        public Input<string> AlarmName { get; set; } = null!;

        public GetCompositeAlarmInvokeArgs()
        {
        }
        public static new GetCompositeAlarmInvokeArgs Empty => new GetCompositeAlarmInvokeArgs();
    }


    [OutputType]
    public sealed class GetCompositeAlarmResult
    {
        /// <summary>
        /// Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
        /// </summary>
        public readonly bool? ActionsEnabled;
        /// <summary>
        /// Actions will be suppressed if the suppressor alarm is in the ALARM state. ActionsSuppressor can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm. 
        /// </summary>
        public readonly string? ActionsSuppressor;
        /// <summary>
        /// Actions will be suppressed if WaitPeriod is active. The length of time that actions are suppressed is in seconds.
        /// </summary>
        public readonly int? ActionsSuppressorExtensionPeriod;
        /// <summary>
        /// Actions will be suppressed if ExtensionPeriod is active. The length of time that actions are suppressed is in seconds.
        /// </summary>
        public readonly int? ActionsSuppressorWaitPeriod;
        /// <summary>
        /// The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).
        /// </summary>
        public readonly ImmutableArray<string> AlarmActions;
        /// <summary>
        /// The description of the alarm
        /// </summary>
        public readonly string? AlarmDescription;
        /// <summary>
        /// Expression which aggregates the state of other Alarms (Metric or Composite Alarms)
        /// </summary>
        public readonly string? AlarmRule;
        /// <summary>
        /// Amazon Resource Name (ARN) of the alarm
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        /// </summary>
        public readonly ImmutableArray<string> InsufficientDataActions;
        /// <summary>
        /// The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        /// </summary>
        public readonly ImmutableArray<string> OKActions;

        [OutputConstructor]
        private GetCompositeAlarmResult(
            bool? actionsEnabled,

            string? actionsSuppressor,

            int? actionsSuppressorExtensionPeriod,

            int? actionsSuppressorWaitPeriod,

            ImmutableArray<string> alarmActions,

            string? alarmDescription,

            string? alarmRule,

            string? arn,

            ImmutableArray<string> insufficientDataActions,

            ImmutableArray<string> oKActions)
        {
            ActionsEnabled = actionsEnabled;
            ActionsSuppressor = actionsSuppressor;
            ActionsSuppressorExtensionPeriod = actionsSuppressorExtensionPeriod;
            ActionsSuppressorWaitPeriod = actionsSuppressorWaitPeriod;
            AlarmActions = alarmActions;
            AlarmDescription = alarmDescription;
            AlarmRule = alarmRule;
            Arn = arn;
            InsufficientDataActions = insufficientDataActions;
            OKActions = oKActions;
        }
    }
}