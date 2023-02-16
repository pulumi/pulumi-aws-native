// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    public sealed class ClientVpnEndpointFederatedAuthenticationRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("sAMLProviderArn", required: true)]
        public Input<string> SAMLProviderArn { get; set; } = null!;

        [Input("selfServiceSAMLProviderArn")]
        public Input<string>? SelfServiceSAMLProviderArn { get; set; }

        public ClientVpnEndpointFederatedAuthenticationRequestArgs()
        {
        }
        public static new ClientVpnEndpointFederatedAuthenticationRequestArgs Empty => new ClientVpnEndpointFederatedAuthenticationRequestArgs();
    }
}