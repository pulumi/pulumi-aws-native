// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Account-level metrics configurations.
    /// </summary>
    public sealed class StorageLensAccountLevelArgs : global::Pulumi.ResourceArgs
    {
        [Input("activityMetrics")]
        public Input<Inputs.StorageLensActivityMetricsArgs>? ActivityMetrics { get; set; }

        [Input("advancedCostOptimizationMetrics")]
        public Input<Inputs.StorageLensAdvancedCostOptimizationMetricsArgs>? AdvancedCostOptimizationMetrics { get; set; }

        [Input("advancedDataProtectionMetrics")]
        public Input<Inputs.StorageLensAdvancedDataProtectionMetricsArgs>? AdvancedDataProtectionMetrics { get; set; }

        [Input("bucketLevel", required: true)]
        public Input<Inputs.StorageLensBucketLevelArgs> BucketLevel { get; set; } = null!;

        [Input("detailedStatusCodesMetrics")]
        public Input<Inputs.StorageLensDetailedStatusCodesMetricsArgs>? DetailedStatusCodesMetrics { get; set; }

        public StorageLensAccountLevelArgs()
        {
        }
        public static new StorageLensAccountLevelArgs Empty => new StorageLensAccountLevelArgs();
    }
}