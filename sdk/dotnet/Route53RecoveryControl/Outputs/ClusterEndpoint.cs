// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53RecoveryControl.Outputs
{

    [OutputType]
    public sealed class ClusterEndpoint
    {
        /// <summary>
        /// A cluster endpoint URL for one of the five redundant clusters that you specify to set or retrieve a routing control state.
        /// </summary>
        public readonly string? Endpoint;
        /// <summary>
        /// The AWS Region for a cluster endpoint.
        /// </summary>
        public readonly string? Region;

        [OutputConstructor]
        private ClusterEndpoint(
            string? endpoint,

            string? region)
        {
            Endpoint = endpoint;
            Region = region;
        }
    }
}
