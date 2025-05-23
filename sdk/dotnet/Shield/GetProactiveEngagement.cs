// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Shield
{
    public static class GetProactiveEngagement
    {
        /// <summary>
        /// Authorizes the Shield Response Team (SRT) to use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
        /// </summary>
        public static Task<GetProactiveEngagementResult> InvokeAsync(GetProactiveEngagementArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProactiveEngagementResult>("aws-native:shield:getProactiveEngagement", args ?? new GetProactiveEngagementArgs(), options.WithDefaults());

        /// <summary>
        /// Authorizes the Shield Response Team (SRT) to use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
        /// </summary>
        public static Output<GetProactiveEngagementResult> Invoke(GetProactiveEngagementInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProactiveEngagementResult>("aws-native:shield:getProactiveEngagement", args ?? new GetProactiveEngagementInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Authorizes the Shield Response Team (SRT) to use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
        /// </summary>
        public static Output<GetProactiveEngagementResult> Invoke(GetProactiveEngagementInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetProactiveEngagementResult>("aws-native:shield:getProactiveEngagement", args ?? new GetProactiveEngagementInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProactiveEngagementArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the account that submitted the template.
        /// </summary>
        [Input("accountId", required: true)]
        public string AccountId { get; set; } = null!;

        public GetProactiveEngagementArgs()
        {
        }
        public static new GetProactiveEngagementArgs Empty => new GetProactiveEngagementArgs();
    }

    public sealed class GetProactiveEngagementInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the account that submitted the template.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        public GetProactiveEngagementInvokeArgs()
        {
        }
        public static new GetProactiveEngagementInvokeArgs Empty => new GetProactiveEngagementInvokeArgs();
    }


    [OutputType]
    public sealed class GetProactiveEngagementResult
    {
        /// <summary>
        /// The ID of the account that submitted the template.
        /// </summary>
        public readonly string? AccountId;
        /// <summary>
        /// A list of email addresses and phone numbers that the Shield Response Team (SRT) can use to contact you for escalations to the SRT and to initiate proactive customer support.
        /// To enable proactive engagement, the contact list must include at least one phone number.
        /// </summary>
        public readonly ImmutableArray<Outputs.ProactiveEngagementEmergencyContact> EmergencyContactList;
        /// <summary>
        /// If `ENABLED`, the Shield Response Team (SRT) will use email and phone to notify contacts about escalations to the SRT and to initiate proactive customer support.
        /// If `DISABLED`, the SRT will not proactively notify contacts about escalations or to initiate proactive customer support.
        /// </summary>
        public readonly Pulumi.AwsNative.Shield.ProactiveEngagementStatus? ProactiveEngagementStatus;

        [OutputConstructor]
        private GetProactiveEngagementResult(
            string? accountId,

            ImmutableArray<Outputs.ProactiveEngagementEmergencyContact> emergencyContactList,

            Pulumi.AwsNative.Shield.ProactiveEngagementStatus? proactiveEngagementStatus)
        {
            AccountId = accountId;
            EmergencyContactList = emergencyContactList;
            ProactiveEngagementStatus = proactiveEngagementStatus;
        }
    }
}
