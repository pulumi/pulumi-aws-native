// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Inputs
{

    public sealed class DatasetTriggeringDatasetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the data set whose content generation triggers the new data set content generation.
        /// </summary>
        [Input("datasetName", required: true)]
        public Input<string> DatasetName { get; set; } = null!;

        public DatasetTriggeringDatasetArgs()
        {
        }
        public static new DatasetTriggeringDatasetArgs Empty => new DatasetTriggeringDatasetArgs();
    }
}
