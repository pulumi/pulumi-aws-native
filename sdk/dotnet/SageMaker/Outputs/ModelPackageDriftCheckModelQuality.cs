// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageDriftCheckModelQuality
    {
        /// <summary>
        /// The drift check model quality constraints.
        /// </summary>
        public readonly Outputs.ModelPackageMetricsSource? Constraints;
        /// <summary>
        /// The drift check model quality statistics.
        /// </summary>
        public readonly Outputs.ModelPackageMetricsSource? Statistics;

        [OutputConstructor]
        private ModelPackageDriftCheckModelQuality(
            Outputs.ModelPackageMetricsSource? constraints,

            Outputs.ModelPackageMetricsSource? statistics)
        {
            Constraints = constraints;
            Statistics = statistics;
        }
    }
}
