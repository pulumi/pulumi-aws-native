// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    /// <summary>
    /// Enable a resource that has been published in the CloudFormation Registry.
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudformation:TypeActivation")]
    public partial class TypeActivation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the extension.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.
        /// </summary>
        [Output("autoUpdate")]
        public Output<bool?> AutoUpdate { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
        /// </summary>
        [Output("executionRoleArn")]
        public Output<string?> ExecutionRoleArn { get; private set; } = null!;

        /// <summary>
        /// Specifies logging configuration information for a type.
        /// </summary>
        [Output("loggingConfig")]
        public Output<Outputs.TypeActivationLoggingConfig?> LoggingConfig { get; private set; } = null!;

        /// <summary>
        /// The Major Version of the type you want to enable
        /// </summary>
        [Output("majorVersion")]
        public Output<string?> MajorVersion { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Number (ARN) assigned to the public extension upon publication
        /// </summary>
        [Output("publicTypeArn")]
        public Output<string?> PublicTypeArn { get; private set; } = null!;

        /// <summary>
        /// The publisher id assigned by CloudFormation for publishing in this region.
        /// </summary>
        [Output("publisherId")]
        public Output<string?> PublisherId { get; private set; } = null!;

        /// <summary>
        /// The kind of extension
        /// </summary>
        [Output("type")]
        public Output<Pulumi.AwsNative.CloudFormation.TypeActivationType?> Type { get; private set; } = null!;

        /// <summary>
        /// The name of the type being registered.
        /// 
        /// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
        /// </summary>
        [Output("typeName")]
        public Output<string?> TypeName { get; private set; } = null!;

        /// <summary>
        /// An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.
        /// </summary>
        [Output("typeNameAlias")]
        public Output<string?> TypeNameAlias { get; private set; } = null!;

        /// <summary>
        /// Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled
        /// </summary>
        [Output("versionBump")]
        public Output<Pulumi.AwsNative.CloudFormation.TypeActivationVersionBump?> VersionBump { get; private set; } = null!;


        /// <summary>
        /// Create a TypeActivation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TypeActivation(string name, TypeActivationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:TypeActivation", name, args ?? new TypeActivationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TypeActivation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudformation:TypeActivation", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing TypeActivation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TypeActivation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TypeActivation(name, id, options);
        }
    }

    public sealed class TypeActivationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated.
        /// </summary>
        [Input("autoUpdate")]
        public Input<bool>? AutoUpdate { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials.
        /// </summary>
        [Input("executionRoleArn")]
        public Input<string>? ExecutionRoleArn { get; set; }

        /// <summary>
        /// Specifies logging configuration information for a type.
        /// </summary>
        [Input("loggingConfig")]
        public Input<Inputs.TypeActivationLoggingConfigArgs>? LoggingConfig { get; set; }

        /// <summary>
        /// The Major Version of the type you want to enable
        /// </summary>
        [Input("majorVersion")]
        public Input<string>? MajorVersion { get; set; }

        /// <summary>
        /// The Amazon Resource Number (ARN) assigned to the public extension upon publication
        /// </summary>
        [Input("publicTypeArn")]
        public Input<string>? PublicTypeArn { get; set; }

        /// <summary>
        /// The publisher id assigned by CloudFormation for publishing in this region.
        /// </summary>
        [Input("publisherId")]
        public Input<string>? PublisherId { get; set; }

        /// <summary>
        /// The kind of extension
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AwsNative.CloudFormation.TypeActivationType>? Type { get; set; }

        /// <summary>
        /// The name of the type being registered.
        /// 
        /// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
        /// </summary>
        [Input("typeName")]
        public Input<string>? TypeName { get; set; }

        /// <summary>
        /// An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates.
        /// </summary>
        [Input("typeNameAlias")]
        public Input<string>? TypeNameAlias { get; set; }

        /// <summary>
        /// Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled
        /// </summary>
        [Input("versionBump")]
        public Input<Pulumi.AwsNative.CloudFormation.TypeActivationVersionBump>? VersionBump { get; set; }

        public TypeActivationArgs()
        {
        }
        public static new TypeActivationArgs Empty => new TypeActivationArgs();
    }
}