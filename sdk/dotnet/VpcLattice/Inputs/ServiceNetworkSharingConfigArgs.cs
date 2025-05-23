// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice.Inputs
{

    public sealed class ServiceNetworkSharingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify if the service network should be enabled for sharing.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        public ServiceNetworkSharingConfigArgs()
        {
        }
        public static new ServiceNetworkSharingConfigArgs Empty => new ServiceNetworkSharingConfigArgs();
    }
}
