// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRuleSetRuleBooleanExpressionArgs : global::Pulumi.ResourceArgs
    {
        [Input("evaluate", required: true)]
        public object Evaluate { get; set; } = null!;

        [Input("operator", required: true)]
        public Input<Pulumi.AwsNative.Ses.MailManagerRuleSetRuleBooleanOperator> Operator { get; set; } = null!;

        public MailManagerRuleSetRuleBooleanExpressionArgs()
        {
        }
        public static new MailManagerRuleSetRuleBooleanExpressionArgs Empty => new MailManagerRuleSetRuleBooleanExpressionArgs();
    }
}
