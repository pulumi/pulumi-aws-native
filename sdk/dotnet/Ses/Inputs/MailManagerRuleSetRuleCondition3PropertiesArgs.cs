// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRuleSetRuleCondition3PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("ipExpression", required: true)]
        public Input<Inputs.MailManagerRuleSetRuleIpExpressionArgs> IpExpression { get; set; } = null!;

        public MailManagerRuleSetRuleCondition3PropertiesArgs()
        {
        }
        public static new MailManagerRuleSetRuleCondition3PropertiesArgs Empty => new MailManagerRuleSetRuleCondition3PropertiesArgs();
    }
}
