// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    /// <summary>
    /// An Amazon CloudFront VPC origin configuration.
    /// </summary>
    [OutputType]
    public sealed class DistributionVpcOriginConfig
    {
        /// <summary>
        /// Specifies how long, in seconds, CloudFront persists its connection to the origin. The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don't specify otherwise) is 5 seconds.
        ///  For more information, see [Keep-alive timeout (custom origins only)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginKeepaliveTimeout) in the *Amazon CloudFront Developer Guide*.
        /// </summary>
        public readonly int? OriginKeepaliveTimeout;
        /// <summary>
        /// Specifies how long, in seconds, CloudFront waits for a response from the origin. This is also known as the *origin response timeout*. The minimum timeout is 1 second, the maximum is 60 seconds, and the default (if you don't specify otherwise) is 30 seconds.
        ///  For more information, see [Response timeout (custom origins only)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValuesOriginResponseTimeout) in the *Amazon CloudFront Developer Guide*.
        /// </summary>
        public readonly int? OriginReadTimeout;
        /// <summary>
        /// The VPC origin ID.
        /// </summary>
        public readonly string VpcOriginId;

        [OutputConstructor]
        private DistributionVpcOriginConfig(
            int? originKeepaliveTimeout,

            int? originReadTimeout,

            string vpcOriginId)
        {
            OriginKeepaliveTimeout = originKeepaliveTimeout;
            OriginReadTimeout = originReadTimeout;
            VpcOriginId = vpcOriginId;
        }
    }
}
