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
    /// The type of filtering that you want to apply to certain objects or content of the data source. For example, the PATTERN type is regular expression patterns you can apply to filter your content.
    /// </summary>
    public sealed class DataSourceCrawlFilterConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration of filtering certain objects or content types of the data source.
        /// </summary>
        [Input("patternObjectFilter")]
        public Input<Inputs.DataSourcePatternObjectFilterConfigurationArgs>? PatternObjectFilter { get; set; }

        /// <summary>
        /// The crawl filter type.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.DataSourceCrawlFilterConfigurationType> Type { get; set; } = null!;

        public DataSourceCrawlFilterConfigurationArgs()
        {
        }
        public static new DataSourceCrawlFilterConfigurationArgs Empty => new DataSourceCrawlFilterConfigurationArgs();
    }
}
