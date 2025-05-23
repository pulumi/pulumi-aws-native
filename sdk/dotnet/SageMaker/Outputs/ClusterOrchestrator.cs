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
    /// Specifies parameter(s) specific to the orchestrator, e.g. specify the EKS cluster.
    /// </summary>
    [OutputType]
    public sealed class ClusterOrchestrator
    {
        /// <summary>
        /// The configuration of the Amazon EKS orchestrator cluster for the SageMaker HyperPod cluster.
        /// </summary>
        public readonly Outputs.ClusterOrchestratorEksConfig Eks;

        [OutputConstructor]
        private ClusterOrchestrator(Outputs.ClusterOrchestratorEksConfig eks)
        {
            Eks = eks;
        }
    }
}
