// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew
{
    public static class GetDataset
    {
        /// <summary>
        /// Resource schema for AWS::DataBrew::Dataset.
        /// </summary>
        public static Task<GetDatasetResult> InvokeAsync(GetDatasetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatasetResult>("aws-native:databrew:getDataset", args ?? new GetDatasetArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataBrew::Dataset.
        /// </summary>
        public static Output<GetDatasetResult> Invoke(GetDatasetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatasetResult>("aws-native:databrew:getDataset", args ?? new GetDatasetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatasetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Dataset name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDatasetArgs()
        {
        }
        public static new GetDatasetArgs Empty => new GetDatasetArgs();
    }

    public sealed class GetDatasetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Dataset name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDatasetInvokeArgs()
        {
        }
        public static new GetDatasetInvokeArgs Empty => new GetDatasetInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatasetResult
    {
        /// <summary>
        /// Dataset format
        /// </summary>
        public readonly Pulumi.AwsNative.DataBrew.DatasetFormat? Format;
        /// <summary>
        /// Format options for dataset
        /// </summary>
        public readonly Outputs.DatasetFormatOptions? FormatOptions;
        /// <summary>
        /// Input
        /// </summary>
        public readonly Outputs.DatasetInput? Input;
        /// <summary>
        /// PathOptions
        /// </summary>
        public readonly Outputs.DatasetPathOptions? PathOptions;

        [OutputConstructor]
        private GetDatasetResult(
            Pulumi.AwsNative.DataBrew.DatasetFormat? format,

            Outputs.DatasetFormatOptions? formatOptions,

            Outputs.DatasetInput? input,

            Outputs.DatasetPathOptions? pathOptions)
        {
            Format = format;
            FormatOptions = formatOptions;
            Input = input;
            PathOptions = pathOptions;
        }
    }
}