// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ConnectCampaigns.Outputs
{

    /// <summary>
    /// The possible types of dialer config parameters
    /// </summary>
    [OutputType]
    public sealed class CampaignDialerConfig
    {
        public readonly Outputs.CampaignPredictiveDialerConfig? PredictiveDialerConfig;
        public readonly Outputs.CampaignProgressiveDialerConfig? ProgressiveDialerConfig;

        [OutputConstructor]
        private CampaignDialerConfig(
            Outputs.CampaignPredictiveDialerConfig? predictiveDialerConfig,

            Outputs.CampaignProgressiveDialerConfig? progressiveDialerConfig)
        {
            PredictiveDialerConfig = predictiveDialerConfig;
            ProgressiveDialerConfig = progressiveDialerConfig;
        }
    }
}