// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Inputs
{

    /// <summary>
    /// Todo: add description
    /// </summary>
    public sealed class ClusterComputeConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Todo: add description
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("nodePools")]
        private InputList<string>? _nodePools;

        /// <summary>
        /// Todo: add description
        /// </summary>
        public InputList<string> NodePools
        {
            get => _nodePools ?? (_nodePools = new InputList<string>());
            set => _nodePools = value;
        }

        /// <summary>
        /// Todo: add description
        /// </summary>
        [Input("nodeRoleArn")]
        public Input<string>? NodeRoleArn { get; set; }

        public ClusterComputeConfigArgs()
        {
        }
        public static new ClusterComputeConfigArgs Empty => new ClusterComputeConfigArgs();
    }
}
