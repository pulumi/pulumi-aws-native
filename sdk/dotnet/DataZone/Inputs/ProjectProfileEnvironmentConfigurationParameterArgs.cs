// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    public sealed class ProjectProfileEnvironmentConfigurationParameterArgs : global::Pulumi.ResourceArgs
    {
        [Input("isEditable")]
        public Input<bool>? IsEditable { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ProjectProfileEnvironmentConfigurationParameterArgs()
        {
        }
        public static new ProjectProfileEnvironmentConfigurationParameterArgs Empty => new ProjectProfileEnvironmentConfigurationParameterArgs();
    }
}
