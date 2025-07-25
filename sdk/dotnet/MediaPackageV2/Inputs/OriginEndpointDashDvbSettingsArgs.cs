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
    /// &lt;p&gt;For endpoints that use the DVB-DASH profile only. The font download and error reporting information that you want MediaPackage to pass through to the manifest.&lt;/p&gt;
    /// </summary>
    public sealed class OriginEndpointDashDvbSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("errorMetrics")]
        private InputList<Inputs.OriginEndpointDashDvbMetricsReportingArgs>? _errorMetrics;

        /// <summary>
        /// &lt;p&gt;Playback device error reporting settings.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.OriginEndpointDashDvbMetricsReportingArgs> ErrorMetrics
        {
            get => _errorMetrics ?? (_errorMetrics = new InputList<Inputs.OriginEndpointDashDvbMetricsReportingArgs>());
            set => _errorMetrics = value;
        }

        [Input("fontDownload")]
        public Input<Inputs.OriginEndpointDashDvbFontDownloadArgs>? FontDownload { get; set; }

        public OriginEndpointDashDvbSettingsArgs()
        {
        }
        public static new OriginEndpointDashDvbSettingsArgs Empty => new OriginEndpointDashDvbSettingsArgs();
    }
}
