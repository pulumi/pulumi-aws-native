// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisKpiConditionalFormattingOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The conditional formatting for the actual value of a KPI visual.
        /// </summary>
        [Input("actualValue")]
        public Input<Inputs.AnalysisKpiActualValueConditionalFormattingArgs>? ActualValue { get; set; }

        /// <summary>
        /// The conditional formatting for the comparison value of a KPI visual.
        /// </summary>
        [Input("comparisonValue")]
        public Input<Inputs.AnalysisKpiComparisonValueConditionalFormattingArgs>? ComparisonValue { get; set; }

        /// <summary>
        /// The conditional formatting for the primary value of a KPI visual.
        /// </summary>
        [Input("primaryValue")]
        public Input<Inputs.AnalysisKpiPrimaryValueConditionalFormattingArgs>? PrimaryValue { get; set; }

        /// <summary>
        /// The conditional formatting for the progress bar of a KPI visual.
        /// </summary>
        [Input("progressBar")]
        public Input<Inputs.AnalysisKpiProgressBarConditionalFormattingArgs>? ProgressBar { get; set; }

        public AnalysisKpiConditionalFormattingOptionArgs()
        {
        }
        public static new AnalysisKpiConditionalFormattingOptionArgs Empty => new AnalysisKpiConditionalFormattingOptionArgs();
    }
}
