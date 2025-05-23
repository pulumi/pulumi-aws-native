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
    /// Describes the key format for server access log file in the target bucket. You can choose between SimplePrefix and PartitionedPrefix.
    /// </summary>
    public sealed class BucketTargetObjectKeyFormatArgs : global::Pulumi.ResourceArgs
    {
        public BucketTargetObjectKeyFormatArgs()
        {
        }
        public static new BucketTargetObjectKeyFormatArgs Empty => new BucketTargetObjectKeyFormatArgs();
    }
}
