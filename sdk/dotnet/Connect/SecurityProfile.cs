// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    /// <summary>
    /// Resource Type definition for AWS::Connect::SecurityProfile
    /// </summary>
    [AwsNativeResourceType("aws-native:connect:SecurityProfile")]
    public partial class SecurityProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The list of tags that a security profile uses to restrict access to resources in Amazon Connect.
        /// </summary>
        [Output("allowedAccessControlTags")]
        public Output<ImmutableArray<Outputs.SecurityProfileTag>> AllowedAccessControlTags { get; private set; } = null!;

        /// <summary>
        /// The description of the security profile.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        [Output("instanceArn")]
        public Output<string> InstanceArn { get; private set; } = null!;

        /// <summary>
        /// Permissions assigned to the security profile.
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<string>> Permissions { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the security profile.
        /// </summary>
        [Output("securityProfileArn")]
        public Output<string> SecurityProfileArn { get; private set; } = null!;

        /// <summary>
        /// The name of the security profile.
        /// </summary>
        [Output("securityProfileName")]
        public Output<string> SecurityProfileName { get; private set; } = null!;

        /// <summary>
        /// The list of resources that a security profile applies tag restrictions to in Amazon Connect.
        /// </summary>
        [Output("tagRestrictedResources")]
        public Output<ImmutableArray<string>> TagRestrictedResources { get; private set; } = null!;

        /// <summary>
        /// The tags used to organize, track, or control access for this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.SecurityProfileTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityProfile(string name, SecurityProfileArgs args, CustomResourceOptions? options = null)
            : base("aws-native:connect:SecurityProfile", name, args ?? new SecurityProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:connect:SecurityProfile", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "instanceArn",
                    "securityProfileName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityProfile(name, id, options);
        }
    }

    public sealed class SecurityProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedAccessControlTags")]
        private InputList<Inputs.SecurityProfileTagArgs>? _allowedAccessControlTags;

        /// <summary>
        /// The list of tags that a security profile uses to restrict access to resources in Amazon Connect.
        /// </summary>
        public InputList<Inputs.SecurityProfileTagArgs> AllowedAccessControlTags
        {
            get => _allowedAccessControlTags ?? (_allowedAccessControlTags = new InputList<Inputs.SecurityProfileTagArgs>());
            set => _allowedAccessControlTags = value;
        }

        /// <summary>
        /// The description of the security profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The identifier of the Amazon Connect instance.
        /// </summary>
        [Input("instanceArn", required: true)]
        public Input<string> InstanceArn { get; set; } = null!;

        [Input("permissions")]
        private InputList<string>? _permissions;

        /// <summary>
        /// Permissions assigned to the security profile.
        /// </summary>
        public InputList<string> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        /// <summary>
        /// The name of the security profile.
        /// </summary>
        [Input("securityProfileName")]
        public Input<string>? SecurityProfileName { get; set; }

        [Input("tagRestrictedResources")]
        private InputList<string>? _tagRestrictedResources;

        /// <summary>
        /// The list of resources that a security profile applies tag restrictions to in Amazon Connect.
        /// </summary>
        public InputList<string> TagRestrictedResources
        {
            get => _tagRestrictedResources ?? (_tagRestrictedResources = new InputList<string>());
            set => _tagRestrictedResources = value;
        }

        [Input("tags")]
        private InputList<Inputs.SecurityProfileTagArgs>? _tags;

        /// <summary>
        /// The tags used to organize, track, or control access for this resource.
        /// </summary>
        public InputList<Inputs.SecurityProfileTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.SecurityProfileTagArgs>());
            set => _tags = value;
        }

        public SecurityProfileArgs()
        {
        }
        public static new SecurityProfileArgs Empty => new SecurityProfileArgs();
    }
}