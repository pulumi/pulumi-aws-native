// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Action override for rules in the rule group.
    /// </summary>
    [OutputType]
    public sealed class WebAclRuleActionOverride
    {
        /// <summary>
        /// The override action to use, in place of the configured action of the rule in the rule group.
        /// </summary>
        public readonly Outputs.WebAclRuleAction ActionToUse;
        /// <summary>
        /// The name of the rule to override.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private WebAclRuleActionOverride(
            Outputs.WebAclRuleAction actionToUse,

            string name)
        {
            ActionToUse = actionToUse;
            Name = name;
        }
    }
}
