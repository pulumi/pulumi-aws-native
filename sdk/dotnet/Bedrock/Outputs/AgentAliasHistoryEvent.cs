// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// History event for an alias for an Agent.
    /// </summary>
    [OutputType]
    public sealed class AgentAliasHistoryEvent
    {
        /// <summary>
        /// Time Stamp.
        /// </summary>
        public readonly string? EndDate;
        /// <summary>
        /// Routing configuration for an Agent alias.
        /// </summary>
        public readonly ImmutableArray<Outputs.AgentAliasRoutingConfigurationListItem> RoutingConfiguration;
        /// <summary>
        /// Time Stamp.
        /// </summary>
        public readonly string? StartDate;

        [OutputConstructor]
        private AgentAliasHistoryEvent(
            string? endDate,

            ImmutableArray<Outputs.AgentAliasRoutingConfigurationListItem> routingConfiguration,

            string? startDate)
        {
            EndDate = endDate;
            RoutingConfiguration = routingConfiguration;
            StartDate = startDate;
        }
    }
}