// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sdc
{
    public static class GetDeployment
    {
        /// <summary>
        /// Resource Type definition for AMZN::SDC::Deployment
        /// </summary>
        public static Task<GetDeploymentResult> InvokeAsync(GetDeploymentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentResult>("aws-native:sdc:getDeployment", args ?? new GetDeploymentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AMZN::SDC::Deployment
        /// </summary>
        public static Output<GetDeploymentResult> Invoke(GetDeploymentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentResult>("aws-native:sdc:getDeployment", args ?? new GetDeploymentInvokeArgs(), options.WithDefaults());
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
        public readonly string? ConfigName;
        public readonly string? Dimension;
        public readonly string? Id;
        public readonly string? PipelineId;
        public readonly string? S3Bucket;
        public readonly string? Stage;
        public readonly string? TargetRegionOverride;

        [OutputConstructor]
        private GetDeploymentResult(
            string? configName,

            string? dimension,

            string? id,

            string? pipelineId,

            string? s3Bucket,

            string? stage,

            string? targetRegionOverride)
        {
            ConfigName = configName;
            Dimension = dimension;
            Id = id;
            PipelineId = pipelineId;
            S3Bucket = s3Bucket;
            Stage = stage;
            TargetRegionOverride = targetRegionOverride;
        }
    }
}