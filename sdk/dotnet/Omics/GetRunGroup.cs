// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Omics
{
    public static class GetRunGroup
    {
        /// <summary>
        /// Definition of AWS::Omics::RunGroup Resource Type
        /// </summary>
        public static Task<GetRunGroupResult> InvokeAsync(GetRunGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRunGroupResult>("aws-native:omics:getRunGroup", args ?? new GetRunGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Omics::RunGroup Resource Type
        /// </summary>
        public static Output<GetRunGroupResult> Invoke(GetRunGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRunGroupResult>("aws-native:omics:getRunGroup", args ?? new GetRunGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Omics::RunGroup Resource Type
        /// </summary>
        public static Output<GetRunGroupResult> Invoke(GetRunGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRunGroupResult>("aws-native:omics:getRunGroup", args ?? new GetRunGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRunGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The run group's ID.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetRunGroupArgs()
        {
        }
        public static new GetRunGroupArgs Empty => new GetRunGroupArgs();
    }

    public sealed class GetRunGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The run group's ID.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetRunGroupInvokeArgs()
        {
        }
        public static new GetRunGroupInvokeArgs Empty => new GetRunGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetRunGroupResult
    {
        /// <summary>
        /// The run group's ARN.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// When the run group was created.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// The run group's ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The group's maximum CPU count setting.
        /// </summary>
        public readonly double? MaxCpus;
        /// <summary>
        /// The group's maximum duration setting in minutes.
        /// </summary>
        public readonly double? MaxDuration;
        /// <summary>
        /// The maximum GPUs that can be used by a run group.
        /// </summary>
        public readonly double? MaxGpus;
        /// <summary>
        /// The group's maximum concurrent run setting.
        /// </summary>
        public readonly double? MaxRuns;
        /// <summary>
        /// The group's name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Tags for the group.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetRunGroupResult(
            string? arn,

            string? creationTime,

            string? id,

            double? maxCpus,

            double? maxDuration,

            double? maxGpus,

            double? maxRuns,

            string? name,

            ImmutableDictionary<string, string>? tags)
        {
            Arn = arn;
            CreationTime = creationTime;
            Id = id;
            MaxCpus = maxCpus;
            MaxDuration = maxDuration;
            MaxGpus = maxGpus;
            MaxRuns = maxRuns;
            Name = name;
            Tags = tags;
        }
    }
}
