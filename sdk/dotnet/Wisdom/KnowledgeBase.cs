// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
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
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("knowledgeBaseArn")]
        public Output<string> KnowledgeBaseArn { get; private set; } = null!;

        [Output("knowledgeBaseId")]
        public Output<string> KnowledgeBaseId { get; private set; } = null!;

        [Output("knowledgeBaseType")]
        public Output<Pulumi.AwsNative.Wisdom.KnowledgeBaseType> KnowledgeBaseType { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("renderingConfiguration")]
        public Output<Outputs.KnowledgeBaseRenderingConfiguration?> RenderingConfiguration { get; private set; } = null!;

        [Output("serverSideEncryptionConfiguration")]
        public Output<Outputs.KnowledgeBaseServerSideEncryptionConfiguration?> ServerSideEncryptionConfiguration { get; private set; } = null!;

        [Output("sourceConfiguration")]
        public Output<Outputs.KnowledgeBaseSourceConfiguration?> SourceConfiguration { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.KnowledgeBaseTag>> Tags { get; private set; } = null!;


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
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("knowledgeBaseType", required: true)]
        public Input<Pulumi.AwsNative.Wisdom.KnowledgeBaseType> KnowledgeBaseType { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("renderingConfiguration")]
        public Input<Inputs.KnowledgeBaseRenderingConfigurationArgs>? RenderingConfiguration { get; set; }

        [Input("serverSideEncryptionConfiguration")]
        public Input<Inputs.KnowledgeBaseServerSideEncryptionConfigurationArgs>? ServerSideEncryptionConfiguration { get; set; }

        [Input("sourceConfiguration")]
        public Input<Inputs.KnowledgeBaseSourceConfigurationArgs>? SourceConfiguration { get; set; }

        [Input("tags")]
        private InputList<Inputs.KnowledgeBaseTagArgs>? _tags;
        public InputList<Inputs.KnowledgeBaseTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.KnowledgeBaseTagArgs>());
            set => _tags = value;
        }

        public KnowledgeBaseArgs()
        {
        }
        public static new KnowledgeBaseArgs Empty => new KnowledgeBaseArgs();
    }
}