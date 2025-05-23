// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CustomerProfiles.Outputs
{

    [OutputType]
    public sealed class IntegrationSourceFlowConfig
    {
        /// <summary>
        /// The name of the Amazon AppFlow connector profile. This name must be unique for each connector profile in the AWS account .
        /// </summary>
        public readonly string? ConnectorProfileName;
        /// <summary>
        /// The type of connector, such as Salesforce, Marketo, and so on.
        /// </summary>
        public readonly Pulumi.AwsNative.CustomerProfiles.IntegrationConnectorType ConnectorType;
        /// <summary>
        /// Defines the configuration for a scheduled incremental data pull. If a valid configuration is provided, the fields specified in the configuration are used when querying for the incremental data pull.
        /// </summary>
        public readonly Outputs.IntegrationIncrementalPullConfig? IncrementalPullConfig;
        /// <summary>
        /// Specifies the information that is required to query a particular source connector.
        /// </summary>
        public readonly Outputs.IntegrationSourceConnectorProperties SourceConnectorProperties;

        [OutputConstructor]
        private IntegrationSourceFlowConfig(
            string? connectorProfileName,

            Pulumi.AwsNative.CustomerProfiles.IntegrationConnectorType connectorType,

            Outputs.IntegrationIncrementalPullConfig? incrementalPullConfig,

            Outputs.IntegrationSourceConnectorProperties sourceConnectorProperties)
        {
            ConnectorProfileName = connectorProfileName;
            ConnectorType = connectorType;
            IncrementalPullConfig = incrementalPullConfig;
            SourceConnectorProperties = sourceConnectorProperties;
        }
    }
}
