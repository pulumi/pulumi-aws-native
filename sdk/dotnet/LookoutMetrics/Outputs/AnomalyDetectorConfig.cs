// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Outputs
{

    [OutputType]
    public sealed class AnomalyDetectorConfig
    {
        /// <summary>
        /// Frequency of anomaly detection
        /// </summary>
        public readonly Pulumi.AwsNative.LookoutMetrics.AnomalyDetectorFrequency AnomalyDetectorFrequency;

        [OutputConstructor]
        private AnomalyDetectorConfig(Pulumi.AwsNative.LookoutMetrics.AnomalyDetectorFrequency anomalyDetectorFrequency)
        {
            AnomalyDetectorFrequency = anomalyDetectorFrequency;
        }
    }
}
