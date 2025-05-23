// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// Details about the routing configuration for a Flow alias.
    /// </summary>
    [OutputType]
    public sealed class FlowAliasRoutingConfigurationListItem
    {
        /// <summary>
        /// Version.
        /// </summary>
        public readonly string? FlowVersion;

        [OutputConstructor]
        private FlowAliasRoutingConfigurationListItem(string? flowVersion)
        {
            FlowVersion = flowVersion;
        }
    }
}
