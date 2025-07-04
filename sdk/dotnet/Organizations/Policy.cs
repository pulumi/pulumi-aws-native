// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Organizations
{
    /// <summary>
    /// Policies in AWS Organizations enable you to manage different features of the AWS accounts in your organization.  You can use policies when all features are enabled in your organization.
    /// </summary>
    [AwsNativeResourceType("aws-native:organizations:Policy")]
    public partial class Policy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ARN of the Policy
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Id of the Policy
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// A boolean value that indicates whether the specified policy is an AWS managed policy. If true, then you can attach the policy to roots, OUs, or accounts, but you cannot edit it.
        /// </summary>
        [Output("awsManaged")]
        public Output<bool> AwsManaged { get; private set; } = null!;

        /// <summary>
        /// The Policy text content. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Organizations::Policy` for more information about the expected schema for this property.
        /// </summary>
        [Output("content")]
        public Output<object> Content { get; private set; } = null!;

        /// <summary>
        /// Human readable description of the policy
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the Policy
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to
        /// </summary>
        [Output("targetIds")]
        public Output<ImmutableArray<string>> TargetIds { get; private set; } = null!;

        /// <summary>
        /// The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY, CHATBOT_POLICY, RESOURCE_CONTROL_POLICY,DECLARATIVE_POLICY_EC2, SECURITYHUB_POLICY
        /// </summary>
        [Output("type")]
        public Output<Pulumi.AwsNative.Organizations.PolicyType> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:organizations:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:organizations:Policy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "type",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, options);
        }
    }

    public sealed class PolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Policy text content. For AWS CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. AWS CloudFormation always converts a YAML policy to JSON format before submitting it.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Organizations::Policy` for more information about the expected schema for this property.
        /// </summary>
        [Input("content", required: true)]
        public Input<object> Content { get; set; } = null!;

        /// <summary>
        /// Human readable description of the policy
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the Policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// A list of tags that you want to attach to the newly created policy. For each tag in the list, you must specify both a tag key and a value. You can set the value to an empty string, but you can't set it to null.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("targetIds")]
        private InputList<string>? _targetIds;

        /// <summary>
        /// List of unique identifiers (IDs) of the root, OU, or account that you want to attach the policy to
        /// </summary>
        public InputList<string> TargetIds
        {
            get => _targetIds ?? (_targetIds = new InputList<string>());
            set => _targetIds = value;
        }

        /// <summary>
        /// The type of policy to create. You can specify one of the following values: AISERVICES_OPT_OUT_POLICY, BACKUP_POLICY, SERVICE_CONTROL_POLICY, TAG_POLICY, CHATBOT_POLICY, RESOURCE_CONTROL_POLICY,DECLARATIVE_POLICY_EC2, SECURITYHUB_POLICY
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.Organizations.PolicyType> Type { get; set; } = null!;

        public PolicyArgs()
        {
        }
        public static new PolicyArgs Empty => new PolicyArgs();
    }
}
