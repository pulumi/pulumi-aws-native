// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    [OutputType]
    public sealed class FeatureGroupThroughputConfig
    {
        /// <summary>
        /// For provisioned feature groups with online store enabled, this indicates the read throughput you are billed for and can consume without throttling.
        /// </summary>
        public readonly int? ProvisionedReadCapacityUnits;
        /// <summary>
        /// For provisioned feature groups, this indicates the write throughput you are billed for and can consume without throttling.
        /// </summary>
        public readonly int? ProvisionedWriteCapacityUnits;
        /// <summary>
        /// The mode used for your feature group throughput: `ON_DEMAND` or `PROVISIONED` .
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.FeatureGroupThroughputMode ThroughputMode;

        [OutputConstructor]
        private FeatureGroupThroughputConfig(
            int? provisionedReadCapacityUnits,

            int? provisionedWriteCapacityUnits,

            Pulumi.AwsNative.SageMaker.FeatureGroupThroughputMode throughputMode)
        {
            ProvisionedReadCapacityUnits = provisionedReadCapacityUnits;
            ProvisionedWriteCapacityUnits = provisionedWriteCapacityUnits;
            ThroughputMode = throughputMode;
        }
    }
}
