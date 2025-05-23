// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppConfig
{
    /// <summary>
    /// Resource Type definition for AWS::AppConfig::Extension
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
    ///     var basicExtension = new AwsNative.AppConfig.Extension("basicExtension", new()
    ///     {
    ///         Name = "My Test Extension",
    ///         Description = "My test extension",
    ///         LatestVersionNumber = 0,
    ///         Actions = 
    ///         {
    ///             { "pre_create_hosted_configuration_version", new[]
    ///             {
    ///                 new AwsNative.AppConfig.Inputs.ExtensionActionArgs
    ///                 {
    ///                     Name = "My Test Action",
    ///                     Uri = "DependentLambda.Arn",
    ///                     RoleArn = "DependentRole.Arn",
    ///                     Description = "My test action point",
    ///                 },
    ///             } },
    ///         },
    ///         Parameters = 
    ///         {
    ///             { "myTestParam", new AwsNative.AppConfig.Inputs.ExtensionParameterArgs
    ///             {
    ///                 Required = false,
    ///                 Description = "My test parameter",
    ///             } },
    ///         },
    ///         Tags = new[]
    ///         {
    ///             new AwsNative.Inputs.TagArgs
    ///             {
    ///                 Key = "Ext",
    ///                 Value = "Test",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:appconfig:Extension")]
    public partial class Extension : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The actions defined in the extension.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableDictionary<string, ImmutableArray<Outputs.ExtensionAction>>> Actions { get; private set; } = null!;

        /// <summary>
        /// The system-generated Amazon Resource Name (ARN) for the extension.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The system-generated ID of the extension.
        /// </summary>
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        /// <summary>
        /// Description of the extension.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// You can omit this field when you create an extension. When you create a new version, specify the most recent current version number. For example, you create version 3, enter 2 for this field.
        /// </summary>
        [Output("latestVersionNumber")]
        public Output<int?> LatestVersionNumber { get; private set; } = null!;

        /// <summary>
        /// Name of the extension.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parameters accepted by the extension. You specify parameter values when you associate the extension to an AWS AppConfig resource by using the `CreateExtensionAssociation` API action. For AWS Lambda extension actions, these parameters are included in the Lambda request object.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, Outputs.ExtensionParameter>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// An array of key-value tags to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The extension version number.
        /// </summary>
        [Output("versionNumber")]
        public Output<int> VersionNumber { get; private set; } = null!;


        /// <summary>
        /// Create a Extension resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Extension(string name, ExtensionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appconfig:Extension", name, args ?? new ExtensionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Extension(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appconfig:Extension", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Extension resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Extension Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Extension(name, id, options);
        }
    }

    public sealed class ExtensionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputMap<ImmutableArray<Inputs.ExtensionActionArgs>>? _actions;

        /// <summary>
        /// The actions defined in the extension.
        /// </summary>
        public InputMap<ImmutableArray<Inputs.ExtensionActionArgs>> Actions
        {
            get => _actions ?? (_actions = new InputMap<ImmutableArray<Inputs.ExtensionActionArgs>>());
            set => _actions = value;
        }

        /// <summary>
        /// Description of the extension.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// You can omit this field when you create an extension. When you create a new version, specify the most recent current version number. For example, you create version 3, enter 2 for this field.
        /// </summary>
        [Input("latestVersionNumber")]
        public Input<int>? LatestVersionNumber { get; set; }

        /// <summary>
        /// Name of the extension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<Inputs.ExtensionParameterArgs>? _parameters;

        /// <summary>
        /// The parameters accepted by the extension. You specify parameter values when you associate the extension to an AWS AppConfig resource by using the `CreateExtensionAssociation` API action. For AWS Lambda extension actions, these parameters are included in the Lambda request object.
        /// </summary>
        public InputMap<Inputs.ExtensionParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<Inputs.ExtensionParameterArgs>());
            set => _parameters = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value tags to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ExtensionArgs()
        {
        }
        public static new ExtensionArgs Empty => new ExtensionArgs();
    }
}
