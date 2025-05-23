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
    /// Details of what logs are delivered and where they are delivered.
    /// </summary>
    [OutputType]
    public sealed class ConnectorLogDelivery
    {
        /// <summary>
        /// The workers can send worker logs to different destination types. This configuration specifies the details of these destinations.
        /// </summary>
        public readonly Outputs.ConnectorWorkerLogDelivery WorkerLogDelivery;

        [OutputConstructor]
        private ConnectorLogDelivery(Outputs.ConnectorWorkerLogDelivery workerLogDelivery)
        {
            WorkerLogDelivery = workerLogDelivery;
        }
    }
}
