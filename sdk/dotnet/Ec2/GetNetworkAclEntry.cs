// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetNetworkAclEntry
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::NetworkAclEntry
        /// </summary>
        public static Task<GetNetworkAclEntryResult> InvokeAsync(GetNetworkAclEntryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkAclEntryResult>("aws-native:ec2:getNetworkAclEntry", args ?? new GetNetworkAclEntryArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::NetworkAclEntry
        /// </summary>
        public static Output<GetNetworkAclEntryResult> Invoke(GetNetworkAclEntryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkAclEntryResult>("aws-native:ec2:getNetworkAclEntry", args ?? new GetNetworkAclEntryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkAclEntryArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetNetworkAclEntryArgs()
        {
        }
        public static new GetNetworkAclEntryArgs Empty => new GetNetworkAclEntryArgs();
    }

    public sealed class GetNetworkAclEntryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetNetworkAclEntryInvokeArgs()
        {
        }
        public static new GetNetworkAclEntryInvokeArgs Empty => new GetNetworkAclEntryInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkAclEntryResult
    {
        public readonly string? CidrBlock;
        public readonly Outputs.NetworkAclEntryIcmp? Icmp;
        public readonly string? Id;
        public readonly string? Ipv6CidrBlock;
        public readonly Outputs.NetworkAclEntryPortRange? PortRange;
        public readonly int? Protocol;
        public readonly string? RuleAction;

        [OutputConstructor]
        private GetNetworkAclEntryResult(
            string? cidrBlock,

            Outputs.NetworkAclEntryIcmp? icmp,

            string? id,

            string? ipv6CidrBlock,

            Outputs.NetworkAclEntryPortRange? portRange,

            int? protocol,

            string? ruleAction)
        {
            CidrBlock = cidrBlock;
            Icmp = icmp;
            Id = id;
            Ipv6CidrBlock = ipv6CidrBlock;
            PortRange = portRange;
            Protocol = protocol;
            RuleAction = ruleAction;
        }
    }
}