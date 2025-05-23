// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecr
{
    /// <summary>
    /// The ``AWS::ECR::PullThroughCacheRule`` resource creates or updates a pull through cache rule. A pull through cache rule provides a way to cache images from an upstream registry in your Amazon ECR private registry.
    /// 
    /// ## Example Usage
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myECRPullThroughCacheRule = new AwsNative.Ecr.PullThroughCacheRule("myECRPullThroughCacheRule", new()
    ///     {
    ///         EcrRepositoryPrefix = "my-ecr",
    ///         UpstreamRegistryUrl = "public.ecr.aws",
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// ### Example
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AwsNative = Pulumi.AwsNative;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myECRPullThroughCacheRule = new AwsNative.Ecr.PullThroughCacheRule("myECRPullThroughCacheRule", new()
    ///     {
    ///         EcrRepositoryPrefix = "my-ecr",
    ///         UpstreamRegistryUrl = "public.ecr.aws",
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:ecr:PullThroughCacheRule")]
    public partial class PullThroughCacheRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Secrets Manager secret associated with the pull through cache rule.
        /// </summary>
        [Output("credentialArn")]
        public Output<string?> CredentialArn { get; private set; } = null!;

        /// <summary>
        /// The ARN of the IAM role associated with the pull through cache rule.
        /// </summary>
        [Output("customRoleArn")]
        public Output<string?> CustomRoleArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon ECR repository prefix associated with the pull through cache rule.
        /// </summary>
        [Output("ecrRepositoryPrefix")]
        public Output<string?> EcrRepositoryPrefix { get; private set; } = null!;

        /// <summary>
        /// The name of the upstream source registry associated with the pull through cache rule.
        /// </summary>
        [Output("upstreamRegistry")]
        public Output<string?> UpstreamRegistry { get; private set; } = null!;

        /// <summary>
        /// The upstream registry URL associated with the pull through cache rule.
        /// </summary>
        [Output("upstreamRegistryUrl")]
        public Output<string?> UpstreamRegistryUrl { get; private set; } = null!;

        /// <summary>
        /// The upstream repository prefix associated with the pull through cache rule.
        /// </summary>
        [Output("upstreamRepositoryPrefix")]
        public Output<string?> UpstreamRepositoryPrefix { get; private set; } = null!;


        /// <summary>
        /// Create a PullThroughCacheRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PullThroughCacheRule(string name, PullThroughCacheRuleArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ecr:PullThroughCacheRule", name, args ?? new PullThroughCacheRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PullThroughCacheRule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ecr:PullThroughCacheRule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "credentialArn",
                    "customRoleArn",
                    "ecrRepositoryPrefix",
                    "upstreamRegistry",
                    "upstreamRegistryUrl",
                    "upstreamRepositoryPrefix",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PullThroughCacheRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PullThroughCacheRule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PullThroughCacheRule(name, id, options);
        }
    }

    public sealed class PullThroughCacheRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Secrets Manager secret associated with the pull through cache rule.
        /// </summary>
        [Input("credentialArn")]
        public Input<string>? CredentialArn { get; set; }

        /// <summary>
        /// The ARN of the IAM role associated with the pull through cache rule.
        /// </summary>
        [Input("customRoleArn")]
        public Input<string>? CustomRoleArn { get; set; }

        /// <summary>
        /// The Amazon ECR repository prefix associated with the pull through cache rule.
        /// </summary>
        [Input("ecrRepositoryPrefix")]
        public Input<string>? EcrRepositoryPrefix { get; set; }

        /// <summary>
        /// The name of the upstream source registry associated with the pull through cache rule.
        /// </summary>
        [Input("upstreamRegistry")]
        public Input<string>? UpstreamRegistry { get; set; }

        /// <summary>
        /// The upstream registry URL associated with the pull through cache rule.
        /// </summary>
        [Input("upstreamRegistryUrl")]
        public Input<string>? UpstreamRegistryUrl { get; set; }

        /// <summary>
        /// The upstream repository prefix associated with the pull through cache rule.
        /// </summary>
        [Input("upstreamRepositoryPrefix")]
        public Input<string>? UpstreamRepositoryPrefix { get; set; }

        public PullThroughCacheRuleArgs()
        {
        }
        public static new PullThroughCacheRuleArgs Empty => new PullThroughCacheRuleArgs();
    }
}
