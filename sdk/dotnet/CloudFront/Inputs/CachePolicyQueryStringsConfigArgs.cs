// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class CachePolicyQueryStringsConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("queryStringBehavior", required: true)]
        public Input<string> QueryStringBehavior { get; set; } = null!;

        [Input("queryStrings")]
        private InputList<string>? _queryStrings;
        public InputList<string> QueryStrings
        {
            get => _queryStrings ?? (_queryStrings = new InputList<string>());
            set => _queryStrings = value;
        }

        public CachePolicyQueryStringsConfigArgs()
        {
        }
        public static new CachePolicyQueryStringsConfigArgs Empty => new CachePolicyQueryStringsConfigArgs();
    }
}