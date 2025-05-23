// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudTrail
{
    public static class GetDashboard
    {
        /// <summary>
        /// The Amazon CloudTrail dashboard resource allows customers to manage managed dashboards and create custom dashboards. You can manually refresh custom and managed dashboards. For custom dashboards, you can also set up an automatic refresh schedule and modify dashboard widgets.
        /// </summary>
        public static Task<GetDashboardResult> InvokeAsync(GetDashboardArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDashboardResult>("aws-native:cloudtrail:getDashboard", args ?? new GetDashboardArgs(), options.WithDefaults());

        /// <summary>
        /// The Amazon CloudTrail dashboard resource allows customers to manage managed dashboards and create custom dashboards. You can manually refresh custom and managed dashboards. For custom dashboards, you can also set up an automatic refresh schedule and modify dashboard widgets.
        /// </summary>
        public static Output<GetDashboardResult> Invoke(GetDashboardInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDashboardResult>("aws-native:cloudtrail:getDashboard", args ?? new GetDashboardInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The Amazon CloudTrail dashboard resource allows customers to manage managed dashboards and create custom dashboards. You can manually refresh custom and managed dashboards. For custom dashboards, you can also set up an automatic refresh schedule and modify dashboard widgets.
        /// </summary>
        public static Output<GetDashboardResult> Invoke(GetDashboardInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDashboardResult>("aws-native:cloudtrail:getDashboard", args ?? new GetDashboardInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDashboardArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the dashboard.
        /// </summary>
        [Input("dashboardArn", required: true)]
        public string DashboardArn { get; set; } = null!;

        public GetDashboardArgs()
        {
        }
        public static new GetDashboardArgs Empty => new GetDashboardArgs();
    }

    public sealed class GetDashboardInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the dashboard.
        /// </summary>
        [Input("dashboardArn", required: true)]
        public Input<string> DashboardArn { get; set; } = null!;

        public GetDashboardInvokeArgs()
        {
        }
        public static new GetDashboardInvokeArgs Empty => new GetDashboardInvokeArgs();
    }


    [OutputType]
    public sealed class GetDashboardResult
    {
        /// <summary>
        /// The timestamp of the dashboard creation.
        /// </summary>
        public readonly string? CreatedTimestamp;
        /// <summary>
        /// The ARN of the dashboard.
        /// </summary>
        public readonly string? DashboardArn;
        /// <summary>
        /// The name of the dashboard.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Configures the automatic refresh schedule for the dashboard. Includes the frequency unit (DAYS or HOURS) and value, as well as the status (ENABLED or DISABLED) of the refresh schedule.
        /// </summary>
        public readonly Outputs.DashboardRefreshSchedule? RefreshSchedule;
        /// <summary>
        /// The status of the dashboard. Values are CREATING, CREATED, UPDATING, UPDATED and DELETING.
        /// </summary>
        public readonly Pulumi.AwsNative.CloudTrail.DashboardStatus? Status;
        /// <summary>
        /// A list of tags.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// Indicates whether the dashboard is protected from termination.
        /// </summary>
        public readonly bool? TerminationProtectionEnabled;
        /// <summary>
        /// The type of the dashboard. Values are CUSTOM and MANAGED.
        /// </summary>
        public readonly Pulumi.AwsNative.CloudTrail.DashboardType? Type;
        /// <summary>
        /// The timestamp showing when the dashboard was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.
        /// </summary>
        public readonly string? UpdatedTimestamp;
        /// <summary>
        /// List of widgets on the dashboard
        /// </summary>
        public readonly ImmutableArray<Outputs.DashboardWidget> Widgets;

        [OutputConstructor]
        private GetDashboardResult(
            string? createdTimestamp,

            string? dashboardArn,

            string? name,

            Outputs.DashboardRefreshSchedule? refreshSchedule,

            Pulumi.AwsNative.CloudTrail.DashboardStatus? status,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            bool? terminationProtectionEnabled,

            Pulumi.AwsNative.CloudTrail.DashboardType? type,

            string? updatedTimestamp,

            ImmutableArray<Outputs.DashboardWidget> widgets)
        {
            CreatedTimestamp = createdTimestamp;
            DashboardArn = dashboardArn;
            Name = name;
            RefreshSchedule = refreshSchedule;
            Status = status;
            Tags = tags;
            TerminationProtectionEnabled = terminationProtectionEnabled;
            Type = type;
            UpdatedTimestamp = updatedTimestamp;
            Widgets = widgets;
        }
    }
}
