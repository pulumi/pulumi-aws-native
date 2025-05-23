// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecr.Outputs
{

    /// <summary>
    /// The image scanning configuration for a repository.
    /// </summary>
    [OutputType]
    public sealed class RepositoryImageScanningConfiguration
    {
        /// <summary>
        /// The setting that determines whether images are scanned after being pushed to a repository. If set to ``true``, images will be scanned after being pushed. If this parameter is not specified, it will default to ``false`` and images will not be scanned unless a scan is manually started.
        /// </summary>
        public readonly bool? ScanOnPush;

        [OutputConstructor]
        private RepositoryImageScanningConfiguration(bool? scanOnPush)
        {
            ScanOnPush = scanOnPush;
        }
    }
}
