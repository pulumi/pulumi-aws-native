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
    /// &lt;p&gt;Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest. &lt;/p&gt;
    /// </summary>
    public sealed class OriginEndpointFilterConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;Optionally specify the clip start time for all of your manifest egress requests. When you include clip start time, note that you cannot use clip start time query parameters for this manifest's endpoint URL.&lt;/p&gt;
        /// </summary>
        [Input("clipStartTime")]
        public Input<string>? ClipStartTime { get; set; }

        /// <summary>
        /// &lt;p&gt;Optionally specify the end time for all of your manifest egress requests. When you include end time, note that you cannot use end time query parameters for this manifest's endpoint URL.&lt;/p&gt;
        /// </summary>
        [Input("end")]
        public Input<string>? End { get; set; }

        /// <summary>
        /// &lt;p&gt;Optionally specify one or more manifest filters for all of your manifest egress requests. When you include a manifest filter, note that you cannot use an identical manifest filter query parameter for this manifest's endpoint URL.&lt;/p&gt;
        /// </summary>
        [Input("manifestFilter")]
        public Input<string>? ManifestFilter { get; set; }

        /// <summary>
        /// &lt;p&gt;Optionally specify the start time for all of your manifest egress requests. When you include start time, note that you cannot use start time query parameters for this manifest's endpoint URL.&lt;/p&gt;
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        /// <summary>
        /// &lt;p&gt;Optionally specify the time delay for all of your manifest egress requests. Enter a value that is smaller than your endpoint's startover window. When you include time delay, note that you cannot use time delay query parameters for this manifest's endpoint URL.&lt;/p&gt;
        /// </summary>
        [Input("timeDelaySeconds")]
        public Input<int>? TimeDelaySeconds { get; set; }

        public OriginEndpointFilterConfigurationArgs()
        {
        }
        public static new OriginEndpointFilterConfigurationArgs Empty => new OriginEndpointFilterConfigurationArgs();
    }
}
