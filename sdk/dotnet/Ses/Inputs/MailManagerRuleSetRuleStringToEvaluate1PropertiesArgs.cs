// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRuleSetRuleStringToEvaluate1PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("mimeHeaderAttribute", required: true)]
        public Input<string> MimeHeaderAttribute { get; set; } = null!;

        public MailManagerRuleSetRuleStringToEvaluate1PropertiesArgs()
        {
        }
        public static new MailManagerRuleSetRuleStringToEvaluate1PropertiesArgs Empty => new MailManagerRuleSetRuleStringToEvaluate1PropertiesArgs();
    }
}