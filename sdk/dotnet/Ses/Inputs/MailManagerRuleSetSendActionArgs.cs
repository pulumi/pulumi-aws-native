// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRuleSetSendActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionFailurePolicy")]
        public Input<Pulumi.AwsNative.Ses.MailManagerRuleSetActionFailurePolicy>? ActionFailurePolicy { get; set; }

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public MailManagerRuleSetSendActionArgs()
        {
        }
        public static new MailManagerRuleSetSendActionArgs Empty => new MailManagerRuleSetSendActionArgs();
    }
}