// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    /// <summary>
    /// Tags to be applied to Input.
    /// </summary>
    public sealed class DetectorModelTagArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key of the Tag.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Value of the Tag.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public DetectorModelTagArgs()
        {
        }
        public static new DetectorModelTagArgs Empty => new DetectorModelTagArgs();
    }
}