// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudTrail.Outputs
{

    /// <summary>
    /// The type of email sending events to publish to the event destination.
    /// </summary>
    [OutputType]
    public sealed class TrailEventSelector
    {
        public readonly ImmutableArray<Outputs.TrailDataResource> DataResources;
        /// <summary>
        /// An optional list of service event sources from which you do not want management events to be logged on your trail. In this release, the list can be empty (disables the filter), or it can filter out AWS Key Management Service events by containing "kms.amazonaws.com". By default, ExcludeManagementEventSources is empty, and AWS KMS events are included in events that are logged to your trail.
        /// </summary>
        public readonly ImmutableArray<string> ExcludeManagementEventSources;
        /// <summary>
        /// Specify if you want your event selector to include management events for your trail.
        /// </summary>
        public readonly bool? IncludeManagementEvents;
        /// <summary>
        /// Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.
        /// </summary>
        public readonly Pulumi.AwsNative.CloudTrail.TrailEventSelectorReadWriteType? ReadWriteType;

        [OutputConstructor]
        private TrailEventSelector(
            ImmutableArray<Outputs.TrailDataResource> dataResources,

            ImmutableArray<string> excludeManagementEventSources,

            bool? includeManagementEvents,

            Pulumi.AwsNative.CloudTrail.TrailEventSelectorReadWriteType? readWriteType)
        {
            DataResources = dataResources;
            ExcludeManagementEventSources = excludeManagementEventSources;
            IncludeManagementEvents = includeManagementEvents;
            ReadWriteType = readWriteType;
        }
    }
}