// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisConditionalFormattingColorArgs : global::Pulumi.ResourceArgs
    {
        [Input("gradient")]
        public Input<Inputs.AnalysisConditionalFormattingGradientColorArgs>? Gradient { get; set; }

        [Input("solid")]
        public Input<Inputs.AnalysisConditionalFormattingSolidColorArgs>? Solid { get; set; }

        public AnalysisConditionalFormattingColorArgs()
        {
        }
        public static new AnalysisConditionalFormattingColorArgs Empty => new AnalysisConditionalFormattingColorArgs();
    }
}