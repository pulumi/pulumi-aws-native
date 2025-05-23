// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sns
{
    /// <summary>
    /// Schema for AWS::SNS::TopicInlinePolicy
    /// </summary>
    [AwsNativeResourceType("aws-native:sns:TopicInlinePolicy")]
    public partial class TopicInlinePolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A policy document that contains permissions to add to the specified SNS topics.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::TopicInlinePolicy` for more information about the expected schema for this property.
        /// </summary>
        [Output("policyDocument")]
        public Output<object> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the topic to which you want to add the policy.
        /// </summary>
        [Output("topicArn")]
        public Output<string> TopicArn { get; private set; } = null!;


        /// <summary>
        /// Create a TopicInlinePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TopicInlinePolicy(string name, TopicInlinePolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:sns:TopicInlinePolicy", name, args ?? new TopicInlinePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TopicInlinePolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:sns:TopicInlinePolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "topicArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TopicInlinePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TopicInlinePolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TopicInlinePolicy(name, id, options);
        }
    }

    public sealed class TopicInlinePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A policy document that contains permissions to add to the specified SNS topics.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::SNS::TopicInlinePolicy` for more information about the expected schema for this property.
        /// </summary>
        [Input("policyDocument", required: true)]
        public Input<object> PolicyDocument { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the topic to which you want to add the policy.
        /// </summary>
        [Input("topicArn", required: true)]
        public Input<string> TopicArn { get; set; } = null!;

        public TopicInlinePolicyArgs()
        {
        }
        public static new TopicInlinePolicyArgs Empty => new TopicInlinePolicyArgs();
    }
}
