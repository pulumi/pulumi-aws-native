// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Waf.Outputs
{

    [OutputType]
    public sealed class WebAclActivatedRule
    {
        public readonly Outputs.WebAclWafAction? Action;
        public readonly int Priority;
        public readonly string RuleId;

        [OutputConstructor]
        private WebAclActivatedRule(
            Outputs.WebAclWafAction? action,

            int priority,

            string ruleId)
        {
            Action = action;
            Priority = priority;
            RuleId = ruleId;
        }
    }
}