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
    /// Resource Type definition for AWS::EC2::CapacityReservationFleet
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:CapacityReservationFleet")]
    public partial class CapacityReservationFleet : global::Pulumi.CustomResource
    {
        [Output("allocationStrategy")]
        public Output<string?> AllocationStrategy { get; private set; } = null!;

        [Output("capacityReservationFleetId")]
        public Output<string> CapacityReservationFleetId { get; private set; } = null!;

        [Output("endDate")]
        public Output<string?> EndDate { get; private set; } = null!;

        [Output("instanceMatchCriteria")]
        public Output<Pulumi.AwsNative.Ec2.CapacityReservationFleetInstanceMatchCriteria?> InstanceMatchCriteria { get; private set; } = null!;

        [Output("instanceTypeSpecifications")]
        public Output<ImmutableArray<Outputs.CapacityReservationFleetInstanceTypeSpecification>> InstanceTypeSpecifications { get; private set; } = null!;

        [Output("noRemoveEndDate")]
        public Output<bool?> NoRemoveEndDate { get; private set; } = null!;

        [Output("removeEndDate")]
        public Output<bool?> RemoveEndDate { get; private set; } = null!;

        [Output("tagSpecifications")]
        public Output<ImmutableArray<Outputs.CapacityReservationFleetTagSpecification>> TagSpecifications { get; private set; } = null!;

        [Output("tenancy")]
        public Output<Pulumi.AwsNative.Ec2.CapacityReservationFleetTenancy?> Tenancy { get; private set; } = null!;

        [Output("totalTargetCapacity")]
        public Output<int?> TotalTargetCapacity { get; private set; } = null!;


        /// <summary>
        /// Create a CapacityReservationFleet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CapacityReservationFleet(string name, CapacityReservationFleetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:CapacityReservationFleet", name, args ?? new CapacityReservationFleetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CapacityReservationFleet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:CapacityReservationFleet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "allocationStrategy",
                    "endDate",
                    "instanceMatchCriteria",
                    "instanceTypeSpecifications[*]",
                    "tagSpecifications[*]",
                    "tenancy",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CapacityReservationFleet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CapacityReservationFleet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CapacityReservationFleet(name, id, options);
        }
    }

    public sealed class CapacityReservationFleetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allocationStrategy")]
        public Input<string>? AllocationStrategy { get; set; }

        [Input("endDate")]
        public Input<string>? EndDate { get; set; }

        [Input("instanceMatchCriteria")]
        public Input<Pulumi.AwsNative.Ec2.CapacityReservationFleetInstanceMatchCriteria>? InstanceMatchCriteria { get; set; }

        [Input("instanceTypeSpecifications")]
        private InputList<Inputs.CapacityReservationFleetInstanceTypeSpecificationArgs>? _instanceTypeSpecifications;
        public InputList<Inputs.CapacityReservationFleetInstanceTypeSpecificationArgs> InstanceTypeSpecifications
        {
            get => _instanceTypeSpecifications ?? (_instanceTypeSpecifications = new InputList<Inputs.CapacityReservationFleetInstanceTypeSpecificationArgs>());
            set => _instanceTypeSpecifications = value;
        }

        [Input("noRemoveEndDate")]
        public Input<bool>? NoRemoveEndDate { get; set; }

        [Input("removeEndDate")]
        public Input<bool>? RemoveEndDate { get; set; }

        [Input("tagSpecifications")]
        private InputList<Inputs.CapacityReservationFleetTagSpecificationArgs>? _tagSpecifications;
        public InputList<Inputs.CapacityReservationFleetTagSpecificationArgs> TagSpecifications
        {
            get => _tagSpecifications ?? (_tagSpecifications = new InputList<Inputs.CapacityReservationFleetTagSpecificationArgs>());
            set => _tagSpecifications = value;
        }

        [Input("tenancy")]
        public Input<Pulumi.AwsNative.Ec2.CapacityReservationFleetTenancy>? Tenancy { get; set; }

        [Input("totalTargetCapacity")]
        public Input<int>? TotalTargetCapacity { get; set; }

        public CapacityReservationFleetArgs()
        {
        }
        public static new CapacityReservationFleetArgs Empty => new CapacityReservationFleetArgs();
    }
}