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
    public sealed class RouteGrpcRouteMatch
    {
        public readonly ImmutableArray<Outputs.RouteGrpcRouteMetadata> Metadata;
        public readonly string? MethodName;
        public readonly int? Port;
        public readonly string? ServiceName;

        [OutputConstructor]
        private RouteGrpcRouteMatch(
            ImmutableArray<Outputs.RouteGrpcRouteMetadata> metadata,

            string? methodName,

            int? port,

            string? serviceName)
        {
            Metadata = metadata;
            MethodName = methodName;
            Port = port;
            ServiceName = serviceName;
        }
    }
}