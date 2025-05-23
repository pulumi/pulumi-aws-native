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
    public sealed class AnalysisGeospatialCircleRadius
    {
        /// <summary>
        /// The positive value for the radius of a circle.
        /// </summary>
        public readonly double? Radius;

        [OutputConstructor]
        private AnalysisGeospatialCircleRadius(double? radius)
        {
            Radius = radius;
        }
    }
}
