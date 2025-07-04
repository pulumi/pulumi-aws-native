// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall
{
    /// <summary>
    /// Resource type definition for AWS::NetworkFirewall::Firewall
    /// </summary>
    [AwsNativeResourceType("aws-native:networkfirewall:Firewall")]
    public partial class Firewall : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A flag indicating whether it is possible to delete the firewall. A setting of `TRUE` indicates that the firewall is protected against deletion. Use this setting to protect against accidentally deleting a firewall that is in use. When you create a firewall, the operation initializes this flag to `TRUE` .
        /// </summary>
        [Output("deleteProtection")]
        public Output<bool?> DeleteProtection { get; private set; } = null!;

        /// <summary>
        /// A description of the firewall.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The types of analysis to enable for the firewall. Can be TLS_SNI, HTTP_HOST, or both.
        /// </summary>
        [Output("enabledAnalysisTypes")]
        public Output<ImmutableArray<Pulumi.AwsNative.NetworkFirewall.FirewallEnabledAnalysisType>> EnabledAnalysisTypes { get; private set; } = null!;

        /// <summary>
        /// The unique IDs of the firewall endpoints for all of the subnets that you attached to the firewall. The subnets are not listed in any particular order. For example: `["us-west-2c:vpce-111122223333", "us-west-2a:vpce-987654321098", "us-west-2b:vpce-012345678901"]` .
        /// </summary>
        [Output("endpointIds")]
        public Output<ImmutableArray<string>> EndpointIds { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the `Firewall` .
        /// </summary>
        [Output("firewallArn")]
        public Output<string> FirewallArn { get; private set; } = null!;

        /// <summary>
        /// The name of the `Firewall` resource.
        /// </summary>
        [Output("firewallId")]
        public Output<string> FirewallId { get; private set; } = null!;

        /// <summary>
        /// The descriptive name of the firewall. You can't change the name of a firewall after you create it.
        /// </summary>
        [Output("firewallName")]
        public Output<string> FirewallName { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the firewall policy.
        /// 
        /// The relationship of firewall to firewall policy is many to one. Each firewall requires one firewall policy association, and you can use the same firewall policy for multiple firewalls.
        /// </summary>
        [Output("firewallPolicyArn")]
        public Output<string> FirewallPolicyArn { get; private set; } = null!;

        /// <summary>
        /// A setting indicating whether the firewall is protected against a change to the firewall policy association. Use this setting to protect against accidentally modifying the firewall policy for a firewall that is in use. When you create a firewall, the operation initializes this setting to `TRUE` .
        /// </summary>
        [Output("firewallPolicyChangeProtection")]
        public Output<bool?> FirewallPolicyChangeProtection { get; private set; } = null!;

        /// <summary>
        /// A setting indicating whether the firewall is protected against changes to the subnet associations. Use this setting to protect against accidentally modifying the subnet associations for a firewall that is in use. When you create a firewall, the operation initializes this setting to `TRUE` .
        /// </summary>
        [Output("subnetChangeProtection")]
        public Output<bool?> SubnetChangeProtection { get; private set; } = null!;

        /// <summary>
        /// The primary public subnets that Network Firewall is using for the firewall. Network Firewall creates a firewall endpoint in each subnet. Create a subnet mapping for each Availability Zone where you want to use the firewall.
        /// 
        /// These subnets are all defined for a single, primary VPC, and each must belong to a different Availability Zone. Each of these subnets establishes the availability of the firewall in its Availability Zone.
        /// 
        /// In addition to these subnets, you can define other endpoints for the firewall in `VpcEndpointAssociation` resources. You can define these additional endpoints for any VPC, and for any of the Availability Zones where the firewall resource already has a subnet mapping. VPC endpoint associations give you the ability to protect multiple VPCs using a single firewall, and to define multiple firewall endpoints for a VPC in a single Availability Zone.
        /// </summary>
        [Output("subnetMappings")]
        public Output<ImmutableArray<Outputs.FirewallSubnetMapping>> SubnetMappings { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// 
        /// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the VPC where the firewall is in use. You can't change the VPC of a firewall after you create the firewall.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Firewall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Firewall(string name, FirewallArgs args, CustomResourceOptions? options = null)
            : base("aws-native:networkfirewall:Firewall", name, args ?? new FirewallArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Firewall(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:networkfirewall:Firewall", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "firewallName",
                    "vpcId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Firewall resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Firewall Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Firewall(name, id, options);
        }
    }

    public sealed class FirewallArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A flag indicating whether it is possible to delete the firewall. A setting of `TRUE` indicates that the firewall is protected against deletion. Use this setting to protect against accidentally deleting a firewall that is in use. When you create a firewall, the operation initializes this flag to `TRUE` .
        /// </summary>
        [Input("deleteProtection")]
        public Input<bool>? DeleteProtection { get; set; }

        /// <summary>
        /// A description of the firewall.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabledAnalysisTypes")]
        private InputList<Pulumi.AwsNative.NetworkFirewall.FirewallEnabledAnalysisType>? _enabledAnalysisTypes;

        /// <summary>
        /// The types of analysis to enable for the firewall. Can be TLS_SNI, HTTP_HOST, or both.
        /// </summary>
        public InputList<Pulumi.AwsNative.NetworkFirewall.FirewallEnabledAnalysisType> EnabledAnalysisTypes
        {
            get => _enabledAnalysisTypes ?? (_enabledAnalysisTypes = new InputList<Pulumi.AwsNative.NetworkFirewall.FirewallEnabledAnalysisType>());
            set => _enabledAnalysisTypes = value;
        }

        /// <summary>
        /// The descriptive name of the firewall. You can't change the name of a firewall after you create it.
        /// </summary>
        [Input("firewallName")]
        public Input<string>? FirewallName { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the firewall policy.
        /// 
        /// The relationship of firewall to firewall policy is many to one. Each firewall requires one firewall policy association, and you can use the same firewall policy for multiple firewalls.
        /// </summary>
        [Input("firewallPolicyArn", required: true)]
        public Input<string> FirewallPolicyArn { get; set; } = null!;

        /// <summary>
        /// A setting indicating whether the firewall is protected against a change to the firewall policy association. Use this setting to protect against accidentally modifying the firewall policy for a firewall that is in use. When you create a firewall, the operation initializes this setting to `TRUE` .
        /// </summary>
        [Input("firewallPolicyChangeProtection")]
        public Input<bool>? FirewallPolicyChangeProtection { get; set; }

        /// <summary>
        /// A setting indicating whether the firewall is protected against changes to the subnet associations. Use this setting to protect against accidentally modifying the subnet associations for a firewall that is in use. When you create a firewall, the operation initializes this setting to `TRUE` .
        /// </summary>
        [Input("subnetChangeProtection")]
        public Input<bool>? SubnetChangeProtection { get; set; }

        [Input("subnetMappings", required: true)]
        private InputList<Inputs.FirewallSubnetMappingArgs>? _subnetMappings;

        /// <summary>
        /// The primary public subnets that Network Firewall is using for the firewall. Network Firewall creates a firewall endpoint in each subnet. Create a subnet mapping for each Availability Zone where you want to use the firewall.
        /// 
        /// These subnets are all defined for a single, primary VPC, and each must belong to a different Availability Zone. Each of these subnets establishes the availability of the firewall in its Availability Zone.
        /// 
        /// In addition to these subnets, you can define other endpoints for the firewall in `VpcEndpointAssociation` resources. You can define these additional endpoints for any VPC, and for any of the Availability Zones where the firewall resource already has a subnet mapping. VPC endpoint associations give you the ability to protect multiple VPCs using a single firewall, and to define multiple firewall endpoints for a VPC in a single Availability Zone.
        /// </summary>
        public InputList<Inputs.FirewallSubnetMappingArgs> SubnetMappings
        {
            get => _subnetMappings ?? (_subnetMappings = new InputList<Inputs.FirewallSubnetMappingArgs>());
            set => _subnetMappings = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// 
        /// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The unique identifier of the VPC where the firewall is in use. You can't change the VPC of a firewall after you create the firewall.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public FirewallArgs()
        {
        }
        public static new FirewallArgs Empty => new FirewallArgs();
    }
}
