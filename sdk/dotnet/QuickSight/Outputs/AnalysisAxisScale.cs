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
    public sealed class AnalysisAxisScale
    {
        /// <summary>
        /// The linear axis scale setup.
        /// </summary>
        public readonly Outputs.AnalysisAxisLinearScale? Linear;
        /// <summary>
        /// The logarithmic axis scale setup.
        /// </summary>
        public readonly Outputs.AnalysisAxisLogarithmicScale? Logarithmic;

        [OutputConstructor]
        private AnalysisAxisScale(
            Outputs.AnalysisAxisLinearScale? linear,

            Outputs.AnalysisAxisLogarithmicScale? logarithmic)
        {
            Linear = linear;
            Logarithmic = logarithmic;
        }
    }
}
