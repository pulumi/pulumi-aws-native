// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    /// <summary>
    /// Transport stream service descriptor configuration for the Multiplex program.
    /// </summary>
    [OutputType]
    public sealed class MultiplexprogramMultiplexProgramServiceDescriptor
    {
        /// <summary>
        /// Name of the provider.
        /// </summary>
        public readonly string ProviderName;
        /// <summary>
        /// Name of the service.
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private MultiplexprogramMultiplexProgramServiceDescriptor(
            string providerName,

            string serviceName)
        {
            ProviderName = providerName;
            ServiceName = serviceName;
        }
    }
}