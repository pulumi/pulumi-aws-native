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
    public sealed class AnalysisTableConditionalFormatting
    {
        /// <summary>
        /// Conditional formatting options for a `PivotTableVisual` .
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisTableConditionalFormattingOption> ConditionalFormattingOptions;

        [OutputConstructor]
        private AnalysisTableConditionalFormatting(ImmutableArray<Outputs.AnalysisTableConditionalFormattingOption> conditionalFormattingOptions)
        {
            ConditionalFormattingOptions = conditionalFormattingOptions;
        }
    }
}
