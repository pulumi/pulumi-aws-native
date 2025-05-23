// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ConnectCampaignsV2.Outputs
{

    /// <summary>
    /// Time range in 24 hour format
    /// </summary>
    [OutputType]
    public sealed class CampaignTimeRange
    {
        public readonly string EndTime;
        public readonly string StartTime;

        [OutputConstructor]
        private CampaignTimeRange(
            string endTime,

            string startTime)
        {
            EndTime = endTime;
            StartTime = startTime;
        }
    }
}
