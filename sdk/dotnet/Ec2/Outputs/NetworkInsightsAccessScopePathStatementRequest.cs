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
    public sealed class NetworkInsightsAccessScopePathStatementRequest
    {
        /// <summary>
        /// The packet header statement.
        /// </summary>
        public readonly Outputs.NetworkInsightsAccessScopePacketHeaderStatementRequest? PacketHeaderStatement;
        /// <summary>
        /// The resource statement.
        /// </summary>
        public readonly Outputs.NetworkInsightsAccessScopeResourceStatementRequest? ResourceStatement;

        [OutputConstructor]
        private NetworkInsightsAccessScopePathStatementRequest(
            Outputs.NetworkInsightsAccessScopePacketHeaderStatementRequest? packetHeaderStatement,

            Outputs.NetworkInsightsAccessScopeResourceStatementRequest? resourceStatement)
        {
            PacketHeaderStatement = packetHeaderStatement;
            ResourceStatement = resourceStatement;
        }
    }
}
