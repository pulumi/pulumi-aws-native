// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppRunner.Outputs
{

    [OutputType]
    public sealed class ServiceKeyValuePair
    {
        /// <summary>
        /// The key name string to map to a value.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The value string to which the key name is mapped.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ServiceKeyValuePair(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
