// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation.Outputs
{

    [OutputType]
    public sealed class StackOutput
    {
        /// <summary>
        /// User defined description associated with the output.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The name of the export associated with the output.
        /// </summary>
        public readonly string? ExportName;
        /// <summary>
        /// The key associated with the output.
        /// </summary>
        public readonly string? OutputKey;
        /// <summary>
        /// The value associated with the output.
        /// </summary>
        public readonly string? OutputValue;

        [OutputConstructor]
        private StackOutput(
            string? description,

            string? exportName,

            string? outputKey,

            string? outputValue)
        {
            Description = description;
            ExportName = exportName;
            OutputKey = outputKey;
            OutputValue = outputValue;
        }
    }
}
