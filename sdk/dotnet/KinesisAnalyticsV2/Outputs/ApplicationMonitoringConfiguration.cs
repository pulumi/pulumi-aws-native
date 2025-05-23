// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Outputs
{

    /// <summary>
    /// Describes configuration parameters for Amazon CloudWatch logging for a Java-based Kinesis Data Analytics application. For more information about CloudWatch logging, see Monitoring.
    /// </summary>
    [OutputType]
    public sealed class ApplicationMonitoringConfiguration
    {
        /// <summary>
        /// Describes whether to use the default CloudWatch logging configuration for an application. You must set this property to CUSTOM in order to set the LogLevel or MetricsLevel parameters.
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationMonitoringConfigurationConfigurationType ConfigurationType;
        /// <summary>
        /// Describes the verbosity of the CloudWatch Logs for an application.
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationMonitoringConfigurationLogLevel? LogLevel;
        /// <summary>
        /// Describes the granularity of the CloudWatch Logs for an application. The Parallelism level is not recommended for applications with a Parallelism over 64 due to excessive costs.
        /// </summary>
        public readonly Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationMonitoringConfigurationMetricsLevel? MetricsLevel;

        [OutputConstructor]
        private ApplicationMonitoringConfiguration(
            Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationMonitoringConfigurationConfigurationType configurationType,

            Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationMonitoringConfigurationLogLevel? logLevel,

            Pulumi.AwsNative.KinesisAnalyticsV2.ApplicationMonitoringConfigurationMetricsLevel? metricsLevel)
        {
            ConfigurationType = configurationType;
            LogLevel = logLevel;
            MetricsLevel = metricsLevel;
        }
    }
}
