// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudTrail.Outputs
{

    /// <summary>
    /// A string that contains Insights types that are logged on an event data store.
    /// </summary>
    [OutputType]
    public sealed class EventDataStoreInsightSelector
    {
        /// <summary>
        /// The type of Insights to log on an event data store.
        /// </summary>
        public readonly string? InsightType;

        [OutputConstructor]
        private EventDataStoreInsightSelector(string? insightType)
        {
            InsightType = insightType;
        }
    }
}
