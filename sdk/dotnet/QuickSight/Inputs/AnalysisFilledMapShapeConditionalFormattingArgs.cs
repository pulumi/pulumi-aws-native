// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisFilledMapShapeConditionalFormattingArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        [Input("format")]
        public Input<Inputs.AnalysisShapeConditionalFormatArgs>? Format { get; set; }

        public AnalysisFilledMapShapeConditionalFormattingArgs()
        {
        }
        public static new AnalysisFilledMapShapeConditionalFormattingArgs Empty => new AnalysisFilledMapShapeConditionalFormattingArgs();
    }
}