// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect.Inputs
{

    /// <summary>
    /// Information about the scale out policy of the connector.
    /// </summary>
    public sealed class ConnectorScaleOutPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the CPU utilization percentage threshold at which connector scale out should trigger.
        /// </summary>
        [Input("cpuUtilizationPercentage", required: true)]
        public Input<int> CpuUtilizationPercentage { get; set; } = null!;

        public ConnectorScaleOutPolicyArgs()
        {
        }
        public static new ConnectorScaleOutPolicyArgs Empty => new ConnectorScaleOutPolicyArgs();
    }
}
