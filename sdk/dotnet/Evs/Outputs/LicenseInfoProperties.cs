// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evs.Outputs
{

    /// <summary>
    /// The license information for an EVS environment
    /// </summary>
    [OutputType]
    public sealed class LicenseInfoProperties
    {
        /// <summary>
        /// The VCF solution key. This license unlocks VMware VCF product features, including vSphere, NSX, SDDC Manager, and vCenter Server. The VCF solution key must cover a minimum of 256 cores.
        /// </summary>
        public readonly string SolutionKey;
        /// <summary>
        /// The VSAN license key. This license unlocks vSAN features. The vSAN license key must provide at least 110 TiB of vSAN capacity.
        /// </summary>
        public readonly string VsanKey;

        [OutputConstructor]
        private LicenseInfoProperties(
            string solutionKey,

            string vsanKey)
        {
            SolutionKey = solutionKey;
            VsanKey = vsanKey;
        }
    }
}
