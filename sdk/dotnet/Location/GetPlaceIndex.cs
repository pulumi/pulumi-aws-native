// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Location
{
    public static class GetPlaceIndex
    {
        /// <summary>
        /// Definition of AWS::Location::PlaceIndex Resource Type
        /// </summary>
        public static Task<GetPlaceIndexResult> InvokeAsync(GetPlaceIndexArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPlaceIndexResult>("aws-native:location:getPlaceIndex", args ?? new GetPlaceIndexArgs(), options.WithDefaults());

        /// <summary>
        /// Definition of AWS::Location::PlaceIndex Resource Type
        /// </summary>
        public static Output<GetPlaceIndexResult> Invoke(GetPlaceIndexInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlaceIndexResult>("aws-native:location:getPlaceIndex", args ?? new GetPlaceIndexInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPlaceIndexArgs : global::Pulumi.InvokeArgs
    {
        [Input("indexName", required: true)]
        public string IndexName { get; set; } = null!;

        public GetPlaceIndexArgs()
        {
        }
        public static new GetPlaceIndexArgs Empty => new GetPlaceIndexArgs();
    }

    public sealed class GetPlaceIndexInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("indexName", required: true)]
        public Input<string> IndexName { get; set; } = null!;

        public GetPlaceIndexInvokeArgs()
        {
        }
        public static new GetPlaceIndexInvokeArgs Empty => new GetPlaceIndexInvokeArgs();
    }


    [OutputType]
    public sealed class GetPlaceIndexResult
    {
        public readonly string? Arn;
        public readonly string? CreateTime;
        public readonly string? IndexArn;
        public readonly string? UpdateTime;

        [OutputConstructor]
        private GetPlaceIndexResult(
            string? arn,

            string? createTime,

            string? indexArn,

            string? updateTime)
        {
            Arn = arn;
            CreateTime = createTime;
            IndexArn = indexArn;
            UpdateTime = updateTime;
        }
    }
}