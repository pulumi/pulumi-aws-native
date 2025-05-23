// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    /// <summary>
    /// Output Csv options
    /// </summary>
    [OutputType]
    public sealed class JobCsvOutputOptions
    {
        /// <summary>
        /// A single character that specifies the delimiter used to create CSV job output.
        /// </summary>
        public readonly string? Delimiter;

        [OutputConstructor]
        private JobCsvOutputOptions(string? delimiter)
        {
            Delimiter = delimiter;
        }
    }
}
