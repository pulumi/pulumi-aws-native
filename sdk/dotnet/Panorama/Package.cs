// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Panorama
{
    /// <summary>
    /// Schema for Package CloudFormation Resource
    /// </summary>
    [AwsNativeResourceType("aws-native:panorama:Package")]
    public partial class Package : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("createdTime")]
        public Output<int> CreatedTime { get; private set; } = null!;

        [Output("packageId")]
        public Output<string> PackageId { get; private set; } = null!;

        [Output("packageName")]
        public Output<string> PackageName { get; private set; } = null!;

        [Output("storageLocation")]
        public Output<Outputs.PackageStorageLocation?> StorageLocation { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.PackageTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Package resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Package(string name, PackageArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:panorama:Package", name, args ?? new PackageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Package(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:panorama:Package", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Package resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Package Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Package(name, id, options);
        }
    }

    public sealed class PackageArgs : global::Pulumi.ResourceArgs
    {
        [Input("packageName")]
        public Input<string>? PackageName { get; set; }

        [Input("storageLocation")]
        public Input<Inputs.PackageStorageLocationArgs>? StorageLocation { get; set; }

        [Input("tags")]
        private InputList<Inputs.PackageTagArgs>? _tags;
        public InputList<Inputs.PackageTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.PackageTagArgs>());
            set => _tags = value;
        }

        public PackageArgs()
        {
        }
        public static new PackageArgs Empty => new PackageArgs();
    }
}