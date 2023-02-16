// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog
{
    public static class GetLaunchNotificationConstraint
    {
        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::LaunchNotificationConstraint
        /// </summary>
        public static Task<GetLaunchNotificationConstraintResult> InvokeAsync(GetLaunchNotificationConstraintArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLaunchNotificationConstraintResult>("aws-native:servicecatalog:getLaunchNotificationConstraint", args ?? new GetLaunchNotificationConstraintArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::LaunchNotificationConstraint
        /// </summary>
        public static Output<GetLaunchNotificationConstraintResult> Invoke(GetLaunchNotificationConstraintInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLaunchNotificationConstraintResult>("aws-native:servicecatalog:getLaunchNotificationConstraint", args ?? new GetLaunchNotificationConstraintInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLaunchNotificationConstraintArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetLaunchNotificationConstraintArgs()
        {
        }
        public static new GetLaunchNotificationConstraintArgs Empty => new GetLaunchNotificationConstraintArgs();
    }

    public sealed class GetLaunchNotificationConstraintInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetLaunchNotificationConstraintInvokeArgs()
        {
        }
        public static new GetLaunchNotificationConstraintInvokeArgs Empty => new GetLaunchNotificationConstraintInvokeArgs();
    }


    [OutputType]
    public sealed class GetLaunchNotificationConstraintResult
    {
        public readonly string? AcceptLanguage;
        public readonly string? Description;
        public readonly string? Id;
        public readonly ImmutableArray<string> NotificationArns;

        [OutputConstructor]
        private GetLaunchNotificationConstraintResult(
            string? acceptLanguage,

            string? description,

            string? id,

            ImmutableArray<string> notificationArns)
        {
            AcceptLanguage = acceptLanguage;
            Description = description;
            Id = id;
            NotificationArns = notificationArns;
        }
    }
}