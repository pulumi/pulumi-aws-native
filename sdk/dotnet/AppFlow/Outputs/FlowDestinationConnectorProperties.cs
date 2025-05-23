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
    /// Destination connector details
    /// </summary>
    [OutputType]
    public sealed class FlowDestinationConnectorProperties
    {
        /// <summary>
        /// The properties that are required to query the custom Connector.
        /// </summary>
        public readonly Outputs.FlowCustomConnectorDestinationProperties? CustomConnector;
        /// <summary>
        /// The properties required to query Amazon EventBridge.
        /// </summary>
        public readonly Outputs.FlowEventBridgeDestinationProperties? EventBridge;
        /// <summary>
        /// The properties required to query Amazon Lookout for Metrics.
        /// </summary>
        public readonly Outputs.FlowLookoutMetricsDestinationProperties? LookoutMetrics;
        /// <summary>
        /// The properties required to query Marketo.
        /// </summary>
        public readonly Outputs.FlowMarketoDestinationProperties? Marketo;
        /// <summary>
        /// The properties required to query Amazon Redshift.
        /// </summary>
        public readonly Outputs.FlowRedshiftDestinationProperties? Redshift;
        /// <summary>
        /// The properties required to query Amazon S3.
        /// </summary>
        public readonly Outputs.FlowS3DestinationProperties? S3;
        /// <summary>
        /// The properties required to query Salesforce.
        /// </summary>
        public readonly Outputs.FlowSalesforceDestinationProperties? Salesforce;
        /// <summary>
        /// The properties required to query SAPOData.
        /// </summary>
        public readonly Outputs.FlowSapoDataDestinationProperties? SapoData;
        /// <summary>
        /// The properties required to query Snowflake.
        /// </summary>
        public readonly Outputs.FlowSnowflakeDestinationProperties? Snowflake;
        /// <summary>
        /// The properties required to query Upsolver.
        /// </summary>
        public readonly Outputs.FlowUpsolverDestinationProperties? Upsolver;
        /// <summary>
        /// The properties required to query Zendesk.
        /// </summary>
        public readonly Outputs.FlowZendeskDestinationProperties? Zendesk;

        [OutputConstructor]
        private FlowDestinationConnectorProperties(
            Outputs.FlowCustomConnectorDestinationProperties? customConnector,

            Outputs.FlowEventBridgeDestinationProperties? eventBridge,

            Outputs.FlowLookoutMetricsDestinationProperties? lookoutMetrics,

            Outputs.FlowMarketoDestinationProperties? marketo,

            Outputs.FlowRedshiftDestinationProperties? redshift,

            Outputs.FlowS3DestinationProperties? s3,

            Outputs.FlowSalesforceDestinationProperties? salesforce,

            Outputs.FlowSapoDataDestinationProperties? sapoData,

            Outputs.FlowSnowflakeDestinationProperties? snowflake,

            Outputs.FlowUpsolverDestinationProperties? upsolver,

            Outputs.FlowZendeskDestinationProperties? zendesk)
        {
            CustomConnector = customConnector;
            EventBridge = eventBridge;
            LookoutMetrics = lookoutMetrics;
            Marketo = marketo;
            Redshift = redshift;
            S3 = s3;
            Salesforce = salesforce;
            SapoData = sapoData;
            Snowflake = snowflake;
            Upsolver = upsolver;
            Zendesk = zendesk;
        }
    }
}
