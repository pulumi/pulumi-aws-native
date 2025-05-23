// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise
{
    public static class GetDataset
    {
        /// <summary>
        /// Resource schema for AWS::IoTSiteWise::Dataset.
        /// </summary>
        public static Task<GetDatasetResult> InvokeAsync(GetDatasetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatasetResult>("aws-native:iotsitewise:getDataset", args ?? new GetDatasetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::IoTSiteWise::Dataset.
        /// </summary>
        public static Output<GetDatasetResult> Invoke(GetDatasetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatasetResult>("aws-native:iotsitewise:getDataset", args ?? new GetDatasetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::IoTSiteWise::Dataset.
        /// </summary>
        public static Output<GetDatasetResult> Invoke(GetDatasetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatasetResult>("aws-native:iotsitewise:getDataset", args ?? new GetDatasetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatasetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the dataset.
        /// </summary>
        [Input("datasetId", required: true)]
        public string DatasetId { get; set; } = null!;

        public GetDatasetArgs()
        {
        }
        public static new GetDatasetArgs Empty => new GetDatasetArgs();
    }

    public sealed class GetDatasetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the dataset.
        /// </summary>
        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        public GetDatasetInvokeArgs()
        {
        }
        public static new GetDatasetInvokeArgs Empty => new GetDatasetInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatasetResult
    {
        /// <summary>
        /// The ARN of the dataset.
        /// </summary>
        public readonly string? DatasetArn;
        /// <summary>
        /// A description about the dataset, and its functionality.
        /// </summary>
        public readonly string? DatasetDescription;
        /// <summary>
        /// The ID of the dataset.
        /// </summary>
        public readonly string? DatasetId;
        /// <summary>
        /// The name of the dataset.
        /// </summary>
        public readonly string? DatasetName;
        /// <summary>
        /// The data source for the dataset.
        /// </summary>
        public readonly Outputs.DatasetSource? DatasetSource;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetDatasetResult(
            string? datasetArn,

            string? datasetDescription,

            string? datasetId,

            string? datasetName,

            Outputs.DatasetSource? datasetSource,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            DatasetArn = datasetArn;
            DatasetDescription = datasetDescription;
            DatasetId = datasetId;
            DatasetName = datasetName;
            DatasetSource = datasetSource;
            Tags = tags;
        }
    }
}
