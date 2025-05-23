// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaConnect.Inputs
{

    /// <summary>
    /// The priority you want to assign to a source. You can have a primary stream and a backup stream or two equally prioritized streams.
    /// </summary>
    public sealed class BridgeSourcePriorityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the source you choose as the primary source for this flow.
        /// </summary>
        [Input("primarySource")]
        public Input<string>? PrimarySource { get; set; }

        public BridgeSourcePriorityArgs()
        {
        }
        public static new BridgeSourcePriorityArgs Empty => new BridgeSourcePriorityArgs();
    }
}
