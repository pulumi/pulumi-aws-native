// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ses.Inputs
{

    public sealed class MailManagerTrafficPolicyIngressStringExpressionArgs : global::Pulumi.ResourceArgs
    {
        [Input("evaluate", required: true)]
        public Input<Inputs.MailManagerTrafficPolicyIngressStringToEvaluatePropertiesArgs> Evaluate { get; set; } = null!;

        [Input("operator", required: true)]
        public Input<Pulumi.AwsNative.Ses.MailManagerTrafficPolicyIngressStringOperator> Operator { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public MailManagerTrafficPolicyIngressStringExpressionArgs()
        {
        }
        public static new MailManagerTrafficPolicyIngressStringExpressionArgs Empty => new MailManagerTrafficPolicyIngressStringExpressionArgs();
    }
}