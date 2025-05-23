// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder
{
    /// <summary>
    /// Resource schema for AWS::ImageBuilder::ImagePipeline
    /// </summary>
    [AwsNativeResourceType("aws-native:imagebuilder:ImagePipeline")]
    public partial class ImagePipeline : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the image pipeline.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
        /// </summary>
        [Output("containerRecipeArn")]
        public Output<string?> ContainerRecipeArn { get; private set; } = null!;

        /// <summary>
        /// The description of the image pipeline.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.
        /// </summary>
        [Output("distributionConfigurationArn")]
        public Output<string?> DistributionConfigurationArn { get; private set; } = null!;

        /// <summary>
        /// Collects additional information about the image being created, including the operating system (OS) version and package list.
        /// </summary>
        [Output("enhancedImageMetadataEnabled")]
        public Output<bool?> EnhancedImageMetadataEnabled { get; private set; } = null!;

        /// <summary>
        /// The execution role name/ARN for the image build, if provided
        /// </summary>
        [Output("executionRole")]
        public Output<string?> ExecutionRole { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
        /// </summary>
        [Output("imageRecipeArn")]
        public Output<string?> ImageRecipeArn { get; private set; } = null!;

        /// <summary>
        /// Contains settings for vulnerability scans.
        /// </summary>
        [Output("imageScanningConfiguration")]
        public Output<Outputs.ImagePipelineImageScanningConfiguration?> ImageScanningConfiguration { get; private set; } = null!;

        /// <summary>
        /// The image tests configuration of the image pipeline.
        /// </summary>
        [Output("imageTestsConfiguration")]
        public Output<Outputs.ImagePipelineImageTestsConfiguration?> ImageTestsConfiguration { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
        /// </summary>
        [Output("infrastructureConfigurationArn")]
        public Output<string?> InfrastructureConfigurationArn { get; private set; } = null!;

        /// <summary>
        /// The name of the image pipeline.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The schedule of the image pipeline.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.ImagePipelineSchedule?> Schedule { get; private set; } = null!;

        /// <summary>
        /// The status of the image pipeline.
        /// </summary>
        [Output("status")]
        public Output<Pulumi.AwsNative.ImageBuilder.ImagePipelineStatus?> Status { get; private set; } = null!;

        /// <summary>
        /// The tags of this image pipeline.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Workflows to define the image build process
        /// </summary>
        [Output("workflows")]
        public Output<ImmutableArray<Outputs.ImagePipelineWorkflowConfiguration>> Workflows { get; private set; } = null!;


        /// <summary>
        /// Create a ImagePipeline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImagePipeline(string name, ImagePipelineArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:imagebuilder:ImagePipeline", name, args ?? new ImagePipelineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImagePipeline(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:imagebuilder:ImagePipeline", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ImagePipeline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImagePipeline Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ImagePipeline(name, id, options);
        }
    }

    public sealed class ImagePipelineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.
        /// </summary>
        [Input("containerRecipeArn")]
        public Input<string>? ContainerRecipeArn { get; set; }

        /// <summary>
        /// The description of the image pipeline.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the distribution configuration associated with this image pipeline.
        /// </summary>
        [Input("distributionConfigurationArn")]
        public Input<string>? DistributionConfigurationArn { get; set; }

        /// <summary>
        /// Collects additional information about the image being created, including the operating system (OS) version and package list.
        /// </summary>
        [Input("enhancedImageMetadataEnabled")]
        public Input<bool>? EnhancedImageMetadataEnabled { get; set; }

        /// <summary>
        /// The execution role name/ARN for the image build, if provided
        /// </summary>
        [Input("executionRole")]
        public Input<string>? ExecutionRole { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.
        /// </summary>
        [Input("imageRecipeArn")]
        public Input<string>? ImageRecipeArn { get; set; }

        /// <summary>
        /// Contains settings for vulnerability scans.
        /// </summary>
        [Input("imageScanningConfiguration")]
        public Input<Inputs.ImagePipelineImageScanningConfigurationArgs>? ImageScanningConfiguration { get; set; }

        /// <summary>
        /// The image tests configuration of the image pipeline.
        /// </summary>
        [Input("imageTestsConfiguration")]
        public Input<Inputs.ImagePipelineImageTestsConfigurationArgs>? ImageTestsConfiguration { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the infrastructure configuration associated with this image pipeline.
        /// </summary>
        [Input("infrastructureConfigurationArn")]
        public Input<string>? InfrastructureConfigurationArn { get; set; }

        /// <summary>
        /// The name of the image pipeline.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The schedule of the image pipeline.
        /// </summary>
        [Input("schedule")]
        public Input<Inputs.ImagePipelineScheduleArgs>? Schedule { get; set; }

        /// <summary>
        /// The status of the image pipeline.
        /// </summary>
        [Input("status")]
        public Input<Pulumi.AwsNative.ImageBuilder.ImagePipelineStatus>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags of this image pipeline.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("workflows")]
        private InputList<Inputs.ImagePipelineWorkflowConfigurationArgs>? _workflows;

        /// <summary>
        /// Workflows to define the image build process
        /// </summary>
        public InputList<Inputs.ImagePipelineWorkflowConfigurationArgs> Workflows
        {
            get => _workflows ?? (_workflows = new InputList<Inputs.ImagePipelineWorkflowConfigurationArgs>());
            set => _workflows = value;
        }

        public ImagePipelineArgs()
        {
        }
        public static new ImagePipelineArgs Empty => new ImagePipelineArgs();
    }
}
