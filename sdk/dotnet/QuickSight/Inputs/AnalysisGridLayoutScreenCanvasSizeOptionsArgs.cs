// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisGridLayoutScreenCanvasSizeOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        [Input("optimizedViewPortWidth")]
        public Input<string>? OptimizedViewPortWidth { get; set; }

        [Input("resizeOption", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisResizeOption> ResizeOption { get; set; } = null!;

        public AnalysisGridLayoutScreenCanvasSizeOptionsArgs()
        {
        }
        public static new AnalysisGridLayoutScreenCanvasSizeOptionsArgs Empty => new AnalysisGridLayoutScreenCanvasSizeOptionsArgs();
    }
}