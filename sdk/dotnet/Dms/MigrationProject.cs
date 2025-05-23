// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms
{
    /// <summary>
    /// Resource schema for AWS::DMS::MigrationProject
    /// </summary>
    [AwsNativeResourceType("aws-native:dms:MigrationProject")]
    public partial class MigrationProject : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The optional description of the migration project.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The property describes an instance profile arn for the migration project. For read
        /// </summary>
        [Output("instanceProfileArn")]
        public Output<string?> InstanceProfileArn { get; private set; } = null!;

        /// <summary>
        /// The property describes an instance profile identifier for the migration project. For create
        /// </summary>
        [Output("instanceProfileIdentifier")]
        public Output<string?> InstanceProfileIdentifier { get; private set; } = null!;

        /// <summary>
        /// The property describes an instance profile name for the migration project. For read
        /// </summary>
        [Output("instanceProfileName")]
        public Output<string?> InstanceProfileName { get; private set; } = null!;

        /// <summary>
        /// The property describes an ARN of the migration project.
        /// </summary>
        [Output("migrationProjectArn")]
        public Output<string> MigrationProjectArn { get; private set; } = null!;

        /// <summary>
        /// The property describes a creating time of the migration project.
        /// </summary>
        [Output("migrationProjectCreationTime")]
        public Output<string?> MigrationProjectCreationTime { get; private set; } = null!;

        /// <summary>
        /// The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn
        /// </summary>
        [Output("migrationProjectIdentifier")]
        public Output<string?> MigrationProjectIdentifier { get; private set; } = null!;

        /// <summary>
        /// The property describes a name to identify the migration project.
        /// </summary>
        [Output("migrationProjectName")]
        public Output<string?> MigrationProjectName { get; private set; } = null!;

        /// <summary>
        /// The property describes schema conversion application attributes for the migration project.
        /// </summary>
        [Output("schemaConversionApplicationAttributes")]
        public Output<Outputs.SchemaConversionApplicationAttributesProperties?> SchemaConversionApplicationAttributes { get; private set; } = null!;

        /// <summary>
        /// The property describes source data provider descriptors for the migration project.
        /// </summary>
        [Output("sourceDataProviderDescriptors")]
        public Output<ImmutableArray<Outputs.MigrationProjectDataProviderDescriptor>> SourceDataProviderDescriptors { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The property describes target data provider descriptors for the migration project.
        /// </summary>
        [Output("targetDataProviderDescriptors")]
        public Output<ImmutableArray<Outputs.MigrationProjectDataProviderDescriptor>> TargetDataProviderDescriptors { get; private set; } = null!;

        /// <summary>
        /// The property describes transformation rules for the migration project.
        /// </summary>
        [Output("transformationRules")]
        public Output<string?> TransformationRules { get; private set; } = null!;


        /// <summary>
        /// Create a MigrationProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MigrationProject(string name, MigrationProjectArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:dms:MigrationProject", name, args ?? new MigrationProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MigrationProject(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:dms:MigrationProject", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MigrationProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MigrationProject Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MigrationProject(name, id, options);
        }
    }

    public sealed class MigrationProjectArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The optional description of the migration project.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The property describes an instance profile arn for the migration project. For read
        /// </summary>
        [Input("instanceProfileArn")]
        public Input<string>? InstanceProfileArn { get; set; }

        /// <summary>
        /// The property describes an instance profile identifier for the migration project. For create
        /// </summary>
        [Input("instanceProfileIdentifier")]
        public Input<string>? InstanceProfileIdentifier { get; set; }

        /// <summary>
        /// The property describes an instance profile name for the migration project. For read
        /// </summary>
        [Input("instanceProfileName")]
        public Input<string>? InstanceProfileName { get; set; }

        /// <summary>
        /// The property describes a creating time of the migration project.
        /// </summary>
        [Input("migrationProjectCreationTime")]
        public Input<string>? MigrationProjectCreationTime { get; set; }

        /// <summary>
        /// The property describes an identifier for the migration project. It is used for describing/deleting/modifying can be name/arn
        /// </summary>
        [Input("migrationProjectIdentifier")]
        public Input<string>? MigrationProjectIdentifier { get; set; }

        /// <summary>
        /// The property describes a name to identify the migration project.
        /// </summary>
        [Input("migrationProjectName")]
        public Input<string>? MigrationProjectName { get; set; }

        /// <summary>
        /// The property describes schema conversion application attributes for the migration project.
        /// </summary>
        [Input("schemaConversionApplicationAttributes")]
        public Input<Inputs.SchemaConversionApplicationAttributesPropertiesArgs>? SchemaConversionApplicationAttributes { get; set; }

        [Input("sourceDataProviderDescriptors")]
        private InputList<Inputs.MigrationProjectDataProviderDescriptorArgs>? _sourceDataProviderDescriptors;

        /// <summary>
        /// The property describes source data provider descriptors for the migration project.
        /// </summary>
        public InputList<Inputs.MigrationProjectDataProviderDescriptorArgs> SourceDataProviderDescriptors
        {
            get => _sourceDataProviderDescriptors ?? (_sourceDataProviderDescriptors = new InputList<Inputs.MigrationProjectDataProviderDescriptorArgs>());
            set => _sourceDataProviderDescriptors = value;
        }

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

        [Input("targetDataProviderDescriptors")]
        private InputList<Inputs.MigrationProjectDataProviderDescriptorArgs>? _targetDataProviderDescriptors;

        /// <summary>
        /// The property describes target data provider descriptors for the migration project.
        /// </summary>
        public InputList<Inputs.MigrationProjectDataProviderDescriptorArgs> TargetDataProviderDescriptors
        {
            get => _targetDataProviderDescriptors ?? (_targetDataProviderDescriptors = new InputList<Inputs.MigrationProjectDataProviderDescriptorArgs>());
            set => _targetDataProviderDescriptors = value;
        }

        /// <summary>
        /// The property describes transformation rules for the migration project.
        /// </summary>
        [Input("transformationRules")]
        public Input<string>? TransformationRules { get; set; }

        public MigrationProjectArgs()
        {
        }
        public static new MigrationProjectArgs Empty => new MigrationProjectArgs();
    }
}
