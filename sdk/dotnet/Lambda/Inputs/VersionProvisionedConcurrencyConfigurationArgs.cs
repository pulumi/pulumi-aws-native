// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    /// <summary>
    /// A provisioned concurrency configuration for a function's version.
    /// </summary>
    public sealed class VersionProvisionedConcurrencyConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of provisioned concurrency to allocate for the version.
        /// </summary>
        [Input("provisionedConcurrentExecutions", required: true)]
        public Input<int> ProvisionedConcurrentExecutions { get; set; } = null!;

        public VersionProvisionedConcurrencyConfigurationArgs()
        {
        }
        public static new VersionProvisionedConcurrencyConfigurationArgs Empty => new VersionProvisionedConcurrencyConfigurationArgs();
    }
}
