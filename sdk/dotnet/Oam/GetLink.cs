// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Oam
{
    public static class GetLink
    {
        /// <summary>
        /// Definition of AWS::Oam::Link Resource Type
        /// </summary>
        public static Task<GetLinkResult> InvokeAsync(GetLinkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLinkResult>("aws-native:oam:getLink", args ?? new GetLinkArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Oam::Link Resource Type
        /// </summary>
        public static Output<GetLinkResult> Invoke(GetLinkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLinkResult>("aws-native:oam:getLink", args ?? new GetLinkInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Oam::Link Resource Type
        /// </summary>
        public static Output<GetLinkResult> Invoke(GetLinkInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLinkResult>("aws-native:oam:getLink", args ?? new GetLinkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLinkArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the link. For example, `arn:aws:oam:us-west-1:111111111111:link:abcd1234-a123-456a-a12b-a123b456c789`
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetLinkArgs()
        {
        }
        public static new GetLinkArgs Empty => new GetLinkArgs();
    }

    public sealed class GetLinkInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the link. For example, `arn:aws:oam:us-west-1:111111111111:link:abcd1234-a123-456a-a12b-a123b456c789`
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetLinkInvokeArgs()
        {
        }
        public static new GetLinkInvokeArgs Empty => new GetLinkInvokeArgs();
    }


    [OutputType]
    public sealed class GetLinkResult
    {
        /// <summary>
        /// The ARN of the link. For example, `arn:aws:oam:us-west-1:111111111111:link:abcd1234-a123-456a-a12b-a123b456c789`
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The friendly human-readable name used to identify this source account when it is viewed from the monitoring account. For example, `my-account1` .
        /// </summary>
        public readonly string? Label;
        /// <summary>
        /// Use this structure to optionally create filters that specify that only some metric namespaces or log groups are to be shared from the source account to the monitoring account.
        /// </summary>
        public readonly Outputs.LinkConfiguration? LinkConfiguration;
        /// <summary>
        /// An array of strings that define which types of data that the source account shares with the monitoring account. Valid values are `AWS::CloudWatch::Metric | AWS::Logs::LogGroup | AWS::XRay::Trace | AWS::ApplicationInsights::Application | AWS::InternetMonitor::Monitor` .
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Oam.LinkResourceType> ResourceTypes;
        /// <summary>
        /// Tags to apply to the link
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetLinkResult(
            string? arn,

            string? label,

            Outputs.LinkConfiguration? linkConfiguration,

            ImmutableArray<Pulumi.AwsNative.Oam.LinkResourceType> resourceTypes,

            ImmutableDictionary<string, string>? tags)
        {
            Arn = arn;
            Label = label;
            LinkConfiguration = linkConfiguration;
            ResourceTypes = resourceTypes;
            Tags = tags;
        }
    }
}
