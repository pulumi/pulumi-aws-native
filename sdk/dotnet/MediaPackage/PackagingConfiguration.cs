// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage
{
    /// <summary>
    /// Resource schema for AWS::MediaPackage::PackagingConfiguration
    /// </summary>
    [AwsNativeResourceType("aws-native:mediapackage:PackagingConfiguration")]
    public partial class PackagingConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the PackagingConfiguration.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A CMAF packaging configuration.
        /// </summary>
        [Output("cmafPackage")]
        public Output<Outputs.PackagingConfigurationCmafPackage?> CmafPackage { get; private set; } = null!;

        /// <summary>
        /// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
        /// </summary>
        [Output("dashPackage")]
        public Output<Outputs.PackagingConfigurationDashPackage?> DashPackage { get; private set; } = null!;

        /// <summary>
        /// An HTTP Live Streaming (HLS) packaging configuration.
        /// </summary>
        [Output("hlsPackage")]
        public Output<Outputs.PackagingConfigurationHlsPackage?> HlsPackage { get; private set; } = null!;

        /// <summary>
        /// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
        /// </summary>
        [Output("mssPackage")]
        public Output<Outputs.PackagingConfigurationMssPackage?> MssPackage { get; private set; } = null!;

        /// <summary>
        /// The ID of a PackagingGroup.
        /// </summary>
        [Output("packagingGroupId")]
        public Output<string> PackagingGroupId { get; private set; } = null!;

        /// <summary>
        /// A collection of tags associated with a resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.PackagingConfigurationTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a PackagingConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PackagingConfiguration(string name, PackagingConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:mediapackage:PackagingConfiguration", name, args ?? new PackagingConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PackagingConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:mediapackage:PackagingConfiguration", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing PackagingConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PackagingConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PackagingConfiguration(name, id, options);
        }
    }

    public sealed class PackagingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A CMAF packaging configuration.
        /// </summary>
        [Input("cmafPackage")]
        public Input<Inputs.PackagingConfigurationCmafPackageArgs>? CmafPackage { get; set; }

        /// <summary>
        /// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
        /// </summary>
        [Input("dashPackage")]
        public Input<Inputs.PackagingConfigurationDashPackageArgs>? DashPackage { get; set; }

        /// <summary>
        /// An HTTP Live Streaming (HLS) packaging configuration.
        /// </summary>
        [Input("hlsPackage")]
        public Input<Inputs.PackagingConfigurationHlsPackageArgs>? HlsPackage { get; set; }

        /// <summary>
        /// A Microsoft Smooth Streaming (MSS) PackagingConfiguration.
        /// </summary>
        [Input("mssPackage")]
        public Input<Inputs.PackagingConfigurationMssPackageArgs>? MssPackage { get; set; }

        /// <summary>
        /// The ID of a PackagingGroup.
        /// </summary>
        [Input("packagingGroupId", required: true)]
        public Input<string> PackagingGroupId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.PackagingConfigurationTagArgs>? _tags;

        /// <summary>
        /// A collection of tags associated with a resource
        /// </summary>
        public InputList<Inputs.PackagingConfigurationTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.PackagingConfigurationTagArgs>());
            set => _tags = value;
        }

        public PackagingConfigurationArgs()
        {
        }
        public static new PackagingConfigurationArgs Empty => new PackagingConfigurationArgs();
    }
}