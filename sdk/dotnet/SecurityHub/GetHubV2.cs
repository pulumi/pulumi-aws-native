// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub
{
    public static class GetHubV2
    {
        /// <summary>
        /// The AWS::SecurityHub::HubV2 resource represents the implementation of the AWS Security Hub V2 service in your account. Only one hubv2 resource can created in each region in which you enable Security Hub V2.
        /// </summary>
        public static Task<GetHubV2Result> InvokeAsync(GetHubV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHubV2Result>("aws-native:securityhub:getHubV2", args ?? new GetHubV2Args(), options.WithDefaults());

        /// <summary>
        /// The AWS::SecurityHub::HubV2 resource represents the implementation of the AWS Security Hub V2 service in your account. Only one hubv2 resource can created in each region in which you enable Security Hub V2.
        /// </summary>
        public static Output<GetHubV2Result> Invoke(GetHubV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHubV2Result>("aws-native:securityhub:getHubV2", args ?? new GetHubV2InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::SecurityHub::HubV2 resource represents the implementation of the AWS Security Hub V2 service in your account. Only one hubv2 resource can created in each region in which you enable Security Hub V2.
        /// </summary>
        public static Output<GetHubV2Result> Invoke(GetHubV2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetHubV2Result>("aws-native:securityhub:getHubV2", args ?? new GetHubV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHubV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name of the Security Hub V2 resource.
        /// </summary>
        [Input("hubV2Arn", required: true)]
        public string HubV2Arn { get; set; } = null!;

        public GetHubV2Args()
        {
        }
        public static new GetHubV2Args Empty => new GetHubV2Args();
    }

    public sealed class GetHubV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name of the Security Hub V2 resource.
        /// </summary>
        [Input("hubV2Arn", required: true)]
        public Input<string> HubV2Arn { get; set; } = null!;

        public GetHubV2InvokeArgs()
        {
        }
        public static new GetHubV2InvokeArgs Empty => new GetHubV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetHubV2Result
    {
        /// <summary>
        /// The Amazon Resource Name of the Security Hub V2 resource.
        /// </summary>
        public readonly string? HubV2Arn;
        /// <summary>
        /// The date and time when the service was enabled in the account.
        /// </summary>
        public readonly string? SubscribedAt;
        /// <summary>
        /// The tags to add to the hub V2 resource when you enable Security Hub.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetHubV2Result(
            string? hubV2Arn,

            string? subscribedAt,

            ImmutableDictionary<string, string>? tags)
        {
            HubV2Arn = hubV2Arn;
            SubscribedAt = subscribedAt;
            Tags = tags;
        }
    }
}
