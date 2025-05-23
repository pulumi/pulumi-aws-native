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
    /// A cache policy configuration.
    ///  This configuration determines the following:
    ///   +  The values that CloudFront includes in the cache key. These values can include HTTP headers, cookies, and URL query strings. CloudFront uses the cache key to find an object in its cache that it can return to the viewer.
    ///   +  The default, minimum, and maximum time to live (TTL) values that you want objects to stay in the CloudFront cache.
    ///   
    ///  The headers, cookies, and query strings that are included in the cache key are also included in requests that CloudFront sends to the origin. CloudFront sends a request when it can't find a valid object in its cache that matches the request's cache key. If you want to send values to the origin but *not* include them in the cache key, use ``OriginRequestPolicy``.
    /// </summary>
    public sealed class CachePolicyConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A comment to describe the cache policy. The comment cannot be longer than 128 characters.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. CloudFront uses this value as the object's time to live (TTL) only when the origin does *not* send ``Cache-Control`` or ``Expires`` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.
        ///  The default value for this field is 86400 seconds (one day). If the value of ``MinTTL`` is more than 86400 seconds, then the default value for this field is the same as the value of ``MinTTL``.
        /// </summary>
        [Input("defaultTtl", required: true)]
        public Input<double> DefaultTtl { get; set; } = null!;

        /// <summary>
        /// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. CloudFront uses this value only when the origin sends ``Cache-Control`` or ``Expires`` headers with the object. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.
        ///  The default value for this field is 31536000 seconds (one year). If the value of ``MinTTL`` or ``DefaultTTL`` is more than 31536000 seconds, then the default value for this field is the same as the value of ``DefaultTTL``.
        /// </summary>
        [Input("maxTtl", required: true)]
        public Input<double> MaxTtl { get; set; } = null!;

        /// <summary>
        /// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Expiration.html) in the *Amazon CloudFront Developer Guide*.
        /// </summary>
        [Input("minTtl", required: true)]
        public Input<double> MinTtl { get; set; } = null!;

        /// <summary>
        /// A unique name to identify the cache policy.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The HTTP headers, cookies, and URL query strings to include in the cache key. The values included in the cache key are also included in requests that CloudFront sends to the origin.
        /// </summary>
        [Input("parametersInCacheKeyAndForwardedToOrigin", required: true)]
        public Input<Inputs.CachePolicyParametersInCacheKeyAndForwardedToOriginArgs> ParametersInCacheKeyAndForwardedToOrigin { get; set; } = null!;

        public CachePolicyConfigArgs()
        {
        }
        public static new CachePolicyConfigArgs Empty => new CachePolicyConfigArgs();
    }
}
