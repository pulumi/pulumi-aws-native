// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect.Outputs
{

    /// <summary>
    /// Details about a fixed capacity allocated to a connector.
    /// </summary>
    [OutputType]
    public sealed class ConnectorProvisionedCapacity
    {
        /// <summary>
        /// Specifies how many MSK Connect Units (MCU) are allocated to the connector.
        /// </summary>
        public readonly int? McuCount;
        /// <summary>
        /// Number of workers for a connector.
        /// </summary>
        public readonly int WorkerCount;

        [OutputConstructor]
        private ConnectorProvisionedCapacity(
            int? mcuCount,

            int workerCount)
        {
            McuCount = mcuCount;
            WorkerCount = workerCount;
        }
    }
}
