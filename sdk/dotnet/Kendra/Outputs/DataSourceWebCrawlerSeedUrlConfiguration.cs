// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceWebCrawlerSeedUrlConfiguration
    {
        /// <summary>
        /// The list of seed or starting point URLs of the websites you want to crawl.
        /// 
        /// The list can include a maximum of 100 seed URLs.
        /// </summary>
        public readonly ImmutableArray<string> SeedUrls;
        /// <summary>
        /// You can choose one of the following modes:
        /// 
        /// - `HOST_ONLY` —crawl only the website host names. For example, if the seed URL is "abc.example.com", then only URLs with host name "abc.example.com" are crawled.
        /// - `SUBDOMAINS` —crawl the website host names with subdomains. For example, if the seed URL is "abc.example.com", then "a.abc.example.com" and "b.abc.example.com" are also crawled.
        /// - `EVERYTHING` —crawl the website host names with subdomains and other domains that the web pages link to.
        /// 
        /// The default mode is set to `HOST_ONLY` .
        /// </summary>
        public readonly Pulumi.AwsNative.Kendra.DataSourceWebCrawlerSeedUrlConfigurationWebCrawlerMode? WebCrawlerMode;

        [OutputConstructor]
        private DataSourceWebCrawlerSeedUrlConfiguration(
            ImmutableArray<string> seedUrls,

            Pulumi.AwsNative.Kendra.DataSourceWebCrawlerSeedUrlConfigurationWebCrawlerMode? webCrawlerMode)
        {
            SeedUrls = seedUrls;
            WebCrawlerMode = webCrawlerMode;
        }
    }
}
