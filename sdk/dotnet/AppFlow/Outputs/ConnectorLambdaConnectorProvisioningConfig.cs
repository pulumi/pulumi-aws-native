// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    /// <summary>
    /// Contains information about the configuration of the lambda which is being registered as the connector.
    /// </summary>
    [OutputType]
    public sealed class ConnectorLambdaConnectorProvisioningConfig
    {
        /// <summary>
        /// Lambda ARN of the connector being registered.
        /// </summary>
        public readonly string LambdaArn;

        [OutputConstructor]
        private ConnectorLambdaConnectorProvisioningConfig(string lambdaArn)
        {
            LambdaArn = lambdaArn;
        }
    }
}