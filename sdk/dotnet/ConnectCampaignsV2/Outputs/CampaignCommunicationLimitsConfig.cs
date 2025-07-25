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
    /// Communication limits config
    /// </summary>
    [OutputType]
    public sealed class CampaignCommunicationLimitsConfig
    {
        /// <summary>
        /// The CommunicationLimits that apply to all channel subtypes defined in an outbound campaign.
        /// </summary>
        public readonly Outputs.CampaignCommunicationLimits? AllChannelsSubtypes;
        /// <summary>
        /// Opt-in or Opt-out from instance-level limits.
        /// </summary>
        public readonly Pulumi.AwsNative.ConnectCampaignsV2.CampaignInstanceLimitsHandling? InstanceLimitsHandling;

        [OutputConstructor]
        private CampaignCommunicationLimitsConfig(
            Outputs.CampaignCommunicationLimits? allChannelsSubtypes,

            Pulumi.AwsNative.ConnectCampaignsV2.CampaignInstanceLimitsHandling? instanceLimitsHandling)
        {
            AllChannelsSubtypes = allChannelsSubtypes;
            InstanceLimitsHandling = instanceLimitsHandling;
        }
    }
}
