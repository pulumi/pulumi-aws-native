// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk
{
    /// <summary>
    /// Resource Type definition for AWS::MSK::BatchScramSecret
    /// </summary>
    [AwsNativeResourceType("aws-native:msk:BatchScramSecret")]
    public partial class BatchScramSecret : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
        /// </summary>
        [Output("clusterArn")]
        public Output<string> ClusterArn { get; private set; } = null!;

        /// <summary>
        /// List of Amazon Resource Name (ARN)s of Secrets Manager secrets.
        /// </summary>
        [Output("secretArnList")]
        public Output<ImmutableArray<string>> SecretArnList { get; private set; } = null!;


        /// <summary>
        /// Create a BatchScramSecret resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BatchScramSecret(string name, BatchScramSecretArgs args, CustomResourceOptions? options = null)
            : base("aws-native:msk:BatchScramSecret", name, args ?? new BatchScramSecretArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BatchScramSecret(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:msk:BatchScramSecret", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "clusterArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BatchScramSecret resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BatchScramSecret Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BatchScramSecret(name, id, options);
        }
    }

    public sealed class BatchScramSecretArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
        /// </summary>
        [Input("clusterArn", required: true)]
        public Input<string> ClusterArn { get; set; } = null!;

        [Input("secretArnList")]
        private InputList<string>? _secretArnList;

        /// <summary>
        /// List of Amazon Resource Name (ARN)s of Secrets Manager secrets.
        /// </summary>
        public InputList<string> SecretArnList
        {
            get => _secretArnList ?? (_secretArnList = new InputList<string>());
            set => _secretArnList = value;
        }

        public BatchScramSecretArgs()
        {
        }
        public static new BatchScramSecretArgs Empty => new BatchScramSecretArgs();
    }
}
