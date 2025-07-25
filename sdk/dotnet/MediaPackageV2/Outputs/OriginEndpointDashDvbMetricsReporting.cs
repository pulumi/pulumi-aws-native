// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackageV2.Outputs
{

    /// <summary>
    /// &lt;p&gt;For use with DVB-DASH profiles only. The settings for error reporting from the playback device that you want Elemental MediaPackage to pass through to the manifest.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class OriginEndpointDashDvbMetricsReporting
    {
        /// <summary>
        /// &lt;p&gt;The number of playback devices per 1000 that will send error reports to the reporting URL. This represents the probability that a playback device will be a reporting player for this session.&lt;/p&gt;
        /// </summary>
        public readonly int? Probability;
        /// <summary>
        /// &lt;p&gt;The URL where playback devices send error reports.&lt;/p&gt;
        /// </summary>
        public readonly string ReportingUrl;

        [OutputConstructor]
        private OriginEndpointDashDvbMetricsReporting(
            int? probability,

            string reportingUrl)
        {
            Probability = probability;
            ReportingUrl = reportingUrl;
        }
    }
}
