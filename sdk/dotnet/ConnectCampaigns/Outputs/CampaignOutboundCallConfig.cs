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
    /// The configuration used for outbound calls.
    /// </summary>
    [OutputType]
    public sealed class CampaignOutboundCallConfig
    {
        /// <summary>
        /// The identifier of the contact flow for the outbound call.
        /// </summary>
        public readonly string ConnectContactFlowArn;
        /// <summary>
        /// The queue for the call. If you specify a queue, the phone displayed for caller ID is the phone number specified in the queue. If you do not specify a queue, the queue defined in the contact flow is used. If you do not specify a queue, you must specify a source phone number.
        /// </summary>
        public readonly string ConnectQueueArn;
        /// <summary>
        /// The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.
        /// </summary>
        public readonly string? ConnectSourcePhoneNumber;

        [OutputConstructor]
        private CampaignOutboundCallConfig(
            string connectContactFlowArn,

            string connectQueueArn,

            string? connectSourcePhoneNumber)
        {
            ConnectContactFlowArn = connectContactFlowArn;
            ConnectQueueArn = connectQueueArn;
            ConnectSourcePhoneNumber = connectSourcePhoneNumber;
        }
    }
}