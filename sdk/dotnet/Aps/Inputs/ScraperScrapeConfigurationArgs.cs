// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Aps.Inputs
{

    /// <summary>
    /// Scraper configuration
    /// </summary>
    public sealed class ScraperScrapeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Prometheus compatible scrape configuration in base64 encoded blob format
        /// </summary>
        [Input("configurationBlob")]
        public Input<string>? ConfigurationBlob { get; set; }

        public ScraperScrapeConfigurationArgs()
        {
        }
        public static new ScraperScrapeConfigurationArgs Empty => new ScraperScrapeConfigurationArgs();
    }
}