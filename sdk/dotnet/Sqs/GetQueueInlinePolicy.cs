// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sqs
{
    public static class GetQueueInlinePolicy
    {
        /// <summary>
        /// Schema for SQS QueueInlinePolicy
        /// </summary>
        public static Task<GetQueueInlinePolicyResult> InvokeAsync(GetQueueInlinePolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQueueInlinePolicyResult>("aws-native:sqs:getQueueInlinePolicy", args ?? new GetQueueInlinePolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Schema for SQS QueueInlinePolicy
        /// </summary>
        public static Output<GetQueueInlinePolicyResult> Invoke(GetQueueInlinePolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQueueInlinePolicyResult>("aws-native:sqs:getQueueInlinePolicy", args ?? new GetQueueInlinePolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQueueInlinePolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The URL of the SQS queue.
        /// </summary>
        [Input("queue", required: true)]
        public string Queue { get; set; } = null!;

        public GetQueueInlinePolicyArgs()
        {
        }
        public static new GetQueueInlinePolicyArgs Empty => new GetQueueInlinePolicyArgs();
    }

    public sealed class GetQueueInlinePolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The URL of the SQS queue.
        /// </summary>
        [Input("queue", required: true)]
        public Input<string> Queue { get; set; } = null!;

        public GetQueueInlinePolicyInvokeArgs()
        {
        }
        public static new GetQueueInlinePolicyInvokeArgs Empty => new GetQueueInlinePolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetQueueInlinePolicyResult
    {
        /// <summary>
        /// A policy document that contains permissions to add to the specified SQS queue
        /// </summary>
        public readonly object? PolicyDocument;

        [OutputConstructor]
        private GetQueueInlinePolicyResult(object? policyDocument)
        {
            PolicyDocument = policyDocument;
        }
    }
}