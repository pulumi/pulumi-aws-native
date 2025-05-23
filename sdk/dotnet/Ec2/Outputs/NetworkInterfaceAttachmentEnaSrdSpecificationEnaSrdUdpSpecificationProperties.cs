// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    /// <summary>
    /// Configures ENA Express for UDP network traffic.
    /// </summary>
    [OutputType]
    public sealed class NetworkInterfaceAttachmentEnaSrdSpecificationEnaSrdUdpSpecificationProperties
    {
        public readonly bool? EnaSrdUdpEnabled;

        [OutputConstructor]
        private NetworkInterfaceAttachmentEnaSrdSpecificationEnaSrdUdpSpecificationProperties(bool? enaSrdUdpEnabled)
        {
            EnaSrdUdpEnabled = enaSrdUdpEnabled;
        }
    }
}
