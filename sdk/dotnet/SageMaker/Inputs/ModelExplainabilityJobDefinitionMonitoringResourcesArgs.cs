// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Identifies the resources to deploy for a monitoring job.
    /// </summary>
    public sealed class ModelExplainabilityJobDefinitionMonitoringResourcesArgs : global::Pulumi.ResourceArgs
    {
        [Input("clusterConfig", required: true)]
        public Input<Inputs.ModelExplainabilityJobDefinitionClusterConfigArgs> ClusterConfig { get; set; } = null!;

        public ModelExplainabilityJobDefinitionMonitoringResourcesArgs()
        {
        }
        public static new ModelExplainabilityJobDefinitionMonitoringResourcesArgs Empty => new ModelExplainabilityJobDefinitionMonitoringResourcesArgs();
    }
}