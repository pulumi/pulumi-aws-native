// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    /// <summary>
    /// Resource schema for AWS::SageMaker::DeviceFleet
    /// </summary>
    [AwsNativeResourceType("aws-native:sagemaker:DeviceFleet")]
    public partial class DeviceFleet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description for the edge device fleet
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the edge device fleet
        /// </summary>
        [Output("deviceFleetName")]
        public Output<string> DeviceFleetName { get; private set; } = null!;

        /// <summary>
        /// S3 bucket and an ecryption key id (if available) to store outputs for the fleet
        /// </summary>
        [Output("outputConfig")]
        public Output<Outputs.DeviceFleetEdgeOutputConfig> OutputConfig { get; private set; } = null!;

        /// <summary>
        /// Role associated with the device fleet
        /// </summary>
        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        /// <summary>
        /// Associate tags with the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.DeviceFleetTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DeviceFleet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeviceFleet(string name, DeviceFleetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:DeviceFleet", name, args ?? new DeviceFleetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeviceFleet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:DeviceFleet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DeviceFleet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeviceFleet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeviceFleet(name, id, options);
        }
    }

    public sealed class DeviceFleetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for the edge device fleet
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the edge device fleet
        /// </summary>
        [Input("deviceFleetName")]
        public Input<string>? DeviceFleetName { get; set; }

        /// <summary>
        /// S3 bucket and an ecryption key id (if available) to store outputs for the fleet
        /// </summary>
        [Input("outputConfig", required: true)]
        public Input<Inputs.DeviceFleetEdgeOutputConfigArgs> OutputConfig { get; set; } = null!;

        /// <summary>
        /// Role associated with the device fleet
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.DeviceFleetTagArgs>? _tags;

        /// <summary>
        /// Associate tags with the resource
        /// </summary>
        public InputList<Inputs.DeviceFleetTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DeviceFleetTagArgs>());
            set => _tags = value;
        }

        public DeviceFleetArgs()
        {
        }
        public static new DeviceFleetArgs Empty => new DeviceFleetArgs();
    }
}