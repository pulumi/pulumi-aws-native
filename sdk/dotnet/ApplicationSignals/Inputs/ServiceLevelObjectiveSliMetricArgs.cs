// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationSignals.Inputs
{

    /// <summary>
    /// A structure that contains information about the metric that the SLO monitors.
    /// </summary>
    public sealed class ServiceLevelObjectiveSliMetricArgs : global::Pulumi.ResourceArgs
    {
        [Input("keyAttributes")]
        private InputMap<string>? _keyAttributes;

        /// <summary>
        /// If this SLO is related to a metric collected by Application Signals, you must use this field to specify which service the SLO metric is related to. To do so, you must specify at least the `Type` , `Name` , and `Environment` attributes.
        /// 
        /// This is a string-to-string map. It can include the following fields.
        /// 
        /// - `Type` designates the type of object this is.
        /// - `ResourceType` specifies the type of the resource. This field is used only when the value of the `Type` field is `Resource` or `AWS::Resource` .
        /// - `Name` specifies the name of the object. This is used only if the value of the `Type` field is `Service` , `RemoteService` , or `AWS::Service` .
        /// - `Identifier` identifies the resource objects of this resource. This is used only if the value of the `Type` field is `Resource` or `AWS::Resource` .
        /// - `Environment` specifies the location where this object is hosted, or what it belongs to.
        /// </summary>
        public InputMap<string> KeyAttributes
        {
            get => _keyAttributes ?? (_keyAttributes = new InputMap<string>());
            set => _keyAttributes = value;
        }

        [Input("metricDataQueries")]
        private InputList<Inputs.ServiceLevelObjectiveMetricDataQueryArgs>? _metricDataQueries;

        /// <summary>
        /// If this SLO monitors a CloudWatch metric or the result of a CloudWatch metric math expression, use this structure to specify that metric or expression.
        /// </summary>
        public InputList<Inputs.ServiceLevelObjectiveMetricDataQueryArgs> MetricDataQueries
        {
            get => _metricDataQueries ?? (_metricDataQueries = new InputList<Inputs.ServiceLevelObjectiveMetricDataQueryArgs>());
            set => _metricDataQueries = value;
        }

        /// <summary>
        /// If the SLO monitors either the LATENCY or AVAILABILITY metric that Application Signals collects, this field displays which of those metrics is used.
        /// </summary>
        [Input("metricType")]
        public Input<Pulumi.AwsNative.ApplicationSignals.ServiceLevelObjectiveSliMetricMetricType>? MetricType { get; set; }

        /// <summary>
        /// If the SLO monitors a specific operation of the service, this field displays that operation name.
        /// </summary>
        [Input("operationName")]
        public Input<string>? OperationName { get; set; }

        /// <summary>
        /// The number of seconds to use as the period for SLO evaluation. Your application's performance is compared to the SLI during each period. For each period, the application is determined to have either achieved or not achieved the necessary performance.
        /// </summary>
        [Input("periodSeconds")]
        public Input<int>? PeriodSeconds { get; set; }

        /// <summary>
        /// The statistic to use for comparison to the threshold. It can be any CloudWatch statistic or extended statistic
        /// </summary>
        [Input("statistic")]
        public Input<string>? Statistic { get; set; }

        public ServiceLevelObjectiveSliMetricArgs()
        {
        }
        public static new ServiceLevelObjectiveSliMetricArgs Empty => new ServiceLevelObjectiveSliMetricArgs();
    }
}