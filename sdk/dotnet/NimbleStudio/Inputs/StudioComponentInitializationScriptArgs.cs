// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    public sealed class StudioComponentInitializationScriptArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The version number of the protocol that is used by the launch profile. The only valid version is "2021-03-31".
        /// </summary>
        [Input("launchProfileProtocolVersion")]
        public Input<string>? LaunchProfileProtocolVersion { get; set; }

        /// <summary>
        /// The platform of the initialization script, either Windows or Linux.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The method to use when running the initialization script.
        /// </summary>
        [Input("runContext")]
        public Input<string>? RunContext { get; set; }

        /// <summary>
        /// The initialization script.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        public StudioComponentInitializationScriptArgs()
        {
        }
        public static new StudioComponentInitializationScriptArgs Empty => new StudioComponentInitializationScriptArgs();
    }
}
