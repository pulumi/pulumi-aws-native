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
    /// The ``AWS::ECR::Repository`` resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For more information, see [Amazon ECR private repositories](https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html) in the *Amazon ECR User Guide*.
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
    ///     var config = new Config();
    ///     var lifecyclePolicyText = config.Require("lifecyclePolicyText");
    ///     var repositoryName = config.Require("repositoryName");
    ///     var registryId = config.Require("registryId");
    ///     var myRepository = new AwsNative.Ecr.Repository("myRepository", new()
    ///     {
    ///         LifecyclePolicy = new AwsNative.Ecr.Inputs.RepositoryLifecyclePolicyArgs
    ///         {
    ///             LifecyclePolicyText = lifecyclePolicyText,
    ///             RegistryId = registryId,
    ///         },
    ///         RepositoryName = repositoryName,
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["arn"] = myRepository.Arn,
    ///     };
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
    ///     var config = new Config();
    ///     var lifecyclePolicyText = config.Require("lifecyclePolicyText");
    ///     var repositoryName = config.Require("repositoryName");
    ///     var registryId = config.Require("registryId");
    ///     var myRepository = new AwsNative.Ecr.Repository("myRepository", new()
    ///     {
    ///         LifecyclePolicy = new AwsNative.Ecr.Inputs.RepositoryLifecyclePolicyArgs
    ///         {
    ///             LifecyclePolicyText = lifecyclePolicyText,
    ///             RegistryId = registryId,
    ///         },
    ///         RepositoryName = repositoryName,
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["arn"] = myRepository.Arn,
    ///     };
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:ecr:Repository")]
    public partial class Repository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Returns the Amazon Resource Name (ARN) for the specified `AWS::ECR::Repository` resource. For example, `arn:aws:ecr: *eu-west-1* : *123456789012* :repository/ *test-repository*` .
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// If true, deleting the repository force deletes the contents of the repository. If false, the repository must be empty before attempting to delete it.
        /// </summary>
        [Output("emptyOnDelete")]
        public Output<bool?> EmptyOnDelete { get; private set; } = null!;

        /// <summary>
        /// The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.
        /// </summary>
        [Output("encryptionConfiguration")]
        public Output<Outputs.RepositoryEncryptionConfiguration?> EncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.
        /// </summary>
        [Output("imageScanningConfiguration")]
        public Output<Outputs.RepositoryImageScanningConfiguration?> ImageScanningConfiguration { get; private set; } = null!;

        /// <summary>
        /// The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
        /// </summary>
        [Output("imageTagMutability")]
        public Output<Pulumi.AwsNative.Ecr.RepositoryImageTagMutability?> ImageTagMutability { get; private set; } = null!;

        /// <summary>
        /// Creates or updates a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html).
        /// </summary>
        [Output("lifecyclePolicy")]
        public Output<Outputs.RepositoryLifecyclePolicy?> LifecyclePolicy { get; private set; } = null!;

        /// <summary>
        /// The name to use for the repository. The repository name may be specified on its own (such as ``nginx-web-app``) or it can be prepended with a namespace to group the repository into a category (such as ``project-a/nginx-web-app``). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
        ///  The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.
        ///   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Output("repositoryName")]
        public Output<string?> RepositoryName { get; private set; } = null!;

        /// <summary>
        /// The JSON repository policy text to apply to the repository. For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide*.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ECR::Repository` for more information about the expected schema for this property.
        /// </summary>
        [Output("repositoryPolicyText")]
        public Output<object?> RepositoryPolicyText { get; private set; } = null!;

        /// <summary>
        /// Returns the URI for the specified `AWS::ECR::Repository` resource. For example, `*123456789012* .dkr.ecr. *us-west-2* .amazonaws.com/repository` .
        /// </summary>
        [Output("repositoryUri")]
        public Output<string> RepositoryUri { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ecr:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ecr:Repository", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "encryptionConfiguration",
                    "repositoryName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Repository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repository Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Repository(name, id, options);
        }
    }

    public sealed class RepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, deleting the repository force deletes the contents of the repository. If false, the repository must be empty before attempting to delete it.
        /// </summary>
        [Input("emptyOnDelete")]
        public Input<bool>? EmptyOnDelete { get; set; }

        /// <summary>
        /// The encryption configuration for the repository. This determines how the contents of your repository are encrypted at rest.
        /// </summary>
        [Input("encryptionConfiguration")]
        public Input<Inputs.RepositoryEncryptionConfigurationArgs>? EncryptionConfiguration { get; set; }

        /// <summary>
        /// The image scanning configuration for the repository. This determines whether images are scanned for known vulnerabilities after being pushed to the repository.
        /// </summary>
        [Input("imageScanningConfiguration")]
        public Input<Inputs.RepositoryImageScanningConfigurationArgs>? ImageScanningConfiguration { get; set; }

        /// <summary>
        /// The tag mutability setting for the repository. If this parameter is omitted, the default setting of ``MUTABLE`` will be used which will allow image tags to be overwritten. If ``IMMUTABLE`` is specified, all image tags within the repository will be immutable which will prevent them from being overwritten.
        /// </summary>
        [Input("imageTagMutability")]
        public Input<Pulumi.AwsNative.Ecr.RepositoryImageTagMutability>? ImageTagMutability { get; set; }

        /// <summary>
        /// Creates or updates a lifecycle policy. For information about lifecycle policy syntax, see [Lifecycle policy template](https://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html).
        /// </summary>
        [Input("lifecyclePolicy")]
        public Input<Inputs.RepositoryLifecyclePolicyArgs>? LifecyclePolicy { get; set; }

        /// <summary>
        /// The name to use for the repository. The repository name may be specified on its own (such as ``nginx-web-app``) or it can be prepended with a namespace to group the repository into a category (such as ``project-a/nginx-web-app``). If you don't specify a name, CFNlong generates a unique physical ID and uses that ID for the repository name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
        ///  The repository name must start with a letter and can only contain lowercase letters, numbers, hyphens, underscores, and forward slashes.
        ///   If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

        /// <summary>
        /// The JSON repository policy text to apply to the repository. For more information, see [Amazon ECR repository policies](https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-policy-examples.html) in the *Amazon Elastic Container Registry User Guide*.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::ECR::Repository` for more information about the expected schema for this property.
        /// </summary>
        [Input("repositoryPolicyText")]
        public Input<object>? RepositoryPolicyText { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public RepositoryArgs()
        {
        }
        public static new RepositoryArgs Empty => new RepositoryArgs();
    }
}
