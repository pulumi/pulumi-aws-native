// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// Resource Type definition for AWS::EC2::LaunchTemplate
    /// </summary>
    [Obsolete(@"LaunchTemplate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:LaunchTemplate")]
    public partial class LaunchTemplate : global::Pulumi.CustomResource
    {
        [Output("defaultVersionNumber")]
        public Output<string> DefaultVersionNumber { get; private set; } = null!;

        [Output("latestVersionNumber")]
        public Output<string> LatestVersionNumber { get; private set; } = null!;

        [Output("launchTemplateData")]
        public Output<Outputs.LaunchTemplateData> LaunchTemplateData { get; private set; } = null!;

        [Output("launchTemplateName")]
        public Output<string?> LaunchTemplateName { get; private set; } = null!;

        [Output("tagSpecifications")]
        public Output<ImmutableArray<Outputs.LaunchTemplateTagSpecification>> TagSpecifications { get; private set; } = null!;

        [Output("versionDescription")]
        public Output<string?> VersionDescription { get; private set; } = null!;


        /// <summary>
        /// Create a LaunchTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LaunchTemplate(string name, LaunchTemplateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:LaunchTemplate", name, args ?? new LaunchTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LaunchTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:LaunchTemplate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LaunchTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LaunchTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LaunchTemplate(name, id, options);
        }
    }

    public sealed class LaunchTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("launchTemplateData", required: true)]
        public Input<Inputs.LaunchTemplateDataArgs> LaunchTemplateData { get; set; } = null!;

        [Input("launchTemplateName")]
        public Input<string>? LaunchTemplateName { get; set; }

        [Input("tagSpecifications")]
        private InputList<Inputs.LaunchTemplateTagSpecificationArgs>? _tagSpecifications;
        public InputList<Inputs.LaunchTemplateTagSpecificationArgs> TagSpecifications
        {
            get => _tagSpecifications ?? (_tagSpecifications = new InputList<Inputs.LaunchTemplateTagSpecificationArgs>());
            set => _tagSpecifications = value;
        }

        [Input("versionDescription")]
        public Input<string>? VersionDescription { get; set; }

        public LaunchTemplateArgs()
        {
        }
        public static new LaunchTemplateArgs Empty => new LaunchTemplateArgs();
    }
}