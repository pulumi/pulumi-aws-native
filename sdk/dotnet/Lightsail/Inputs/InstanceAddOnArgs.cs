// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// A addon associate with a resource.
    /// </summary>
    public sealed class InstanceAddOnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The add-on type
        /// </summary>
        [Input("addOnType", required: true)]
        public Input<string> AddOnType { get; set; } = null!;

        /// <summary>
        /// The parameters for the automatic snapshot add-on, such as the daily time when an automatic snapshot will be created.
        /// </summary>
        [Input("autoSnapshotAddOnRequest")]
        public Input<Inputs.InstanceAutoSnapshotAddOnArgs>? AutoSnapshotAddOnRequest { get; set; }

        /// <summary>
        /// Status of the Addon
        /// </summary>
        [Input("status")]
        public Input<Pulumi.AwsNative.Lightsail.InstanceAddOnStatus>? Status { get; set; }

        public InstanceAddOnArgs()
        {
        }
        public static new InstanceAddOnArgs Empty => new InstanceAddOnArgs();
    }
}
