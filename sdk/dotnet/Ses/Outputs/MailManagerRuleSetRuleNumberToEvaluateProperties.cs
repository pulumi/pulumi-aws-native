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
    public sealed class MailManagerRuleSetRuleNumberToEvaluateProperties
    {
        public readonly Pulumi.AwsNative.Ses.MailManagerRuleSetRuleNumberEmailAttribute Attribute;

        [OutputConstructor]
        private MailManagerRuleSetRuleNumberToEvaluateProperties(Pulumi.AwsNative.Ses.MailManagerRuleSetRuleNumberEmailAttribute attribute)
        {
            Attribute = attribute;
        }
    }
}
