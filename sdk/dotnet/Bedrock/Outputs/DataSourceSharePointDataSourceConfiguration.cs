// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Outputs
{

    /// <summary>
    /// The configuration information to connect to SharePoint as your data source.
    /// </summary>
    [OutputType]
    public sealed class DataSourceSharePointDataSourceConfiguration
    {
        public readonly Outputs.DataSourceSharePointCrawlerConfiguration? CrawlerConfiguration;
        public readonly Outputs.DataSourceSharePointSourceConfiguration SourceConfiguration;

        [OutputConstructor]
        private DataSourceSharePointDataSourceConfiguration(
            Outputs.DataSourceSharePointCrawlerConfiguration? crawlerConfiguration,

            Outputs.DataSourceSharePointSourceConfiguration sourceConfiguration)
        {
            CrawlerConfiguration = crawlerConfiguration;
            SourceConfiguration = sourceConfiguration;
        }
    }
}