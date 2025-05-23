// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetMlflowTrackingServer
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::MlflowTrackingServer
        /// </summary>
        public static Task<GetMlflowTrackingServerResult> InvokeAsync(GetMlflowTrackingServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMlflowTrackingServerResult>("aws-native:sagemaker:getMlflowTrackingServer", args ?? new GetMlflowTrackingServerArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::MlflowTrackingServer
        /// </summary>
        public static Output<GetMlflowTrackingServerResult> Invoke(GetMlflowTrackingServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMlflowTrackingServerResult>("aws-native:sagemaker:getMlflowTrackingServer", args ?? new GetMlflowTrackingServerInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::MlflowTrackingServer
        /// </summary>
        public static Output<GetMlflowTrackingServerResult> Invoke(GetMlflowTrackingServerInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMlflowTrackingServerResult>("aws-native:sagemaker:getMlflowTrackingServer", args ?? new GetMlflowTrackingServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMlflowTrackingServerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the MLFlow Tracking Server.
        /// </summary>
        [Input("trackingServerName", required: true)]
        public string TrackingServerName { get; set; } = null!;

        public GetMlflowTrackingServerArgs()
        {
        }
        public static new GetMlflowTrackingServerArgs Empty => new GetMlflowTrackingServerArgs();
    }

    public sealed class GetMlflowTrackingServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the MLFlow Tracking Server.
        /// </summary>
        [Input("trackingServerName", required: true)]
        public Input<string> TrackingServerName { get; set; } = null!;

        public GetMlflowTrackingServerInvokeArgs()
        {
        }
        public static new GetMlflowTrackingServerInvokeArgs Empty => new GetMlflowTrackingServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetMlflowTrackingServerResult
    {
        /// <summary>
        /// The Amazon S3 URI for MLFlow Tracking Server artifacts.
        /// </summary>
        public readonly string? ArtifactStoreUri;
        /// <summary>
        /// A flag to enable Automatic SageMaker Model Registration.
        /// </summary>
        public readonly bool? AutomaticModelRegistration;
        /// <summary>
        /// The MLFlow Version used on the MLFlow Tracking Server.
        /// </summary>
        public readonly string? MlflowVersion;
        /// <summary>
        /// The Amazon Resource Name (ARN) of an IAM role that enables Amazon SageMaker to perform tasks on behalf of the customer.
        /// </summary>
        public readonly string? RoleArn;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the MLFlow Tracking Server.
        /// </summary>
        public readonly string? TrackingServerArn;
        /// <summary>
        /// The size of the MLFlow Tracking Server.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.MlflowTrackingServerTrackingServerSize? TrackingServerSize;
        /// <summary>
        /// The start of the time window for maintenance of the MLFlow Tracking Server in UTC time.
        /// </summary>
        public readonly string? WeeklyMaintenanceWindowStart;

        [OutputConstructor]
        private GetMlflowTrackingServerResult(
            string? artifactStoreUri,

            bool? automaticModelRegistration,

            string? mlflowVersion,

            string? roleArn,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? trackingServerArn,

            Pulumi.AwsNative.SageMaker.MlflowTrackingServerTrackingServerSize? trackingServerSize,

            string? weeklyMaintenanceWindowStart)
        {
            ArtifactStoreUri = artifactStoreUri;
            AutomaticModelRegistration = automaticModelRegistration;
            MlflowVersion = mlflowVersion;
            RoleArn = roleArn;
            Tags = tags;
            TrackingServerArn = trackingServerArn;
            TrackingServerSize = trackingServerSize;
            WeeklyMaintenanceWindowStart = weeklyMaintenanceWindowStart;
        }
    }
}
