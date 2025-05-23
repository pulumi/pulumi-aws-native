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
    /// S3 buckets and Regions to include/exclude in the Amazon S3 Storage Lens configuration.
    /// </summary>
    public sealed class StorageLensBucketsAndRegionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("buckets")]
        private InputList<string>? _buckets;

        /// <summary>
        /// This property contains the details of the buckets for the Amazon S3 Storage Lens configuration. This should be the bucket Amazon Resource Name(ARN). For valid values, see [Buckets ARN format here](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_Include.html#API_control_Include_Contents) in the *Amazon S3 API Reference* .
        /// </summary>
        public InputList<string> Buckets
        {
            get => _buckets ?? (_buckets = new InputList<string>());
            set => _buckets = value;
        }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// This property contains the details of the Regions for the S3 Storage Lens configuration.
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        public StorageLensBucketsAndRegionsArgs()
        {
        }
        public static new StorageLensBucketsAndRegionsArgs Empty => new StorageLensBucketsAndRegionsArgs();
    }
}
