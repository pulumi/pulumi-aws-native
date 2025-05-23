// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardGaugeChartConditionalFormattingOption
    {
        /// <summary>
        /// The options that determine the presentation of the arc of a `GaugeChartVisual` .
        /// </summary>
        public readonly Outputs.DashboardGaugeChartArcConditionalFormatting? Arc;
        /// <summary>
        /// The conditional formatting for the primary value of a `GaugeChartVisual` .
        /// </summary>
        public readonly Outputs.DashboardGaugeChartPrimaryValueConditionalFormatting? PrimaryValue;

        [OutputConstructor]
        private DashboardGaugeChartConditionalFormattingOption(
            Outputs.DashboardGaugeChartArcConditionalFormatting? arc,

            Outputs.DashboardGaugeChartPrimaryValueConditionalFormatting? primaryValue)
        {
            Arc = arc;
            PrimaryValue = primaryValue;
        }
    }
}
