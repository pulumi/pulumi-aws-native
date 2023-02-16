// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class PartitionInput
    {
        public readonly object? Parameters;
        public readonly Outputs.PartitionStorageDescriptor? StorageDescriptor;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private PartitionInput(
            object? parameters,

            Outputs.PartitionStorageDescriptor? storageDescriptor,

            ImmutableArray<string> values)
        {
            Parameters = parameters;
            StorageDescriptor = storageDescriptor;
            Values = values;
        }
    }
}