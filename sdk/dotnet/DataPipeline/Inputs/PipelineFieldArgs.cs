// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataPipeline.Inputs
{

    public sealed class PipelineFieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the name of a field for a particular object. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// A field value that you specify as an identifier of another object in the same pipeline definition.
        /// </summary>
        [Input("refValue")]
        public Input<string>? RefValue { get; set; }

        /// <summary>
        /// A field value that you specify as a string. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.
        /// </summary>
        [Input("stringValue")]
        public Input<string>? StringValue { get; set; }

        public PipelineFieldArgs()
        {
        }
        public static new PipelineFieldArgs Empty => new PipelineFieldArgs();
    }
}
