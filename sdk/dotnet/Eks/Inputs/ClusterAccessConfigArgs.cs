// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Inputs
{

    /// <summary>
    /// An object representing the Access Config to use for the cluster.
    /// </summary>
    public sealed class ClusterAccessConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify the authentication mode that should be used to create your cluster.
        /// </summary>
        [Input("authenticationMode")]
        public Input<Pulumi.AwsNative.Eks.ClusterAccessConfigAuthenticationMode>? AuthenticationMode { get; set; }

        /// <summary>
        /// Set this value to false to avoid creating a default cluster admin Access Entry using the IAM principal used to create the cluster.
        /// </summary>
        [Input("bootstrapClusterCreatorAdminPermissions")]
        public Input<bool>? BootstrapClusterCreatorAdminPermissions { get; set; }

        public ClusterAccessConfigArgs()
        {
        }
        public static new ClusterAccessConfigArgs Empty => new ClusterAccessConfigArgs();
    }
}