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
    /// Specifies a raw data source location to ingest.
    /// </summary>
    public sealed class DataSourceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("s3Configuration", required: true)]
        public Input<Inputs.DataSourceS3DataSourceConfigurationArgs> S3Configuration { get; set; } = null!;

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.DataSourceType> Type { get; set; } = null!;

        public DataSourceConfigurationArgs()
        {
        }
        public static new DataSourceConfigurationArgs Empty => new DataSourceConfigurationArgs();
    }
}