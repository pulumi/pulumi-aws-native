// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Action taken when Rule matches its condition.
    /// </summary>
    [OutputType]
    public sealed class WebAclRuleAction
    {
        public readonly Outputs.WebAclAllowAction? Allow;
        public readonly Outputs.WebAclBlockAction? Block;
        public readonly Outputs.WebAclCaptchaAction? Captcha;
        public readonly Outputs.WebAclChallengeAction? Challenge;
        public readonly Outputs.WebAclCountAction? Count;

        [OutputConstructor]
        private WebAclRuleAction(
            Outputs.WebAclAllowAction? allow,

            Outputs.WebAclBlockAction? block,

            Outputs.WebAclCaptchaAction? captcha,

            Outputs.WebAclChallengeAction? challenge,

            Outputs.WebAclCountAction? count)
        {
            Allow = allow;
            Block = block;
            Captcha = captcha;
            Challenge = challenge;
            Count = count;
        }
    }
}