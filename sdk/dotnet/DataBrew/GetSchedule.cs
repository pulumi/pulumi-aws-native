// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew
{
    public static class GetSchedule
    {
        /// <summary>
        /// Resource schema for AWS::DataBrew::Schedule.
        /// </summary>
        public static Task<GetScheduleResult> InvokeAsync(GetScheduleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetScheduleResult>("aws-native:databrew:getSchedule", args ?? new GetScheduleArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataBrew::Schedule.
        /// </summary>
        public static Output<GetScheduleResult> Invoke(GetScheduleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetScheduleResult>("aws-native:databrew:getSchedule", args ?? new GetScheduleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataBrew::Schedule.
        /// </summary>
        public static Output<GetScheduleResult> Invoke(GetScheduleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetScheduleResult>("aws-native:databrew:getSchedule", args ?? new GetScheduleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetScheduleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Schedule Name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetScheduleArgs()
        {
        }
        public static new GetScheduleArgs Empty => new GetScheduleArgs();
    }

    public sealed class GetScheduleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Schedule Name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetScheduleInvokeArgs()
        {
        }
        public static new GetScheduleInvokeArgs Empty => new GetScheduleInvokeArgs();
    }


    [OutputType]
    public sealed class GetScheduleResult
    {
        /// <summary>
        /// Schedule cron
        /// </summary>
        public readonly string? CronExpression;
        /// <summary>
        /// A list of jobs to be run, according to the schedule.
        /// </summary>
        public readonly ImmutableArray<string> JobNames;
        /// <summary>
        /// Metadata tags that have been applied to the schedule.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetScheduleResult(
            string? cronExpression,

            ImmutableArray<string> jobNames,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            CronExpression = cronExpression;
            JobNames = jobNames;
            Tags = tags;
        }
    }
}
