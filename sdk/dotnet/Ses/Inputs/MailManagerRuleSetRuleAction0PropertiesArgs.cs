// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRuleSetRuleAction0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("drop", required: true)]
        public Input<Inputs.MailManagerRuleSetDropActionArgs> Drop { get; set; } = null!;

        public MailManagerRuleSetRuleAction0PropertiesArgs()
        {
        }
        public static new MailManagerRuleSetRuleAction0PropertiesArgs Empty => new MailManagerRuleSetRuleAction0PropertiesArgs();
    }
}