// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackageV2.Inputs
{

    /// <summary>
    /// &lt;p&gt;The encryption type.&lt;/p&gt;
    /// </summary>
    public sealed class OriginEndpointEncryptionMethodArgs : global::Pulumi.ResourceArgs
    {
        [Input("cmafEncryptionMethod")]
        public Input<Pulumi.AwsNative.MediaPackageV2.OriginEndpointCmafEncryptionMethod>? CmafEncryptionMethod { get; set; }

        [Input("tsEncryptionMethod")]
        public Input<Pulumi.AwsNative.MediaPackageV2.OriginEndpointTsEncryptionMethod>? TsEncryptionMethod { get; set; }

        public OriginEndpointEncryptionMethodArgs()
        {
        }
        public static new OriginEndpointEncryptionMethodArgs Empty => new OriginEndpointEncryptionMethodArgs();
    }
}