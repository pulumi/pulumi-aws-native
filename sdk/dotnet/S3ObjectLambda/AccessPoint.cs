// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3ObjectLambda
{
    /// <summary>
    /// The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions
    /// </summary>
    [AwsNativeResourceType("aws-native:s3objectlambda:AccessPoint")]
    public partial class AccessPoint : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The date and time when the Object lambda Access Point was created.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The name you want to assign to this Object lambda Access Point.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
        /// </summary>
        [Output("objectLambdaConfiguration")]
        public Output<Outputs.AccessPointObjectLambdaConfiguration> ObjectLambdaConfiguration { get; private set; } = null!;

        [Output("policyStatus")]
        public Output<Outputs.PolicyStatusProperties> PolicyStatus { get; private set; } = null!;

        /// <summary>
        /// The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide.
        /// </summary>
        [Output("publicAccessBlockConfiguration")]
        public Output<Outputs.AccessPointPublicAccessBlockConfiguration> PublicAccessBlockConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPoint(string name, AccessPointArgs args, CustomResourceOptions? options = null)
            : base("aws-native:s3objectlambda:AccessPoint", name, args ?? new AccessPointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:s3objectlambda:AccessPoint", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccessPoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessPoint(name, id, options);
        }
    }

    public sealed class AccessPointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name you want to assign to this Object lambda Access Point.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions
        /// </summary>
        [Input("objectLambdaConfiguration", required: true)]
        public Input<Inputs.AccessPointObjectLambdaConfigurationArgs> ObjectLambdaConfiguration { get; set; } = null!;

        public AccessPointArgs()
        {
        }
        public static new AccessPointArgs Empty => new AccessPointArgs();
    }
}