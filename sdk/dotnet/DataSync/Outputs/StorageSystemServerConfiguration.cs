// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync.Outputs
{

    /// <summary>
    /// The server name and network port required to connect with the management interface of the on-premises storage system.
    /// </summary>
    [OutputType]
    public sealed class StorageSystemServerConfiguration
    {
        /// <summary>
        /// The domain name or IP address of the storage system's management interface.
        /// </summary>
        public readonly string ServerHostname;
        /// <summary>
        /// The network port needed to access the system's management interface
        /// </summary>
        public readonly int? ServerPort;

        [OutputConstructor]
        private StorageSystemServerConfiguration(
            string serverHostname,

            int? serverPort)
        {
            ServerHostname = serverHostname;
            ServerPort = serverPort;
        }
    }
}