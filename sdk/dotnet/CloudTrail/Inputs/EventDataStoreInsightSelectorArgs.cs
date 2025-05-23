// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudTrail.Inputs
{

    /// <summary>
    /// A string that contains Insights types that are logged on an event data store.
    /// </summary>
    public sealed class EventDataStoreInsightSelectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of Insights to log on an event data store.
        /// </summary>
        [Input("insightType")]
        public Input<string>? InsightType { get; set; }

        public EventDataStoreInsightSelectorArgs()
        {
        }
        public static new EventDataStoreInsightSelectorArgs Empty => new EventDataStoreInsightSelectorArgs();
    }
}
