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
    public sealed class MailManagerTrafficPolicyPolicyCondition2Properties
    {
        public readonly Outputs.MailManagerTrafficPolicyIngressTlsProtocolExpression TlsExpression;

        [OutputConstructor]
        private MailManagerTrafficPolicyPolicyCondition2Properties(Outputs.MailManagerTrafficPolicyIngressTlsProtocolExpression tlsExpression)
        {
            TlsExpression = tlsExpression;
        }
    }
}
