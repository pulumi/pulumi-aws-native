// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataPipeline.Inputs
{

    public sealed class PipelineObjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("fields", required: true)]
        private InputList<Inputs.PipelineFieldArgs>? _fields;

        /// <summary>
        /// Key-value pairs that define the properties of the object.
        /// </summary>
        public InputList<Inputs.PipelineFieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.PipelineFieldArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// The ID of the object.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public PipelineObjectArgs()
        {
        }
        public static new PipelineObjectArgs Empty => new PipelineObjectArgs();
    }
}