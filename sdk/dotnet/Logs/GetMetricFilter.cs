// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs
{
    public static class GetMetricFilter
    {
        /// <summary>
        /// Specifies a metric filter that describes how CloudWatch Logs extracts information from logs and transforms it into Amazon CloudWatch metrics.
        /// </summary>
        public static Task<GetMetricFilterResult> InvokeAsync(GetMetricFilterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMetricFilterResult>("aws-native:logs:getMetricFilter", args ?? new GetMetricFilterArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies a metric filter that describes how CloudWatch Logs extracts information from logs and transforms it into Amazon CloudWatch metrics.
        /// </summary>
        public static Output<GetMetricFilterResult> Invoke(GetMetricFilterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMetricFilterResult>("aws-native:logs:getMetricFilter", args ?? new GetMetricFilterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMetricFilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the metric filter.
        /// </summary>
        [Input("filterName", required: true)]
        public string FilterName { get; set; } = null!;

        /// <summary>
        /// Existing log group that you want to associate with this filter.
        /// </summary>
        [Input("logGroupName", required: true)]
        public string LogGroupName { get; set; } = null!;

        public GetMetricFilterArgs()
        {
        }
        public static new GetMetricFilterArgs Empty => new GetMetricFilterArgs();
    }

    public sealed class GetMetricFilterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A name for the metric filter.
        /// </summary>
        [Input("filterName", required: true)]
        public Input<string> FilterName { get; set; } = null!;

        /// <summary>
        /// Existing log group that you want to associate with this filter.
        /// </summary>
        [Input("logGroupName", required: true)]
        public Input<string> LogGroupName { get; set; } = null!;

        public GetMetricFilterInvokeArgs()
        {
        }
        public static new GetMetricFilterInvokeArgs Empty => new GetMetricFilterInvokeArgs();
    }


    [OutputType]
    public sealed class GetMetricFilterResult
    {
        /// <summary>
        /// Pattern that Logs follows to interpret each entry in a log.
        /// </summary>
        public readonly string? FilterPattern;
        /// <summary>
        /// A collection of information that defines how metric data gets emitted.
        /// </summary>
        public readonly ImmutableArray<Outputs.MetricFilterMetricTransformation> MetricTransformations;

        [OutputConstructor]
        private GetMetricFilterResult(
            string? filterPattern,

            ImmutableArray<Outputs.MetricFilterMetricTransformation> metricTransformations)
        {
            FilterPattern = filterPattern;
            MetricTransformations = metricTransformations;
        }
    }
}