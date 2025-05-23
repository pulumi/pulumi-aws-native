// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Outputs
{

    /// <summary>
    /// The parameter details of an environment profile.
    /// </summary>
    [OutputType]
    public sealed class EnvironmentProfileEnvironmentParameter
    {
        /// <summary>
        /// The name of an environment profile parameter.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The value of an environment profile parameter.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private EnvironmentProfileEnvironmentParameter(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
