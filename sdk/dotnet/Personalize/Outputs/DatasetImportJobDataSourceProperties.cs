// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize.Outputs
{

    /// <summary>
    /// The Amazon S3 bucket that contains the training data to import.
    /// </summary>
    [OutputType]
    public sealed class DatasetImportJobDataSourceProperties
    {
        /// <summary>
        /// The path to the Amazon S3 bucket where the data that you want to upload to your dataset is stored.
        /// </summary>
        public readonly string? DataLocation;

        [OutputConstructor]
        private DatasetImportJobDataSourceProperties(string? dataLocation)
        {
            DataLocation = dataLocation;
        }
    }
}
