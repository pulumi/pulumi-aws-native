// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceGoogleDriveConfiguration
    {
        public readonly ImmutableArray<string> ExcludeMimeTypes;
        public readonly ImmutableArray<string> ExcludeSharedDrives;
        public readonly ImmutableArray<string> ExcludeUserAccounts;
        public readonly ImmutableArray<string> ExclusionPatterns;
        public readonly ImmutableArray<Outputs.DataSourceToIndexFieldMapping> FieldMappings;
        public readonly ImmutableArray<string> InclusionPatterns;
        public readonly string SecretArn;

        [OutputConstructor]
        private DataSourceGoogleDriveConfiguration(
            ImmutableArray<string> excludeMimeTypes,

            ImmutableArray<string> excludeSharedDrives,

            ImmutableArray<string> excludeUserAccounts,

            ImmutableArray<string> exclusionPatterns,

            ImmutableArray<Outputs.DataSourceToIndexFieldMapping> fieldMappings,

            ImmutableArray<string> inclusionPatterns,

            string secretArn)
        {
            ExcludeMimeTypes = excludeMimeTypes;
            ExcludeSharedDrives = excludeSharedDrives;
            ExcludeUserAccounts = excludeUserAccounts;
            ExclusionPatterns = exclusionPatterns;
            FieldMappings = fieldMappings;
            InclusionPatterns = inclusionPatterns;
            SecretArn = secretArn;
        }
    }
}