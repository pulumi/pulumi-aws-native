// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Outputs
{

    [OutputType]
    public sealed class MailManagerRuleSetRuleCondition0Properties
    {
        public readonly Outputs.MailManagerRuleSetRuleBooleanExpression BooleanExpression;

        [OutputConstructor]
        private MailManagerRuleSetRuleCondition0Properties(Outputs.MailManagerRuleSetRuleBooleanExpression booleanExpression)
        {
            BooleanExpression = booleanExpression;
        }
    }
}