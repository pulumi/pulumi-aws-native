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
    public sealed class AnalysisGradientColor
    {
        /// <summary>
        /// The list of gradient color stops.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisGradientStop> Stops;

        [OutputConstructor]
        private AnalysisGradientColor(ImmutableArray<Outputs.AnalysisGradientStop> stops)
        {
            Stops = stops;
        }
    }
}
