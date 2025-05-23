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
    public sealed class AnalysisUniqueValuesComputation
    {
        /// <summary>
        /// The category field that is used in a computation.
        /// </summary>
        public readonly Outputs.AnalysisDimensionField? Category;
        /// <summary>
        /// The ID for a computation.
        /// </summary>
        public readonly string ComputationId;
        /// <summary>
        /// The name of a computation.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private AnalysisUniqueValuesComputation(
            Outputs.AnalysisDimensionField? category,

            string computationId,

            string? name)
        {
            Category = category;
            ComputationId = computationId;
            Name = name;
        }
    }
}
