// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs.Outputs
{

    [OutputType]
    public sealed class TransformerParseCloudfront
    {
        /// <summary>
        /// Omit this parameter and the whole log message will be processed by this processor. No other value than `@message` is allowed for `source` .
        /// </summary>
        public readonly string? Source;

        [OutputConstructor]
        private TransformerParseCloudfront(string? source)
        {
            Source = source;
        }
    }
}
