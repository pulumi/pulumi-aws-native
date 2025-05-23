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
    /// &lt;p&gt;The SCTE configuration.&lt;/p&gt;
    /// </summary>
    public sealed class OriginEndpointScteDashArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Choose how ad markers are included in the packaged content. If you include ad markers in the content stream in your upstream encoders, then you need to inform MediaPackage what to do with the ad markers in the output.
        /// 
        /// Value description:
        /// 
        /// - `Binary` - The SCTE-35 marker is expressed as a hex-string (Base64 string) rather than full XML.
        /// - `XML` - The SCTE marker is expressed fully in XML.
        /// </summary>
        [Input("adMarkerDash")]
        public Input<Pulumi.AwsNative.MediaPackageV2.OriginEndpointAdMarkerDash>? AdMarkerDash { get; set; }

        public OriginEndpointScteDashArgs()
        {
        }
        public static new OriginEndpointScteDashArgs Empty => new OriginEndpointScteDashArgs();
    }
}
