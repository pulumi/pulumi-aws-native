// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync.Outputs
{

    /// <summary>
    /// Configuration information for RPC Protection and Data Transfer Protection. These parameters can be set to AUTHENTICATION, INTEGRITY, or PRIVACY. The default value is PRIVACY.
    /// </summary>
    [OutputType]
    public sealed class LocationHdfsQopConfiguration
    {
        /// <summary>
        /// Configuration for Data Transfer Protection.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.LocationHdfsQopConfigurationDataTransferProtection? DataTransferProtection;
        /// <summary>
        /// Configuration for RPC Protection.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.LocationHdfsQopConfigurationRpcProtection? RpcProtection;

        [OutputConstructor]
        private LocationHdfsQopConfiguration(
            Pulumi.AwsNative.DataSync.LocationHdfsQopConfigurationDataTransferProtection? dataTransferProtection,

            Pulumi.AwsNative.DataSync.LocationHdfsQopConfigurationRpcProtection? rpcProtection)
        {
            DataTransferProtection = dataTransferProtection;
            RpcProtection = rpcProtection;
        }
    }
}
