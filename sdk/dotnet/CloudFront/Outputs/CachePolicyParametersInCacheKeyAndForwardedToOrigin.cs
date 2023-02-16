// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Outputs
{

    [OutputType]
    public sealed class CachePolicyParametersInCacheKeyAndForwardedToOrigin
    {
        public readonly Outputs.CachePolicyCookiesConfig CookiesConfig;
        public readonly bool? EnableAcceptEncodingBrotli;
        public readonly bool EnableAcceptEncodingGzip;
        public readonly Outputs.CachePolicyHeadersConfig HeadersConfig;
        public readonly Outputs.CachePolicyQueryStringsConfig QueryStringsConfig;

        [OutputConstructor]
        private CachePolicyParametersInCacheKeyAndForwardedToOrigin(
            Outputs.CachePolicyCookiesConfig cookiesConfig,

            bool? enableAcceptEncodingBrotli,

            bool enableAcceptEncodingGzip,

            Outputs.CachePolicyHeadersConfig headersConfig,

            Outputs.CachePolicyQueryStringsConfig queryStringsConfig)
        {
            CookiesConfig = cookiesConfig;
            EnableAcceptEncodingBrotli = enableAcceptEncodingBrotli;
            EnableAcceptEncodingGzip = enableAcceptEncodingGzip;
            HeadersConfig = headersConfig;
            QueryStringsConfig = queryStringsConfig;
        }
    }
}