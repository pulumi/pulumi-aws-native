// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Outputs
{

    [OutputType]
    public sealed class ExperimentMetricGoalObject
    {
        /// <summary>
        /// `INCREASE` means that a variation with a higher number for this metric is performing better.
        /// 
        /// `DECREASE` means that a variation with a lower number for this metric is performing better.
        /// </summary>
        public readonly Pulumi.AwsNative.Evidently.ExperimentMetricGoalObjectDesiredChange DesiredChange;
        /// <summary>
        /// The JSON path to reference the entity id in the event.
        /// </summary>
        public readonly string EntityIdKey;
        /// <summary>
        /// Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.
        /// </summary>
        public readonly string? EventPattern;
        /// <summary>
        /// A name for the metric. It can include up to 255 characters.
        /// </summary>
        public readonly string MetricName;
        /// <summary>
        /// A label for the units that the metric is measuring.
        /// </summary>
        public readonly string? UnitLabel;
        /// <summary>
        /// The JSON path to reference the numerical metric value in the event.
        /// </summary>
        public readonly string ValueKey;

        [OutputConstructor]
        private ExperimentMetricGoalObject(
            Pulumi.AwsNative.Evidently.ExperimentMetricGoalObjectDesiredChange desiredChange,

            string entityIdKey,

            string? eventPattern,

            string metricName,

            string? unitLabel,

            string valueKey)
        {
            DesiredChange = desiredChange;
            EntityIdKey = entityIdKey;
            EventPattern = eventPattern;
            MetricName = metricName;
            UnitLabel = unitLabel;
            ValueKey = valueKey;
        }
    }
}
