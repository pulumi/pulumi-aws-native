// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRoomsMl.Outputs
{

    [OutputType]
    public sealed class TrainingDatasetDataset
    {
        /// <summary>
        /// A DatasetInputConfig object that defines the data source and schema mapping.
        /// </summary>
        public readonly Outputs.TrainingDatasetDatasetInputConfig InputConfig;
        /// <summary>
        /// What type of information is found in the dataset.
        /// </summary>
        public readonly Pulumi.AwsNative.CleanRoomsMl.TrainingDatasetDatasetType Type;

        [OutputConstructor]
        private TrainingDatasetDataset(
            Outputs.TrainingDatasetDatasetInputConfig inputConfig,

            Pulumi.AwsNative.CleanRoomsMl.TrainingDatasetDatasetType type)
        {
            InputConfig = inputConfig;
            Type = type;
        }
    }
}
