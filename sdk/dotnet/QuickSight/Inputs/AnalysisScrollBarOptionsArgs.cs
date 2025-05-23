// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisScrollBarOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The visibility of the data zoom scroll bar.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? Visibility { get; set; }

        /// <summary>
        /// The visibility range for the data zoom scroll bar.
        /// </summary>
        [Input("visibleRange")]
        public Input<Inputs.AnalysisVisibleRangeOptionsArgs>? VisibleRange { get; set; }

        public AnalysisScrollBarOptionsArgs()
        {
        }
        public static new AnalysisScrollBarOptionsArgs Empty => new AnalysisScrollBarOptionsArgs();
    }
}
