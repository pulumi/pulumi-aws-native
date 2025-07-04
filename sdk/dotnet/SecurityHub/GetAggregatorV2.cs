// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub
{
    public static class GetAggregatorV2
    {
        /// <summary>
        /// The AWS::SecurityHub::AggregatorV2 resource represents the AWS Security Hub AggregatorV2 in your account. One aggregatorv2 resource is created for each account in non opt-in region in which you configure region linking mode.
        /// </summary>
        public static Task<GetAggregatorV2Result> InvokeAsync(GetAggregatorV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAggregatorV2Result>("aws-native:securityhub:getAggregatorV2", args ?? new GetAggregatorV2Args(), options.WithDefaults());

        /// <summary>
        /// The AWS::SecurityHub::AggregatorV2 resource represents the AWS Security Hub AggregatorV2 in your account. One aggregatorv2 resource is created for each account in non opt-in region in which you configure region linking mode.
        /// </summary>
        public static Output<GetAggregatorV2Result> Invoke(GetAggregatorV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAggregatorV2Result>("aws-native:securityhub:getAggregatorV2", args ?? new GetAggregatorV2InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::SecurityHub::AggregatorV2 resource represents the AWS Security Hub AggregatorV2 in your account. One aggregatorv2 resource is created for each account in non opt-in region in which you configure region linking mode.
        /// </summary>
        public static Output<GetAggregatorV2Result> Invoke(GetAggregatorV2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAggregatorV2Result>("aws-native:securityhub:getAggregatorV2", args ?? new GetAggregatorV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAggregatorV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the AggregatorV2 being created and assigned as the unique identifier
        /// </summary>
        [Input("aggregatorV2Arn", required: true)]
        public string AggregatorV2Arn { get; set; } = null!;

        public GetAggregatorV2Args()
        {
        }
        public static new GetAggregatorV2Args Empty => new GetAggregatorV2Args();
    }

    public sealed class GetAggregatorV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the AggregatorV2 being created and assigned as the unique identifier
        /// </summary>
        [Input("aggregatorV2Arn", required: true)]
        public Input<string> AggregatorV2Arn { get; set; } = null!;

        public GetAggregatorV2InvokeArgs()
        {
        }
        public static new GetAggregatorV2InvokeArgs Empty => new GetAggregatorV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetAggregatorV2Result
    {
        /// <summary>
        /// The aggregation Region of the AggregatorV2
        /// </summary>
        public readonly string? AggregationRegion;
        /// <summary>
        /// The ARN of the AggregatorV2 being created and assigned as the unique identifier
        /// </summary>
        public readonly string? AggregatorV2Arn;
        /// <summary>
        /// The list of included Regions
        /// </summary>
        public readonly ImmutableArray<string> LinkedRegions;
        /// <summary>
        /// Indicates to link a list of included Regions
        /// </summary>
        public readonly Pulumi.AwsNative.SecurityHub.AggregatorV2RegionLinkingMode? RegionLinkingMode;
        /// <summary>
        /// A list of key-value pairs to be applied to the AggregatorV2.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private GetAggregatorV2Result(
            string? aggregationRegion,

            string? aggregatorV2Arn,

            ImmutableArray<string> linkedRegions,

            Pulumi.AwsNative.SecurityHub.AggregatorV2RegionLinkingMode? regionLinkingMode,

            ImmutableDictionary<string, string>? tags)
        {
            AggregationRegion = aggregationRegion;
            AggregatorV2Arn = aggregatorV2Arn;
            LinkedRegions = linkedRegions;
            RegionLinkingMode = regionLinkingMode;
            Tags = tags;
        }
    }
}
