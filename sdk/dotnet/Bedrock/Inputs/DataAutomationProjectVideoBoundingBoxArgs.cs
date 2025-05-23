// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    public sealed class DataAutomationProjectVideoBoundingBoxArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether bounding boxes are enabled for video.
        /// </summary>
        [Input("state", required: true)]
        public Input<Pulumi.AwsNative.Bedrock.DataAutomationProjectState> State { get; set; } = null!;

        public DataAutomationProjectVideoBoundingBoxArgs()
        {
        }
        public static new DataAutomationProjectVideoBoundingBoxArgs Empty => new DataAutomationProjectVideoBoundingBoxArgs();
    }
}
