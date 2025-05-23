// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppIntegrations.Outputs
{

    /// <summary>
    /// Application source config
    /// </summary>
    [OutputType]
    public sealed class ApplicationSourceConfigProperties
    {
        /// <summary>
        /// The external URL source for the application.
        /// </summary>
        public readonly Outputs.ApplicationExternalUrlConfig ExternalUrlConfig;

        [OutputConstructor]
        private ApplicationSourceConfigProperties(Outputs.ApplicationExternalUrlConfig externalUrlConfig)
        {
            ExternalUrlConfig = externalUrlConfig;
        }
    }
}
