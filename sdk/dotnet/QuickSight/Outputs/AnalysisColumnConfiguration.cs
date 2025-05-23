// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisColumnConfiguration
    {
        /// <summary>
        /// The color configurations of the column.
        /// </summary>
        public readonly Outputs.AnalysisColorsConfiguration? ColorsConfiguration;
        /// <summary>
        /// The column.
        /// </summary>
        public readonly Outputs.AnalysisColumnIdentifier Column;
        /// <summary>
        /// The format configuration of a column.
        /// </summary>
        public readonly Outputs.AnalysisFormatConfiguration? FormatConfiguration;
        /// <summary>
        /// The role of the column.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisColumnRole? Role;

        [OutputConstructor]
        private AnalysisColumnConfiguration(
            Outputs.AnalysisColorsConfiguration? colorsConfiguration,

            Outputs.AnalysisColumnIdentifier column,

            Outputs.AnalysisFormatConfiguration? formatConfiguration,

            Pulumi.AwsNative.QuickSight.AnalysisColumnRole? role)
        {
            ColorsConfiguration = colorsConfiguration;
            Column = column;
            FormatConfiguration = formatConfiguration;
            Role = role;
        }
    }
}
