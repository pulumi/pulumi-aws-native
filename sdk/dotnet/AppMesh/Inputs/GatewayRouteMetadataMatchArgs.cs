// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class GatewayRouteMetadataMatchArgs : global::Pulumi.ResourceArgs
    {
        [Input("exact")]
        public Input<string>? Exact { get; set; }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("range")]
        public Input<Inputs.GatewayRouteRangeMatchArgs>? Range { get; set; }

        [Input("regex")]
        public Input<string>? Regex { get; set; }

        [Input("suffix")]
        public Input<string>? Suffix { get; set; }

        public GatewayRouteMetadataMatchArgs()
        {
        }
        public static new GatewayRouteMetadataMatchArgs Empty => new GatewayRouteMetadataMatchArgs();
    }
}