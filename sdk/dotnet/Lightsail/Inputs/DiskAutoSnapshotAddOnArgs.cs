// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// An object that represents additional parameters when enabling or modifying the automatic snapshot add-on
    /// </summary>
    public sealed class DiskAutoSnapshotAddOnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The daily time when an automatic snapshot will be created.
        /// </summary>
        [Input("snapshotTimeOfDay")]
        public Input<string>? SnapshotTimeOfDay { get; set; }

        public DiskAutoSnapshotAddOnArgs()
        {
        }
        public static new DiskAutoSnapshotAddOnArgs Empty => new DiskAutoSnapshotAddOnArgs();
    }
}