// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RedshiftServerless.Inputs
{

    public sealed class NamespaceSnapshotCopyConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationKmsKeyId")]
        public Input<string>? DestinationKmsKeyId { get; set; }

        [Input("destinationRegion", required: true)]
        public Input<string> DestinationRegion { get; set; } = null!;

        [Input("snapshotRetentionPeriod")]
        public Input<int>? SnapshotRetentionPeriod { get; set; }

        public NamespaceSnapshotCopyConfigurationArgs()
        {
        }
        public static new NamespaceSnapshotCopyConfigurationArgs Empty => new NamespaceSnapshotCopyConfigurationArgs();
    }
}