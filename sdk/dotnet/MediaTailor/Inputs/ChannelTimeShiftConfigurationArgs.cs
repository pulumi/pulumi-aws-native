// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor.Inputs
{

    /// <summary>
    /// &lt;p&gt;The configuration for time-shifted viewing.&lt;/p&gt;
    /// </summary>
    public sealed class ChannelTimeShiftConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The maximum time delay for time-shifted viewing. The minimum allowed maximum time delay is 0 seconds, and the maximum allowed maximum time delay is 21600 seconds (6 hours).&lt;/p&gt;
        /// </summary>
        [Input("maxTimeDelaySeconds", required: true)]
        public Input<double> MaxTimeDelaySeconds { get; set; } = null!;

        public ChannelTimeShiftConfigurationArgs()
        {
        }
        public static new ChannelTimeShiftConfigurationArgs Empty => new ChannelTimeShiftConfigurationArgs();
    }
}
