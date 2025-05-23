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
    public sealed class AnalysisGeospatialCategoricalColor
    {
        /// <summary>
        /// A list of categorical data colors for each category.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisGeospatialCategoricalDataColor> CategoryDataColors;
        /// <summary>
        /// The default opacity of a categorical color.
        /// </summary>
        public readonly double? DefaultOpacity;
        /// <summary>
        /// The null data visualization settings.
        /// </summary>
        public readonly Outputs.AnalysisGeospatialNullDataSettings? NullDataSettings;
        /// <summary>
        /// The state of visibility for null data.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisVisibility? NullDataVisibility;

        [OutputConstructor]
        private AnalysisGeospatialCategoricalColor(
            ImmutableArray<Outputs.AnalysisGeospatialCategoricalDataColor> categoryDataColors,

            double? defaultOpacity,

            Outputs.AnalysisGeospatialNullDataSettings? nullDataSettings,

            Pulumi.AwsNative.QuickSight.AnalysisVisibility? nullDataVisibility)
        {
            CategoryDataColors = categoryDataColors;
            DefaultOpacity = defaultOpacity;
            NullDataSettings = nullDataSettings;
            NullDataVisibility = nullDataVisibility;
        }
    }
}
