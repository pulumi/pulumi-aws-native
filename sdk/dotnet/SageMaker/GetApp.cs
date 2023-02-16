// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    public static class GetApp
    {
        /// <summary>
        /// Resource Type definition for AWS::SageMaker::App
        /// </summary>
        public static Task<GetAppResult> InvokeAsync(GetAppArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppResult>("aws-native:sagemaker:getApp", args ?? new GetAppArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SageMaker::App
        /// </summary>
        public static Output<GetAppResult> Invoke(GetAppInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppResult>("aws-native:sagemaker:getApp", args ?? new GetAppInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the app.
        /// </summary>
        [Input("appName", required: true)]
        public string AppName { get; set; } = null!;

        /// <summary>
        /// The type of app.
        /// </summary>
        [Input("appType", required: true)]
        public Pulumi.AwsNative.SageMaker.AppType AppType { get; set; }

        /// <summary>
        /// The domain ID.
        /// </summary>
        [Input("domainId", required: true)]
        public string DomainId { get; set; } = null!;

        /// <summary>
        /// The user profile name.
        /// </summary>
        [Input("userProfileName", required: true)]
        public string UserProfileName { get; set; } = null!;

        public GetAppArgs()
        {
        }
        public static new GetAppArgs Empty => new GetAppArgs();
    }

    public sealed class GetAppInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the app.
        /// </summary>
        [Input("appName", required: true)]
        public Input<string> AppName { get; set; } = null!;

        /// <summary>
        /// The type of app.
        /// </summary>
        [Input("appType", required: true)]
        public Input<Pulumi.AwsNative.SageMaker.AppType> AppType { get; set; } = null!;

        /// <summary>
        /// The domain ID.
        /// </summary>
        [Input("domainId", required: true)]
        public Input<string> DomainId { get; set; } = null!;

        /// <summary>
        /// The user profile name.
        /// </summary>
        [Input("userProfileName", required: true)]
        public Input<string> UserProfileName { get; set; } = null!;

        public GetAppInvokeArgs()
        {
        }
        public static new GetAppInvokeArgs Empty => new GetAppInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the app.
        /// </summary>
        public readonly string? AppArn;
        /// <summary>
        /// The instance type and the Amazon Resource Name (ARN) of the SageMaker image created on the instance.
        /// </summary>
        public readonly Outputs.AppResourceSpec? ResourceSpec;

        [OutputConstructor]
        private GetAppResult(
            string? appArn,

            Outputs.AppResourceSpec? resourceSpec)
        {
            AppArn = appArn;
            ResourceSpec = resourceSpec;
        }
    }
}