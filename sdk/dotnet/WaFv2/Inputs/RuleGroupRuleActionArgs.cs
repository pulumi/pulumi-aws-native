// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// Action taken when Rule matches its condition.
    /// </summary>
    public sealed class RuleGroupRuleActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("allow")]
        public Input<Inputs.RuleGroupAllowActionArgs>? Allow { get; set; }

        [Input("block")]
        public Input<Inputs.RuleGroupBlockActionArgs>? Block { get; set; }

        [Input("captcha")]
        public Input<Inputs.RuleGroupCaptchaActionArgs>? Captcha { get; set; }

        [Input("challenge")]
        public Input<Inputs.RuleGroupChallengeActionArgs>? Challenge { get; set; }

        [Input("count")]
        public Input<Inputs.RuleGroupCountActionArgs>? Count { get; set; }

        public RuleGroupRuleActionArgs()
        {
        }
        public static new RuleGroupRuleActionArgs Empty => new RuleGroupRuleActionArgs();
    }
}