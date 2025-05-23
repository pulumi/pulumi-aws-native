// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeArtifact
{
    /// <summary>
    /// The resource schema to create a CodeArtifact repository.
    /// </summary>
    [AwsNativeResourceType("aws-native:codeartifact:Repository")]
    public partial class Repository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the repository.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A text description of the repository.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the domain that contains the repository.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The 12-digit account ID of the AWS account that owns the domain.
        /// </summary>
        [Output("domainOwner")]
        public Output<string> DomainOwner { get; private set; } = null!;

        /// <summary>
        /// A list of external connections associated with the repository.
        /// </summary>
        [Output("externalConnections")]
        public Output<ImmutableArray<string>> ExternalConnections { get; private set; } = null!;

        /// <summary>
        /// The name of the repository. This is used for GetAtt
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The access control resource policy on the provided repository.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CodeArtifact::Repository` for more information about the expected schema for this property.
        /// </summary>
        [Output("permissionsPolicyDocument")]
        public Output<object?> PermissionsPolicyDocument { get; private set; } = null!;

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Output("repositoryName")]
        public Output<string> RepositoryName { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// A list of upstream repositories associated with the repository.
        /// </summary>
        [Output("upstreams")]
        public Output<ImmutableArray<string>> Upstreams { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs args, CustomResourceOptions? options = null)
            : base("aws-native:codeartifact:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:codeartifact:Repository", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "domainName",
                    "domainOwner",
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
        /// A text description of the repository.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the domain that contains the repository.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The 12-digit account ID of the AWS account that owns the domain.
        /// </summary>
        [Input("domainOwner")]
        public Input<string>? DomainOwner { get; set; }

        [Input("externalConnections")]
        private InputList<string>? _externalConnections;

        /// <summary>
        /// A list of external connections associated with the repository.
        /// </summary>
        public InputList<string> ExternalConnections
        {
            get => _externalConnections ?? (_externalConnections = new InputList<string>());
            set => _externalConnections = value;
        }

        /// <summary>
        /// The access control resource policy on the provided repository.
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::CodeArtifact::Repository` for more information about the expected schema for this property.
        /// </summary>
        [Input("permissionsPolicyDocument")]
        public Input<object>? PermissionsPolicyDocument { get; set; }

        /// <summary>
        /// The name of the repository.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

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

        [Input("upstreams")]
        private InputList<string>? _upstreams;

        /// <summary>
        /// A list of upstream repositories associated with the repository.
        /// </summary>
        public InputList<string> Upstreams
        {
            get => _upstreams ?? (_upstreams = new InputList<string>());
            set => _upstreams = value;
        }

        public RepositoryArgs()
        {
        }
        public static new RepositoryArgs Empty => new RepositoryArgs();
    }
}
