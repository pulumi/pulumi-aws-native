// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRuleSetRuleAction3PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("writeToS3", required: true)]
        public Input<Inputs.MailManagerRuleSetS3ActionArgs> WriteToS3 { get; set; } = null!;

        public MailManagerRuleSetRuleAction3PropertiesArgs()
        {
        }
        public static new MailManagerRuleSetRuleAction3PropertiesArgs Empty => new MailManagerRuleSetRuleAction3PropertiesArgs();
    }
}