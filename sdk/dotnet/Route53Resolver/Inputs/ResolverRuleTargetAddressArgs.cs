// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53Resolver.Inputs
{

    public sealed class ResolverRuleTargetAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// One IP address that you want to forward DNS queries to. You can specify only IPv4 addresses. 
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// The port at Ip that you want to forward DNS queries to. 
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        public ResolverRuleTargetAddressArgs()
        {
        }
        public static new ResolverRuleTargetAddressArgs Empty => new ResolverRuleTargetAddressArgs();
    }
}