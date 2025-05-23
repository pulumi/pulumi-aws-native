// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    /// <summary>
    /// Creates a continuous deployment policy that routes a subset of production traffic from a primary distribution to a staging distribution.
    ///  After you create and update a staging distribution, you can use a continuous deployment policy to incrementally move traffic to the staging distribution. This enables you to test changes to a distribution's configuration before moving all of your production traffic to the new configuration.
    ///  For more information, see [Using CloudFront continuous deployment to safely test CDN configuration changes](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/continuous-deployment.html) in the *Amazon CloudFront Developer Guide*.
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudfront:ContinuousDeploymentPolicy")]
    public partial class ContinuousDeploymentPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The identifier of the cotinuous deployment policy.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// Contains the configuration for a continuous deployment policy.
        /// </summary>
        [Output("continuousDeploymentPolicyConfig")]
        public Output<Outputs.ContinuousDeploymentPolicyConfig> ContinuousDeploymentPolicyConfig { get; private set; } = null!;

        /// <summary>
        /// The date and time when the continuous deployment policy was last modified.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;


        /// <summary>
        /// Create a ContinuousDeploymentPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContinuousDeploymentPolicy(string name, ContinuousDeploymentPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:ContinuousDeploymentPolicy", name, args ?? new ContinuousDeploymentPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContinuousDeploymentPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:ContinuousDeploymentPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ContinuousDeploymentPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContinuousDeploymentPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ContinuousDeploymentPolicy(name, id, options);
        }
    }

    public sealed class ContinuousDeploymentPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains the configuration for a continuous deployment policy.
        /// </summary>
        [Input("continuousDeploymentPolicyConfig", required: true)]
        public Input<Inputs.ContinuousDeploymentPolicyConfigArgs> ContinuousDeploymentPolicyConfig { get; set; } = null!;

        public ContinuousDeploymentPolicyArgs()
        {
        }
        public static new ContinuousDeploymentPolicyArgs Empty => new ContinuousDeploymentPolicyArgs();
    }
}
