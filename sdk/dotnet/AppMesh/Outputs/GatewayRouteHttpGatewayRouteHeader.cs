// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    [OutputType]
    public sealed class GatewayRouteHttpGatewayRouteHeader
    {
        public readonly bool? Invert;
        public readonly Outputs.GatewayRouteHttpGatewayRouteHeaderMatch? Match;
        public readonly string Name;

        [OutputConstructor]
        private GatewayRouteHttpGatewayRouteHeader(
            bool? invert,

            Outputs.GatewayRouteHttpGatewayRouteHeaderMatch? match,

            string name)
        {
            Invert = invert;
            Match = match;
            Name = name;
        }
    }
}