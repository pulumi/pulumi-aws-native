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
    /// Resource Type definition for AWS::Route53Profiles::ProfileResourceAssociation
    /// </summary>
    [AwsNativeResourceType("aws-native:route53profiles:ProfileResourceAssociation")]
    public partial class ProfileResourceAssociation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Primary Identifier for  Profile Resource Association
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// The name of an association between the  Profile and resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the  profile that you associated the resource to that is specified by ResourceArn.
        /// </summary>
        [Output("profileId")]
        public Output<string> ProfileId { get; private set; } = null!;

        /// <summary>
        /// The arn of the resource that you associated to the  Profile.
        /// </summary>
        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;

        /// <summary>
        /// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
        /// </summary>
        [Output("resourceProperties")]
        public Output<string?> ResourceProperties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource associated to the  Profile.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;


        /// <summary>
        /// Create a ProfileResourceAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfileResourceAssociation(string name, ProfileResourceAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:route53profiles:ProfileResourceAssociation", name, args ?? new ProfileResourceAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfileResourceAssociation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:route53profiles:ProfileResourceAssociation", name, null, MakeResourceOptions(options, id))
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
                    "resourceArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProfileResourceAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfileResourceAssociation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProfileResourceAssociation(name, id, options);
        }
    }

    public sealed class ProfileResourceAssociationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of an association between the  Profile and resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the  profile that you associated the resource to that is specified by ResourceArn.
        /// </summary>
        [Input("profileId", required: true)]
        public Input<string> ProfileId { get; set; } = null!;

        /// <summary>
        /// The arn of the resource that you associated to the  Profile.
        /// </summary>
        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        /// <summary>
        /// A JSON-formatted string with key-value pairs specifying the properties of the associated resource.
        /// </summary>
        [Input("resourceProperties")]
        public Input<string>? ResourceProperties { get; set; }

        public ProfileResourceAssociationArgs()
        {
        }
        public static new ProfileResourceAssociationArgs Empty => new ProfileResourceAssociationArgs();
    }
}