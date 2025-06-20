// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    /// <summary>
    /// Configures the options for on-source DDoS protection provided by supported resource type.
    /// </summary>
    public sealed class WebAclOnSourceDDoSProtectionConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The level of DDoS protection that applies to web ACLs associated with Application Load Balancers. `ACTIVE_UNDER_DDOS` protection is enabled by default whenever a web ACL is associated with an Application Load Balancer. In the event that an Application Load Balancer experiences high-load conditions or suspected DDoS attacks, the `ACTIVE_UNDER_DDOS` protection automatically rate limits traffic from known low reputation sources without disrupting Application Load Balancer availability. `ALWAYS_ON` protection provides constant, always-on monitoring of known low reputation sources for suspected DDoS attacks. While this provides a higher level of protection, there may be potential impacts on legitimate traffic.
        /// </summary>
        [Input("albLowReputationMode", required: true)]
        public Input<Pulumi.AwsNative.WaFv2.WebAclOnSourceDDoSProtectionConfigAlbLowReputationMode> AlbLowReputationMode { get; set; } = null!;

        public WebAclOnSourceDDoSProtectionConfigArgs()
        {
        }
        public static new WebAclOnSourceDDoSProtectionConfigArgs Empty => new WebAclOnSourceDDoSProtectionConfigArgs();
    }
}
