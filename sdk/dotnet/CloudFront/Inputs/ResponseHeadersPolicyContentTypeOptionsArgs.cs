// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class ResponseHeadersPolicyContentTypeOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("override", required: true)]
        public Input<bool> Override { get; set; } = null!;

        public ResponseHeadersPolicyContentTypeOptionsArgs()
        {
        }
        public static new ResponseHeadersPolicyContentTypeOptionsArgs Empty => new ResponseHeadersPolicyContentTypeOptionsArgs();
    }
}