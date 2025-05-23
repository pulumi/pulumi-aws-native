// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize.Outputs
{

    /// <summary>
    /// Describes the properties for hyperparameter optimization (HPO)
    /// </summary>
    [OutputType]
    public sealed class SolutionConfigHpoConfigProperties
    {
        /// <summary>
        /// The hyperparameters and their allowable ranges
        /// </summary>
        public readonly Outputs.SolutionConfigHpoConfigPropertiesAlgorithmHyperParameterRangesProperties? AlgorithmHyperParameterRanges;
        /// <summary>
        /// The metric to optimize during HPO.
        /// </summary>
        public readonly Outputs.SolutionConfigHpoConfigPropertiesHpoObjectiveProperties? HpoObjective;
        /// <summary>
        /// Describes the resource configuration for hyperparameter optimization (HPO).
        /// </summary>
        public readonly Outputs.SolutionConfigHpoConfigPropertiesHpoResourceConfigProperties? HpoResourceConfig;

        [OutputConstructor]
        private SolutionConfigHpoConfigProperties(
            Outputs.SolutionConfigHpoConfigPropertiesAlgorithmHyperParameterRangesProperties? algorithmHyperParameterRanges,

            Outputs.SolutionConfigHpoConfigPropertiesHpoObjectiveProperties? hpoObjective,

            Outputs.SolutionConfigHpoConfigPropertiesHpoResourceConfigProperties? hpoResourceConfig)
        {
            AlgorithmHyperParameterRanges = algorithmHyperParameterRanges;
            HpoObjective = hpoObjective;
            HpoResourceConfig = hpoResourceConfig;
        }
    }
}
