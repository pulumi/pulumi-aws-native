// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Scheduler
{
    /// <summary>
    /// Definition of AWS::Scheduler::Schedule Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:scheduler:Schedule")]
    public partial class Schedule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the schedule.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the schedule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The date, in UTC, before which the schedule can invoke its target. Depending on the schedule's recurrence expression, invocations might stop on, or before, the EndDate you specify.
        /// </summary>
        [Output("endDate")]
        public Output<string?> EndDate { get; private set; } = null!;

        /// <summary>
        /// Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.
        /// </summary>
        [Output("flexibleTimeWindow")]
        public Output<Outputs.ScheduleFlexibleTimeWindow> FlexibleTimeWindow { get; private set; } = null!;

        /// <summary>
        /// The name of the schedule group to associate with this schedule. If you omit this, the default schedule group is used.
        /// </summary>
        [Output("groupName")]
        public Output<string?> GroupName { get; private set; } = null!;

        /// <summary>
        /// The ARN for a KMS Key that will be used to encrypt customer data.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string?> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// The name of the schedule.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The scheduling expression.
        /// </summary>
        [Output("scheduleExpression")]
        public Output<string> ScheduleExpression { get; private set; } = null!;

        /// <summary>
        /// The timezone in which the scheduling expression is evaluated.
        /// </summary>
        [Output("scheduleExpressionTimezone")]
        public Output<string?> ScheduleExpressionTimezone { get; private set; } = null!;

        /// <summary>
        /// The date, in UTC, after which the schedule can begin invoking its target. Depending on the schedule's recurrence expression, invocations might occur on, or after, the StartDate you specify.
        /// </summary>
        [Output("startDate")]
        public Output<string?> StartDate { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the schedule is enabled or disabled.
        /// 
        /// *Allowed Values* : `ENABLED` | `DISABLED`
        /// </summary>
        [Output("state")]
        public Output<Pulumi.AwsNative.Scheduler.ScheduleState?> State { get; private set; } = null!;

        /// <summary>
        /// The schedule's target details.
        /// </summary>
        [Output("target")]
        public Output<Outputs.ScheduleTarget> Target { get; private set; } = null!;


        /// <summary>
        /// Create a Schedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Schedule(string name, ScheduleArgs args, CustomResourceOptions? options = null)
            : base("aws-native:scheduler:Schedule", name, args ?? new ScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Schedule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:scheduler:Schedule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Schedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Schedule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Schedule(name, id, options);
        }
    }

    public sealed class ScheduleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the schedule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The date, in UTC, before which the schedule can invoke its target. Depending on the schedule's recurrence expression, invocations might stop on, or before, the EndDate you specify.
        /// </summary>
        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        /// <summary>
        /// Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.
        /// </summary>
        [Input("flexibleTimeWindow", required: true)]
        public Input<Inputs.ScheduleFlexibleTimeWindowArgs> FlexibleTimeWindow { get; set; } = null!;

        /// <summary>
        /// The name of the schedule group to associate with this schedule. If you omit this, the default schedule group is used.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        /// <summary>
        /// The ARN for a KMS Key that will be used to encrypt customer data.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// The name of the schedule.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The scheduling expression.
        /// </summary>
        [Input("scheduleExpression", required: true)]
        public Input<string> ScheduleExpression { get; set; } = null!;

        /// <summary>
        /// The timezone in which the scheduling expression is evaluated.
        /// </summary>
        [Input("scheduleExpressionTimezone")]
        public Input<string>? ScheduleExpressionTimezone { get; set; }

        /// <summary>
        /// The date, in UTC, after which the schedule can begin invoking its target. Depending on the schedule's recurrence expression, invocations might occur on, or after, the StartDate you specify.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// Specifies whether the schedule is enabled or disabled.
        /// 
        /// *Allowed Values* : `ENABLED` | `DISABLED`
        /// </summary>
        [Input("state")]
        public Input<Pulumi.AwsNative.Scheduler.ScheduleState>? State { get; set; }

        /// <summary>
        /// The schedule's target details.
        /// </summary>
        [Input("target", required: true)]
        public Input<Inputs.ScheduleTargetArgs> Target { get; set; } = null!;

        public ScheduleArgs()
        {
        }
        public static new ScheduleArgs Empty => new ScheduleArgs();
    }
}
