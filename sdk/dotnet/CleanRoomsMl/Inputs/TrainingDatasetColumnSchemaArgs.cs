// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRoomsMl.Inputs
{

    public sealed class TrainingDatasetColumnSchemaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of a column.
        /// </summary>
        [Input("columnName", required: true)]
        public Input<string> ColumnName { get; set; } = null!;

        [Input("columnTypes", required: true)]
        private InputList<Pulumi.AwsNative.CleanRoomsMl.TrainingDatasetColumnType>? _columnTypes;

        /// <summary>
        /// The data type of column.
        /// </summary>
        public InputList<Pulumi.AwsNative.CleanRoomsMl.TrainingDatasetColumnType> ColumnTypes
        {
            get => _columnTypes ?? (_columnTypes = new InputList<Pulumi.AwsNative.CleanRoomsMl.TrainingDatasetColumnType>());
            set => _columnTypes = value;
        }

        public TrainingDatasetColumnSchemaArgs()
        {
        }
        public static new TrainingDatasetColumnSchemaArgs Empty => new TrainingDatasetColumnSchemaArgs();
    }
}
