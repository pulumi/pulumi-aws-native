// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// Contains settings for the network performance options for the instance.
    /// </summary>
    public sealed class LaunchTemplateNetworkPerformanceOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify the bandwidth weighting option to boost the associated type of baseline bandwidth, as follows:
        ///   + default This option uses the standard bandwidth configuration for your instance type. + vpc-1 This option boosts your networking baseline bandwidth and reduces your EBS baseline bandwidth. + ebs-1 This option boosts your EBS baseline bandwidth and reduces your networking baseline bandwidth.
        /// </summary>
        [Input("bandwidthWeighting")]
        public Input<string>? BandwidthWeighting { get; set; }

        public LaunchTemplateNetworkPerformanceOptionsArgs()
        {
        }
        public static new LaunchTemplateNetworkPerformanceOptionsArgs Empty => new LaunchTemplateNetworkPerformanceOptionsArgs();
    }
}
