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
    /// Specifies the encryption algorithm for the VPN tunnel for phase 2 IKE negotiations.
    /// </summary>
    [OutputType]
    public sealed class VpnConnectionPhase2EncryptionAlgorithmsRequestListValue
    {
        /// <summary>
        /// The encryption algorithm.
        /// </summary>
        public readonly Pulumi.AwsNative.Ec2.VpnConnectionPhase2EncryptionAlgorithmsRequestListValueValue? Value;

        [OutputConstructor]
        private VpnConnectionPhase2EncryptionAlgorithmsRequestListValue(Pulumi.AwsNative.Ec2.VpnConnectionPhase2EncryptionAlgorithmsRequestListValueValue? value)
        {
            Value = value;
        }
    }
}
