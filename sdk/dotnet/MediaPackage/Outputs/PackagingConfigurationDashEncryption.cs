// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    /// <summary>
    /// A Dynamic Adaptive Streaming over HTTP (DASH) encryption configuration.
    /// </summary>
    [OutputType]
    public sealed class PackagingConfigurationDashEncryption
    {
        public readonly Outputs.PackagingConfigurationSpekeKeyProvider SpekeKeyProvider;

        [OutputConstructor]
        private PackagingConfigurationDashEncryption(Outputs.PackagingConfigurationSpekeKeyProvider spekeKeyProvider)
        {
            SpekeKeyProvider = spekeKeyProvider;
        }
    }
}