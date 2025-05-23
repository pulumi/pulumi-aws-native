// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms.Outputs
{

    [OutputType]
    public sealed class DataMigrationSettings
    {
        /// <summary>
        /// The property specifies whether to enable the Cloudwatch log.
        /// </summary>
        public readonly bool? CloudwatchLogsEnabled;
        /// <summary>
        /// The number of parallel jobs that trigger parallel threads to unload the tables from the source, and then load them to the target.
        /// </summary>
        public readonly int? NumberOfJobs;
        /// <summary>
        /// The property specifies the rules of selecting objects for data migration.
        /// </summary>
        public readonly string? SelectionRules;

        [OutputConstructor]
        private DataMigrationSettings(
            bool? cloudwatchLogsEnabled,

            int? numberOfJobs,

            string? selectionRules)
        {
            CloudwatchLogsEnabled = cloudwatchLogsEnabled;
            NumberOfJobs = numberOfJobs;
            SelectionRules = selectionRules;
        }
    }
}
