// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio
{
    /// <summary>
    /// Represents a launch profile which delegates access to a collection of studio components to studio users
    /// </summary>
    [AwsNativeResourceType("aws-native:nimblestudio:LaunchProfile")]
    public partial class LaunchProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// &lt;p&gt;The description.&lt;/p&gt;
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from.
        ///             These subnets must support the specified instance types. &lt;/p&gt;
        /// </summary>
        [Output("ec2SubnetIds")]
        public Output<ImmutableArray<string>> Ec2SubnetIds { get; private set; } = null!;

        [Output("launchProfileId")]
        public Output<string> LaunchProfileId { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The version number of the protocol that is used by the launch profile. The only valid
        ///             version is "2021-03-31".&lt;/p&gt;
        /// </summary>
        [Output("launchProfileProtocolVersions")]
        public Output<ImmutableArray<string>> LaunchProfileProtocolVersions { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The name for the launch profile.&lt;/p&gt;
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("streamConfiguration")]
        public Output<Outputs.LaunchProfileStreamConfiguration> StreamConfiguration { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;Unique identifiers for a collection of studio components that can be used with this
        ///             launch profile.&lt;/p&gt;
        /// </summary>
        [Output("studioComponentIds")]
        public Output<ImmutableArray<string>> StudioComponentIds { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The studio ID. &lt;/p&gt;
        /// </summary>
        [Output("studioId")]
        public Output<string> StudioId { get; private set; } = null!;

        [Output("tags")]
        public Output<Outputs.LaunchProfileTags?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a LaunchProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LaunchProfile(string name, LaunchProfileArgs args, CustomResourceOptions? options = null)
            : base("aws-native:nimblestudio:LaunchProfile", name, args ?? new LaunchProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LaunchProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:nimblestudio:LaunchProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LaunchProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LaunchProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LaunchProfile(name, id, options);
        }
    }

    public sealed class LaunchProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The description.&lt;/p&gt;
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ec2SubnetIds", required: true)]
        private InputList<string>? _ec2SubnetIds;

        /// <summary>
        /// &lt;p&gt;Specifies the IDs of the EC2 subnets where streaming sessions will be accessible from.
        ///             These subnets must support the specified instance types. &lt;/p&gt;
        /// </summary>
        public InputList<string> Ec2SubnetIds
        {
            get => _ec2SubnetIds ?? (_ec2SubnetIds = new InputList<string>());
            set => _ec2SubnetIds = value;
        }

        [Input("launchProfileProtocolVersions", required: true)]
        private InputList<string>? _launchProfileProtocolVersions;

        /// <summary>
        /// &lt;p&gt;The version number of the protocol that is used by the launch profile. The only valid
        ///             version is "2021-03-31".&lt;/p&gt;
        /// </summary>
        public InputList<string> LaunchProfileProtocolVersions
        {
            get => _launchProfileProtocolVersions ?? (_launchProfileProtocolVersions = new InputList<string>());
            set => _launchProfileProtocolVersions = value;
        }

        /// <summary>
        /// &lt;p&gt;The name for the launch profile.&lt;/p&gt;
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("streamConfiguration", required: true)]
        public Input<Inputs.LaunchProfileStreamConfigurationArgs> StreamConfiguration { get; set; } = null!;

        [Input("studioComponentIds", required: true)]
        private InputList<string>? _studioComponentIds;

        /// <summary>
        /// &lt;p&gt;Unique identifiers for a collection of studio components that can be used with this
        ///             launch profile.&lt;/p&gt;
        /// </summary>
        public InputList<string> StudioComponentIds
        {
            get => _studioComponentIds ?? (_studioComponentIds = new InputList<string>());
            set => _studioComponentIds = value;
        }

        /// <summary>
        /// &lt;p&gt;The studio ID. &lt;/p&gt;
        /// </summary>
        [Input("studioId", required: true)]
        public Input<string> StudioId { get; set; } = null!;

        [Input("tags")]
        public Input<Inputs.LaunchProfileTagsArgs>? Tags { get; set; }

        public LaunchProfileArgs()
        {
        }
        public static new LaunchProfileArgs Empty => new LaunchProfileArgs();
    }
}