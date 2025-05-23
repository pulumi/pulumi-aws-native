// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Omics
{
    public static class GetSequenceStore
    {
        /// <summary>
        /// Resource Type definition for AWS::Omics::SequenceStore
        /// </summary>
        public static Task<GetSequenceStoreResult> InvokeAsync(GetSequenceStoreArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSequenceStoreResult>("aws-native:omics:getSequenceStore", args ?? new GetSequenceStoreArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Omics::SequenceStore
        /// </summary>
        public static Output<GetSequenceStoreResult> Invoke(GetSequenceStoreInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSequenceStoreResult>("aws-native:omics:getSequenceStore", args ?? new GetSequenceStoreInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Omics::SequenceStore
        /// </summary>
        public static Output<GetSequenceStoreResult> Invoke(GetSequenceStoreInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSequenceStoreResult>("aws-native:omics:getSequenceStore", args ?? new GetSequenceStoreInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSequenceStoreArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The store's ID.
        /// </summary>
        [Input("sequenceStoreId", required: true)]
        public string SequenceStoreId { get; set; } = null!;

        public GetSequenceStoreArgs()
        {
        }
        public static new GetSequenceStoreArgs Empty => new GetSequenceStoreArgs();
    }

    public sealed class GetSequenceStoreInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The store's ID.
        /// </summary>
        [Input("sequenceStoreId", required: true)]
        public Input<string> SequenceStoreId { get; set; } = null!;

        public GetSequenceStoreInvokeArgs()
        {
        }
        public static new GetSequenceStoreInvokeArgs Empty => new GetSequenceStoreInvokeArgs();
    }


    [OutputType]
    public sealed class GetSequenceStoreResult
    {
        /// <summary>
        /// Location of the access logs.
        /// </summary>
        public readonly string? AccessLogLocation;
        /// <summary>
        /// The store's ARN.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// When the store was created.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// A description for the store.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// An S3 location that is used to store files that have failed a direct upload.
        /// </summary>
        public readonly string? FallbackLocation;
        /// <summary>
        /// A name for the store.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The tags keys to propagate to the S3 objects associated with read sets in the sequence store.
        /// </summary>
        public readonly ImmutableArray<string> PropagatedSetLevelTags;
        /// <summary>
        /// This is ARN of the access point associated with the S3 bucket storing read sets.
        /// </summary>
        public readonly string? S3AccessPointArn;
        /// <summary>
        /// The resource policy that controls S3 access on the store
        /// 
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Omics::SequenceStore` for more information about the expected schema for this property.
        /// </summary>
        public readonly object? S3AccessPolicy;
        /// <summary>
        /// The S3 URI of the sequence store.
        /// </summary>
        public readonly string? S3Uri;
        /// <summary>
        /// The store's ID.
        /// </summary>
        public readonly string? SequenceStoreId;
        /// <summary>
        /// Status of the sequence store.
        /// </summary>
        public readonly Pulumi.AwsNative.Omics.SequenceStoreStatus? Status;
        /// <summary>
        /// The status message of the sequence store.
        /// </summary>
        public readonly string? StatusMessage;
        /// <summary>
        /// Tags for the store.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The last-updated time of the sequence store.
        /// </summary>
        public readonly string? UpdateTime;

        [OutputConstructor]
        private GetSequenceStoreResult(
            string? accessLogLocation,

            string? arn,

            string? creationTime,

            string? description,

            string? fallbackLocation,

            string? name,

            ImmutableArray<string> propagatedSetLevelTags,

            string? s3AccessPointArn,

            object? s3AccessPolicy,

            string? s3Uri,

            string? sequenceStoreId,

            Pulumi.AwsNative.Omics.SequenceStoreStatus? status,

            string? statusMessage,

            ImmutableDictionary<string, string>? tags,

            string? updateTime)
        {
            AccessLogLocation = accessLogLocation;
            Arn = arn;
            CreationTime = creationTime;
            Description = description;
            FallbackLocation = fallbackLocation;
            Name = name;
            PropagatedSetLevelTags = propagatedSetLevelTags;
            S3AccessPointArn = s3AccessPointArn;
            S3AccessPolicy = s3AccessPolicy;
            S3Uri = s3Uri;
            SequenceStoreId = sequenceStoreId;
            Status = status;
            StatusMessage = statusMessage;
            Tags = tags;
            UpdateTime = updateTime;
        }
    }
}
