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
    /// &lt;p&gt;The SCTE configuration.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class OriginEndpointScteHls
    {
        public readonly Pulumi.AwsNative.MediaPackageV2.OriginEndpointAdMarkerHls? AdMarkerHls;

        [OutputConstructor]
        private OriginEndpointScteHls(Pulumi.AwsNative.MediaPackageV2.OriginEndpointAdMarkerHls? adMarkerHls)
        {
            AdMarkerHls = adMarkerHls;
        }
    }
}