// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Outputs
{

    /// <summary>
    /// An environment variable to set inside a container, in the form of a key-value pair.
    /// </summary>
    [OutputType]
    public sealed class ContainerGroupDefinitionContainerEnvironment
    {
        /// <summary>
        /// The environment variable name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The environment variable value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ContainerGroupDefinitionContainerEnvironment(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}