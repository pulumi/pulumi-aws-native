// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise
{
    /// <summary>
    /// Resource schema for AWS::IoTSiteWise::AccessPolicy
    /// </summary>
    [AwsNativeResourceType("aws-native:iotsitewise:AccessPolicy")]
    public partial class AccessPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the access policy.
        /// </summary>
        [Output("accessPolicyArn")]
        public Output<string> AccessPolicyArn { get; private set; } = null!;

        /// <summary>
        /// The ID of the access policy.
        /// </summary>
        [Output("accessPolicyId")]
        public Output<string> AccessPolicyId { get; private set; } = null!;

        /// <summary>
        /// The identity for this access policy. Choose either a user or a group but not both.
        /// </summary>
        [Output("accessPolicyIdentity")]
        public Output<Outputs.AccessPolicyIdentity> AccessPolicyIdentity { get; private set; } = null!;

        /// <summary>
        /// The permission level for this access policy. Valid values are ADMINISTRATOR or VIEWER.
        /// </summary>
        [Output("accessPolicyPermission")]
        public Output<string> AccessPolicyPermission { get; private set; } = null!;

        /// <summary>
        /// The AWS IoT SiteWise Monitor resource for this access policy. Choose either portal or project but not both.
        /// </summary>
        [Output("accessPolicyResource")]
        public Output<Outputs.AccessPolicyResource> AccessPolicyResource { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPolicy(string name, AccessPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:iotsitewise:AccessPolicy", name, args ?? new AccessPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iotsitewise:AccessPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccessPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessPolicy(name, id, options);
        }
    }

    public sealed class AccessPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity for this access policy. Choose either a user or a group but not both.
        /// </summary>
        [Input("accessPolicyIdentity", required: true)]
        public Input<Inputs.AccessPolicyIdentityArgs> AccessPolicyIdentity { get; set; } = null!;

        /// <summary>
        /// The permission level for this access policy. Valid values are ADMINISTRATOR or VIEWER.
        /// </summary>
        [Input("accessPolicyPermission", required: true)]
        public Input<string> AccessPolicyPermission { get; set; } = null!;

        /// <summary>
        /// The AWS IoT SiteWise Monitor resource for this access policy. Choose either portal or project but not both.
        /// </summary>
        [Input("accessPolicyResource", required: true)]
        public Input<Inputs.AccessPolicyResourceArgs> AccessPolicyResource { get; set; } = null!;

        public AccessPolicyArgs()
        {
        }
        public static new AccessPolicyArgs Empty => new AccessPolicyArgs();
    }
}
