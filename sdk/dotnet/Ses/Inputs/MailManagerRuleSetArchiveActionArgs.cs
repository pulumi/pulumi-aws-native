// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRuleSetArchiveActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionFailurePolicy")]
        public Input<Pulumi.AwsNative.Ses.MailManagerRuleSetActionFailurePolicy>? ActionFailurePolicy { get; set; }

        [Input("targetArchive", required: true)]
        public Input<string> TargetArchive { get; set; } = null!;

        public MailManagerRuleSetArchiveActionArgs()
        {
        }
        public static new MailManagerRuleSetArchiveActionArgs Empty => new MailManagerRuleSetArchiveActionArgs();
    }
}