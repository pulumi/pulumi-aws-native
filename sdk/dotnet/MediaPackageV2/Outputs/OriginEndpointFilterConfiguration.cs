// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackageV2.Outputs
{

    /// <summary>
    /// &lt;p&gt;Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest. &lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class OriginEndpointFilterConfiguration
    {
        /// <summary>
        /// &lt;p&gt;Optionally specify the end time for all of your manifest egress requests. When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.&lt;/p&gt;
        /// </summary>
        public readonly string? End;
        /// <summary>
        /// &lt;p&gt;Optionally specify one or more manifest filters for all of your manifest egress requests. When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.&lt;/p&gt;
        /// </summary>
        public readonly string? ManifestFilter;
        /// <summary>
        /// &lt;p&gt;Optionally specify the start time for all of your manifest egress requests. When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.&lt;/p&gt;
        /// </summary>
        public readonly string? Start;
        /// <summary>
        /// &lt;p&gt;Optionally specify the time delay for all of your manifest egress requests. Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.&lt;/p&gt;
        /// </summary>
        public readonly int? TimeDelaySeconds;

        [OutputConstructor]
        private OriginEndpointFilterConfiguration(
            string? end,

            string? manifestFilter,

            string? start,

            int? timeDelaySeconds)
        {
            End = end;
            ManifestFilter = manifestFilter;
            Start = start;
            TimeDelaySeconds = timeDelaySeconds;
        }
    }
}