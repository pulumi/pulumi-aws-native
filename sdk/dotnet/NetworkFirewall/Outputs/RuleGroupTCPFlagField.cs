// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class RuleGroupTCPFlagField
    {
        public readonly ImmutableArray<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag> Flags;
        public readonly ImmutableArray<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag> Masks;

        [OutputConstructor]
        private RuleGroupTCPFlagField(
            ImmutableArray<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag> flags,

            ImmutableArray<Pulumi.AwsNative.NetworkFirewall.RuleGroupTCPFlag> masks)
        {
            Flags = flags;
            Masks = masks;
        }
    }
}