// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTableRowConditionalFormattingArgs : global::Pulumi.ResourceArgs
    {
        [Input("backgroundColor")]
        public Input<Inputs.TemplateConditionalFormattingColorArgs>? BackgroundColor { get; set; }

        [Input("textColor")]
        public Input<Inputs.TemplateConditionalFormattingColorArgs>? TextColor { get; set; }

        public TemplateTableRowConditionalFormattingArgs()
        {
        }
        public static new TemplateTableRowConditionalFormattingArgs Empty => new TemplateTableRowConditionalFormattingArgs();
    }
}