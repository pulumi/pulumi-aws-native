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
    public sealed class AnalysisKPIConditionalFormattingOption
    {
        public readonly Outputs.AnalysisKPIPrimaryValueConditionalFormatting? PrimaryValue;
        public readonly Outputs.AnalysisKPIProgressBarConditionalFormatting? ProgressBar;

        [OutputConstructor]
        private AnalysisKPIConditionalFormattingOption(
            Outputs.AnalysisKPIPrimaryValueConditionalFormatting? primaryValue,

            Outputs.AnalysisKPIProgressBarConditionalFormatting? progressBar)
        {
            PrimaryValue = primaryValue;
            ProgressBar = progressBar;
        }
    }
}