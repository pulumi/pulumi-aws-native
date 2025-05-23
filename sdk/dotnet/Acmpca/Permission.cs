// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca
{
    /// <summary>
    /// Permission set on private certificate authority
    /// </summary>
    [AwsNativeResourceType("aws-native:acmpca:Permission")]
    public partial class Permission : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<string>> Actions { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.
        /// </summary>
        [Output("certificateAuthorityArn")]
        public Output<string> CertificateAuthorityArn { get; private set; } = null!;

        /// <summary>
        /// The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// The ID of the calling account.
        /// </summary>
        [Output("sourceAccount")]
        public Output<string?> SourceAccount { get; private set; } = null!;


        /// <summary>
        /// Create a Permission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Permission(string name, PermissionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:Permission", name, args ?? new PermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Permission(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:acmpca:Permission", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "actions[*]",
                    "certificateAuthorityArn",
                    "principal",
                    "sourceAccount",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Permission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Permission Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Permission(name, id, options);
        }
    }

    public sealed class PermissionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<string>? _actions;

        /// <summary>
        /// The actions that the specified AWS service principal can use. Actions IssueCertificate, GetCertificate and ListPermissions must be provided.
        /// </summary>
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Private Certificate Authority that grants the permission.
        /// </summary>
        [Input("certificateAuthorityArn", required: true)]
        public Input<string> CertificateAuthorityArn { get; set; } = null!;

        /// <summary>
        /// The AWS service or identity that receives the permission. At this time, the only valid principal is acm.amazonaws.com.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// The ID of the calling account.
        /// </summary>
        [Input("sourceAccount")]
        public Input<string>? SourceAccount { get; set; }

        public PermissionArgs()
        {
        }
        public static new PermissionArgs Empty => new PermissionArgs();
    }
}
