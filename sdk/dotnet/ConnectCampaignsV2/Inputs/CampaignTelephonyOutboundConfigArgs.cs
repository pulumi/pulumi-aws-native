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
    /// Default Telephone Outbound config
    /// </summary>
    public sealed class CampaignTelephonyOutboundConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The answering machine detection configuration.
        /// </summary>
        [Input("answerMachineDetectionConfig")]
        public Input<Inputs.CampaignAnswerMachineDetectionConfigArgs>? AnswerMachineDetectionConfig { get; set; }

        /// <summary>
        /// The identifier of the published Amazon Connect contact flow.
        /// </summary>
        [Input("connectContactFlowId", required: true)]
        public Input<string> ConnectContactFlowId { get; set; } = null!;

        /// <summary>
        /// The Amazon Connect source phone number.
        /// </summary>
        [Input("connectSourcePhoneNumber")]
        public Input<string>? ConnectSourcePhoneNumber { get; set; }

        public CampaignTelephonyOutboundConfigArgs()
        {
        }
        public static new CampaignTelephonyOutboundConfigArgs Empty => new CampaignTelephonyOutboundConfigArgs();
    }
}
