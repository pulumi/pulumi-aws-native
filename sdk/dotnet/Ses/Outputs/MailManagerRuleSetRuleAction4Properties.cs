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
    public sealed class MailManagerRuleSetRuleAction4Properties
    {
        public readonly Outputs.MailManagerRuleSetSendAction Send;

        [OutputConstructor]
        private MailManagerRuleSetRuleAction4Properties(Outputs.MailManagerRuleSetSendAction send)
        {
            Send = send;
        }
    }
}
