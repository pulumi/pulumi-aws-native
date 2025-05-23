// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SystemsManagerSap
{
    /// <summary>
    /// Resource schema for AWS::SystemsManagerSAP::Application
    /// </summary>
    [AwsNativeResourceType("aws-native:systemsmanagersap:Application")]
    public partial class Application : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the application.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// The type of the application.
        /// </summary>
        [Output("applicationType")]
        public Output<Pulumi.AwsNative.SystemsManagerSap.ApplicationType> ApplicationType { get; private set; } = null!;

        /// <summary>
        /// The ARN of the SSM-SAP application
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// This is an optional parameter for component details to which the SAP ABAP application is attached, such as Web Dispatcher.
        /// </summary>
        [Output("componentsInfo")]
        public Output<ImmutableArray<Outputs.ApplicationComponentInfo>> ComponentsInfo { get; private set; } = null!;

        /// <summary>
        /// The credentials of the SAP application.
        /// </summary>
        [Output("credentials")]
        public Output<ImmutableArray<Outputs.ApplicationCredential>> Credentials { get; private set; } = null!;

        /// <summary>
        /// The ARN of the SAP HANA database
        /// </summary>
        [Output("databaseArn")]
        public Output<string?> DatabaseArn { get; private set; } = null!;

        /// <summary>
        /// The Amazon EC2 instances on which your SAP application is running.
        /// </summary>
        [Output("instances")]
        public Output<ImmutableArray<string>> Instances { get; private set; } = null!;

        /// <summary>
        /// The SAP instance number of the application.
        /// </summary>
        [Output("sapInstanceNumber")]
        public Output<string?> SapInstanceNumber { get; private set; } = null!;

        /// <summary>
        /// The System ID of the application.
        /// </summary>
        [Output("sid")]
        public Output<string?> Sid { get; private set; } = null!;

        /// <summary>
        /// The tags of a SystemsManagerSAP application.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:systemsmanagersap:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:systemsmanagersap:Application", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "componentsInfo[*]",
                    "credentials[*]",
                    "databaseArn",
                    "instances[*]",
                    "sapInstanceNumber",
                    "sid",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Application(name, id, options);
        }
    }

    public sealed class ApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the application.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// The type of the application.
        /// </summary>
        [Input("applicationType", required: true)]
        public Input<Pulumi.AwsNative.SystemsManagerSap.ApplicationType> ApplicationType { get; set; } = null!;

        [Input("componentsInfo")]
        private InputList<Inputs.ApplicationComponentInfoArgs>? _componentsInfo;

        /// <summary>
        /// This is an optional parameter for component details to which the SAP ABAP application is attached, such as Web Dispatcher.
        /// </summary>
        public InputList<Inputs.ApplicationComponentInfoArgs> ComponentsInfo
        {
            get => _componentsInfo ?? (_componentsInfo = new InputList<Inputs.ApplicationComponentInfoArgs>());
            set => _componentsInfo = value;
        }

        [Input("credentials")]
        private InputList<Inputs.ApplicationCredentialArgs>? _credentials;

        /// <summary>
        /// The credentials of the SAP application.
        /// </summary>
        public InputList<Inputs.ApplicationCredentialArgs> Credentials
        {
            get => _credentials ?? (_credentials = new InputList<Inputs.ApplicationCredentialArgs>());
            set => _credentials = value;
        }

        /// <summary>
        /// The ARN of the SAP HANA database
        /// </summary>
        [Input("databaseArn")]
        public Input<string>? DatabaseArn { get; set; }

        [Input("instances")]
        private InputList<string>? _instances;

        /// <summary>
        /// The Amazon EC2 instances on which your SAP application is running.
        /// </summary>
        public InputList<string> Instances
        {
            get => _instances ?? (_instances = new InputList<string>());
            set => _instances = value;
        }

        /// <summary>
        /// The SAP instance number of the application.
        /// </summary>
        [Input("sapInstanceNumber")]
        public Input<string>? SapInstanceNumber { get; set; }

        /// <summary>
        /// The System ID of the application.
        /// </summary>
        [Input("sid")]
        public Input<string>? Sid { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// The tags of a SystemsManagerSAP application.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public ApplicationArgs()
        {
        }
        public static new ApplicationArgs Empty => new ApplicationArgs();
    }
}
