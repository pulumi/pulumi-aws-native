// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Outputs
{

    [OutputType]
    public sealed class DataSourceGlueRunConfigurationInput
    {
        /// <summary>
        /// Specifies whether to automatically import data quality metrics as part of the data source run.
        /// </summary>
        public readonly bool? AutoImportDataQualityResult;
        /// <summary>
        /// The catalog name in the AWS Glue run configuration.
        /// </summary>
        public readonly string? CatalogName;
        /// <summary>
        /// The data access role included in the configuration details of the AWS Glue data source.
        /// </summary>
        public readonly string? DataAccessRole;
        /// <summary>
        /// The relational filter configurations included in the configuration details of the AWS Glue data source.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceRelationalFilterConfiguration> RelationalFilterConfigurations;

        [OutputConstructor]
        private DataSourceGlueRunConfigurationInput(
            bool? autoImportDataQualityResult,

            string? catalogName,

            string? dataAccessRole,

            ImmutableArray<Outputs.DataSourceRelationalFilterConfiguration> relationalFilterConfigurations)
        {
            AutoImportDataQualityResult = autoImportDataQualityResult;
            CatalogName = catalogName;
            DataAccessRole = dataAccessRole;
            RelationalFilterConfigurations = relationalFilterConfigurations;
        }
    }
}
