// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Deadline
{
    /// <summary>
    /// Definition of AWS::Deadline::Fleet Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:deadline:Fleet")]
    public partial class Fleet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned to the fleet.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("capabilities")]
        public Output<Outputs.FleetCapabilities> Capabilities { get; private set; } = null!;

        /// <summary>
        /// The configuration details for the fleet.
        /// </summary>
        [Output("configuration")]
        public Output<Union<Outputs.FleetConfiguration0Properties, Outputs.FleetConfiguration1Properties>> Configuration { get; private set; } = null!;

        /// <summary>
        /// A description that helps identify what the fleet is used for.
        /// 
        /// &gt; This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the fleet summary to update.
        /// 
        /// &gt; This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The farm ID.
        /// </summary>
        [Output("farmId")]
        public Output<string> FarmId { get; private set; } = null!;

        /// <summary>
        /// The fleet ID.
        /// </summary>
        [Output("fleetId")]
        public Output<string> FleetId { get; private set; } = null!;

        /// <summary>
        /// Provides a script that runs as a worker is starting up that you can use to provide additional configuration for workers in your fleet.
        /// 
        /// To remove a script from a fleet, use the [UpdateFleet](https://docs.aws.amazon.com/deadline-cloud/latest/APIReference/API_UpdateFleet.html) operation with the `hostConfiguration` `scriptBody` parameter set to an empty string ("").
        /// </summary>
        [Output("hostConfiguration")]
        public Output<Outputs.FleetHostConfiguration?> HostConfiguration { get; private set; } = null!;

        /// <summary>
        /// The maximum number of workers specified in the fleet.
        /// </summary>
        [Output("maxWorkerCount")]
        public Output<int> MaxWorkerCount { get; private set; } = null!;

        /// <summary>
        /// The minimum number of workers in the fleet.
        /// </summary>
        [Output("minWorkerCount")]
        public Output<int?> MinWorkerCount { get; private set; } = null!;

        /// <summary>
        /// The IAM role that workers in the fleet use when processing jobs.
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// The status of the fleet.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.Deadline.FleetStatus> Status { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The number of workers in the fleet summary.
        /// </summary>
        [Output("workerCount")]
        public Output<int> WorkerCount { get; private set; } = null!;


        /// <summary>
        /// Create a Fleet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Fleet(string name, FleetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:deadline:Fleet", name, args ?? new FleetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Fleet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:deadline:Fleet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "farmId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Fleet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Fleet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Fleet(name, id, options);
        }
    }

    public sealed class FleetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration details for the fleet.
        /// </summary>
        [Input("configuration", required: true)]
        public InputUnion<Inputs.FleetConfiguration0PropertiesArgs, Inputs.FleetConfiguration1PropertiesArgs> Configuration { get; set; } = null!;

        /// <summary>
        /// A description that helps identify what the fleet is used for.
        /// 
        /// &gt; This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the fleet summary to update.
        /// 
        /// &gt; This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The farm ID.
        /// </summary>
        [Input("farmId", required: true)]
        public Input<string> FarmId { get; set; } = null!;

        /// <summary>
        /// Provides a script that runs as a worker is starting up that you can use to provide additional configuration for workers in your fleet.
        /// 
        /// To remove a script from a fleet, use the [UpdateFleet](https://docs.aws.amazon.com/deadline-cloud/latest/APIReference/API_UpdateFleet.html) operation with the `hostConfiguration` `scriptBody` parameter set to an empty string ("").
        /// </summary>
        [Input("hostConfiguration")]
        public Input<Inputs.FleetHostConfigurationArgs>? HostConfiguration { get; set; }

        /// <summary>
        /// The maximum number of workers specified in the fleet.
        /// </summary>
        [Input("maxWorkerCount", required: true)]
        public Input<int> MaxWorkerCount { get; set; } = null!;

        /// <summary>
        /// The minimum number of workers in the fleet.
        /// </summary>
        [Input("minWorkerCount")]
        public Input<int>? MinWorkerCount { get; set; }

        /// <summary>
        /// The IAM role that workers in the fleet use when processing jobs.
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public FleetArgs()
        {
        }
        public static new FleetArgs Empty => new FleetArgs();
    }
}
