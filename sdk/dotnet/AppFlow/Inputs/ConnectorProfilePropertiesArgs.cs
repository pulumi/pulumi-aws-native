// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    /// <summary>
    /// Connector specific properties needed to create connector profile - currently not needed for Amplitude, Trendmicro, Googleanalytics and Singular
    /// </summary>
    public sealed class ConnectorProfilePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The properties required by the custom connector.
        /// </summary>
        [Input("customConnector")]
        public Input<Inputs.ConnectorProfileCustomConnectorProfilePropertiesArgs>? CustomConnector { get; set; }

        /// <summary>
        /// The connector-specific properties required by Datadog.
        /// </summary>
        [Input("datadog")]
        public Input<Inputs.ConnectorProfileDatadogConnectorProfilePropertiesArgs>? Datadog { get; set; }

        /// <summary>
        /// The connector-specific properties required by Dynatrace.
        /// </summary>
        [Input("dynatrace")]
        public Input<Inputs.ConnectorProfileDynatraceConnectorProfilePropertiesArgs>? Dynatrace { get; set; }

        /// <summary>
        /// The connector-specific properties required by Infor Nexus.
        /// </summary>
        [Input("inforNexus")]
        public Input<Inputs.ConnectorProfileInforNexusConnectorProfilePropertiesArgs>? InforNexus { get; set; }

        /// <summary>
        /// The connector-specific properties required by Marketo.
        /// </summary>
        [Input("marketo")]
        public Input<Inputs.ConnectorProfileMarketoConnectorProfilePropertiesArgs>? Marketo { get; set; }

        /// <summary>
        /// The connector-specific properties required by Salesforce Pardot.
        /// </summary>
        [Input("pardot")]
        public Input<Inputs.ConnectorProfilePardotConnectorProfilePropertiesArgs>? Pardot { get; set; }

        /// <summary>
        /// The connector-specific properties required by Amazon Redshift.
        /// </summary>
        [Input("redshift")]
        public Input<Inputs.ConnectorProfileRedshiftConnectorProfilePropertiesArgs>? Redshift { get; set; }

        /// <summary>
        /// The connector-specific properties required by Salesforce.
        /// </summary>
        [Input("salesforce")]
        public Input<Inputs.ConnectorProfileSalesforceConnectorProfilePropertiesArgs>? Salesforce { get; set; }

        /// <summary>
        /// The connector-specific profile properties required when using SAPOData.
        /// </summary>
        [Input("sapoData")]
        public Input<Inputs.ConnectorProfileSapoDataConnectorProfilePropertiesArgs>? SapoData { get; set; }

        /// <summary>
        /// The connector-specific properties required by serviceNow.
        /// </summary>
        [Input("serviceNow")]
        public Input<Inputs.ConnectorProfileServiceNowConnectorProfilePropertiesArgs>? ServiceNow { get; set; }

        /// <summary>
        /// The connector-specific properties required by Slack.
        /// </summary>
        [Input("slack")]
        public Input<Inputs.ConnectorProfileSlackConnectorProfilePropertiesArgs>? Slack { get; set; }

        /// <summary>
        /// The connector-specific properties required by Snowflake.
        /// </summary>
        [Input("snowflake")]
        public Input<Inputs.ConnectorProfileSnowflakeConnectorProfilePropertiesArgs>? Snowflake { get; set; }

        /// <summary>
        /// The connector-specific properties required by Veeva.
        /// </summary>
        [Input("veeva")]
        public Input<Inputs.ConnectorProfileVeevaConnectorProfilePropertiesArgs>? Veeva { get; set; }

        /// <summary>
        /// The connector-specific properties required by Zendesk.
        /// </summary>
        [Input("zendesk")]
        public Input<Inputs.ConnectorProfileZendeskConnectorProfilePropertiesArgs>? Zendesk { get; set; }

        public ConnectorProfilePropertiesArgs()
        {
        }
        public static new ConnectorProfilePropertiesArgs Empty => new ConnectorProfilePropertiesArgs();
    }
}
