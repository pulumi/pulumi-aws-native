// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class RouteTcpRouteArgs : global::Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<Inputs.RouteTcpRouteActionArgs> Action { get; set; } = null!;

        [Input("match")]
        public Input<Inputs.RouteTcpRouteMatchArgs>? Match { get; set; }

        [Input("timeout")]
        public Input<Inputs.RouteTcpTimeoutArgs>? Timeout { get; set; }

        public RouteTcpRouteArgs()
        {
        }
        public static new RouteTcpRouteArgs Empty => new RouteTcpRouteArgs();
    }
}