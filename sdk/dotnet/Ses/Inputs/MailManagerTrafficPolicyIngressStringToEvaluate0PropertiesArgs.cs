// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerTrafficPolicyIngressStringToEvaluate0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("attribute", required: true)]
        public Input<Pulumi.AwsNative.Ses.MailManagerTrafficPolicyIngressStringEmailAttribute> Attribute { get; set; } = null!;

        public MailManagerTrafficPolicyIngressStringToEvaluate0PropertiesArgs()
        {
        }
        public static new MailManagerTrafficPolicyIngressStringToEvaluate0PropertiesArgs Empty => new MailManagerTrafficPolicyIngressStringToEvaluate0PropertiesArgs();
    }
}
