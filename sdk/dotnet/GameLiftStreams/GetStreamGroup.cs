// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLiftStreams
{
    public static class GetStreamGroup
    {
        /// <summary>
        /// Definition of AWS::GameLiftStreams::StreamGroup Resource Type
        /// </summary>
        public static Task<GetStreamGroupResult> InvokeAsync(GetStreamGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStreamGroupResult>("aws-native:gameliftstreams:getStreamGroup", args ?? new GetStreamGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::GameLiftStreams::StreamGroup Resource Type
        /// </summary>
        public static Output<GetStreamGroupResult> Invoke(GetStreamGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStreamGroupResult>("aws-native:gameliftstreams:getStreamGroup", args ?? new GetStreamGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::GameLiftStreams::StreamGroup Resource Type
        /// </summary>
        public static Output<GetStreamGroupResult> Invoke(GetStreamGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetStreamGroupResult>("aws-native:gameliftstreams:getStreamGroup", args ?? new GetStreamGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStreamGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the stream group resource. For example: `arn:aws:gameliftstreams:us-west-2:123456789012:streamgroup/sg-1AB2C3De4` .
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetStreamGroupArgs()
        {
        }
        public static new GetStreamGroupArgs Empty => new GetStreamGroupArgs();
    }

    public sealed class GetStreamGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the stream group resource. For example: `arn:aws:gameliftstreams:us-west-2:123456789012:streamgroup/sg-1AB2C3De4` .
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetStreamGroupInvokeArgs()
        {
        }
        public static new GetStreamGroupInvokeArgs Empty => new GetStreamGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetStreamGroupResult
    {
        /// <summary>
        /// An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely identifies the stream group resource. For example: `arn:aws:gameliftstreams:us-west-2:123456789012:streamgroup/sg-1AB2C3De4` .
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Object that identifies the Amazon GameLift Streams application to stream with this stream group.
        /// </summary>
        public readonly Outputs.StreamGroupDefaultApplication? DefaultApplication;
        /// <summary>
        /// A descriptive label for the stream group.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// An ID that uniquely identifies the stream group resource. For example: `sg-1AB2C3De4` .
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// A set of one or more locations and the streaming capacity for each location. One of the locations MUST be your primary location, which is the AWS Region where you are specifying this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.StreamGroupLocationConfiguration> LocationConfigurations;
        /// <summary>
        /// A list of labels to assign to the new stream group resource. Tags are developer-defined key-value pairs. Tagging AWS resources is useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* .
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetStreamGroupResult(
            string? arn,

            Outputs.StreamGroupDefaultApplication? defaultApplication,

            string? description,

            string? id,

            ImmutableArray<Outputs.StreamGroupLocationConfiguration> locationConfigurations,

            ImmutableDictionary<string, string>? tags)
        {
            Arn = arn;
            DefaultApplication = defaultApplication;
            Description = description;
            Id = id;
            LocationConfigurations = locationConfigurations;
            Tags = tags;
        }
    }
}
