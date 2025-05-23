// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    /// <summary>
    /// An object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in requests that CloudFront sends to the origin.
    /// </summary>
    public sealed class OriginRequestPolicyQueryStringsConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether any URL query strings in viewer requests are included in requests that CloudFront sends to the origin. Valid values are:
        ///   +  ``none`` – No query strings in viewer requests are included in requests that CloudFront sends to the origin. Even when this field is set to ``none``, any query strings that are listed in a ``CachePolicy``*are* included in origin requests.
        ///   +  ``whitelist`` – Only the query strings in viewer requests that are listed in the ``QueryStringNames`` type are included in requests that CloudFront sends to the origin.
        ///   +  ``all`` – All query strings in viewer requests are included in requests that CloudFront sends to the origin.
        ///   +  ``allExcept`` – All query strings in viewer requests are included in requests that CloudFront sends to the origin, *except* for those listed in the ``QueryStringNames`` type, which are not included.
        /// </summary>
        [Input("queryStringBehavior", required: true)]
        public Input<string> QueryStringBehavior { get; set; } = null!;

        [Input("queryStrings")]
        private InputList<string>? _queryStrings;

        /// <summary>
        /// Contains a list of query string names.
        /// </summary>
        public InputList<string> QueryStrings
        {
            get => _queryStrings ?? (_queryStrings = new InputList<string>());
            set => _queryStrings = value;
        }

        public OriginRequestPolicyQueryStringsConfigArgs()
        {
        }
        public static new OriginRequestPolicyQueryStringsConfigArgs Empty => new OriginRequestPolicyQueryStringsConfigArgs();
    }
}
