// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight
{
    public static class GetTopic
    {
        /// <summary>
        /// Definition of the AWS::QuickSight::Topic Resource Type.
        /// </summary>
        public static Task<GetTopicResult> InvokeAsync(GetTopicArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTopicResult>("aws-native:quicksight:getTopic", args ?? new GetTopicArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of the AWS::QuickSight::Topic Resource Type.
        /// </summary>
        public static Output<GetTopicResult> Invoke(GetTopicInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTopicResult>("aws-native:quicksight:getTopic", args ?? new GetTopicInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of the AWS::QuickSight::Topic Resource Type.
        /// </summary>
        public static Output<GetTopicResult> Invoke(GetTopicInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTopicResult>("aws-native:quicksight:getTopic", args ?? new GetTopicInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTopicArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the AWS account that you want to create a topic in.
        /// </summary>
        [Input("awsAccountId", required: true)]
        public string AwsAccountId { get; set; } = null!;

        /// <summary>
        /// The ID for the topic. This ID is unique per AWS Region for each AWS account.
        /// </summary>
        [Input("topicId", required: true)]
        public string TopicId { get; set; } = null!;

        public GetTopicArgs()
        {
        }
        public static new GetTopicArgs Empty => new GetTopicArgs();
    }

    public sealed class GetTopicInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the AWS account that you want to create a topic in.
        /// </summary>
        [Input("awsAccountId", required: true)]
        public Input<string> AwsAccountId { get; set; } = null!;

        /// <summary>
        /// The ID for the topic. This ID is unique per AWS Region for each AWS account.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public GetTopicInvokeArgs()
        {
        }
        public static new GetTopicInvokeArgs Empty => new GetTopicInvokeArgs();
    }


    [OutputType]
    public sealed class GetTopicResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the topic.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Configuration options for a `Topic` .
        /// </summary>
        public readonly Outputs.TopicConfigOptions? ConfigOptions;
        /// <summary>
        /// The data sets that the topic is associated with.
        /// </summary>
        public readonly ImmutableArray<Outputs.TopicDatasetMetadata> DataSets;
        /// <summary>
        /// The description of the topic.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The name of the topic.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The user experience version of the topic.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TopicUserExperienceVersion? UserExperienceVersion;

        [OutputConstructor]
        private GetTopicResult(
            string? arn,

            Outputs.TopicConfigOptions? configOptions,

            ImmutableArray<Outputs.TopicDatasetMetadata> dataSets,

            string? description,

            string? name,

            Pulumi.AwsNative.QuickSight.TopicUserExperienceVersion? userExperienceVersion)
        {
            Arn = arn;
            ConfigOptions = configOptions;
            DataSets = dataSets;
            Description = description;
            Name = name;
            UserExperienceVersion = userExperienceVersion;
        }
    }
}
