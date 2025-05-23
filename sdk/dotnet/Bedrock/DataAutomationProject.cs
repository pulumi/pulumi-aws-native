// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock
{
    /// <summary>
    /// Definition of AWS::Bedrock::DataAutomationProject Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:bedrock:DataAutomationProject")]
    public partial class DataAutomationProject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time Stamp
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Blueprints to apply to objects processed by the project.
        /// </summary>
        [Output("customOutputConfiguration")]
        public Output<Outputs.DataAutomationProjectCustomOutputConfiguration?> CustomOutputConfiguration { get; private set; } = null!;

        /// <summary>
        /// KMS encryption context
        /// </summary>
        [Output("kmsEncryptionContext")]
        public Output<ImmutableDictionary<string, string>?> KmsEncryptionContext { get; private set; } = null!;

        /// <summary>
        /// KMS key identifier
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Time Stamp
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// Additional settings for the project.
        /// </summary>
        [Output("overrideConfiguration")]
        public Output<Outputs.DataAutomationProjectOverrideConfiguration?> OverrideConfiguration { get; private set; } = null!;

        /// <summary>
        /// ARN of a DataAutomationProject
        /// </summary>
        [Output("projectArn")]
        public Output<string> ProjectArn { get; private set; } = null!;

        /// <summary>
        /// Description of the DataAutomationProject
        /// </summary>
        [Output("projectDescription")]
        public Output<string?> ProjectDescription { get; private set; } = null!;

        /// <summary>
        /// Name of the DataAutomationProject
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The project's stage.
        /// </summary>
        [Output("projectStage")]
        public Output<Pulumi.AwsNative.Bedrock.DataAutomationProjectStage> ProjectStage { get; private set; } = null!;

        /// <summary>
        /// The project's standard output configuration.
        /// </summary>
        [Output("standardOutputConfiguration")]
        public Output<Outputs.DataAutomationProjectStandardOutputConfiguration?> StandardOutputConfiguration { get; private set; } = null!;

        /// <summary>
        /// The project's status.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.Bedrock.DataAutomationProjectStatus> Status { get; private set; } = null!;

        /// <summary>
        /// List of Tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DataAutomationProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataAutomationProject(string name, DataAutomationProjectArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:bedrock:DataAutomationProject", name, args ?? new DataAutomationProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataAutomationProject(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:bedrock:DataAutomationProject", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "projectName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DataAutomationProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataAutomationProject Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DataAutomationProject(name, id, options);
        }
    }

    public sealed class DataAutomationProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Blueprints to apply to objects processed by the project.
        /// </summary>
        [Input("customOutputConfiguration")]
        public Input<Inputs.DataAutomationProjectCustomOutputConfigurationArgs>? CustomOutputConfiguration { get; set; }

        [Input("kmsEncryptionContext")]
        private InputMap<string>? _kmsEncryptionContext;

        /// <summary>
        /// KMS encryption context
        /// </summary>
        public InputMap<string> KmsEncryptionContext
        {
            get => _kmsEncryptionContext ?? (_kmsEncryptionContext = new InputMap<string>());
            set => _kmsEncryptionContext = value;
        }

        /// <summary>
        /// KMS key identifier
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Additional settings for the project.
        /// </summary>
        [Input("overrideConfiguration")]
        public Input<Inputs.DataAutomationProjectOverrideConfigurationArgs>? OverrideConfiguration { get; set; }

        /// <summary>
        /// Description of the DataAutomationProject
        /// </summary>
        [Input("projectDescription")]
        public Input<string>? ProjectDescription { get; set; }

        /// <summary>
        /// Name of the DataAutomationProject
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The project's standard output configuration.
        /// </summary>
        [Input("standardOutputConfiguration")]
        public Input<Inputs.DataAutomationProjectStandardOutputConfigurationArgs>? StandardOutputConfiguration { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// List of Tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public DataAutomationProjectArgs()
        {
        }
        public static new DataAutomationProjectArgs Empty => new DataAutomationProjectArgs();
    }
}
