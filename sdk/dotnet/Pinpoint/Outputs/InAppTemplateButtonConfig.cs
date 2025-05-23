// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class InAppTemplateButtonConfig
    {
        /// <summary>
        /// Optional button configuration to use for in-app messages sent to Android devices. This button configuration overrides the default button configuration.
        /// </summary>
        public readonly Outputs.InAppTemplateOverrideButtonConfiguration? Android;
        /// <summary>
        /// Specifies the default behavior of a button that appears in an in-app message. You can optionally add button configurations that specifically apply to iOS, Android, or web browser users.
        /// </summary>
        public readonly Outputs.InAppTemplateDefaultButtonConfiguration? DefaultConfig;
        /// <summary>
        /// Optional button configuration to use for in-app messages sent to iOS devices. This button configuration overrides the default button configuration.
        /// </summary>
        public readonly Outputs.InAppTemplateOverrideButtonConfiguration? Ios;
        /// <summary>
        /// Optional button configuration to use for in-app messages sent to web applications. This button configuration overrides the default button configuration.
        /// </summary>
        public readonly Outputs.InAppTemplateOverrideButtonConfiguration? Web;

        [OutputConstructor]
        private InAppTemplateButtonConfig(
            Outputs.InAppTemplateOverrideButtonConfiguration? android,

            Outputs.InAppTemplateDefaultButtonConfiguration? defaultConfig,

            Outputs.InAppTemplateOverrideButtonConfiguration? ios,

            Outputs.InAppTemplateOverrideButtonConfiguration? web)
        {
            Android = android;
            DefaultConfig = defaultConfig;
            Ios = ios;
            Web = web;
        }
    }
}
