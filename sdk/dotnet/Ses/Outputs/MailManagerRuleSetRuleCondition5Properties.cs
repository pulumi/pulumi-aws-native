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
    public sealed class MailManagerRuleSetRuleCondition5Properties
    {
        public readonly Outputs.MailManagerRuleSetRuleDmarcExpression DmarcExpression;

        [OutputConstructor]
        private MailManagerRuleSetRuleCondition5Properties(Outputs.MailManagerRuleSetRuleDmarcExpression dmarcExpression)
        {
            DmarcExpression = dmarcExpression;
        }
    }
}
