// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Outputs
{

    /// <summary>
    /// Datasource arn and description
    /// </summary>
    [OutputType]
    public sealed class ApplicationDataSource
    {
        /// <summary>
        /// The ARN of the data source.
        /// </summary>
        public readonly object DataSourceArn;
        /// <summary>
        /// Description of the data source.
        /// </summary>
        public readonly string? DataSourceDescription;

        [OutputConstructor]
        private ApplicationDataSource(
            object dataSourceArn,

            string? dataSourceDescription)
        {
            DataSourceArn = dataSourceArn;
            DataSourceDescription = dataSourceDescription;
        }
    }
}
