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
    /// Configuration for Session Summarization
    /// </summary>
    [OutputType]
    public sealed class AgentSessionSummaryConfiguration
    {
        /// <summary>
        /// Maximum number of Sessions to Summarize
        /// </summary>
        public readonly double? MaxRecentSessions;

        [OutputConstructor]
        private AgentSessionSummaryConfiguration(double? maxRecentSessions)
        {
            MaxRecentSessions = maxRecentSessions;
        }
    }
}
