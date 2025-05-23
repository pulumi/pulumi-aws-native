// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisColumnConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The color configurations of the column.
        /// </summary>
        [Input("colorsConfiguration")]
        public Input<Inputs.AnalysisColorsConfigurationArgs>? ColorsConfiguration { get; set; }

        /// <summary>
        /// The column.
        /// </summary>
        [Input("column", required: true)]
        public Input<Inputs.AnalysisColumnIdentifierArgs> Column { get; set; } = null!;

        /// <summary>
        /// The format configuration of a column.
        /// </summary>
        [Input("formatConfiguration")]
        public Input<Inputs.AnalysisFormatConfigurationArgs>? FormatConfiguration { get; set; }

        /// <summary>
        /// The role of the column.
        /// </summary>
        [Input("role")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisColumnRole>? Role { get; set; }

        public AnalysisColumnConfigurationArgs()
        {
        }
        public static new AnalysisColumnConfigurationArgs Empty => new AnalysisColumnConfigurationArgs();
    }
}
