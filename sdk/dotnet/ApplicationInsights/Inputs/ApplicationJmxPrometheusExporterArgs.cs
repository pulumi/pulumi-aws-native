// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApplicationInsights.Inputs
{

    /// <summary>
    /// The JMX Prometheus Exporter settings.
    /// </summary>
    public sealed class ApplicationJmxPrometheusExporterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Java agent host port
        /// </summary>
        [Input("hostPort")]
        public Input<string>? HostPort { get; set; }

        /// <summary>
        /// JMX service URL.
        /// </summary>
        [Input("jmxurl")]
        public Input<string>? Jmxurl { get; set; }

        /// <summary>
        /// Prometheus exporter port.
        /// </summary>
        [Input("prometheusPort")]
        public Input<string>? PrometheusPort { get; set; }

        public ApplicationJmxPrometheusExporterArgs()
        {
        }
        public static new ApplicationJmxPrometheusExporterArgs Empty => new ApplicationJmxPrometheusExporterArgs();
    }
}