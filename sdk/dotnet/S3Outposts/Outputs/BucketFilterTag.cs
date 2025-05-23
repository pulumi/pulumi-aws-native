// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Outposts.Outputs
{

    /// <summary>
    /// Tag used to identify a subset of objects for an Amazon S3Outposts bucket.
    /// </summary>
    [OutputType]
    public sealed class BucketFilterTag
    {
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private BucketFilterTag(
            string key,

            string value)
        {
            Key = key;
            Value = value;
        }
    }
}
