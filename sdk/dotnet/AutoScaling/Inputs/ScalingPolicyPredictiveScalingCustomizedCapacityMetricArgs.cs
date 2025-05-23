// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Inputs
{

    public sealed class ScalingPolicyPredictiveScalingCustomizedCapacityMetricArgs : global::Pulumi.ResourceArgs
    {
        [Input("metricDataQueries", required: true)]
        private InputList<Inputs.ScalingPolicyMetricDataQueryArgs>? _metricDataQueries;

        /// <summary>
        /// One or more metric data queries to provide the data points for a capacity metric. Use multiple metric data queries only if you are performing a math expression on returned data.
        /// </summary>
        public InputList<Inputs.ScalingPolicyMetricDataQueryArgs> MetricDataQueries
        {
            get => _metricDataQueries ?? (_metricDataQueries = new InputList<Inputs.ScalingPolicyMetricDataQueryArgs>());
            set => _metricDataQueries = value;
        }

        public ScalingPolicyPredictiveScalingCustomizedCapacityMetricArgs()
        {
        }
        public static new ScalingPolicyPredictiveScalingCustomizedCapacityMetricArgs Empty => new ScalingPolicyPredictiveScalingCustomizedCapacityMetricArgs();
    }
}
