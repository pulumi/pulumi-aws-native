// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub
{
    public static class GetFindingAggregator
    {
        /// <summary>
        /// The ``AWS::SecurityHub::FindingAggregator`` resource enables cross-Region aggregation. When cross-Region aggregation is enabled, you can aggregate findings, finding updates, insights, control compliance statuses, and security scores from one or more linked Regions to a single aggregation Region. You can then view and manage all of this data from the aggregation Region. For more details about cross-Region aggregation, see [Cross-Region aggregation](https://docs.aws.amazon.com/securityhub/latest/userguide/finding-aggregation.html) in the *User Guide*
        ///  This resource must be created in the Region that you want to designate as your aggregation Region.
        ///  Cross-Region aggregation is also a prerequisite for using [central configuration](https://docs.aws.amazon.com/securityhub/latest/userguide/central-configuration-intro.html) in ASH.
        /// </summary>
        public static Task<GetFindingAggregatorResult> InvokeAsync(GetFindingAggregatorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFindingAggregatorResult>("aws-native:securityhub:getFindingAggregator", args ?? new GetFindingAggregatorArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::SecurityHub::FindingAggregator`` resource enables cross-Region aggregation. When cross-Region aggregation is enabled, you can aggregate findings, finding updates, insights, control compliance statuses, and security scores from one or more linked Regions to a single aggregation Region. You can then view and manage all of this data from the aggregation Region. For more details about cross-Region aggregation, see [Cross-Region aggregation](https://docs.aws.amazon.com/securityhub/latest/userguide/finding-aggregation.html) in the *User Guide*
        ///  This resource must be created in the Region that you want to designate as your aggregation Region.
        ///  Cross-Region aggregation is also a prerequisite for using [central configuration](https://docs.aws.amazon.com/securityhub/latest/userguide/central-configuration-intro.html) in ASH.
        /// </summary>
        public static Output<GetFindingAggregatorResult> Invoke(GetFindingAggregatorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFindingAggregatorResult>("aws-native:securityhub:getFindingAggregator", args ?? new GetFindingAggregatorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::SecurityHub::FindingAggregator`` resource enables cross-Region aggregation. When cross-Region aggregation is enabled, you can aggregate findings, finding updates, insights, control compliance statuses, and security scores from one or more linked Regions to a single aggregation Region. You can then view and manage all of this data from the aggregation Region. For more details about cross-Region aggregation, see [Cross-Region aggregation](https://docs.aws.amazon.com/securityhub/latest/userguide/finding-aggregation.html) in the *User Guide*
        ///  This resource must be created in the Region that you want to designate as your aggregation Region.
        ///  Cross-Region aggregation is also a prerequisite for using [central configuration](https://docs.aws.amazon.com/securityhub/latest/userguide/central-configuration-intro.html) in ASH.
        /// </summary>
        public static Output<GetFindingAggregatorResult> Invoke(GetFindingAggregatorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetFindingAggregatorResult>("aws-native:securityhub:getFindingAggregator", args ?? new GetFindingAggregatorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFindingAggregatorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the finding aggregator. You use the finding aggregator ARN to retrieve details for, update, and delete the finding aggregator.
        /// </summary>
        [Input("findingAggregatorArn", required: true)]
        public string FindingAggregatorArn { get; set; } = null!;

        public GetFindingAggregatorArgs()
        {
        }
        public static new GetFindingAggregatorArgs Empty => new GetFindingAggregatorArgs();
    }

    public sealed class GetFindingAggregatorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the finding aggregator. You use the finding aggregator ARN to retrieve details for, update, and delete the finding aggregator.
        /// </summary>
        [Input("findingAggregatorArn", required: true)]
        public Input<string> FindingAggregatorArn { get; set; } = null!;

        public GetFindingAggregatorInvokeArgs()
        {
        }
        public static new GetFindingAggregatorInvokeArgs Empty => new GetFindingAggregatorInvokeArgs();
    }


    [OutputType]
    public sealed class GetFindingAggregatorResult
    {
        /// <summary>
        /// The home Region. Findings generated in linked Regions are replicated and sent to the home Region.
        /// </summary>
        public readonly string? FindingAggregationRegion;
        /// <summary>
        /// The ARN of the finding aggregator. You use the finding aggregator ARN to retrieve details for, update, and delete the finding aggregator.
        /// </summary>
        public readonly string? FindingAggregatorArn;
        /// <summary>
        /// Indicates whether to aggregate findings from all of the available Regions in the current partition. Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them.
        ///  The selected option also determines how to use the Regions provided in the Regions list.
        ///  In CFN, the options for this property are as follows:
        ///   +  ``ALL_REGIONS`` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them. 
        ///   +  ``ALL_REGIONS_EXCEPT_SPECIFIED`` - Indicates to aggregate findings from all of the Regions where Security Hub is enabled, except for the Regions listed in the ``Regions`` parameter. When you choose this option, Security Hub also automatically aggregates findings from new Regions as Security Hub supports them and you opt into them. 
        ///   +  ``SPECIFIED_REGIONS`` - Indicates to aggregate findings only from the Regions listed in the ``Regions`` parameter. Security Hub does not automatically aggregate findings from new Regions.
        /// </summary>
        public readonly Pulumi.AwsNative.SecurityHub.FindingAggregatorRegionLinkingMode? RegionLinkingMode;
        /// <summary>
        /// If ``RegionLinkingMode`` is ``ALL_REGIONS_EXCEPT_SPECIFIED``, then this is a space-separated list of Regions that do not aggregate findings to the aggregation Region.
        ///  If ``RegionLinkingMode`` is ``SPECIFIED_REGIONS``, then this is a space-separated list of Regions that do aggregate findings to the aggregation Region.
        /// </summary>
        public readonly ImmutableArray<string> Regions;

        [OutputConstructor]
        private GetFindingAggregatorResult(
            string? findingAggregationRegion,

            string? findingAggregatorArn,

            Pulumi.AwsNative.SecurityHub.FindingAggregatorRegionLinkingMode? regionLinkingMode,

            ImmutableArray<string> regions)
        {
            FindingAggregationRegion = findingAggregationRegion;
            FindingAggregatorArn = findingAggregatorArn;
            RegionLinkingMode = regionLinkingMode;
            Regions = regions;
        }
    }
}
