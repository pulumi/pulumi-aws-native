// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    /// <summary>
    /// The AWS::EC2::SubnetCidrBlock resource creates association between subnet and IPv6 CIDR
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:SubnetCidrBlock")]
    public partial class SubnetCidrBlock : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length
        /// </summary>
        [Output("ipv6CidrBlock")]
        public Output<string> Ipv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a SubnetCidrBlock resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubnetCidrBlock(string name, SubnetCidrBlockArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:SubnetCidrBlock", name, args ?? new SubnetCidrBlockArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubnetCidrBlock(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:SubnetCidrBlock", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SubnetCidrBlock resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubnetCidrBlock Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SubnetCidrBlock(name, id, options);
        }
    }

    public sealed class SubnetCidrBlockArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length
        /// </summary>
        [Input("ipv6CidrBlock", required: true)]
        public Input<string> Ipv6CidrBlock { get; set; } = null!;

        /// <summary>
        /// The ID of the subnet
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public SubnetCidrBlockArgs()
        {
        }
        public static new SubnetCidrBlockArgs Empty => new SubnetCidrBlockArgs();
    }
}