// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// Describes an association between a local gateway route table and a VPC.
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:LocalGatewayRouteTableVPCAssociation")]
    public partial class LocalGatewayRouteTableVPCAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the local gateway.
        /// </summary>
        [Output("localGatewayId")]
        public Output<string> LocalGatewayId { get; private set; } = null!;

        /// <summary>
        /// The ID of the local gateway route table.
        /// </summary>
        [Output("localGatewayRouteTableId")]
        public Output<string> LocalGatewayRouteTableId { get; private set; } = null!;

        /// <summary>
        /// The ID of the association.
        /// </summary>
        [Output("localGatewayRouteTableVpcAssociationId")]
        public Output<string> LocalGatewayRouteTableVpcAssociationId { get; private set; } = null!;

        /// <summary>
        /// The state of the association.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The tags for the association.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.LocalGatewayRouteTableVPCAssociationTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a LocalGatewayRouteTableVPCAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocalGatewayRouteTableVPCAssociation(string name, LocalGatewayRouteTableVPCAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:LocalGatewayRouteTableVPCAssociation", name, args ?? new LocalGatewayRouteTableVPCAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocalGatewayRouteTableVPCAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:LocalGatewayRouteTableVPCAssociation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LocalGatewayRouteTableVPCAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocalGatewayRouteTableVPCAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LocalGatewayRouteTableVPCAssociation(name, id, options);
        }
    }

    public sealed class LocalGatewayRouteTableVPCAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the local gateway route table.
        /// </summary>
        [Input("localGatewayRouteTableId", required: true)]
        public Input<string> LocalGatewayRouteTableId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.LocalGatewayRouteTableVPCAssociationTagArgs>? _tags;

        /// <summary>
        /// The tags for the association.
        /// </summary>
        public InputList<Inputs.LocalGatewayRouteTableVPCAssociationTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LocalGatewayRouteTableVPCAssociationTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public LocalGatewayRouteTableVPCAssociationArgs()
        {
        }
        public static new LocalGatewayRouteTableVPCAssociationArgs Empty => new LocalGatewayRouteTableVPCAssociationArgs();
    }
}