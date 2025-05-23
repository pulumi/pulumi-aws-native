// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GlobalAccelerator
{
    public static class GetAccelerator
    {
        /// <summary>
        /// Resource Type definition for AWS::GlobalAccelerator::Accelerator
        /// </summary>
        public static Task<GetAcceleratorResult> InvokeAsync(GetAcceleratorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAcceleratorResult>("aws-native:globalaccelerator:getAccelerator", args ?? new GetAcceleratorArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::GlobalAccelerator::Accelerator
        /// </summary>
        public static Output<GetAcceleratorResult> Invoke(GetAcceleratorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAcceleratorResult>("aws-native:globalaccelerator:getAccelerator", args ?? new GetAcceleratorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::GlobalAccelerator::Accelerator
        /// </summary>
        public static Output<GetAcceleratorResult> Invoke(GetAcceleratorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAcceleratorResult>("aws-native:globalaccelerator:getAccelerator", args ?? new GetAcceleratorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAcceleratorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the accelerator.
        /// </summary>
        [Input("acceleratorArn", required: true)]
        public string AcceleratorArn { get; set; } = null!;

        public GetAcceleratorArgs()
        {
        }
        public static new GetAcceleratorArgs Empty => new GetAcceleratorArgs();
    }

    public sealed class GetAcceleratorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the accelerator.
        /// </summary>
        [Input("acceleratorArn", required: true)]
        public Input<string> AcceleratorArn { get; set; } = null!;

        public GetAcceleratorInvokeArgs()
        {
        }
        public static new GetAcceleratorInvokeArgs Empty => new GetAcceleratorInvokeArgs();
    }


    [OutputType]
    public sealed class GetAcceleratorResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the accelerator.
        /// </summary>
        public readonly string? AcceleratorArn;
        /// <summary>
        /// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 addresses.
        /// </summary>
        public readonly string? DnsName;
        /// <summary>
        /// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IPv4 and IPv6 addresses.
        /// </summary>
        public readonly string? DualStackDnsName;
        /// <summary>
        /// Indicates whether an accelerator is enabled. The value is true or false.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// IP Address type.
        /// </summary>
        public readonly Pulumi.AwsNative.GlobalAccelerator.AcceleratorIpAddressType? IpAddressType;
        /// <summary>
        /// The IP addresses from BYOIP Prefix pool.
        /// </summary>
        public readonly ImmutableArray<string> IpAddresses;
        /// <summary>
        /// The IPv4 addresses assigned to the accelerator.
        /// </summary>
        public readonly ImmutableArray<string> Ipv4Addresses;
        /// <summary>
        /// The IPv6 addresses assigned if the accelerator is dualstack
        /// </summary>
        public readonly ImmutableArray<string> Ipv6Addresses;
        /// <summary>
        /// Name of accelerator.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Create tags for an accelerator.
        /// 
        /// For more information, see [Tagging](https://docs.aws.amazon.com/global-accelerator/latest/dg/tagging-in-global-accelerator.html) in the *AWS Global Accelerator Developer Guide* .
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetAcceleratorResult(
            string? acceleratorArn,

            string? dnsName,

            string? dualStackDnsName,

            bool? enabled,

            Pulumi.AwsNative.GlobalAccelerator.AcceleratorIpAddressType? ipAddressType,

            ImmutableArray<string> ipAddresses,

            ImmutableArray<string> ipv4Addresses,

            ImmutableArray<string> ipv6Addresses,

            string? name,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AcceleratorArn = acceleratorArn;
            DnsName = dnsName;
            DualStackDnsName = dualStackDnsName;
            Enabled = enabled;
            IpAddressType = ipAddressType;
            IpAddresses = ipAddresses;
            Ipv4Addresses = ipv4Addresses;
            Ipv6Addresses = ipv6Addresses;
            Name = name;
            Tags = tags;
        }
    }
}
