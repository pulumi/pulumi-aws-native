// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager.Outputs
{

    /// <summary>
    /// Connect attachment options for protocol
    /// </summary>
    [OutputType]
    public sealed class ConnectAttachmentOptions
    {
        /// <summary>
        /// Tunnel protocol for connect attachment
        /// </summary>
        public readonly string? Protocol;

        [OutputConstructor]
        private ConnectAttachmentOptions(string? protocol)
        {
            Protocol = protocol;
        }
    }
}
