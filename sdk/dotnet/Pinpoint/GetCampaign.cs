// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    public static class GetCampaign
    {
        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::Campaign
        /// </summary>
        public static Task<GetCampaignResult> InvokeAsync(GetCampaignArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCampaignResult>("aws-native:pinpoint:getCampaign", args ?? new GetCampaignArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::Campaign
        /// </summary>
        public static Output<GetCampaignResult> Invoke(GetCampaignInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCampaignResult>("aws-native:pinpoint:getCampaign", args ?? new GetCampaignInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCampaignArgs : global::Pulumi.InvokeArgs
    {
        [Input("campaignId", required: true)]
        public string CampaignId { get; set; } = null!;

        public GetCampaignArgs()
        {
        }
        public static new GetCampaignArgs Empty => new GetCampaignArgs();
    }

    public sealed class GetCampaignInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("campaignId", required: true)]
        public Input<string> CampaignId { get; set; } = null!;

        public GetCampaignInvokeArgs()
        {
        }
        public static new GetCampaignInvokeArgs Empty => new GetCampaignInvokeArgs();
    }


    [OutputType]
    public sealed class GetCampaignResult
    {
        public readonly ImmutableArray<Outputs.CampaignWriteTreatmentResource> AdditionalTreatments;
        public readonly string? Arn;
        public readonly Outputs.CampaignHook? CampaignHook;
        public readonly string? CampaignId;
        public readonly Outputs.CampaignCustomDeliveryConfiguration? CustomDeliveryConfiguration;
        public readonly string? Description;
        public readonly int? HoldoutPercent;
        public readonly bool? IsPaused;
        public readonly Outputs.CampaignLimits? Limits;
        public readonly Outputs.CampaignMessageConfiguration? MessageConfiguration;
        public readonly string? Name;
        public readonly int? Priority;
        public readonly Outputs.CampaignSchedule? Schedule;
        public readonly string? SegmentId;
        public readonly int? SegmentVersion;
        public readonly object? Tags;
        public readonly Outputs.CampaignTemplateConfiguration? TemplateConfiguration;
        public readonly string? TreatmentDescription;
        public readonly string? TreatmentName;

        [OutputConstructor]
        private GetCampaignResult(
            ImmutableArray<Outputs.CampaignWriteTreatmentResource> additionalTreatments,

            string? arn,

            Outputs.CampaignHook? campaignHook,

            string? campaignId,

            Outputs.CampaignCustomDeliveryConfiguration? customDeliveryConfiguration,

            string? description,

            int? holdoutPercent,

            bool? isPaused,

            Outputs.CampaignLimits? limits,

            Outputs.CampaignMessageConfiguration? messageConfiguration,

            string? name,

            int? priority,

            Outputs.CampaignSchedule? schedule,

            string? segmentId,

            int? segmentVersion,

            object? tags,

            Outputs.CampaignTemplateConfiguration? templateConfiguration,

            string? treatmentDescription,

            string? treatmentName)
        {
            AdditionalTreatments = additionalTreatments;
            Arn = arn;
            CampaignHook = campaignHook;
            CampaignId = campaignId;
            CustomDeliveryConfiguration = customDeliveryConfiguration;
            Description = description;
            HoldoutPercent = holdoutPercent;
            IsPaused = isPaused;
            Limits = limits;
            MessageConfiguration = messageConfiguration;
            Name = name;
            Priority = priority;
            Schedule = schedule;
            SegmentId = segmentId;
            SegmentVersion = segmentVersion;
            Tags = tags;
            TemplateConfiguration = templateConfiguration;
            TreatmentDescription = treatmentDescription;
            TreatmentName = treatmentName;
        }
    }
}