// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Outputs
{

    [OutputType]
    public sealed class MailManagerRuleSetDeliverToMailboxAction
    {
        public readonly Pulumi.AwsNative.Ses.MailManagerRuleSetActionFailurePolicy? ActionFailurePolicy;
        public readonly string MailboxArn;
        public readonly string RoleArn;

        [OutputConstructor]
        private MailManagerRuleSetDeliverToMailboxAction(
            Pulumi.AwsNative.Ses.MailManagerRuleSetActionFailurePolicy? actionFailurePolicy,

            string mailboxArn,

            string roleArn)
        {
            ActionFailurePolicy = actionFailurePolicy;
            MailboxArn = mailboxArn;
            RoleArn = roleArn;
        }
    }
}