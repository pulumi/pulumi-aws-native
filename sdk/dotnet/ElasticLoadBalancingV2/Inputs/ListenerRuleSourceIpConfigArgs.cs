// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Inputs
{

    public sealed class ListenerRuleSourceIpConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("values")]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ListenerRuleSourceIpConfigArgs()
        {
        }
        public static new ListenerRuleSourceIpConfigArgs Empty => new ListenerRuleSourceIpConfigArgs();
    }
}