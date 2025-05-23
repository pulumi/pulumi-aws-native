// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Outputs
{

    [OutputType]
    public sealed class ClusterCloudWatchLogs
    {
        /// <summary>
        /// Specifies whether broker logs get sent to the specified CloudWatch Logs destination.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The CloudWatch log group that is the destination for broker logs.
        /// </summary>
        public readonly string? LogGroup;

        [OutputConstructor]
        private ClusterCloudWatchLogs(
            bool enabled,

            string? logGroup)
        {
            Enabled = enabled;
            LogGroup = logGroup;
        }
    }
}
