// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class ClientVpnEndpointCertificateAuthenticationRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientRootCertificateChainArn", required: true)]
        public Input<string> ClientRootCertificateChainArn { get; set; } = null!;

        public ClientVpnEndpointCertificateAuthenticationRequestArgs()
        {
        }
        public static new ClientVpnEndpointCertificateAuthenticationRequestArgs Empty => new ClientVpnEndpointCertificateAuthenticationRequestArgs();
    }
}