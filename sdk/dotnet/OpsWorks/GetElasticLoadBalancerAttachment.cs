// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks
{
    public static class GetElasticLoadBalancerAttachment
    {
        /// <summary>
        /// Resource Type definition for AWS::OpsWorks::ElasticLoadBalancerAttachment
        /// </summary>
        public static Task<GetElasticLoadBalancerAttachmentResult> InvokeAsync(GetElasticLoadBalancerAttachmentArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetElasticLoadBalancerAttachmentResult>("aws-native:opsworks:getElasticLoadBalancerAttachment", args ?? new GetElasticLoadBalancerAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::OpsWorks::ElasticLoadBalancerAttachment
        /// </summary>
        public static Output<GetElasticLoadBalancerAttachmentResult> Invoke(GetElasticLoadBalancerAttachmentInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetElasticLoadBalancerAttachmentResult>("aws-native:opsworks:getElasticLoadBalancerAttachment", args ?? new GetElasticLoadBalancerAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetElasticLoadBalancerAttachmentArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetElasticLoadBalancerAttachmentArgs()
        {
        }
        public static new GetElasticLoadBalancerAttachmentArgs Empty => new GetElasticLoadBalancerAttachmentArgs();
    }

    public sealed class GetElasticLoadBalancerAttachmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetElasticLoadBalancerAttachmentInvokeArgs()
        {
        }
        public static new GetElasticLoadBalancerAttachmentInvokeArgs Empty => new GetElasticLoadBalancerAttachmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetElasticLoadBalancerAttachmentResult
    {
        public readonly string? ElasticLoadBalancerName;
        public readonly string? Id;
        public readonly string? LayerId;

        [OutputConstructor]
        private GetElasticLoadBalancerAttachmentResult(
            string? elasticLoadBalancerName,

            string? id,

            string? layerId)
        {
            ElasticLoadBalancerName = elasticLoadBalancerName;
            Id = id;
            LayerId = layerId;
        }
    }
}