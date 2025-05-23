// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Inputs
{

    /// <summary>
    /// An environment variable to set inside a container, in the form of a key-value pair.
    /// </summary>
    public sealed class ContainerGroupDefinitionContainerEnvironmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The environment variable name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The environment variable value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ContainerGroupDefinitionContainerEnvironmentArgs()
        {
        }
        public static new ContainerGroupDefinitionContainerEnvironmentArgs Empty => new ContainerGroupDefinitionContainerEnvironmentArgs();
    }
}
