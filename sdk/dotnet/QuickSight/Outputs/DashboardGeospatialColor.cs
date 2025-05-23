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
    public sealed class DashboardGeospatialColor
    {
        /// <summary>
        /// The visualization properties for the categorical color.
        /// </summary>
        public readonly Outputs.DashboardGeospatialCategoricalColor? Categorical;
        /// <summary>
        /// The visualization properties for the gradient color.
        /// </summary>
        public readonly Outputs.DashboardGeospatialGradientColor? Gradient;
        /// <summary>
        /// The visualization properties for the solid color.
        /// </summary>
        public readonly Outputs.DashboardGeospatialSolidColor? Solid;

        [OutputConstructor]
        private DashboardGeospatialColor(
            Outputs.DashboardGeospatialCategoricalColor? categorical,

            Outputs.DashboardGeospatialGradientColor? gradient,

            Outputs.DashboardGeospatialSolidColor? solid)
        {
            Categorical = categorical;
            Gradient = gradient;
            Solid = solid;
        }
    }
}
