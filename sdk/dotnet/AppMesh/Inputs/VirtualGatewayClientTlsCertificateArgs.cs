// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class VirtualGatewayClientTlsCertificateArgs : global::Pulumi.ResourceArgs
    {
        [Input("file")]
        public Input<Inputs.VirtualGatewayListenerTlsFileCertificateArgs>? File { get; set; }

        [Input("sDS")]
        public Input<Inputs.VirtualGatewayListenerTlsSdsCertificateArgs>? SDS { get; set; }

        public VirtualGatewayClientTlsCertificateArgs()
        {
        }
        public static new VirtualGatewayClientTlsCertificateArgs Empty => new VirtualGatewayClientTlsCertificateArgs();
    }
}