// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    public static class GetSubscriptionDefinition
    {
        /// <summary>
        /// Resource Type definition for AWS::Greengrass::SubscriptionDefinition
        /// </summary>
        public static Task<GetSubscriptionDefinitionResult> InvokeAsync(GetSubscriptionDefinitionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionDefinitionResult>("aws-native:greengrass:getSubscriptionDefinition", args ?? new GetSubscriptionDefinitionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Greengrass::SubscriptionDefinition
        /// </summary>
        public static Output<GetSubscriptionDefinitionResult> Invoke(GetSubscriptionDefinitionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSubscriptionDefinitionResult>("aws-native:greengrass:getSubscriptionDefinition", args ?? new GetSubscriptionDefinitionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubscriptionDefinitionArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetSubscriptionDefinitionArgs()
        {
        }
        public static new GetSubscriptionDefinitionArgs Empty => new GetSubscriptionDefinitionArgs();
    }

    public sealed class GetSubscriptionDefinitionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetSubscriptionDefinitionInvokeArgs()
        {
        }
        public static new GetSubscriptionDefinitionInvokeArgs Empty => new GetSubscriptionDefinitionInvokeArgs();
    }


    [OutputType]
    public sealed class GetSubscriptionDefinitionResult
    {
        public readonly string? Arn;
        public readonly string? Id;
        public readonly string? LatestVersionArn;
        public readonly string? Name;
        public readonly object? Tags;

        [OutputConstructor]
        private GetSubscriptionDefinitionResult(
            string? arn,

            string? id,

            string? latestVersionArn,

            string? name,

            object? tags)
        {
            Arn = arn;
            Id = id;
            LatestVersionArn = latestVersionArn;
            Name = name;
            Tags = tags;
        }
    }
}