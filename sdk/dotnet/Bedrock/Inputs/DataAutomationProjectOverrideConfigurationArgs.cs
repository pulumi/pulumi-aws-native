// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Override configuration
    /// </summary>
    public sealed class DataAutomationProjectOverrideConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This element declares whether your project will process audio files.
        /// </summary>
        [Input("audio")]
        public Input<Inputs.DataAutomationProjectAudioOverrideConfigurationArgs>? Audio { get; set; }

        /// <summary>
        /// Additional settings for a project.
        /// </summary>
        [Input("document")]
        public Input<Inputs.DataAutomationProjectDocumentOverrideConfigurationArgs>? Document { get; set; }

        /// <summary>
        /// This element declares whether your project will process image files.
        /// </summary>
        [Input("image")]
        public Input<Inputs.DataAutomationProjectImageOverrideConfigurationArgs>? Image { get; set; }

        /// <summary>
        /// Lets you set which modalities certain file types are processed as.
        /// </summary>
        [Input("modalityRouting")]
        public Input<Inputs.DataAutomationProjectModalityRoutingConfigurationArgs>? ModalityRouting { get; set; }

        /// <summary>
        /// This element declares whether your project will process video files.
        /// </summary>
        [Input("video")]
        public Input<Inputs.DataAutomationProjectVideoOverrideConfigurationArgs>? Video { get; set; }

        public DataAutomationProjectOverrideConfigurationArgs()
        {
        }
        public static new DataAutomationProjectOverrideConfigurationArgs Empty => new DataAutomationProjectOverrideConfigurationArgs();
    }
}
