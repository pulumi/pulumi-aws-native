// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi.Inputs
{

    public sealed class CapabilityS3LocationArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        public CapabilityS3LocationArgs()
        {
        }
        public static new CapabilityS3LocationArgs Empty => new CapabilityS3LocationArgs();
    }
}