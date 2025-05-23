// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// Specifies an IPv6 address in an Amazon EC2 launch template.
    ///  ``Ipv6Add`` is a property of [AWS::EC2::LaunchTemplate NetworkInterface](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html).
    /// </summary>
    public sealed class LaunchTemplateIpv6AddArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You can't use this option if you're specifying a number of IPv6 addresses.
        /// </summary>
        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        public LaunchTemplateIpv6AddArgs()
        {
        }
        public static new LaunchTemplateIpv6AddArgs Empty => new LaunchTemplateIpv6AddArgs();
    }
}
