// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Outputs
{

    [OutputType]
    public sealed class MailManagerTrafficPolicyIngressBooleanExpression
    {
        public readonly Outputs.MailManagerTrafficPolicyIngressBooleanToEvaluateProperties Evaluate;
        public readonly Pulumi.AwsNative.Ses.MailManagerTrafficPolicyIngressBooleanOperator Operator;

        [OutputConstructor]
        private MailManagerTrafficPolicyIngressBooleanExpression(
            Outputs.MailManagerTrafficPolicyIngressBooleanToEvaluateProperties evaluate,

            Pulumi.AwsNative.Ses.MailManagerTrafficPolicyIngressBooleanOperator @operator)
        {
            Evaluate = evaluate;
            Operator = @operator;
        }
    }
}
