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
    /// Contains explainability metrics for a model.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageDriftCheckExplainability
    {
        public readonly Outputs.ModelPackageFileSource? ConfigFile;
        public readonly Outputs.ModelPackageMetricsSource? Constraints;

        [OutputConstructor]
        private ModelPackageDriftCheckExplainability(
            Outputs.ModelPackageFileSource? configFile,

            Outputs.ModelPackageMetricsSource? constraints)
        {
            ConfigFile = configFile;
            Constraints = constraints;
        }
    }
}