// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents
{
    public static class GetAlarmModel
    {
        /// <summary>
        /// Represents an alarm model to monitor an ITE input attribute. You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *Developer Guide*.
        /// </summary>
        public static Task<GetAlarmModelResult> InvokeAsync(GetAlarmModelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlarmModelResult>("aws-native:iotevents:getAlarmModel", args ?? new GetAlarmModelArgs(), options.WithDefaults());

        /// <summary>
        /// Represents an alarm model to monitor an ITE input attribute. You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *Developer Guide*.
        /// </summary>
        public static Output<GetAlarmModelResult> Invoke(GetAlarmModelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlarmModelResult>("aws-native:iotevents:getAlarmModel", args ?? new GetAlarmModelInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Represents an alarm model to monitor an ITE input attribute. You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *Developer Guide*.
        /// </summary>
        public static Output<GetAlarmModelResult> Invoke(GetAlarmModelInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlarmModelResult>("aws-native:iotevents:getAlarmModel", args ?? new GetAlarmModelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlarmModelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the alarm model.
        /// </summary>
        [Input("alarmModelName", required: true)]
        public string AlarmModelName { get; set; } = null!;

        public GetAlarmModelArgs()
        {
        }
        public static new GetAlarmModelArgs Empty => new GetAlarmModelArgs();
    }

    public sealed class GetAlarmModelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the alarm model.
        /// </summary>
        [Input("alarmModelName", required: true)]
        public Input<string> AlarmModelName { get; set; } = null!;

        public GetAlarmModelInvokeArgs()
        {
        }
        public static new GetAlarmModelInvokeArgs Empty => new GetAlarmModelInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlarmModelResult
    {
        /// <summary>
        /// Contains the configuration information of alarm state changes.
        /// </summary>
        public readonly Outputs.AlarmModelAlarmCapabilities? AlarmCapabilities;
        /// <summary>
        /// Contains information about one or more alarm actions.
        /// </summary>
        public readonly Outputs.AlarmModelAlarmEventActions? AlarmEventActions;
        /// <summary>
        /// The description of the alarm model.
        /// </summary>
        public readonly string? AlarmModelDescription;
        /// <summary>
        /// Defines when your alarm is invoked.
        /// </summary>
        public readonly Outputs.AlarmModelAlarmRule? AlarmRule;
        /// <summary>
        /// The ARN of the IAM role that allows the alarm to perform actions and access AWS resources. For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// A non-negative integer that reflects the severity level of the alarm.
        /// </summary>
        public readonly int? Severity;
        /// <summary>
        /// A list of key-value pairs that contain metadata for the alarm model. The tags help you manage the alarm model. For more information, see [Tagging your resources](https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html) in the *Developer Guide*.
        ///  You can create up to 50 tags for one alarm model.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetAlarmModelResult(
            Outputs.AlarmModelAlarmCapabilities? alarmCapabilities,

            Outputs.AlarmModelAlarmEventActions? alarmEventActions,

            string? alarmModelDescription,

            Outputs.AlarmModelAlarmRule? alarmRule,

            string? roleArn,

            int? severity,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AlarmCapabilities = alarmCapabilities;
            AlarmEventActions = alarmEventActions;
            AlarmModelDescription = alarmModelDescription;
            AlarmRule = alarmRule;
            RoleArn = roleArn;
            Severity = severity;
            Tags = tags;
        }
    }
}
