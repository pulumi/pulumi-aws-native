// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53Resolver.Outputs
{

    /// <summary>
    /// Firewall Rule associating the Rule Group to a Domain List
    /// </summary>
    [OutputType]
    public sealed class FirewallRuleGroupFirewallRule
    {
        /// <summary>
        /// Rule Action
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleAction Action;
        /// <summary>
        /// BlockOverrideDnsType
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleBlockOverrideDnsType? BlockOverrideDnsType;
        /// <summary>
        /// BlockOverrideDomain
        /// </summary>
        public readonly string? BlockOverrideDomain;
        /// <summary>
        /// BlockOverrideTtl
        /// </summary>
        public readonly int? BlockOverrideTtl;
        /// <summary>
        /// BlockResponse
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleBlockResponse? BlockResponse;
        /// <summary>
        /// FirewallDomainRedirectionAction
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleConfidenceThreshold? ConfidenceThreshold;
        /// <summary>
        /// FirewallDomainRedirectionAction
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleDnsThreatProtection? DnsThreatProtection;
        /// <summary>
        /// ResourceId
        /// </summary>
        public readonly string? FirewallDomainListId;
        /// <summary>
        /// FirewallDomainRedirectionAction
        /// </summary>
        public readonly Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleFirewallDomainRedirectionAction? FirewallDomainRedirectionAction;
        /// <summary>
        /// ResourceId
        /// </summary>
        public readonly string? FirewallThreatProtectionId;
        /// <summary>
        /// Rule Priority
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Qtype
        /// </summary>
        public readonly string? Qtype;

        [OutputConstructor]
        private FirewallRuleGroupFirewallRule(
            Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleAction action,

            Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleBlockOverrideDnsType? blockOverrideDnsType,

            string? blockOverrideDomain,

            int? blockOverrideTtl,

            Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleBlockResponse? blockResponse,

            Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleConfidenceThreshold? confidenceThreshold,

            Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleDnsThreatProtection? dnsThreatProtection,

            string? firewallDomainListId,

            Pulumi.AwsNative.Route53Resolver.FirewallRuleGroupFirewallRuleFirewallDomainRedirectionAction? firewallDomainRedirectionAction,

            string? firewallThreatProtectionId,

            int priority,

            string? qtype)
        {
            Action = action;
            BlockOverrideDnsType = blockOverrideDnsType;
            BlockOverrideDomain = blockOverrideDomain;
            BlockOverrideTtl = blockOverrideTtl;
            BlockResponse = blockResponse;
            ConfidenceThreshold = confidenceThreshold;
            DnsThreatProtection = dnsThreatProtection;
            FirewallDomainListId = firewallDomainListId;
            FirewallDomainRedirectionAction = firewallDomainRedirectionAction;
            FirewallThreatProtectionId = firewallThreatProtectionId;
            Priority = priority;
            Qtype = qtype;
        }
    }
}
