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
    /// A location for storing content from data sources temporarily as it is processed by custom components in the ingestion pipeline.
    /// </summary>
    public sealed class DataSourceIntermediateStorageArgs : global::Pulumi.ResourceArgs
    {
        [Input("s3Location", required: true)]
        public Input<Inputs.DataSourceS3LocationArgs> S3Location { get; set; } = null!;

        public DataSourceIntermediateStorageArgs()
        {
        }
        public static new DataSourceIntermediateStorageArgs Empty => new DataSourceIntermediateStorageArgs();
    }
}