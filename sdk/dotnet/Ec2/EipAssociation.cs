// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    /// <summary>
    /// Associates an Elastic IP address with an instance or a network interface. Before you can use an Elastic IP address, you must allocate it to your account. For more information about working with Elastic IP addresses, see [Elastic IP address concepts and rules](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#vpc-eip-overview).
    ///  You must specify ``AllocationId`` and either ``InstanceId``, ``NetworkInterfaceId``, or ``PrivateIpAddress``.
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:EipAssociation")]
    public partial class EipAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The allocation ID. This is required.
        /// </summary>
        [Output("allocationId")]
        public Output<string?> AllocationId { get; private set; } = null!;

        /// <summary>
        /// The ID of the association.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("eip")]
        public Output<string?> Eip { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance. The instance must have exactly one attached network interface. You can specify either the instance ID or the network interface ID, but not both.
        /// </summary>
        [Output("instanceId")]
        public Output<string?> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the network interface. If the instance has more than one network interface, you must specify a network interface ID.
        ///  You can specify either the instance ID or the network interface ID, but not both.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string?> NetworkInterfaceId { get; private set; } = null!;

        /// <summary>
        /// The primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string?> PrivateIpAddress { get; private set; } = null!;


        /// <summary>
        /// Create a EipAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EipAssociation(string name, EipAssociationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:EipAssociation", name, args ?? new EipAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EipAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:EipAssociation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "allocationId",
                    "eip",
                    "instanceId",
                    "networkInterfaceId",
                    "privateIpAddress",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EipAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EipAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EipAssociation(name, id, options);
        }
    }

    public sealed class EipAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allocation ID. This is required.
        /// </summary>
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        [Input("eip")]
        public Input<string>? Eip { get; set; }

        /// <summary>
        /// The ID of the instance. The instance must have exactly one attached network interface. You can specify either the instance ID or the network interface ID, but not both.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The ID of the network interface. If the instance has more than one network interface, you must specify a network interface ID.
        ///  You can specify either the instance ID or the network interface ID, but not both.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        public EipAssociationArgs()
        {
        }
        public static new EipAssociationArgs Empty => new EipAssociationArgs();
    }
}
