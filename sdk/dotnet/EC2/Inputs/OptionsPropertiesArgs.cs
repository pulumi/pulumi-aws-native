// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    /// <summary>
    /// The options for the transit gateway vpc attachment.
    /// </summary>
    public sealed class OptionsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable
        /// </summary>
        [Input("applianceModeSupport")]
        public Input<string>? ApplianceModeSupport { get; set; }

        /// <summary>
        /// Indicates whether to enable DNS Support for Vpc Attachment. Valid Values: enable | disable
        /// </summary>
        [Input("dnsSupport")]
        public Input<string>? DnsSupport { get; set; }

        /// <summary>
        /// Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable
        /// </summary>
        [Input("ipv6Support")]
        public Input<string>? Ipv6Support { get; set; }

        public OptionsPropertiesArgs()
        {
        }
        public static new OptionsPropertiesArgs Empty => new OptionsPropertiesArgs();
    }
}