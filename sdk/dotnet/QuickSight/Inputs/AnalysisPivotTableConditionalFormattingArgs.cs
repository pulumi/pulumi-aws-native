// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisPivotTableConditionalFormattingArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditionalFormattingOptions")]
        private InputList<Inputs.AnalysisPivotTableConditionalFormattingOptionArgs>? _conditionalFormattingOptions;
        public InputList<Inputs.AnalysisPivotTableConditionalFormattingOptionArgs> ConditionalFormattingOptions
        {
            get => _conditionalFormattingOptions ?? (_conditionalFormattingOptions = new InputList<Inputs.AnalysisPivotTableConditionalFormattingOptionArgs>());
            set => _conditionalFormattingOptions = value;
        }

        public AnalysisPivotTableConditionalFormattingArgs()
        {
        }
        public static new AnalysisPivotTableConditionalFormattingArgs Empty => new AnalysisPivotTableConditionalFormattingArgs();
    }
}