// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTFleetWise.Outputs
{

    [OutputType]
    public sealed class CampaignConditionBasedCollectionScheme
    {
        public readonly int? ConditionLanguageVersion;
        public readonly string Expression;
        public readonly double? MinimumTriggerIntervalMs;
        public readonly Pulumi.AwsNative.IoTFleetWise.CampaignTriggerMode? TriggerMode;

        [OutputConstructor]
        private CampaignConditionBasedCollectionScheme(
            int? conditionLanguageVersion,

            string expression,

            double? minimumTriggerIntervalMs,

            Pulumi.AwsNative.IoTFleetWise.CampaignTriggerMode? triggerMode)
        {
            ConditionLanguageVersion = conditionLanguageVersion;
            Expression = expression;
            MinimumTriggerIntervalMs = minimumTriggerIntervalMs;
            TriggerMode = triggerMode;
        }
    }
}