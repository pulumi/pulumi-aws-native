// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom
{
    /// <summary>
    /// A version for the specified customer-managed message template within the specified knowledge base.
    /// </summary>
    [AwsNativeResourceType("aws-native:wisdom:MessageTemplateVersion")]
    public partial class MessageTemplateVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The unqualified Amazon Resource Name (ARN) of the message template.
        /// </summary>
        [Output("messageTemplateArn")]
        public Output<string> MessageTemplateArn { get; private set; } = null!;

        /// <summary>
        /// The content SHA256 of the message template.
        /// </summary>
        [Output("messageTemplateContentSha256")]
        public Output<string?> MessageTemplateContentSha256 { get; private set; } = null!;

        /// <summary>
        /// The unqualified Amazon Resource Name (ARN) of the message template version.
        /// </summary>
        [Output("messageTemplateVersionArn")]
        public Output<string> MessageTemplateVersionArn { get; private set; } = null!;

        /// <summary>
        /// Current version number of the message template.
        /// </summary>
        [Output("messageTemplateVersionNumber")]
        public Output<double> MessageTemplateVersionNumber { get; private set; } = null!;


        /// <summary>
        /// Create a MessageTemplateVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MessageTemplateVersion(string name, MessageTemplateVersionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:wisdom:MessageTemplateVersion", name, args ?? new MessageTemplateVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MessageTemplateVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:wisdom:MessageTemplateVersion", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "messageTemplateArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MessageTemplateVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MessageTemplateVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MessageTemplateVersion(name, id, options);
        }
    }

    public sealed class MessageTemplateVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unqualified Amazon Resource Name (ARN) of the message template.
        /// </summary>
        [Input("messageTemplateArn", required: true)]
        public Input<string> MessageTemplateArn { get; set; } = null!;

        /// <summary>
        /// The content SHA256 of the message template.
        /// </summary>
        [Input("messageTemplateContentSha256")]
        public Input<string>? MessageTemplateContentSha256 { get; set; }

        public MessageTemplateVersionArgs()
        {
        }
        public static new MessageTemplateVersionArgs Empty => new MessageTemplateVersionArgs();
    }
}
