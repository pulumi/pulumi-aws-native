// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fms.Inputs
{

    /// <summary>
    /// Network ACL entry.
    /// </summary>
    public sealed class PolicyNetworkAclEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// CIDR block.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Whether the entry is an egress entry.
        /// </summary>
        [Input("egress", required: true)]
        public Input<bool> Egress { get; set; } = null!;

        /// <summary>
        /// ICMP type and code.
        /// </summary>
        [Input("icmpTypeCode")]
        public Input<Inputs.PolicyNetworkAclEntryIcmpTypeCodePropertiesArgs>? IcmpTypeCode { get; set; }

        /// <summary>
        /// IPv6 CIDR block.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// Port range.
        /// </summary>
        [Input("portRange")]
        public Input<Inputs.PolicyNetworkAclEntryPortRangePropertiesArgs>? PortRange { get; set; }

        /// <summary>
        /// Protocol.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Rule Action.
        /// </summary>
        [Input("ruleAction", required: true)]
        public Input<Pulumi.AwsNative.Fms.PolicyNetworkAclEntryRuleAction> RuleAction { get; set; } = null!;

        public PolicyNetworkAclEntryArgs()
        {
        }
        public static new PolicyNetworkAclEntryArgs Empty => new PolicyNetworkAclEntryArgs();
    }
}