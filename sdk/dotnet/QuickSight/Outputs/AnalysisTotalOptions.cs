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
    public sealed class AnalysisTotalOptions
    {
        public readonly string? CustomLabel;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisTableTotalsPlacement? Placement;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisTableTotalsScrollStatus? ScrollStatus;
        public readonly Outputs.AnalysisTableCellStyle? TotalCellStyle;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? TotalsVisibility;

        [OutputConstructor]
        private AnalysisTotalOptions(
            string? customLabel,

            Pulumi.AwsNative.QuickSight.AnalysisTableTotalsPlacement? placement,

            Pulumi.AwsNative.QuickSight.AnalysisTableTotalsScrollStatus? scrollStatus,

            Outputs.AnalysisTableCellStyle? totalCellStyle,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? totalsVisibility)
        {
            CustomLabel = customLabel;
            Placement = placement;
            ScrollStatus = scrollStatus;
            TotalCellStyle = totalCellStyle;
            TotalsVisibility = totalsVisibility;
        }
    }
}