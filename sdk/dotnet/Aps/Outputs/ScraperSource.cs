// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Aps.Outputs
{

    /// <summary>
    /// Scraper metrics source
    /// </summary>
    [OutputType]
    public sealed class ScraperSource
    {
        /// <summary>
        /// Configuration for EKS metrics source
        /// </summary>
        public readonly Outputs.ScraperSourceEksConfigurationProperties? EksConfiguration;

        [OutputConstructor]
        private ScraperSource(Outputs.ScraperSourceEksConfigurationProperties? eksConfiguration)
        {
            EksConfiguration = eksConfiguration;
        }
    }
}
