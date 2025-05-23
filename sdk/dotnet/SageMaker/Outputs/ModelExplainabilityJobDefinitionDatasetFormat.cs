// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// The dataset format of the data to monitor
    /// </summary>
    [OutputType]
    public sealed class ModelExplainabilityJobDefinitionDatasetFormat
    {
        public readonly Outputs.ModelExplainabilityJobDefinitionCsv? Csv;
        public readonly Outputs.ModelExplainabilityJobDefinitionJson? Json;
        public readonly bool? Parquet;

        [OutputConstructor]
        private ModelExplainabilityJobDefinitionDatasetFormat(
            Outputs.ModelExplainabilityJobDefinitionCsv? csv,

            Outputs.ModelExplainabilityJobDefinitionJson? json,

            bool? parquet)
        {
            Csv = csv;
            Json = json;
            Parquet = parquet;
        }
    }
}
