// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRuleSetRuleBooleanToEvaluatePropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("attribute", required: true)]
        public Input<Pulumi.AwsNative.Ses.MailManagerRuleSetRuleBooleanEmailAttribute> Attribute { get; set; } = null!;

        public MailManagerRuleSetRuleBooleanToEvaluatePropertiesArgs()
        {
        }
        public static new MailManagerRuleSetRuleBooleanToEvaluatePropertiesArgs Empty => new MailManagerRuleSetRuleBooleanToEvaluatePropertiesArgs();
    }
}
