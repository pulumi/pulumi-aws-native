// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    /// <summary>
    /// Publishes new or first hook version to AWS CloudFormation Registry.
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
    ///     var hookVersion = new AwsNative.CloudFormation.HookVersion("hookVersion", new()
    ///     {
    ///         TypeName = "My::Sample::Hook",
    ///         SchemaHandlerPackage = "s3://my-sample-hookversion-bucket/my-sample-hook.zip",
    ///     });
    /// 
    ///     var hookDefaultVersion = new AwsNative.CloudFormation.HookDefaultVersion("hookDefaultVersion", new()
    ///     {
    ///         TypeVersionArn = hookVersion.Id,
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
    ///     var hookVersion = new AwsNative.CloudFormation.HookVersion("hookVersion", new()
    ///     {
    ///         TypeName = "My::Sample::Hook",
    ///         SchemaHandlerPackage = "s3://my-sample-hookversion-bucket/my-sample-hook.zip",
    ///     });
    /// 
    ///     var hookDefaultVersion = new AwsNative.CloudFormation.HookDefaultVersion("hookDefaultVersion", new()
    ///     {
    ///         TypeVersionArn = hookVersion.Id,
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
    ///     var hookVersion = new AwsNative.CloudFormation.HookVersion("hookVersion", new()
    ///     {
    ///         TypeName = "My::Sample::Hook",
    ///         SchemaHandlerPackage = "s3://my-sample-hookversion-bucket/my-sample-hook.zip",
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
    ///     var hookVersion = new AwsNative.CloudFormation.HookVersion("hookVersion", new()
    ///     {
    ///         TypeName = "My::Sample::Hook",
    ///         SchemaHandlerPackage = "s3://my-sample-hookversion-bucket/my-sample-hook.zip",
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
    ///     var hookVersion = new AwsNative.CloudFormation.HookVersion("hookVersion", new()
    ///     {
    ///         TypeName = "My::Sample::Hook",
    ///         SchemaHandlerPackage = "s3://my-sample-hookversion-bucket/my-sample-hook.zip",
    ///     });
    /// 
    ///     var hookDefaultVersion = new AwsNative.CloudFormation.HookDefaultVersion("hookDefaultVersion", new()
    ///     {
    ///         TypeVersionArn = hookVersion.Id,
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
    ///     var hookVersion = new AwsNative.CloudFormation.HookVersion("hookVersion", new()
    ///     {
    ///         TypeName = "My::Sample::Hook",
    ///         SchemaHandlerPackage = "s3://my-sample-hookversion-bucket/my-sample-hook.zip",
    ///     });
    /// 
    ///     var hookDefaultVersion = new AwsNative.CloudFormation.HookDefaultVersion("hookDefaultVersion", new()
    ///     {
    ///         TypeVersionArn = hookVersion.Id,
    ///     });
    /// 
    /// });
    /// 
    /// 
    /// ```
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudformation:HookVersion")]
    public partial class HookVersion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
        /// </summary>
        [Output("executionRoleArn")]
        public Output<string?> ExecutionRoleArn { get; private set; } = null!;

        /// <summary>
        /// Indicates if this type version is the current default version
        /// </summary>
        [Output("isDefaultVersion")]
        public Output<bool> IsDefaultVersion { get; private set; } = null!;

        /// <summary>
        /// Specifies logging configuration information for a type.
        /// </summary>
        [Output("loggingConfig")]
        public Output<Outputs.HookVersionLoggingConfig?> LoggingConfig { get; private set; } = null!;

        /// <summary>
        /// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
        /// 
        /// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
        /// </summary>
        [Output("schemaHandlerPackage")]
        public Output<string> SchemaHandlerPackage { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the type without the versionID.
        /// </summary>
        [Output("typeArn")]
        public Output<string> TypeArn { get; private set; } = null!;

        /// <summary>
        /// The name of the type being registered.
        /// 
        /// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
        /// </summary>
        [Output("typeName")]
        public Output<string> TypeName { get; private set; } = null!;

        /// <summary>
        /// The ID of the version of the type represented by this hook instance.
        /// </summary>
        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;

        /// <summary>
        /// The scope at which the type is visible and usable in CloudFormation operations.
        /// 
        /// Valid values include:
        /// 
        /// PRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.
        /// 
        /// PUBLIC: The type is publically visible and usable within any Amazon account.
        /// </summary>
        [Output("visibility")]
        public Output<Pulumi.AwsNative.CloudFormation.HookVersionVisibility> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a HookVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HookVersion(string name, HookVersionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:HookVersion", name, args ?? new HookVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HookVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:HookVersion", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "executionRoleArn",
                    "loggingConfig",
                    "schemaHandlerPackage",
                    "typeName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HookVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HookVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HookVersion(name, id, options);
        }
    }

    public sealed class HookVersionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
        /// </summary>
        [Input("executionRoleArn")]
        public Input<string>? ExecutionRoleArn { get; set; }

        /// <summary>
        /// Specifies logging configuration information for a type.
        /// </summary>
        [Input("loggingConfig")]
        public Input<Inputs.HookVersionLoggingConfigArgs>? LoggingConfig { get; set; }

        /// <summary>
        /// A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register.
        /// 
        /// For information on generating a schema handler package for the type you want to register, see submit in the CloudFormation CLI User Guide.
        /// </summary>
        [Input("schemaHandlerPackage", required: true)]
        public Input<string> SchemaHandlerPackage { get; set; } = null!;

        /// <summary>
        /// The name of the type being registered.
        /// 
        /// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
        /// </summary>
        [Input("typeName", required: true)]
        public Input<string> TypeName { get; set; } = null!;

        public HookVersionArgs()
        {
        }
        public static new HookVersionArgs Empty => new HookVersionArgs();
    }
}
