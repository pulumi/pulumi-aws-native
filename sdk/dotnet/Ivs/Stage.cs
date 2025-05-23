// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ivs
{
    /// <summary>
    /// Resource Definition for type AWS::IVS::Stage.
    /// </summary>
    [AwsNativeResourceType("aws-native:ivs:Stage")]
    public partial class Stage : global::Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the active session within the stage.
        /// </summary>
        [Output("activeSessionId")]
        public Output<string> ActiveSessionId { get; private set; } = null!;

        /// <summary>
        /// Stage ARN is automatically generated on creation and assigned as the unique identifier.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("autoParticipantRecordingConfiguration")]
        public Output<Outputs.StageAutoParticipantRecordingConfiguration?> AutoParticipantRecordingConfiguration { get; private set; } = null!;

        /// <summary>
        /// Stage name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Stage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Stage(string name, StageArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ivs:Stage", name, args ?? new StageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Stage(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ivs:Stage", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Stage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Stage Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Stage(name, id, options);
        }
    }

    public sealed class StageArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoParticipantRecordingConfiguration")]
        public Input<Inputs.StageAutoParticipantRecordingConfigurationArgs>? AutoParticipantRecordingConfiguration { get; set; }

        /// <summary>
        /// Stage name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        public StageArgs()
        {
        }
        public static new StageArgs Empty => new StageArgs();
    }
}
