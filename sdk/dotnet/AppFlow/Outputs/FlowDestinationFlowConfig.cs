// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    /// <summary>
    /// Configurations of destination connector.
    /// </summary>
    [OutputType]
    public sealed class FlowDestinationFlowConfig
    {
        /// <summary>
        /// The API version that the destination connector uses.
        /// </summary>
        public readonly string? ApiVersion;
        /// <summary>
        /// Name of destination connector profile
        /// </summary>
        public readonly string? ConnectorProfileName;
        /// <summary>
        /// Destination connector type
        /// </summary>
        public readonly Pulumi.AwsNative.AppFlow.FlowConnectorType ConnectorType;
        /// <summary>
        /// Destination connector details
        /// </summary>
        public readonly Outputs.FlowDestinationConnectorProperties DestinationConnectorProperties;

        [OutputConstructor]
        private FlowDestinationFlowConfig(
            string? apiVersion,

            string? connectorProfileName,

            Pulumi.AwsNative.AppFlow.FlowConnectorType connectorType,

            Outputs.FlowDestinationConnectorProperties destinationConnectorProperties)
        {
            ApiVersion = apiVersion;
            ConnectorProfileName = connectorProfileName;
            ConnectorType = connectorType;
            DestinationConnectorProperties = destinationConnectorProperties;
        }
    }
}
