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
    /// Details of what logs are delivered and where they are delivered.
    /// </summary>
    public sealed class ConnectorLogDeliveryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The workers can send worker logs to different destination types. This configuration specifies the details of these destinations.
        /// </summary>
        [Input("workerLogDelivery", required: true)]
        public Input<Inputs.ConnectorWorkerLogDeliveryArgs> WorkerLogDelivery { get; set; } = null!;

        public ConnectorLogDeliveryArgs()
        {
        }
        public static new ConnectorLogDeliveryArgs Empty => new ConnectorLogDeliveryArgs();
    }
}
