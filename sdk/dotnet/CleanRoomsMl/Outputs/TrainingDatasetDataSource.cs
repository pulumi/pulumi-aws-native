// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRoomsMl.Outputs
{

    [OutputType]
    public sealed class TrainingDatasetDataSource
    {
        /// <summary>
        /// A GlueDataSource object that defines the catalog ID, database name, and table name for the training data.
        /// </summary>
        public readonly Outputs.TrainingDatasetGlueDataSource GlueDataSource;

        [OutputConstructor]
        private TrainingDatasetDataSource(Outputs.TrainingDatasetGlueDataSource glueDataSource)
        {
            GlueDataSource = glueDataSource;
        }
    }
}