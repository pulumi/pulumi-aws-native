// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisVisualPaletteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The chart color options for the visual palette.
        /// </summary>
        [Input("chartColor")]
        public Input<string>? ChartColor { get; set; }

        [Input("colorMap")]
        private InputList<Inputs.AnalysisDataPathColorArgs>? _colorMap;

        /// <summary>
        /// The color map options for the visual palette.
        /// </summary>
        public InputList<Inputs.AnalysisDataPathColorArgs> ColorMap
        {
            get => _colorMap ?? (_colorMap = new InputList<Inputs.AnalysisDataPathColorArgs>());
            set => _colorMap = value;
        }

        public AnalysisVisualPaletteArgs()
        {
        }
        public static new AnalysisVisualPaletteArgs Empty => new AnalysisVisualPaletteArgs();
    }
}
