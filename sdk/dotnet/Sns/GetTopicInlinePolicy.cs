// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sns
{
    public static class GetTopicInlinePolicy
    {
        /// <summary>
        /// Schema for AWS::SNS::TopicInlinePolicy
        /// </summary>
        public static Task<GetTopicInlinePolicyResult> InvokeAsync(GetTopicInlinePolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTopicInlinePolicyResult>("aws-native:sns:getTopicInlinePolicy", args ?? new GetTopicInlinePolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Schema for AWS::SNS::TopicInlinePolicy
        /// </summary>
        public static Output<GetTopicInlinePolicyResult> Invoke(GetTopicInlinePolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTopicInlinePolicyResult>("aws-native:sns:getTopicInlinePolicy", args ?? new GetTopicInlinePolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTopicInlinePolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the topic to which you want to add the policy.
        /// </summary>
        [Input("topicArn", required: true)]
        public string TopicArn { get; set; } = null!;

        public GetTopicInlinePolicyArgs()
        {
        }
        public static new GetTopicInlinePolicyArgs Empty => new GetTopicInlinePolicyArgs();
    }

    public sealed class GetTopicInlinePolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the topic to which you want to add the policy.
        /// </summary>
        [Input("topicArn", required: true)]
        public Input<string> TopicArn { get; set; } = null!;

        public GetTopicInlinePolicyInvokeArgs()
        {
        }
        public static new GetTopicInlinePolicyInvokeArgs Empty => new GetTopicInlinePolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetTopicInlinePolicyResult
    {
        /// <summary>
        /// A policy document that contains permissions to add to the specified SNS topics.
        /// </summary>
        public readonly object? PolicyDocument;

        [OutputConstructor]
        private GetTopicInlinePolicyResult(object? policyDocument)
        {
            PolicyDocument = policyDocument;
        }
    }
}