// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    /// <summary>
    /// A provisioned concurrency configuration for a function's version.
    /// </summary>
    [OutputType]
    public sealed class VersionProvisionedConcurrencyConfiguration
    {
        /// <summary>
        /// The amount of provisioned concurrency to allocate for the version.
        /// </summary>
        public readonly int ProvisionedConcurrentExecutions;

        [OutputConstructor]
        private VersionProvisionedConcurrencyConfiguration(int provisionedConcurrentExecutions)
        {
            ProvisionedConcurrentExecutions = provisionedConcurrentExecutions;
        }
    }
}
