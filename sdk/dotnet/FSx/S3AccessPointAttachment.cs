// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx
{
    /// <summary>
    /// Resource type definition for AWS::FSx::S3AccessPointAttachment
    /// </summary>
    [AwsNativeResourceType("aws-native:fsx:S3AccessPointAttachment")]
    public partial class S3AccessPointAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Name of the S3AccessPointAttachment
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The OpenZFSConfiguration of the S3 access point attachment.
        /// </summary>
        [Output("openZfsConfiguration")]
        public Output<Outputs.S3AccessPointAttachmentS3AccessPointOpenZfsConfiguration> OpenZfsConfiguration { get; private set; } = null!;

        /// <summary>
        /// The S3 access point configuration of the S3 access point attachment.
        /// </summary>
        [Output("s3AccessPoint")]
        public Output<Outputs.S3AccessPointAttachmentS3AccessPoint?> S3AccessPoint { get; private set; } = null!;

        /// <summary>
        /// The type of Amazon FSx volume that the S3 access point is attached to.
        /// </summary>
        [Output("type")]
        public Output<Pulumi.AwsNative.FSx.S3AccessPointAttachmentType> Type { get; private set; } = null!;


        /// <summary>
        /// Create a S3AccessPointAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public S3AccessPointAttachment(string name, S3AccessPointAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws-native:fsx:S3AccessPointAttachment", name, args ?? new S3AccessPointAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private S3AccessPointAttachment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:fsx:S3AccessPointAttachment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                    "openZfsConfiguration",
                    "s3AccessPoint",
                    "type",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing S3AccessPointAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static S3AccessPointAttachment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new S3AccessPointAttachment(name, id, options);
        }
    }

    public sealed class S3AccessPointAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Name of the S3AccessPointAttachment
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OpenZFSConfiguration of the S3 access point attachment.
        /// </summary>
        [Input("openZfsConfiguration", required: true)]
        public Input<Inputs.S3AccessPointAttachmentS3AccessPointOpenZfsConfigurationArgs> OpenZfsConfiguration { get; set; } = null!;

        /// <summary>
        /// The S3 access point configuration of the S3 access point attachment.
        /// </summary>
        [Input("s3AccessPoint")]
        public Input<Inputs.S3AccessPointAttachmentS3AccessPointArgs>? S3AccessPoint { get; set; }

        /// <summary>
        /// The type of Amazon FSx volume that the S3 access point is attached to.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.FSx.S3AccessPointAttachmentType> Type { get; set; } = null!;

        public S3AccessPointAttachmentArgs()
        {
        }
        public static new S3AccessPointAttachmentArgs Empty => new S3AccessPointAttachmentArgs();
    }
}
