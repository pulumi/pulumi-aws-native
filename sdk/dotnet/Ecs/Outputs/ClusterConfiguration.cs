// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Outputs
{

    /// <summary>
    /// The configurations to be set at cluster level.
    /// </summary>
    [OutputType]
    public sealed class ClusterConfiguration
    {
        public readonly Outputs.ClusterExecuteCommandConfiguration? ExecuteCommandConfiguration;

        [OutputConstructor]
        private ClusterConfiguration(Outputs.ClusterExecuteCommandConfiguration? executeCommandConfiguration)
        {
            ExecuteCommandConfiguration = executeCommandConfiguration;
        }
    }
}