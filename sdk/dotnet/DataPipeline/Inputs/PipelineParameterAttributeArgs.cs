// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataPipeline.Inputs
{

    public sealed class PipelineParameterAttributeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field identifier.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The field value, expressed as a String.
        /// </summary>
        [Input("stringValue", required: true)]
        public Input<string> StringValue { get; set; } = null!;

        public PipelineParameterAttributeArgs()
        {
        }
        public static new PipelineParameterAttributeArgs Empty => new PipelineParameterAttributeArgs();
    }
}