// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WafRegional.Inputs
{

    public sealed class IpSetIpSetDescriptorArgs : global::Pulumi.ResourceArgs
    {
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public IpSetIpSetDescriptorArgs()
        {
        }
        public static new IpSetIpSetDescriptorArgs Empty => new IpSetIpSetDescriptorArgs();
    }
}