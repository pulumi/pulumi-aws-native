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
    /// Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:SubnetNetworkAclAssociation")]
    public partial class SubnetNetworkAclAssociation : global::Pulumi.CustomResource
    {
        [Output("associationId")]
        public Output<string> AssociationId { get; private set; } = null!;

        /// <summary>
        /// The ID of the network ACL
        /// </summary>
        [Output("networkAclId")]
        public Output<string> NetworkAclId { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a SubnetNetworkAclAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SubnetNetworkAclAssociation(string name, SubnetNetworkAclAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:SubnetNetworkAclAssociation", name, args ?? new SubnetNetworkAclAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SubnetNetworkAclAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:SubnetNetworkAclAssociation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing SubnetNetworkAclAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SubnetNetworkAclAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SubnetNetworkAclAssociation(name, id, options);
        }
    }

    public sealed class SubnetNetworkAclAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the network ACL
        /// </summary>
        [Input("networkAclId", required: true)]
        public Input<string> NetworkAclId { get; set; } = null!;

        /// <summary>
        /// The ID of the subnet
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        public SubnetNetworkAclAssociationArgs()
        {
        }
        public static new SubnetNetworkAclAssociationArgs Empty => new SubnetNetworkAclAssociationArgs();
    }
}