// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ConnectCampaignsV2.Inputs
{

    /// <summary>
    /// Open Hours config
    /// </summary>
    public sealed class CampaignOpenHoursArgs : global::Pulumi.ResourceArgs
    {
        [Input("dailyHours", required: true)]
        private InputList<Inputs.CampaignDailyHourArgs>? _dailyHours;

        /// <summary>
        /// The daily hours configuration.
        /// </summary>
        public InputList<Inputs.CampaignDailyHourArgs> DailyHours
        {
            get => _dailyHours ?? (_dailyHours = new InputList<Inputs.CampaignDailyHourArgs>());
            set => _dailyHours = value;
        }

        public CampaignOpenHoursArgs()
        {
        }
        public static new CampaignOpenHoursArgs Empty => new CampaignOpenHoursArgs();
    }
}
