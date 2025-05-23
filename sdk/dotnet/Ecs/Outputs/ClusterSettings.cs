// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Outputs
{

    /// <summary>
    /// The settings to use when creating a cluster. This parameter is used to turn on CloudWatch Container Insights with enhanced observability or CloudWatch Container Insights for a cluster.
    ///  Container Insights with enhanced observability provides all the Container Insights metrics, plus additional task and container metrics. This version supports enhanced observability for Amazon ECS clusters using the Amazon EC2 and Fargate launch types. After you configure Container Insights with enhanced observability on Amazon ECS, Container Insights auto-collects detailed infrastructure telemetry from the cluster level down to the container level in your environment and displays these critical performance data in curated dashboards removing the heavy lifting in observability set-up. 
    ///  For more information, see [Monitor Amazon ECS containers using Container Insights with enhanced observability](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cloudwatch-container-insights.html) in the *Amazon Elastic Container Service Developer Guide*.
    /// </summary>
    [OutputType]
    public sealed class ClusterSettings
    {
        /// <summary>
        /// The name of the cluster setting. The value is ``containerInsights`` .
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The value to set for the cluster setting. The supported values are ``enhanced``, ``enabled``, and ``disabled``. 
        ///  To use Container Insights with enhanced observability, set the ``containerInsights`` account setting to ``enhanced``.
        ///  To use Container Insights, set the ``containerInsights`` account setting to ``enabled``.
        ///  If a cluster value is specified, it will override the ``containerInsights`` value set with [PutAccountSetting](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSetting.html) or [PutAccountSettingDefault](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PutAccountSettingDefault.html).
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private ClusterSettings(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
