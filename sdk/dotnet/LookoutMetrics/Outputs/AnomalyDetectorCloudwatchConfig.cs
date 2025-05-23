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
    public sealed class AnomalyDetectorCloudwatchConfig
    {
        /// <summary>
        /// An IAM role that gives Amazon Lookout for Metrics permission to access data in Amazon CloudWatch.
        /// </summary>
        public readonly string RoleArn;

        [OutputConstructor]
        private AnomalyDetectorCloudwatchConfig(string roleArn)
        {
            RoleArn = roleArn;
        }
    }
}
