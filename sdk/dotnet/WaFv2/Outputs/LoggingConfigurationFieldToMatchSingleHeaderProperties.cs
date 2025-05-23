// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Inspect a single header. Provide the name of the header to inspect, for example, User-Agent or Referer. This setting isn't case sensitive.
    /// </summary>
    [OutputType]
    public sealed class LoggingConfigurationFieldToMatchSingleHeaderProperties
    {
        /// <summary>
        /// The name of the query header to inspect.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private LoggingConfigurationFieldToMatchSingleHeaderProperties(string name)
        {
            Name = name;
        }
    }
}
