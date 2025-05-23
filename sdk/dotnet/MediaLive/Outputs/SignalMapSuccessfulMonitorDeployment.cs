// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    /// <summary>
    /// Represents the latest successful monitor deployment of a signal map.
    /// </summary>
    [OutputType]
    public sealed class SignalMapSuccessfulMonitorDeployment
    {
        /// <summary>
        /// URI associated with a signal map's monitor deployment.
        /// </summary>
        public readonly string DetailsUri;
        /// <summary>
        /// A signal map's monitor deployment status.
        /// </summary>
        public readonly Pulumi.AwsNative.MediaLive.SignalMapMonitorDeploymentStatus Status;

        [OutputConstructor]
        private SignalMapSuccessfulMonitorDeployment(
            string detailsUri,

            Pulumi.AwsNative.MediaLive.SignalMapMonitorDeploymentStatus status)
        {
            DetailsUri = detailsUri;
            Status = status;
        }
    }
}
