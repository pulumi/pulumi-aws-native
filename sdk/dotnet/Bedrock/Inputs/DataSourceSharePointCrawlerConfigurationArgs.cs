// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// The configuration of the SharePoint content. For example, configuring specific types of SharePoint content.
    /// </summary>
    public sealed class DataSourceSharePointCrawlerConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("filterConfiguration")]
        public Input<Inputs.DataSourceCrawlFilterConfigurationArgs>? FilterConfiguration { get; set; }

        public DataSourceSharePointCrawlerConfigurationArgs()
        {
        }
        public static new DataSourceSharePointCrawlerConfigurationArgs Empty => new DataSourceSharePointCrawlerConfigurationArgs();
    }
}