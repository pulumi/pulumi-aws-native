// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution.Outputs
{

    [OutputType]
    public sealed class MatchingWorkflowOutputAttribute
    {
        /// <summary>
        /// Enables the ability to hash the column values in the output.
        /// </summary>
        public readonly bool? Hashed;
        /// <summary>
        /// A name of a column to be written to the output. This must be an `InputField` name in the schema mapping.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private MatchingWorkflowOutputAttribute(
            bool? hashed,

            string name)
        {
            Hashed = hashed;
            Name = name;
        }
    }
}
