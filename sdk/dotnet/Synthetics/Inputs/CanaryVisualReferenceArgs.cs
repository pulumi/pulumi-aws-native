// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics.Inputs
{

    public sealed class CanaryVisualReferenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Canary run id to be used as base reference for visual testing
        /// </summary>
        [Input("baseCanaryRunId", required: true)]
        public Input<string> BaseCanaryRunId { get; set; } = null!;

        [Input("baseScreenshots")]
        private InputList<Inputs.CanaryBaseScreenshotArgs>? _baseScreenshots;

        /// <summary>
        /// List of screenshots used as base reference for visual testing
        /// </summary>
        public InputList<Inputs.CanaryBaseScreenshotArgs> BaseScreenshots
        {
            get => _baseScreenshots ?? (_baseScreenshots = new InputList<Inputs.CanaryBaseScreenshotArgs>());
            set => _baseScreenshots = value;
        }

        public CanaryVisualReferenceArgs()
        {
        }
        public static new CanaryVisualReferenceArgs Empty => new CanaryVisualReferenceArgs();
    }
}
