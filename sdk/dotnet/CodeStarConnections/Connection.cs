// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeStarConnections
{
    /// <summary>
    /// Schema for AWS::CodeStarConnections::Connection resource which can be used to connect external source providers with AWS CodePipeline
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
    ///     var sampleConnection = new AwsNative.CodeStarConnections.Connection("sampleConnection", new()
    ///     {
    ///         ConnectionName = "MyConnection",
    ///         ProviderType = "Bitbucket",
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "Project",
    ///                 Value = "ProjectB",
    ///             },
    ///         },
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
    ///     var sampleConnection = new AwsNative.CodeStarConnections.Connection("sampleConnection", new()
    ///     {
    ///         ConnectionName = "MyConnection",
    ///         ProviderType = "Bitbucket",
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "Project",
    ///                 Value = "ProjectB",
    ///             },
    ///         },
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
    ///     var sampleConnection = new AwsNative.CodeStarConnections.Connection("sampleConnection", new()
    ///     {
    ///         ConnectionName = "MyConnection",
    ///         ProviderType = "GitHubEnterpriseServer",
    ///         HostArn = "arn:aws:codestar-connections:us-west-2:123456789123:host/abc123-example",
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "Project",
    ///                 Value = "ProjectB",
    ///             },
    ///         },
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
    ///     var sampleConnection = new AwsNative.CodeStarConnections.Connection("sampleConnection", new()
    ///     {
    ///         ConnectionName = "MyConnection",
    ///         ProviderType = "GitHubEnterpriseServer",
    ///         HostArn = "arn:aws:codestar-connections:us-west-2:123456789123:host/abc123-example",
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "Project",
    ///                 Value = "ProjectB",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:codestarconnections:Connection")]
    public partial class Connection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the  connection. The ARN is used as the connection reference when the connection is shared between AWS services.
        /// </summary>
        [Output("connectionArn")]
        public Output<string> ConnectionArn { get; private set; } = null!;

        /// <summary>
        /// The name of the connection. Connection names must be unique in an AWS user account.
        /// </summary>
        [Output("connectionName")]
        public Output<string> ConnectionName { get; private set; } = null!;

        /// <summary>
        /// The current status of the connection.
        /// </summary>
        [Output("connectionStatus")]
        public Output<string> ConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.
        /// </summary>
        [Output("hostArn")]
        public Output<string?> HostArn { get; private set; } = null!;

        /// <summary>
        /// The name of the external provider where your third-party code repository is configured. For Bitbucket, this is the account ID of the owner of the Bitbucket repository.
        /// </summary>
        [Output("ownerAccountId")]
        public Output<string> OwnerAccountId { get; private set; } = null!;

        /// <summary>
        /// The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.
        /// </summary>
        [Output("providerType")]
        public Output<string?> ProviderType { get; private set; } = null!;

        /// <summary>
        /// Specifies the tags applied to a connection.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Connection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connection(string name, ConnectionArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:codestarconnections:Connection", name, args ?? new ConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Connection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:codestarconnections:Connection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "connectionName",
                    "hostArn",
                    "providerType",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Connection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Connection(name, id, options);
        }
    }

    public sealed class ConnectionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the connection. Connection names must be unique in an AWS user account.
        /// </summary>
        [Input("connectionName")]
        public Input<string>? ConnectionName { get; set; }

        /// <summary>
        /// The host arn configured to represent the infrastructure where your third-party provider is installed. You must specify either a ProviderType or a HostArn.
        /// </summary>
        [Input("hostArn")]
        public Input<string>? HostArn { get; set; }

        /// <summary>
        /// The name of the external provider where your third-party code repository is configured. You must specify either a ProviderType or a HostArn.
        /// </summary>
        [Input("providerType")]
        public Input<string>? ProviderType { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// Specifies the tags applied to a connection.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ConnectionArgs()
        {
        }
        public static new ConnectionArgs Empty => new ConnectionArgs();
    }
}
