// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class StreamingDistributionLoggingArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("prefix", required: true)]
        public Input<string> Prefix { get; set; } = null!;

        public StreamingDistributionLoggingArgs()
        {
        }
        public static new StreamingDistributionLoggingArgs Empty => new StreamingDistributionLoggingArgs();
    }
}