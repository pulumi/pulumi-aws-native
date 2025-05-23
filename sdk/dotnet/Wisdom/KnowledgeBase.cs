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
    /// Definition of AWS::Wisdom::KnowledgeBase Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:wisdom:KnowledgeBase")]
    public partial class KnowledgeBase : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the knowledge base.
        /// </summary>
        [Output("knowledgeBaseArn")]
        public Output<string> KnowledgeBaseArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the knowledge base.
        /// </summary>
        [Output("knowledgeBaseId")]
        public Output<string> KnowledgeBaseId { get; private set; } = null!;

        /// <summary>
        /// The type of knowledge base. Only CUSTOM knowledge bases allow you to upload your own content. EXTERNAL knowledge bases support integrations with third-party systems whose content is synchronized automatically.
        /// </summary>
        [Output("knowledgeBaseType")]
        public Output<Pulumi.AwsNative.Wisdom.KnowledgeBaseType> KnowledgeBaseType { get; private set; } = null!;

        /// <summary>
        /// The name of the knowledge base.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Information about how to render the content.
        /// </summary>
        [Output("renderingConfiguration")]
        public Output<Outputs.KnowledgeBaseRenderingConfiguration?> RenderingConfiguration { get; private set; } = null!;

        /// <summary>
        /// This customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom. For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) . For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) in the *AWS Key Management Service Developer Guide* .
        /// </summary>
        [Output("serverSideEncryptionConfiguration")]
        public Output<Outputs.KnowledgeBaseServerSideEncryptionConfiguration?> ServerSideEncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// The source of the knowledge base content. Only set this argument for EXTERNAL or Managed knowledge bases.
        /// </summary>
        [Output("sourceConfiguration")]
        public Output<Union<Outputs.KnowledgeBaseSourceConfiguration0Properties, Outputs.KnowledgeBaseSourceConfiguration1Properties>?> SourceConfiguration { get; private set; } = null!;

        /// <summary>
        /// The tags used to organize, track, or control access for this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.CreateOnlyTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Contains details about how to ingest the documents in a data source.
        /// </summary>
        [Output("vectorIngestionConfiguration")]
        public Output<Outputs.KnowledgeBaseVectorIngestionConfiguration?> VectorIngestionConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a KnowledgeBase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KnowledgeBase(string name, KnowledgeBaseArgs args, CustomResourceOptions? options = null)
            : base("aws-native:wisdom:KnowledgeBase", name, args ?? new KnowledgeBaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KnowledgeBase(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:wisdom:KnowledgeBase", name, null, MakeResourceOptions(options, id))
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
                    "knowledgeBaseType",
                    "name",
                    "serverSideEncryptionConfiguration",
                    "sourceConfiguration",
                    "tags[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KnowledgeBase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KnowledgeBase Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KnowledgeBase(name, id, options);
        }
    }

    public sealed class KnowledgeBaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The type of knowledge base. Only CUSTOM knowledge bases allow you to upload your own content. EXTERNAL knowledge bases support integrations with third-party systems whose content is synchronized automatically.
        /// </summary>
        [Input("knowledgeBaseType", required: true)]
        public Input<Pulumi.AwsNative.Wisdom.KnowledgeBaseType> KnowledgeBaseType { get; set; } = null!;

        /// <summary>
        /// The name of the knowledge base.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Information about how to render the content.
        /// </summary>
        [Input("renderingConfiguration")]
        public Input<Inputs.KnowledgeBaseRenderingConfigurationArgs>? RenderingConfiguration { get; set; }

        /// <summary>
        /// This customer managed key must have a policy that allows `kms:CreateGrant` and `kms:DescribeKey` permissions to the IAM identity using the key to invoke Wisdom. For more information about setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html) . For information about valid ID values, see [Key identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) in the *AWS Key Management Service Developer Guide* .
        /// </summary>
        [Input("serverSideEncryptionConfiguration")]
        public Input<Inputs.KnowledgeBaseServerSideEncryptionConfigurationArgs>? ServerSideEncryptionConfiguration { get; set; }

        /// <summary>
        /// The source of the knowledge base content. Only set this argument for EXTERNAL or Managed knowledge bases.
        /// </summary>
        [Input("sourceConfiguration")]
        public InputUnion<Inputs.KnowledgeBaseSourceConfiguration0PropertiesArgs, Inputs.KnowledgeBaseSourceConfiguration1PropertiesArgs>? SourceConfiguration { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs>? _tags;

        /// <summary>
        /// The tags used to organize, track, or control access for this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Contains details about how to ingest the documents in a data source.
        /// </summary>
        [Input("vectorIngestionConfiguration")]
        public Input<Inputs.KnowledgeBaseVectorIngestionConfigurationArgs>? VectorIngestionConfiguration { get; set; }

        public KnowledgeBaseArgs()
        {
        }
        public static new KnowledgeBaseArgs Empty => new KnowledgeBaseArgs();
    }
}
