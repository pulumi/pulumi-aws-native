// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Rekognition
{
    /// <summary>
    /// The AWS::Rekognition::Collection type creates an Amazon Rekognition Collection. A collection is a logical grouping of information about detected faces which can later be referenced for searches on the group
    /// </summary>
    [AwsNativeResourceType("aws-native:rekognition:Collection")]
    public partial class Collection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Returns the Amazon Resource Name of the collection.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// ID for the collection that you are creating.
        /// </summary>
        [Output("collectionId")]
        public Output<string> CollectionId { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Collection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Collection(string name, CollectionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:rekognition:Collection", name, args ?? new CollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Collection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:rekognition:Collection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "collectionId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Collection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Collection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Collection(name, id, options);
        }
    }

    public sealed class CollectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID for the collection that you are creating.
        /// </summary>
        [Input("collectionId", required: true)]
        public Input<string> CollectionId { get; set; } = null!;

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

        public CollectionArgs()
        {
        }
        public static new CollectionArgs Empty => new CollectionArgs();
    }
}
