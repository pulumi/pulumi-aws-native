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
    public sealed class MailManagerRuleSetRuleStringExpression
    {
        public readonly Union<Outputs.MailManagerRuleSetRuleStringToEvaluate0Properties, Outputs.MailManagerRuleSetRuleStringToEvaluate1Properties> Evaluate;
        public readonly Pulumi.AwsNative.Ses.MailManagerRuleSetRuleStringOperator Operator;
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private MailManagerRuleSetRuleStringExpression(
            Union<Outputs.MailManagerRuleSetRuleStringToEvaluate0Properties, Outputs.MailManagerRuleSetRuleStringToEvaluate1Properties> evaluate,

            Pulumi.AwsNative.Ses.MailManagerRuleSetRuleStringOperator @operator,

            ImmutableArray<string> values)
        {
            Evaluate = evaluate;
            Operator = @operator;
            Values = values;
        }
    }
}
