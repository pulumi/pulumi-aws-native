// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardKpiConditionalFormattingOption
    {
        public readonly Outputs.DashboardKpiActualValueConditionalFormatting? ActualValue;
        public readonly Outputs.DashboardKpiComparisonValueConditionalFormatting? ComparisonValue;
        public readonly Outputs.DashboardKpiPrimaryValueConditionalFormatting? PrimaryValue;
        public readonly Outputs.DashboardKpiProgressBarConditionalFormatting? ProgressBar;

        [OutputConstructor]
        private DashboardKpiConditionalFormattingOption(
            Outputs.DashboardKpiActualValueConditionalFormatting? actualValue,

            Outputs.DashboardKpiComparisonValueConditionalFormatting? comparisonValue,

            Outputs.DashboardKpiPrimaryValueConditionalFormatting? primaryValue,

            Outputs.DashboardKpiProgressBarConditionalFormatting? progressBar)
        {
            ActualValue = actualValue;
            ComparisonValue = comparisonValue;
            PrimaryValue = primaryValue;
            ProgressBar = progressBar;
        }
    }
}