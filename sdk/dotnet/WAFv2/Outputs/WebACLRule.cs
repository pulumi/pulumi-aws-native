// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Rule of WebACL that contains condition and action.
    /// </summary>
    [OutputType]
    public sealed class WebACLRule
    {
        public readonly Outputs.WebACLRuleAction? Action;
        public readonly Outputs.WebACLCaptchaConfig? CaptchaConfig;
        public readonly Outputs.WebACLChallengeConfig? ChallengeConfig;
        public readonly string Name;
        public readonly Outputs.WebACLOverrideAction? OverrideAction;
        public readonly int Priority;
        /// <summary>
        /// Collection of Rule Labels.
        /// </summary>
        public readonly ImmutableArray<Outputs.WebACLLabel> RuleLabels;
        public readonly Outputs.WebACLStatement Statement;
        public readonly Outputs.WebACLVisibilityConfig VisibilityConfig;

        [OutputConstructor]
        private WebACLRule(
            Outputs.WebACLRuleAction? action,

            Outputs.WebACLCaptchaConfig? captchaConfig,

            Outputs.WebACLChallengeConfig? challengeConfig,

            string name,

            Outputs.WebACLOverrideAction? overrideAction,

            int priority,

            ImmutableArray<Outputs.WebACLLabel> ruleLabels,

            Outputs.WebACLStatement statement,

            Outputs.WebACLVisibilityConfig visibilityConfig)
        {
            Action = action;
            CaptchaConfig = captchaConfig;
            ChallengeConfig = challengeConfig;
            Name = name;
            OverrideAction = overrideAction;
            Priority = priority;
            RuleLabels = ruleLabels;
            Statement = statement;
            VisibilityConfig = visibilityConfig;
        }
    }
}