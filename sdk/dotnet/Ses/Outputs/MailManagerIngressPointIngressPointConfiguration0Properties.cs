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
    public sealed class MailManagerIngressPointIngressPointConfiguration0Properties
    {
        public readonly string SmtpPassword;

        [OutputConstructor]
        private MailManagerIngressPointIngressPointConfiguration0Properties(string smtpPassword)
        {
            SmtpPassword = smtpPassword;
        }
    }
}