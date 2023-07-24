// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MSK
{
    public static class GetVpcConnection
    {
        /// <summary>
        /// Resource Type definition for AWS::MSK::VpcConnection
        /// </summary>
        public static Task<GetVpcConnectionResult> InvokeAsync(GetVpcConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcConnectionResult>("aws-native:msk:getVpcConnection", args ?? new GetVpcConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::MSK::VpcConnection
        /// </summary>
        public static Output<GetVpcConnectionResult> Invoke(GetVpcConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcConnectionResult>("aws-native:msk:getVpcConnection", args ?? new GetVpcConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcConnectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetVpcConnectionArgs()
        {
        }
        public static new GetVpcConnectionArgs Empty => new GetVpcConnectionArgs();
    }

    public sealed class GetVpcConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetVpcConnectionInvokeArgs()
        {
        }
        public static new GetVpcConnectionInvokeArgs Empty => new GetVpcConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcConnectionResult
    {
        public readonly string? Arn;
        public readonly Outputs.VpcConnectionTags? Tags;

        [OutputConstructor]
        private GetVpcConnectionResult(
            string? arn,

            Outputs.VpcConnectionTags? tags)
        {
            Arn = arn;
            Tags = tags;
        }
    }
}