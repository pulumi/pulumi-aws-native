// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class PublicKeyConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("callerReference", required: true)]
        public Input<string> CallerReference { get; set; } = null!;

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("encodedKey", required: true)]
        public Input<string> EncodedKey { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public PublicKeyConfigArgs()
        {
        }
        public static new PublicKeyConfigArgs Empty => new PublicKeyConfigArgs();
    }
}