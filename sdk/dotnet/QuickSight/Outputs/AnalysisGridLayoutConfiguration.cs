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
    public sealed class AnalysisGridLayoutConfiguration
    {
        public readonly Outputs.AnalysisGridLayoutCanvasSizeOptions? CanvasSizeOptions;
        /// <summary>
        /// The elements that are included in a grid layout.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisGridLayoutElement> Elements;

        [OutputConstructor]
        private AnalysisGridLayoutConfiguration(
            Outputs.AnalysisGridLayoutCanvasSizeOptions? canvasSizeOptions,

            ImmutableArray<Outputs.AnalysisGridLayoutElement> elements)
        {
            CanvasSizeOptions = canvasSizeOptions;
            Elements = elements;
        }
    }
}
