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
    /// Resource Type definition for AWS::EC2::VPNConnection
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:VPNConnection")]
    public partial class VPNConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the customer gateway at your end of the VPN connection.
        /// </summary>
        [Output("customerGatewayId")]
        public Output<string> CustomerGatewayId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the VPN connection uses static routes only.
        /// </summary>
        [Output("staticRoutesOnly")]
        public Output<bool?> StaticRoutesOnly { get; private set; } = null!;

        /// <summary>
        /// Any tags assigned to the VPN connection.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.VPNConnectionTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The ID of the transit gateway associated with the VPN connection.
        /// </summary>
        [Output("transitGatewayId")]
        public Output<string?> TransitGatewayId { get; private set; } = null!;

        /// <summary>
        /// The type of VPN connection.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The provider-assigned unique ID for this managed resource
        /// </summary>
        [Output("vpnConnectionId")]
        public Output<string> VpnConnectionId { get; private set; } = null!;

        /// <summary>
        /// The ID of the virtual private gateway at the AWS side of the VPN connection.
        /// </summary>
        [Output("vpnGatewayId")]
        public Output<string?> VpnGatewayId { get; private set; } = null!;

        /// <summary>
        /// The tunnel options for the VPN connection.
        /// </summary>
        [Output("vpnTunnelOptionsSpecifications")]
        public Output<ImmutableArray<Outputs.VPNConnectionVpnTunnelOptionsSpecification>> VpnTunnelOptionsSpecifications { get; private set; } = null!;


        /// <summary>
        /// Create a VPNConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VPNConnection(string name, VPNConnectionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VPNConnection", name, args ?? new VPNConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VPNConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:VPNConnection", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VPNConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VPNConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VPNConnection(name, id, options);
        }
    }

    public sealed class VPNConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the customer gateway at your end of the VPN connection.
        /// </summary>
        [Input("customerGatewayId", required: true)]
        public Input<string> CustomerGatewayId { get; set; } = null!;

        /// <summary>
        /// Indicates whether the VPN connection uses static routes only.
        /// </summary>
        [Input("staticRoutesOnly")]
        public Input<bool>? StaticRoutesOnly { get; set; }

        [Input("tags")]
        private InputList<Inputs.VPNConnectionTagArgs>? _tags;

        /// <summary>
        /// Any tags assigned to the VPN connection.
        /// </summary>
        public InputList<Inputs.VPNConnectionTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.VPNConnectionTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The ID of the transit gateway associated with the VPN connection.
        /// </summary>
        [Input("transitGatewayId")]
        public Input<string>? TransitGatewayId { get; set; }

        /// <summary>
        /// The type of VPN connection.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The ID of the virtual private gateway at the AWS side of the VPN connection.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        [Input("vpnTunnelOptionsSpecifications")]
        private InputList<Inputs.VPNConnectionVpnTunnelOptionsSpecificationArgs>? _vpnTunnelOptionsSpecifications;

        /// <summary>
        /// The tunnel options for the VPN connection.
        /// </summary>
        public InputList<Inputs.VPNConnectionVpnTunnelOptionsSpecificationArgs> VpnTunnelOptionsSpecifications
        {
            get => _vpnTunnelOptionsSpecifications ?? (_vpnTunnelOptionsSpecifications = new InputList<Inputs.VPNConnectionVpnTunnelOptionsSpecificationArgs>());
            set => _vpnTunnelOptionsSpecifications = value;
        }

        public VPNConnectionArgs()
        {
        }
        public static new VPNConnectionArgs Empty => new VPNConnectionArgs();
    }
}