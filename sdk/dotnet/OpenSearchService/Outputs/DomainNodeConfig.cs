// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Outputs
{

    [OutputType]
    public sealed class DomainNodeConfig
    {
        /// <summary>
        /// The number of nodes of a particular node type in the cluster.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// A boolean that indicates whether a particular node type is enabled or not.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// The instance type of a particular node type in the cluster.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DomainNodeConfig(
            int? count,

            bool? enabled,

            string? type)
        {
            Count = count;
            Enabled = enabled;
            Type = type;
        }
    }
}
