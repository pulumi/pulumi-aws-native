// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Represents the drift check bias baselines that can be used when the model monitor is set using the model package.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageDriftCheckBias
    {
        public readonly Outputs.ModelPackageFileSource? ConfigFile;
        public readonly Outputs.ModelPackageMetricsSource? PostTrainingConstraints;
        public readonly Outputs.ModelPackageMetricsSource? PreTrainingConstraints;

        [OutputConstructor]
        private ModelPackageDriftCheckBias(
            Outputs.ModelPackageFileSource? configFile,

            Outputs.ModelPackageMetricsSource? postTrainingConstraints,

            Outputs.ModelPackageMetricsSource? preTrainingConstraints)
        {
            ConfigFile = configFile;
            PostTrainingConstraints = postTrainingConstraints;
            PreTrainingConstraints = preTrainingConstraints;
        }
    }
}