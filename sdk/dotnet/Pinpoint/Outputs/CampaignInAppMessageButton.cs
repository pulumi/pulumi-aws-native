// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Outputs
{

    [OutputType]
    public sealed class CampaignInAppMessageButton
    {
        public readonly Outputs.CampaignOverrideButtonConfiguration? Android;
        public readonly Outputs.CampaignDefaultButtonConfiguration? DefaultConfig;
        public readonly Outputs.CampaignOverrideButtonConfiguration? IOS;
        public readonly Outputs.CampaignOverrideButtonConfiguration? Web;

        [OutputConstructor]
        private CampaignInAppMessageButton(
            Outputs.CampaignOverrideButtonConfiguration? android,

            Outputs.CampaignDefaultButtonConfiguration? defaultConfig,

            Outputs.CampaignOverrideButtonConfiguration? iOS,

            Outputs.CampaignOverrideButtonConfiguration? web)
        {
            Android = android;
            DefaultConfig = defaultConfig;
            IOS = iOS;
            Web = web;
        }
    }
}