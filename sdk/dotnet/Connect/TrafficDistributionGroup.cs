// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    /// <summary>
    /// Resource Type definition for AWS::Connect::TrafficDistributionGroup
    /// </summary>
    [AwsNativeResourceType("aws-native:connect:TrafficDistributionGroup")]
    public partial class TrafficDistributionGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A description for the traffic distribution group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The identifier of the Amazon Connect instance that has been replicated.
        /// </summary>
        [Output("instanceArn")]
        public Output<string> InstanceArn { get; private set; } = null!;

        /// <summary>
        /// If this is the default traffic distribution group.
        /// </summary>
        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The name for the traffic distribution group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The status of the traffic distribution group.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.Connect.TrafficDistributionGroupStatus> Status { get; private set; } = null!;

        /// <summary>
        /// One or more tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.TrafficDistributionGroupTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The identifier of the traffic distribution group.
        /// </summary>
        [Output("trafficDistributionGroupArn")]
        public Output<string> TrafficDistributionGroupArn { get; private set; } = null!;


        /// <summary>
        /// Create a TrafficDistributionGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TrafficDistributionGroup(string name, TrafficDistributionGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:connect:TrafficDistributionGroup", name, args ?? new TrafficDistributionGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TrafficDistributionGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:connect:TrafficDistributionGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "description",
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TrafficDistributionGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TrafficDistributionGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TrafficDistributionGroup(name, id, options);
        }
    }

    public sealed class TrafficDistributionGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the traffic distribution group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The identifier of the Amazon Connect instance that has been replicated.
        /// </summary>
        [Input("instanceArn", required: true)]
        public Input<string> InstanceArn { get; set; } = null!;

        /// <summary>
        /// The name for the traffic distribution group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.TrafficDistributionGroupTagArgs>? _tags;

        /// <summary>
        /// One or more tags.
        /// </summary>
        public InputList<Inputs.TrafficDistributionGroupTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TrafficDistributionGroupTagArgs>());
            set => _tags = value;
        }

        public TrafficDistributionGroupArgs()
        {
        }
        public static new TrafficDistributionGroupArgs Empty => new TrafficDistributionGroupArgs();
    }
}