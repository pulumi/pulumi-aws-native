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
    /// Advanced event selectors let you create fine-grained selectors for the following AWS CloudTrail event record ﬁelds. They help you control costs by logging only those events that are important to you.
    /// </summary>
    [OutputType]
    public sealed class EventDataStoreAdvancedEventSelector
    {
        /// <summary>
        /// Contains all selector statements in an advanced event selector.
        /// </summary>
        public readonly ImmutableArray<Outputs.EventDataStoreAdvancedFieldSelector> FieldSelectors;
        /// <summary>
        /// An optional, descriptive name for an advanced event selector, such as "Log data events for only two S3 buckets".
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private EventDataStoreAdvancedEventSelector(
            ImmutableArray<Outputs.EventDataStoreAdvancedFieldSelector> fieldSelectors,

            string? name)
        {
            FieldSelectors = fieldSelectors;
            Name = name;
        }
    }
}