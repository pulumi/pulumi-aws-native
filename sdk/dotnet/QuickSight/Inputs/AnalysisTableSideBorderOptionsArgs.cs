// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTableSideBorderOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The table border options of the bottom border.
        /// </summary>
        [Input("bottom")]
        public Input<Inputs.AnalysisTableBorderOptionsArgs>? Bottom { get; set; }

        /// <summary>
        /// The table border options of the inner horizontal border.
        /// </summary>
        [Input("innerHorizontal")]
        public Input<Inputs.AnalysisTableBorderOptionsArgs>? InnerHorizontal { get; set; }

        /// <summary>
        /// The table border options of the inner vertical border.
        /// </summary>
        [Input("innerVertical")]
        public Input<Inputs.AnalysisTableBorderOptionsArgs>? InnerVertical { get; set; }

        /// <summary>
        /// The table border options of the left border.
        /// </summary>
        [Input("left")]
        public Input<Inputs.AnalysisTableBorderOptionsArgs>? Left { get; set; }

        /// <summary>
        /// The table border options of the right border.
        /// </summary>
        [Input("right")]
        public Input<Inputs.AnalysisTableBorderOptionsArgs>? Right { get; set; }

        /// <summary>
        /// The table border options of the top border.
        /// </summary>
        [Input("top")]
        public Input<Inputs.AnalysisTableBorderOptionsArgs>? Top { get; set; }

        public AnalysisTableSideBorderOptionsArgs()
        {
        }
        public static new AnalysisTableSideBorderOptionsArgs Empty => new AnalysisTableSideBorderOptionsArgs();
    }
}
