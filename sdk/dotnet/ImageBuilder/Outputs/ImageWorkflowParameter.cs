// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Outputs
{

    /// <summary>
    /// A parameter associated with the workflow
    /// </summary>
    [OutputType]
    public sealed class ImageWorkflowParameter
    {
        /// <summary>
        /// The name of the workflow parameter to set.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Sets the value for the named workflow parameter.
        /// </summary>
        public readonly ImmutableArray<string> Value;

        [OutputConstructor]
        private ImageWorkflowParameter(
            string? name,

            ImmutableArray<string> value)
        {
            Name = name;
            Value = value;
        }
    }
}
