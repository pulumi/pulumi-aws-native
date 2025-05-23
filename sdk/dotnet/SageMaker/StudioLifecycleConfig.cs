// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    /// <summary>
    /// Resource Type definition for AWS::SageMaker::StudioLifecycleConfig
    /// </summary>
    [AwsNativeResourceType("aws-native:sagemaker:StudioLifecycleConfig")]
    public partial class StudioLifecycleConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The App type that the Lifecycle Configuration is attached to.
        /// </summary>
        [Output("studioLifecycleConfigAppType")]
        public Output<Pulumi.AwsNative.SageMaker.StudioLifecycleConfigAppType> StudioLifecycleConfigAppType { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lifecycle Configuration.
        /// </summary>
        [Output("studioLifecycleConfigArn")]
        public Output<string> StudioLifecycleConfigArn { get; private set; } = null!;

        /// <summary>
        /// The content of your Amazon SageMaker Studio Lifecycle Configuration script. This content must be base64 encoded.
        /// </summary>
        [Output("studioLifecycleConfigContent")]
        public Output<string> StudioLifecycleConfigContent { get; private set; } = null!;

        /// <summary>
        /// The name of the Amazon SageMaker Studio Lifecycle Configuration.
        /// </summary>
        [Output("studioLifecycleConfigName")]
        public Output<string> StudioLifecycleConfigName { get; private set; } = null!;

        /// <summary>
        /// Tags to be associated with the Lifecycle Configuration. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.CreateOnlyTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a StudioLifecycleConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StudioLifecycleConfig(string name, StudioLifecycleConfigArgs args, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:StudioLifecycleConfig", name, args ?? new StudioLifecycleConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StudioLifecycleConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:sagemaker:StudioLifecycleConfig", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "studioLifecycleConfigAppType",
                    "studioLifecycleConfigContent",
                    "studioLifecycleConfigName",
                    "tags[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StudioLifecycleConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StudioLifecycleConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StudioLifecycleConfig(name, id, options);
        }
    }

    public sealed class StudioLifecycleConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The App type that the Lifecycle Configuration is attached to.
        /// </summary>
        [Input("studioLifecycleConfigAppType", required: true)]
        public Input<Pulumi.AwsNative.SageMaker.StudioLifecycleConfigAppType> StudioLifecycleConfigAppType { get; set; } = null!;

        /// <summary>
        /// The content of your Amazon SageMaker Studio Lifecycle Configuration script. This content must be base64 encoded.
        /// </summary>
        [Input("studioLifecycleConfigContent", required: true)]
        public Input<string> StudioLifecycleConfigContent { get; set; } = null!;

        /// <summary>
        /// The name of the Amazon SageMaker Studio Lifecycle Configuration.
        /// </summary>
        [Input("studioLifecycleConfigName")]
        public Input<string>? StudioLifecycleConfigName { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs>? _tags;

        /// <summary>
        /// Tags to be associated with the Lifecycle Configuration. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.CreateOnlyTagArgs>());
            set => _tags = value;
        }

        public StudioLifecycleConfigArgs()
        {
        }
        public static new StudioLifecycleConfigArgs Empty => new StudioLifecycleConfigArgs();
    }
}
