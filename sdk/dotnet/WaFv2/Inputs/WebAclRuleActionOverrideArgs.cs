// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// Action override for rules in the rule group.
    /// </summary>
    public sealed class WebAclRuleActionOverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The override action to use, in place of the configured action of the rule in the rule group.
        /// </summary>
        [Input("actionToUse", required: true)]
        public Input<Inputs.WebAclRuleActionArgs> ActionToUse { get; set; } = null!;

        /// <summary>
        /// The name of the rule to override.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public WebAclRuleActionOverrideArgs()
        {
        }
        public static new WebAclRuleActionOverrideArgs Empty => new WebAclRuleActionOverrideArgs();
    }
}
