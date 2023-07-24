// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow
{
    public static class GetConnector
    {
        /// <summary>
        /// Resource schema for AWS::AppFlow::Connector
        /// </summary>
        public static Task<GetConnectorResult> InvokeAsync(GetConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectorResult>("aws-native:appflow:getConnector", args ?? new GetConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::AppFlow::Connector
        /// </summary>
        public static Output<GetConnectorResult> Invoke(GetConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectorResult>("aws-native:appflow:getConnector", args ?? new GetConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        ///  The name of the connector. The name is unique for each ConnectorRegistration in your AWS account.
        /// </summary>
        [Input("connectorLabel", required: true)]
        public string ConnectorLabel { get; set; } = null!;

        public GetConnectorArgs()
        {
        }
        public static new GetConnectorArgs Empty => new GetConnectorArgs();
    }

    public sealed class GetConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        ///  The name of the connector. The name is unique for each ConnectorRegistration in your AWS account.
        /// </summary>
        [Input("connectorLabel", required: true)]
        public Input<string> ConnectorLabel { get; set; } = null!;

        public GetConnectorInvokeArgs()
        {
        }
        public static new GetConnectorInvokeArgs Empty => new GetConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectorResult
    {
        /// <summary>
        ///  The arn of the connector. The arn is unique for each ConnectorRegistration in your AWS account.
        /// </summary>
        public readonly string? ConnectorArn;
        /// <summary>
        /// Contains information about the configuration of the connector being registered.
        /// </summary>
        public readonly Outputs.ConnectorProvisioningConfig? ConnectorProvisioningConfig;
        /// <summary>
        /// The provisioning type of the connector. Currently the only supported value is LAMBDA. 
        /// </summary>
        public readonly string? ConnectorProvisioningType;
        /// <summary>
        /// A description about the connector that's being registered.
        /// </summary>
        public readonly string? Description;

        [OutputConstructor]
        private GetConnectorResult(
            string? connectorArn,

            Outputs.ConnectorProvisioningConfig? connectorProvisioningConfig,

            string? connectorProvisioningType,

            string? description)
        {
            ConnectorArn = connectorArn;
            ConnectorProvisioningConfig = connectorProvisioningConfig;
            ConnectorProvisioningType = connectorProvisioningType;
            Description = description;
        }
    }
}