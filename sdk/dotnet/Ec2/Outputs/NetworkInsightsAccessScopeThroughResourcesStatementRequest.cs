// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    [OutputType]
    public sealed class NetworkInsightsAccessScopeThroughResourcesStatementRequest
    {
        /// <summary>
        /// The resource statement.
        /// </summary>
        public readonly Outputs.NetworkInsightsAccessScopeResourceStatementRequest? ResourceStatement;

        [OutputConstructor]
        private NetworkInsightsAccessScopeThroughResourcesStatementRequest(Outputs.NetworkInsightsAccessScopeResourceStatementRequest? resourceStatement)
        {
            ResourceStatement = resourceStatement;
        }
    }
}
