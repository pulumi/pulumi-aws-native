// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Osis.Inputs
{

    /// <summary>
    /// The destination for OpenSearch Ingestion Service logs sent to Amazon CloudWatch.
    /// </summary>
    public sealed class PipelineLogPublishingOptionsCloudWatchLogDestinationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("logGroup", required: true)]
        public Input<string> LogGroup { get; set; } = null!;

        public PipelineLogPublishingOptionsCloudWatchLogDestinationPropertiesArgs()
        {
        }
        public static new PipelineLogPublishingOptionsCloudWatchLogDestinationPropertiesArgs Empty => new PipelineLogPublishingOptionsCloudWatchLogDestinationPropertiesArgs();
    }
}