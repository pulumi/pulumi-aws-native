// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    /// <summary>
    /// Specifies a metrics configuration for the CloudWatch request metrics (specified by the metrics configuration ID) from an Amazon S3 bucket. If you're updating an existing metrics configuration, note that this is a full replacement of the existing metrics configuration. If you don't include the elements you want to keep, they are erased. For examples, see [AWS::S3::Bucket](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#aws-properties-s3-bucket--examples). For more information, see [PUT Bucket metrics](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTMetricConfiguration.html) in the *Amazon S3 API Reference*.
    /// </summary>
    public sealed class BucketMetricsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access point that was used while performing operations on the object. The metrics configuration only includes objects that meet the filter's criteria.
        /// </summary>
        [Input("accessPointArn")]
        public Input<string>? AccessPointArn { get; set; }

        /// <summary>
        /// The ID used to identify the metrics configuration. This can be any value you choose that helps you identify your metrics configuration.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The prefix that an object must have to be included in the metrics results.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        [Input("tagFilters")]
        private InputList<Inputs.BucketTagFilterArgs>? _tagFilters;

        /// <summary>
        /// Specifies a list of tag filters to use as a metrics configuration filter. The metrics configuration includes only objects that meet the filter's criteria.
        /// </summary>
        public InputList<Inputs.BucketTagFilterArgs> TagFilters
        {
            get => _tagFilters ?? (_tagFilters = new InputList<Inputs.BucketTagFilterArgs>());
            set => _tagFilters = value;
        }

        public BucketMetricsConfigurationArgs()
        {
        }
        public static new BucketMetricsConfigurationArgs Empty => new BucketMetricsConfigurationArgs();
    }
}
