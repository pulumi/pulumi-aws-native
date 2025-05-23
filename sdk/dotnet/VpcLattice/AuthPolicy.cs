// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice
{
    /// <summary>
    /// Creates or updates the auth policy.
    /// </summary>
    [AwsNativeResourceType("aws-native:vpclattice:AuthPolicy")]
    public partial class AuthPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The auth policy.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::VpcLattice::AuthPolicy` for more information about the expected schema for this property.
        /// </summary>
        [Output("policy")]
        public Output<object> Policy { get; private set; } = null!;

        /// <summary>
        /// The ID or ARN of the service network or service for which the policy is created.
        /// </summary>
        [Output("resourceIdentifier")]
        public Output<string> ResourceIdentifier { get; private set; } = null!;

        /// <summary>
        /// The state of the auth policy. The auth policy is only active when the auth type is set to `AWS _IAM` . If you provide a policy, then authentication and authorization decisions are made based on this policy and the client's IAM policy. If the auth type is `NONE` , then any auth policy you provide will remain inactive.
        /// </summary>
        [Output("state")]
        public Output<Pulumi.AwsNative.VpcLattice.AuthPolicyState> State { get; private set; } = null!;


        /// <summary>
        /// Create a AuthPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthPolicy(string name, AuthPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:vpclattice:AuthPolicy", name, args ?? new AuthPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:vpclattice:AuthPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "resourceIdentifier",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AuthPolicy(name, id, options);
        }
    }

    public sealed class AuthPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The auth policy.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::VpcLattice::AuthPolicy` for more information about the expected schema for this property.
        /// </summary>
        [Input("policy", required: true)]
        public Input<object> Policy { get; set; } = null!;

        /// <summary>
        /// The ID or ARN of the service network or service for which the policy is created.
        /// </summary>
        [Input("resourceIdentifier", required: true)]
        public Input<string> ResourceIdentifier { get; set; } = null!;

        public AuthPolicyArgs()
        {
        }
        public static new AuthPolicyArgs Empty => new AuthPolicyArgs();
    }
}
