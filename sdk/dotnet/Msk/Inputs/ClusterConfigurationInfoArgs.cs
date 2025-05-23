// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    public sealed class ClusterConfigurationInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARN of the configuration to use.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        /// <summary>
        /// The revision of the configuration to use.
        /// </summary>
        [Input("revision", required: true)]
        public Input<int> Revision { get; set; } = null!;

        public ClusterConfigurationInfoArgs()
        {
        }
        public static new ClusterConfigurationInfoArgs Empty => new ClusterConfigurationInfoArgs();
    }
}
