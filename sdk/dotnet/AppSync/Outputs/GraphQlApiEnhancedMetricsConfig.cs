// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppSync.Outputs
{

    [OutputType]
    public sealed class GraphQlApiEnhancedMetricsConfig
    {
        public readonly string DataSourceLevelMetricsBehavior;
        public readonly string OperationLevelMetricsConfig;
        public readonly string ResolverLevelMetricsBehavior;

        [OutputConstructor]
        private GraphQlApiEnhancedMetricsConfig(
            string dataSourceLevelMetricsBehavior,

            string operationLevelMetricsConfig,

            string resolverLevelMetricsBehavior)
        {
            DataSourceLevelMetricsBehavior = dataSourceLevelMetricsBehavior;
            OperationLevelMetricsConfig = operationLevelMetricsConfig;
            ResolverLevelMetricsBehavior = resolverLevelMetricsBehavior;
        }
    }
}