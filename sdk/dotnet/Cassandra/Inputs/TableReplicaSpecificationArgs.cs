// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cassandra.Inputs
{

    /// <summary>
    /// Represents replica specifications.
    /// </summary>
    public sealed class TableReplicaSpecificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("readCapacityAutoScaling")]
        public Input<Inputs.TableAutoScalingSettingArgs>? ReadCapacityAutoScaling { get; set; }

        [Input("readCapacityUnits")]
        public Input<int>? ReadCapacityUnits { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public TableReplicaSpecificationArgs()
        {
        }
        public static new TableReplicaSpecificationArgs Empty => new TableReplicaSpecificationArgs();
    }
}