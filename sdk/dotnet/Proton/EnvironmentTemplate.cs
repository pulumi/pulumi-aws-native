// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Proton
{
    /// <summary>
    /// Definition of AWS::Proton::EnvironmentTemplate Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:proton:EnvironmentTemplate")]
    public partial class EnvironmentTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) of the environment template.&lt;/p&gt;
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;A description of the environment template.&lt;/p&gt;
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The environment template name as displayed in the developer interface.&lt;/p&gt;
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;A customer provided encryption key that Proton uses to encrypt data.&lt;/p&gt;
        /// </summary>
        [Output("encryptionKey")]
        public Output<string?> EncryptionKey { get; private set; } = null!;

        /// <summary>
        /// The name of the environment template.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// When included, indicates that the environment template is for customer provisioned and managed infrastructure.
        /// </summary>
        [Output("provisioning")]
        public Output<Pulumi.AwsNative.Proton.EnvironmentTemplateProvisioning?> Provisioning { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;An optional list of metadata items that you can associate with the Proton environment template. A tag is a key-value pair.&lt;/p&gt;
        ///          &lt;p&gt;For more information, see &lt;a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html"&gt;Proton resources and tagging&lt;/a&gt; in the
        ///         &lt;i&gt;Proton User Guide&lt;/i&gt;.&lt;/p&gt;
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EnvironmentTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvironmentTemplate(string name, EnvironmentTemplateArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:proton:EnvironmentTemplate", name, args ?? new EnvironmentTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnvironmentTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:proton:EnvironmentTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "encryptionKey",
                    "name",
                    "provisioning",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EnvironmentTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvironmentTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EnvironmentTemplate(name, id, options);
        }
    }

    public sealed class EnvironmentTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;A description of the environment template.&lt;/p&gt;
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// &lt;p&gt;The environment template name as displayed in the developer interface.&lt;/p&gt;
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// &lt;p&gt;A customer provided encryption key that Proton uses to encrypt data.&lt;/p&gt;
        /// </summary>
        [Input("encryptionKey")]
        public Input<string>? EncryptionKey { get; set; }

        /// <summary>
        /// The name of the environment template.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// When included, indicates that the environment template is for customer provisioned and managed infrastructure.
        /// </summary>
        [Input("provisioning")]
        public Input<Pulumi.AwsNative.Proton.EnvironmentTemplateProvisioning>? Provisioning { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// &lt;p&gt;An optional list of metadata items that you can associate with the Proton environment template. A tag is a key-value pair.&lt;/p&gt;
        ///          &lt;p&gt;For more information, see &lt;a href="https://docs.aws.amazon.com/proton/latest/userguide/resources.html"&gt;Proton resources and tagging&lt;/a&gt; in the
        ///         &lt;i&gt;Proton User Guide&lt;/i&gt;.&lt;/p&gt;
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public EnvironmentTemplateArgs()
        {
        }
        public static new EnvironmentTemplateArgs Empty => new EnvironmentTemplateArgs();
    }
}
