// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class GatewayRouteVirtualServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("virtualServiceName", required: true)]
        public Input<string> VirtualServiceName { get; set; } = null!;

        public GatewayRouteVirtualServiceArgs()
        {
        }
        public static new GatewayRouteVirtualServiceArgs Empty => new GatewayRouteVirtualServiceArgs();
    }
}