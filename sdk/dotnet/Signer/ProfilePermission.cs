// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Signer
{
    /// <summary>
    /// An example resource schema demonstrating some basic constructs and validation rules.
    /// </summary>
    [AwsNativeResourceType("aws-native:signer:ProfilePermission")]
    public partial class ProfilePermission : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS Signer action permitted as part of cross-account permissions.
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The AWS principal receiving cross-account permissions. This may be an IAM role or another AWS account ID.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The human-readable name of the signing profile.
        /// </summary>
        [Output("profileName")]
        public Output<string> ProfileName { get; private set; } = null!;

        /// <summary>
        /// The version of the signing profile.
        /// </summary>
        [Output("profileVersion")]
        public Output<string?> ProfileVersion { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for the cross-account permission statement.
        /// </summary>
        [Output("statementId")]
        public Output<string> StatementId { get; private set; } = null!;


        /// <summary>
        /// Create a ProfilePermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProfilePermission(string name, ProfilePermissionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:signer:ProfilePermission", name, args ?? new ProfilePermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProfilePermission(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:signer:ProfilePermission", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "action",
                    "principal",
                    "profileName",
                    "profileVersion",
                    "statementId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProfilePermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProfilePermission Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProfilePermission(name, id, options);
        }
    }

    public sealed class ProfilePermissionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS Signer action permitted as part of cross-account permissions.
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The AWS principal receiving cross-account permissions. This may be an IAM role or another AWS account ID.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// The human-readable name of the signing profile.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// The version of the signing profile.
        /// </summary>
        [Input("profileVersion")]
        public Input<string>? ProfileVersion { get; set; }

        /// <summary>
        /// A unique identifier for the cross-account permission statement.
        /// </summary>
        [Input("statementId", required: true)]
        public Input<string> StatementId { get; set; } = null!;

        public ProfilePermissionArgs()
        {
        }
        public static new ProfilePermissionArgs Empty => new ProfilePermissionArgs();
    }
}
