// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync.Inputs
{

    /// <summary>
    /// Configures a manifest, which is a list of files or objects that you want DataSync to transfer.
    /// </summary>
    public sealed class TaskManifestConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies what DataSync uses the manifest for.
        /// </summary>
        [Input("action")]
        public Input<Pulumi.AwsNative.DataSync.TaskManifestConfigAction>? Action { get; set; }

        /// <summary>
        /// Specifies the file format of your manifest.
        /// </summary>
        [Input("format")]
        public Input<Pulumi.AwsNative.DataSync.TaskManifestConfigFormat>? Format { get; set; }

        /// <summary>
        /// Specifies the manifest that you want DataSync to use and where it's hosted.
        /// </summary>
        [Input("source", required: true)]
        public Input<Inputs.TaskManifestConfigSourcePropertiesArgs> Source { get; set; } = null!;

        public TaskManifestConfigArgs()
        {
        }
        public static new TaskManifestConfigArgs Empty => new TaskManifestConfigArgs();
    }
}