// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerRuleSetAddHeaderActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("headerName", required: true)]
        public Input<string> HeaderName { get; set; } = null!;

        [Input("headerValue", required: true)]
        public Input<string> HeaderValue { get; set; } = null!;

        public MailManagerRuleSetAddHeaderActionArgs()
        {
        }
        public static new MailManagerRuleSetAddHeaderActionArgs Empty => new MailManagerRuleSetAddHeaderActionArgs();
    }
}