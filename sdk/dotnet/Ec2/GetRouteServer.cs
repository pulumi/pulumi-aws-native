// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetRouteServer
    {
        /// <summary>
        /// VPC Route Server
        /// </summary>
        public static Task<GetRouteServerResult> InvokeAsync(GetRouteServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteServerResult>("aws-native:ec2:getRouteServer", args ?? new GetRouteServerArgs(), options.WithDefaults());

        /// <summary>
        /// VPC Route Server
        /// </summary>
        public static Output<GetRouteServerResult> Invoke(GetRouteServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteServerResult>("aws-native:ec2:getRouteServer", args ?? new GetRouteServerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// VPC Route Server
        /// </summary>
        public static Output<GetRouteServerResult> Invoke(GetRouteServerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteServerResult>("aws-native:ec2:getRouteServer", args ?? new GetRouteServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteServerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Route Server.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetRouteServerArgs()
        {
        }
        public static new GetRouteServerArgs Empty => new GetRouteServerArgs();
    }

    public sealed class GetRouteServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Route Server.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetRouteServerInvokeArgs()
        {
        }
        public static new GetRouteServerInvokeArgs Empty => new GetRouteServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouteServerResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Route Server.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The ID of the Route Server.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Whether to enable persistent routes
        /// </summary>
        public readonly Pulumi.AwsNative.Ec2.RouteServerPersistRoutes? PersistRoutes;
        /// <summary>
        /// Whether to enable SNS notifications
        /// </summary>
        public readonly bool? SnsNotificationsEnabled;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetRouteServerResult(
            string? arn,

            string? id,

            Pulumi.AwsNative.Ec2.RouteServerPersistRoutes? persistRoutes,

            bool? snsNotificationsEnabled,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            Id = id;
            PersistRoutes = persistRoutes;
            SnsNotificationsEnabled = snsNotificationsEnabled;
            Tags = tags;
        }
    }
}
