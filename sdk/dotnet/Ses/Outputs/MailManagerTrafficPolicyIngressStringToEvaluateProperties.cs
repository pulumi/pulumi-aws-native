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
    public sealed class MailManagerTrafficPolicyIngressStringToEvaluateProperties
    {
        public readonly Pulumi.AwsNative.Ses.MailManagerTrafficPolicyIngressStringEmailAttribute Attribute;

        [OutputConstructor]
        private MailManagerTrafficPolicyIngressStringToEvaluateProperties(Pulumi.AwsNative.Ses.MailManagerTrafficPolicyIngressStringEmailAttribute attribute)
        {
            Attribute = attribute;
        }
    }
}
