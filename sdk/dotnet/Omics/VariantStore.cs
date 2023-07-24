// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Omics
{
    /// <summary>
    /// Definition of AWS::Omics::VariantStore Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:omics:VariantStore")]
    public partial class VariantStore : global::Pulumi.CustomResource
    {
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("reference")]
        public Output<Outputs.VariantStoreReferenceItem> Reference { get; private set; } = null!;

        [Output("sseConfig")]
        public Output<Outputs.VariantStoreSseConfig?> SseConfig { get; private set; } = null!;

        [Output("status")]
        public Output<Pulumi.AwsNative.Omics.VariantStoreStoreStatus> Status { get; private set; } = null!;

        [Output("statusMessage")]
        public Output<string> StatusMessage { get; private set; } = null!;

        [Output("storeArn")]
        public Output<string> StoreArn { get; private set; } = null!;

        [Output("storeSizeBytes")]
        public Output<double> StoreSizeBytes { get; private set; } = null!;

        [Output("tags")]
        public Output<Outputs.VariantStoreTagMap?> Tags { get; private set; } = null!;

        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a VariantStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VariantStore(string name, VariantStoreArgs args, CustomResourceOptions? options = null)
            : base("aws-native:omics:VariantStore", name, args ?? new VariantStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VariantStore(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:omics:VariantStore", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VariantStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VariantStore Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VariantStore(name, id, options);
        }
    }

    public sealed class VariantStoreArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("reference", required: true)]
        public Input<Inputs.VariantStoreReferenceItemArgs> Reference { get; set; } = null!;

        [Input("sseConfig")]
        public Input<Inputs.VariantStoreSseConfigArgs>? SseConfig { get; set; }

        [Input("tags")]
        public Input<Inputs.VariantStoreTagMapArgs>? Tags { get; set; }

        public VariantStoreArgs()
        {
        }
        public static new VariantStoreArgs Empty => new VariantStoreArgs();
    }
}