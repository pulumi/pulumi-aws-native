// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3Express
{
    /// <summary>
    /// Resource Type definition for AWS::S3Express::BucketPolicy.
    /// </summary>
    [AwsNativeResourceType("aws-native:s3express:BucketPolicy")]
    public partial class BucketPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the S3 directory bucket to which the policy applies.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3Express::BucketPolicy` for more information about the expected schema for this property.
        /// </summary>
        [Output("policyDocument")]
        public Output<object> PolicyDocument { get; private set; } = null!;


        /// <summary>
        /// Create a BucketPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketPolicy(string name, BucketPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:s3express:BucketPolicy", name, args ?? new BucketPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:s3express:BucketPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "bucket",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BucketPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BucketPolicy(name, id, options);
        }
    }

    public sealed class BucketPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the S3 directory bucket to which the policy applies.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// A policy document containing permissions to add to the specified bucket. In IAM, you must provide policy documents in JSON format. However, in CloudFormation you can provide the policy in JSON or YAML format because CloudFormation converts YAML to JSON before submitting it to IAM.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::S3Express::BucketPolicy` for more information about the expected schema for this property.
        /// </summary>
        [Input("policyDocument", required: true)]
        public Input<object> PolicyDocument { get; set; } = null!;

        public BucketPolicyArgs()
        {
        }
        public static new BucketPolicyArgs Empty => new BucketPolicyArgs();
    }
}
