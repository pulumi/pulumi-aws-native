// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// The CodeEditor app settings.
    /// </summary>
    public sealed class SpaceCodeEditorAppSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings that are used to configure and manage the lifecycle of CodeEditor applications in a space.
        /// </summary>
        [Input("appLifecycleManagement")]
        public Input<Inputs.SpaceAppLifecycleManagementArgs>? AppLifecycleManagement { get; set; }

        /// <summary>
        /// Specifies the ARNs of a SageMaker image and SageMaker image version, and the instance type that the version runs on.
        /// </summary>
        [Input("defaultResourceSpec")]
        public Input<Inputs.SpaceResourceSpecArgs>? DefaultResourceSpec { get; set; }

        public SpaceCodeEditorAppSettingsArgs()
        {
        }
        public static new SpaceCodeEditorAppSettingsArgs Empty => new SpaceCodeEditorAppSettingsArgs();
    }
}
