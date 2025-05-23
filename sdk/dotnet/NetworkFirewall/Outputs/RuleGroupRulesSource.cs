// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class RuleGroupRulesSource
    {
        /// <summary>
        /// Stateful inspection criteria for a domain list rule group.
        /// </summary>
        public readonly Outputs.RuleGroupRulesSourceList? RulesSourceList;
        /// <summary>
        /// Stateful inspection criteria, provided in Suricata compatible rules. Suricata is an open-source threat detection framework that includes a standard rule-based language for network traffic inspection.
        /// 
        /// These rules contain the inspection criteria and the action to take for traffic that matches the criteria, so this type of rule group doesn't have a separate action setting.
        /// 
        /// &gt; You can't use the `priority` keyword if the `RuleOrder` option in `StatefulRuleOptions` is set to `STRICT_ORDER` .
        /// </summary>
        public readonly string? RulesString;
        /// <summary>
        /// An array of individual stateful rules inspection criteria to be used together in a stateful rule group. Use this option to specify simple Suricata rules with protocol, source and destination, ports, direction, and rule options. For information about the Suricata `Rules` format, see [Rules Format](https://docs.aws.amazon.com/https://suricata.readthedocs.io/en/suricata-7.0.3/rules/intro.html) .
        /// </summary>
        public readonly ImmutableArray<Outputs.RuleGroupStatefulRule> StatefulRules;
        /// <summary>
        /// Stateless inspection criteria to be used in a stateless rule group.
        /// </summary>
        public readonly Outputs.RuleGroupStatelessRulesAndCustomActions? StatelessRulesAndCustomActions;

        [OutputConstructor]
        private RuleGroupRulesSource(
            Outputs.RuleGroupRulesSourceList? rulesSourceList,

            string? rulesString,

            ImmutableArray<Outputs.RuleGroupStatefulRule> statefulRules,

            Outputs.RuleGroupStatelessRulesAndCustomActions? statelessRulesAndCustomActions)
        {
            RulesSourceList = rulesSourceList;
            RulesString = rulesString;
            StatefulRules = statefulRules;
            StatelessRulesAndCustomActions = statelessRulesAndCustomActions;
        }
    }
}
