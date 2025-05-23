// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3
{
    public static class GetMultiRegionAccessPoint
    {
        /// <summary>
        /// AWS::S3::MultiRegionAccessPoint is an Amazon S3 resource type that dynamically routes S3 requests to easily satisfy geographic compliance requirements based on customer-defined routing policies.
        /// </summary>
        public static Task<GetMultiRegionAccessPointResult> InvokeAsync(GetMultiRegionAccessPointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMultiRegionAccessPointResult>("aws-native:s3:getMultiRegionAccessPoint", args ?? new GetMultiRegionAccessPointArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::S3::MultiRegionAccessPoint is an Amazon S3 resource type that dynamically routes S3 requests to easily satisfy geographic compliance requirements based on customer-defined routing policies.
        /// </summary>
        public static Output<GetMultiRegionAccessPointResult> Invoke(GetMultiRegionAccessPointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMultiRegionAccessPointResult>("aws-native:s3:getMultiRegionAccessPoint", args ?? new GetMultiRegionAccessPointInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// AWS::S3::MultiRegionAccessPoint is an Amazon S3 resource type that dynamically routes S3 requests to easily satisfy geographic compliance requirements based on customer-defined routing policies.
        /// </summary>
        public static Output<GetMultiRegionAccessPointResult> Invoke(GetMultiRegionAccessPointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetMultiRegionAccessPointResult>("aws-native:s3:getMultiRegionAccessPoint", args ?? new GetMultiRegionAccessPointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMultiRegionAccessPointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name you want to assign to this Multi Region Access Point.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetMultiRegionAccessPointArgs()
        {
        }
        public static new GetMultiRegionAccessPointArgs Empty => new GetMultiRegionAccessPointArgs();
    }

    public sealed class GetMultiRegionAccessPointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name you want to assign to this Multi Region Access Point.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetMultiRegionAccessPointInvokeArgs()
        {
        }
        public static new GetMultiRegionAccessPointInvokeArgs Empty => new GetMultiRegionAccessPointInvokeArgs();
    }


    [OutputType]
    public sealed class GetMultiRegionAccessPointResult
    {
        /// <summary>
        /// The alias is a unique identifier to, and is part of the public DNS name for this Multi Region Access Point
        /// </summary>
        public readonly string? Alias;
        /// <summary>
        /// The timestamp of the when the Multi Region Access Point is created
        /// </summary>
        public readonly string? CreatedAt;

        [OutputConstructor]
        private GetMultiRegionAccessPointResult(
            string? alias,

            string? createdAt)
        {
            Alias = alias;
            CreatedAt = createdAt;
        }
    }
}
