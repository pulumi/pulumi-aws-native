// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3.Inputs
{

    public sealed class BucketReplicationTimeValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("minutes", required: true)]
        public Input<int> Minutes { get; set; } = null!;

        public BucketReplicationTimeValueArgs()
        {
        }
        public static new BucketReplicationTimeValueArgs Empty => new BucketReplicationTimeValueArgs();
    }
}