// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise
{
    public static class GetDashboard
    {
        /// <summary>
        /// Resource schema for AWS::IoTSiteWise::Dashboard
        /// </summary>
        public static Task<GetDashboardResult> InvokeAsync(GetDashboardArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDashboardResult>("aws-native:iotsitewise:getDashboard", args ?? new GetDashboardArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::IoTSiteWise::Dashboard
        /// </summary>
        public static Output<GetDashboardResult> Invoke(GetDashboardInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDashboardResult>("aws-native:iotsitewise:getDashboard", args ?? new GetDashboardInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::IoTSiteWise::Dashboard
        /// </summary>
        public static Output<GetDashboardResult> Invoke(GetDashboardInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDashboardResult>("aws-native:iotsitewise:getDashboard", args ?? new GetDashboardInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDashboardArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the dashboard.
        /// </summary>
        [Input("dashboardId", required: true)]
        public string DashboardId { get; set; } = null!;

        public GetDashboardArgs()
        {
        }
        public static new GetDashboardArgs Empty => new GetDashboardArgs();
    }

    public sealed class GetDashboardInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the dashboard.
        /// </summary>
        [Input("dashboardId", required: true)]
        public Input<string> DashboardId { get; set; } = null!;

        public GetDashboardInvokeArgs()
        {
        }
        public static new GetDashboardInvokeArgs Empty => new GetDashboardInvokeArgs();
    }


    [OutputType]
    public sealed class GetDashboardResult
    {
        /// <summary>
        /// The ARN of the dashboard.
        /// </summary>
        public readonly string? DashboardArn;
        /// <summary>
        /// The dashboard definition specified in a JSON literal.
        /// </summary>
        public readonly string? DashboardDefinition;
        /// <summary>
        /// A description for the dashboard.
        /// </summary>
        public readonly string? DashboardDescription;
        /// <summary>
        /// The ID of the dashboard.
        /// </summary>
        public readonly string? DashboardId;
        /// <summary>
        /// A friendly name for the dashboard.
        /// </summary>
        public readonly string? DashboardName;
        /// <summary>
        /// A list of key-value pairs that contain metadata for the dashboard.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetDashboardResult(
            string? dashboardArn,

            string? dashboardDefinition,

            string? dashboardDescription,

            string? dashboardId,

            string? dashboardName,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            DashboardArn = dashboardArn;
            DashboardDefinition = dashboardDefinition;
            DashboardDescription = dashboardDescription;
            DashboardId = dashboardId;
            DashboardName = dashboardName;
            Tags = tags;
        }
    }
}
