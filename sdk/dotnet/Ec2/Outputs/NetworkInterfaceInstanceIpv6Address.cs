// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    [OutputType]
    public sealed class NetworkInterfaceInstanceIpv6Address
    {
        /// <summary>
        /// An IPv6 address to associate with the network interface.
        /// </summary>
        public readonly string Ipv6Address;

        [OutputConstructor]
        private NetworkInterfaceInstanceIpv6Address(string ipv6Address)
        {
            Ipv6Address = ipv6Address;
        }
    }
}
