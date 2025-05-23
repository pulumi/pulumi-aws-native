// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class DatastoreCustomerManagedS3Storage
    {
        /// <summary>
        /// The name of the Amazon S3 bucket where your data is stored.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// (Optional) The prefix used to create the keys of the data store data objects. Each object in an Amazon S3 bucket has a key that is its unique identifier in the bucket. Each object in a bucket has exactly one key. The prefix must end with a forward slash (/).
        /// </summary>
        public readonly string? KeyPrefix;

        [OutputConstructor]
        private DatastoreCustomerManagedS3Storage(
            string bucket,

            string? keyPrefix)
        {
            Bucket = bucket;
            KeyPrefix = keyPrefix;
        }
    }
}
