// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTThingsGraph.Inputs
{

    public sealed class FlowTemplateDefinitionDocumentArgs : global::Pulumi.ResourceArgs
    {
        [Input("language", required: true)]
        public Input<string> Language { get; set; } = null!;

        [Input("text", required: true)]
        public Input<string> Text { get; set; } = null!;

        public FlowTemplateDefinitionDocumentArgs()
        {
        }
        public static new FlowTemplateDefinitionDocumentArgs Empty => new FlowTemplateDefinitionDocumentArgs();
    }
}