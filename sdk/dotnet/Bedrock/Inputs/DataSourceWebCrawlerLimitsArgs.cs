// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Limit settings for the web crawler.
    /// </summary>
    public sealed class DataSourceWebCrawlerLimitsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of pages the crawler can crawl.
        /// </summary>
        [Input("maxPages")]
        public Input<int>? MaxPages { get; set; }

        /// <summary>
        /// Rate of web URLs retrieved per minute.
        /// </summary>
        [Input("rateLimit")]
        public Input<int>? RateLimit { get; set; }

        public DataSourceWebCrawlerLimitsArgs()
        {
        }
        public static new DataSourceWebCrawlerLimitsArgs Empty => new DataSourceWebCrawlerLimitsArgs();
    }
}
