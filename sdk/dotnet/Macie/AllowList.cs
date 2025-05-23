// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Macie
{
    /// <summary>
    /// Macie AllowList resource schema
    /// </summary>
    [AwsNativeResourceType("aws-native:macie:AllowList")]
    public partial class AllowList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// AllowList ARN.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// AllowList ID.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// AllowList criteria.
        /// </summary>
        [Output("criteria")]
        public Output<Outputs.AllowListCriteria> Criteria { get; private set; } = null!;

        /// <summary>
        /// Description of AllowList.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of AllowList.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// AllowList status.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.Macie.AllowListStatus> Status { get; private set; } = null!;

        /// <summary>
        /// A collection of tags associated with a resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AllowList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AllowList(string name, AllowListArgs args, CustomResourceOptions? options = null)
            : base("aws-native:macie:AllowList", name, args ?? new AllowListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AllowList(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:macie:AllowList", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AllowList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AllowList Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AllowList(name, id, options);
        }
    }

    public sealed class AllowListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// AllowList criteria.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<Inputs.AllowListCriteriaArgs> Criteria { get; set; } = null!;

        /// <summary>
        /// Description of AllowList.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of AllowList.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// A collection of tags associated with a resource
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public AllowListArgs()
        {
        }
        public static new AllowListArgs Empty => new AllowListArgs();
    }
}
