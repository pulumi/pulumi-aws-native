// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FraudDetector
{
    public static class GetDetector
    {
        /// <summary>
        /// A resource schema for a Detector in Amazon Fraud Detector.
        /// </summary>
        public static Task<GetDetectorResult> InvokeAsync(GetDetectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDetectorResult>("aws-native:frauddetector:getDetector", args ?? new GetDetectorArgs(), options.WithDefaults());

        /// <summary>
        /// A resource schema for a Detector in Amazon Fraud Detector.
        /// </summary>
        public static Output<GetDetectorResult> Invoke(GetDetectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDetectorResult>("aws-native:frauddetector:getDetector", args ?? new GetDetectorInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// A resource schema for a Detector in Amazon Fraud Detector.
        /// </summary>
        public static Output<GetDetectorResult> Invoke(GetDetectorInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDetectorResult>("aws-native:frauddetector:getDetector", args ?? new GetDetectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDetectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the detector.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetDetectorArgs()
        {
        }
        public static new GetDetectorArgs Empty => new GetDetectorArgs();
    }

    public sealed class GetDetectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the detector.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetDetectorInvokeArgs()
        {
        }
        public static new GetDetectorInvokeArgs Empty => new GetDetectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetDetectorResult
    {
        /// <summary>
        /// The ARN of the detector.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The models to associate with this detector.
        /// </summary>
        public readonly ImmutableArray<Outputs.DetectorModel> AssociatedModels;
        /// <summary>
        /// The time when the detector was created.
        /// </summary>
        public readonly string? CreatedTime;
        /// <summary>
        /// The description of the detector.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The active version ID of the detector
        /// </summary>
        public readonly string? DetectorVersionId;
        /// <summary>
        /// The desired detector version status for the detector
        /// </summary>
        public readonly Pulumi.AwsNative.FraudDetector.DetectorVersionStatus? DetectorVersionStatus;
        /// <summary>
        /// The event type to associate this detector with.
        /// </summary>
        public readonly Outputs.DetectorEventType? EventType;
        /// <summary>
        /// The time when the detector was last updated.
        /// </summary>
        public readonly string? LastUpdatedTime;
        /// <summary>
        /// The rule execution mode for the rules included in the detector version.
        /// 
        /// Valid values: `FIRST_MATCHED | ALL_MATCHED` Default value: `FIRST_MATCHED`
        /// 
        /// You can define and edit the rule mode at the detector version level, when it is in draft status.
        /// 
        /// If you specify `FIRST_MATCHED` , Amazon Fraud Detector evaluates rules sequentially, first to last, stopping at the first matched rule. Amazon Fraud dectector then provides the outcomes for that single rule.
        /// 
        /// If you specifiy `ALL_MATCHED` , Amazon Fraud Detector evaluates all rules and returns the outcomes for all matched rules.
        /// </summary>
        public readonly Pulumi.AwsNative.FraudDetector.DetectorRuleExecutionMode? RuleExecutionMode;
        /// <summary>
        /// The rules to include in the detector version.
        /// </summary>
        public readonly ImmutableArray<Outputs.DetectorRule> Rules;
        /// <summary>
        /// Tags associated with this detector.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetDetectorResult(
            string? arn,

            ImmutableArray<Outputs.DetectorModel> associatedModels,

            string? createdTime,

            string? description,

            string? detectorVersionId,

            Pulumi.AwsNative.FraudDetector.DetectorVersionStatus? detectorVersionStatus,

            Outputs.DetectorEventType? eventType,

            string? lastUpdatedTime,

            Pulumi.AwsNative.FraudDetector.DetectorRuleExecutionMode? ruleExecutionMode,

            ImmutableArray<Outputs.DetectorRule> rules,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            AssociatedModels = associatedModels;
            CreatedTime = createdTime;
            Description = description;
            DetectorVersionId = detectorVersionId;
            DetectorVersionStatus = detectorVersionStatus;
            EventType = eventType;
            LastUpdatedTime = lastUpdatedTime;
            RuleExecutionMode = ruleExecutionMode;
            Rules = rules;
            Tags = tags;
        }
    }
}
