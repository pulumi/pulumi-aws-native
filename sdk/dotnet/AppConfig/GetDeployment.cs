// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppConfig
{
    public static class GetDeployment
    {
        /// <summary>
        /// Resource Type definition for AWS::AppConfig::Deployment
        /// </summary>
        public static Task<GetDeploymentResult> InvokeAsync(GetDeploymentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentResult>("aws-native:appconfig:getDeployment", args ?? new GetDeploymentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AppConfig::Deployment
        /// </summary>
        public static Output<GetDeploymentResult> Invoke(GetDeploymentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentResult>("aws-native:appconfig:getDeployment", args ?? new GetDeploymentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeploymentArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDeploymentArgs()
        {
        }
        public static new GetDeploymentArgs Empty => new GetDeploymentArgs();
    }

    public sealed class GetDeploymentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDeploymentInvokeArgs()
        {
        }
        public static new GetDeploymentInvokeArgs Empty => new GetDeploymentInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeploymentResult
    {
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.DeploymentTags> Tags;

        [OutputConstructor]
        private GetDeploymentResult(
            string? id,

            ImmutableArray<Outputs.DeploymentTags> tags)
        {
            Id = id;
            Tags = tags;
        }
    }
}