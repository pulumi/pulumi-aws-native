// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class ChannelCustomerManagedS3
    {
        public readonly string Bucket;
        public readonly string? KeyPrefix;
        public readonly string RoleArn;

        [OutputConstructor]
        private ChannelCustomerManagedS3(
            string bucket,

            string? keyPrefix,

            string roleArn)
        {
            Bucket = bucket;
            KeyPrefix = keyPrefix;
            RoleArn = roleArn;
        }
    }
}