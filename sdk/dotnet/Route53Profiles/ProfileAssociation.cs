// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Route53Profiles
{
    /// <summary>
    /// Resource Type definition for AWS::Route53Profiles::ProfileAssociation
    /// </summary>
    [AwsNativeResourceType("aws-native:route53profiles:ProfileAssociation")]
    public partial class ProfileAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the  profile association.
        /// </summary>
        [Output("arn")]
        public Output<string?> Arn { get; private set; } = null!;

        /// <summary>
        /// Primary Identifier for  Profile Association
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The name of an association between a  Profile and a VPC.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the  profile that you associated with the resource that is specified by ResourceId.
        /// </summary>
        [Output("profileId")]
        public Output<string> ProfileId { get; private set; } = null!;

        /// <summary>
        /// The resource that you associated the  profile with.
        /// </summary>
        [Output("resourceId")]
        public Output<string> ResourceId { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileAssociation(string name, ProfileAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:route53profiles:ProfileAssociation", name, args ?? new ProfileAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:route53profiles:ProfileAssociation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                    "profileId",
                    "resourceId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProfileAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProfileAssociation(name, id, options);
        }
    }

    public sealed class ProfileAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the  profile association.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of an association between a  Profile and a VPC.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the  profile that you associated with the resource that is specified by ResourceId.
        /// </summary>
        [Input("profileId", required: true)]
        public Input<string> ProfileId { get; set; } = null!;

        /// <summary>
        /// The resource that you associated the  profile with.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

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

        public ProfileAssociationArgs()
        {
        }
        public static new ProfileAssociationArgs Empty => new ProfileAssociationArgs();
    }
}