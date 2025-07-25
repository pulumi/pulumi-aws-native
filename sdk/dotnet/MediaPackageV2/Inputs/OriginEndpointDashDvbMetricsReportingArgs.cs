// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackageV2.Inputs
{

    /// <summary>
    /// &lt;p&gt;For use with DVB-DASH profiles only. The settings for error reporting from the playback device that you want Elemental MediaPackage to pass through to the manifest.&lt;/p&gt;
    /// </summary>
    public sealed class OriginEndpointDashDvbMetricsReportingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The number of playback devices per 1000 that will send error reports to the reporting URL. This represents the probability that a playback device will be a reporting player for this session.&lt;/p&gt;
        /// </summary>
        [Input("probability")]
        public Input<int>? Probability { get; set; }

        /// <summary>
        /// &lt;p&gt;The URL where playback devices send error reports.&lt;/p&gt;
        /// </summary>
        [Input("reportingUrl", required: true)]
        public Input<string> ReportingUrl { get; set; } = null!;

        public OriginEndpointDashDvbMetricsReportingArgs()
        {
        }
        public static new OriginEndpointDashDvbMetricsReportingArgs Empty => new OriginEndpointDashDvbMetricsReportingArgs();
    }
}
