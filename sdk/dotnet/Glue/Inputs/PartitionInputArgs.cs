// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    public sealed class PartitionInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("parameters")]
        public Input<object>? Parameters { get; set; }

        [Input("storageDescriptor")]
        public Input<Inputs.PartitionStorageDescriptorArgs>? StorageDescriptor { get; set; }

        [Input("values", required: true)]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public PartitionInputArgs()
        {
        }
        public static new PartitionInputArgs Empty => new PartitionInputArgs();
    }
}