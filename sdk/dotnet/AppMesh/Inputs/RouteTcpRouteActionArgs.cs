// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class RouteTcpRouteActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("weightedTargets", required: true)]
        private InputList<Inputs.RouteWeightedTargetArgs>? _weightedTargets;
        public InputList<Inputs.RouteWeightedTargetArgs> WeightedTargets
        {
            get => _weightedTargets ?? (_weightedTargets = new InputList<Inputs.RouteWeightedTargetArgs>());
            set => _weightedTargets = value;
        }

        public RouteTcpRouteActionArgs()
        {
        }
        public static new RouteTcpRouteActionArgs Empty => new RouteTcpRouteActionArgs();
    }
}