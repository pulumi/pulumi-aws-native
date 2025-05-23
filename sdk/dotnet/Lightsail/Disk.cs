// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail
{
    /// <summary>
    /// Resource Type definition for AWS::Lightsail::Disk
    /// </summary>
    [AwsNativeResourceType("aws-native:lightsail:Disk")]
    public partial class Disk : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An array of objects representing the add-ons to enable for the new instance.
        /// </summary>
        [Output("addOns")]
        public Output<ImmutableArray<Outputs.DiskAddOn>> AddOns { get; private set; } = null!;

        /// <summary>
        /// Name of the attached Lightsail Instance
        /// </summary>
        [Output("attachedTo")]
        public Output<string> AttachedTo { get; private set; } = null!;

        /// <summary>
        /// Attachment State of the Lightsail disk
        /// </summary>
        [Output("attachmentState")]
        public Output<string> AttachmentState { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string?> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the disk.
        /// </summary>
        [Output("diskArn")]
        public Output<string> DiskArn { get; private set; } = null!;

        /// <summary>
        /// The names to use for your new Lightsail disk.
        /// </summary>
        [Output("diskName")]
        public Output<string> DiskName { get; private set; } = null!;

        /// <summary>
        /// Iops of the Lightsail disk
        /// </summary>
        [Output("iops")]
        public Output<int> Iops { get; private set; } = null!;

        /// <summary>
        /// Check is Disk is attached state
        /// </summary>
        [Output("isAttached")]
        public Output<bool> IsAttached { get; private set; } = null!;

        /// <summary>
        /// The AWS Region and Availability Zone where the disk is located.
        /// </summary>
        [Output("location")]
        public Output<Outputs.DiskLocation?> Location { get; private set; } = null!;

        /// <summary>
        /// Path of the  attached Disk
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Resource type of Lightsail instance.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        /// <summary>
        /// Size of the Lightsail disk
        /// </summary>
        [Output("sizeInGb")]
        public Output<int> SizeInGb { get; private set; } = null!;

        /// <summary>
        /// State of the Lightsail disk
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Support code to help identify any issues
        /// </summary>
        [Output("supportCode")]
        public Output<string> SupportCode { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Disk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Disk(string name, DiskArgs args, CustomResourceOptions? options = null)
            : base("aws-native:lightsail:Disk", name, args ?? new DiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Disk(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:lightsail:Disk", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "availabilityZone",
                    "diskName",
                    "sizeInGb",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Disk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Disk Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Disk(name, id, options);
        }
    }

    public sealed class DiskArgs : global::Pulumi.ResourceArgs
    {
        [Input("addOns")]
        private InputList<Inputs.DiskAddOnArgs>? _addOns;

        /// <summary>
        /// An array of objects representing the add-ons to enable for the new instance.
        /// </summary>
        public InputList<Inputs.DiskAddOnArgs> AddOns
        {
            get => _addOns ?? (_addOns = new InputList<Inputs.DiskAddOnArgs>());
            set => _addOns = value;
        }

        /// <summary>
        /// The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The names to use for your new Lightsail disk.
        /// </summary>
        [Input("diskName")]
        public Input<string>? DiskName { get; set; }

        /// <summary>
        /// The AWS Region and Availability Zone where the disk is located.
        /// </summary>
        [Input("location")]
        public Input<Inputs.DiskLocationArgs>? Location { get; set; }

        /// <summary>
        /// Size of the Lightsail disk
        /// </summary>
        [Input("sizeInGb", required: true)]
        public Input<int> SizeInGb { get; set; } = null!;

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

        public DiskArgs()
        {
        }
        public static new DiskArgs Empty => new DiskArgs();
    }
}
