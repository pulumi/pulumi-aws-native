// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class FlowS3OutputFormatConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("aggregationConfig")]
        public Input<Inputs.FlowAggregationConfigArgs>? AggregationConfig { get; set; }

        [Input("fileType")]
        public Input<Pulumi.AwsNative.AppFlow.FlowFileType>? FileType { get; set; }

        [Input("prefixConfig")]
        public Input<Inputs.FlowPrefixConfigArgs>? PrefixConfig { get; set; }

        [Input("preserveSourceDataTyping")]
        public Input<bool>? PreserveSourceDataTyping { get; set; }

        public FlowS3OutputFormatConfigArgs()
        {
        }
        public static new FlowS3OutputFormatConfigArgs Empty => new FlowS3OutputFormatConfigArgs();
    }
}