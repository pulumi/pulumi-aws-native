// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodePipeline.Inputs
{

    public sealed class CustomActionTypeArtifactDetailsArgs : global::Pulumi.ResourceArgs
    {
        [Input("maximumCount", required: true)]
        public Input<int> MaximumCount { get; set; } = null!;

        [Input("minimumCount", required: true)]
        public Input<int> MinimumCount { get; set; } = null!;

        public CustomActionTypeArtifactDetailsArgs()
        {
        }
        public static new CustomActionTypeArtifactDetailsArgs Empty => new CustomActionTypeArtifactDetailsArgs();
    }
}