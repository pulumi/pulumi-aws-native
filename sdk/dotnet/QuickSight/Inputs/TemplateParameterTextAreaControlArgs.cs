// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateParameterTextAreaControlArgs : global::Pulumi.ResourceArgs
    {
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        [Input("displayOptions")]
        public Input<Inputs.TemplateTextAreaControlDisplayOptionsArgs>? DisplayOptions { get; set; }

        [Input("parameterControlId", required: true)]
        public Input<string> ParameterControlId { get; set; } = null!;

        [Input("sourceParameterName", required: true)]
        public Input<string> SourceParameterName { get; set; } = null!;

        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public TemplateParameterTextAreaControlArgs()
        {
        }
        public static new TemplateParameterTextAreaControlArgs Empty => new TemplateParameterTextAreaControlArgs();
    }
}