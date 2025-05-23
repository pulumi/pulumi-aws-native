// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecretsManager
{
    /// <summary>
    /// Resource Type definition for AWS::SecretsManager::ResourcePolicy
    /// </summary>
    [AwsNativeResourceType("aws-native:secretsmanager:ResourcePolicy")]
    public partial class ResourcePolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Arn of the secret.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to block resource-based policies that allow broad access to the secret.
        /// </summary>
        [Output("blockPublicPolicy")]
        public Output<bool?> BlockPublicPolicy { get; private set; } = null!;

        /// <summary>
        /// A JSON-formatted string for an AWS resource-based policy.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SecretsManager::ResourcePolicy` for more information about the expected schema for this property.
        /// </summary>
        [Output("resourcePolicy")]
        public Output<object> ResourcePolicyValue { get; private set; } = null!;

        /// <summary>
        /// The ARN or name of the secret to attach the resource-based policy.
        /// </summary>
        [Output("secretId")]
        public Output<string> SecretId { get; private set; } = null!;


        /// <summary>
        /// Create a ResourcePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourcePolicy(string name, ResourcePolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:secretsmanager:ResourcePolicy", name, args ?? new ResourcePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourcePolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:secretsmanager:ResourcePolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "secretId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResourcePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourcePolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourcePolicy(name, id, options);
        }
    }

    public sealed class ResourcePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to block resource-based policies that allow broad access to the secret.
        /// </summary>
        [Input("blockPublicPolicy")]
        public Input<bool>? BlockPublicPolicy { get; set; }

        /// <summary>
        /// A JSON-formatted string for an AWS resource-based policy.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SecretsManager::ResourcePolicy` for more information about the expected schema for this property.
        /// </summary>
        [Input("resourcePolicy", required: true)]
        public Input<object> ResourcePolicyValue { get; set; } = null!;

        /// <summary>
        /// The ARN or name of the secret to attach the resource-based policy.
        /// </summary>
        [Input("secretId", required: true)]
        public Input<string> SecretId { get; set; } = null!;

        public ResourcePolicyArgs()
        {
        }
        public static new ResourcePolicyArgs Empty => new ResourcePolicyArgs();
    }
}
