// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class CrawlerDeltaTarget
    {
        public readonly string? ConnectionName;
        public readonly bool? CreateNativeDeltaTable;
        public readonly ImmutableArray<string> DeltaTables;
        public readonly bool? WriteManifest;

        [OutputConstructor]
        private CrawlerDeltaTarget(
            string? connectionName,

            bool? createNativeDeltaTable,

            ImmutableArray<string> deltaTables,

            bool? writeManifest)
        {
            ConnectionName = connectionName;
            CreateNativeDeltaTable = createNativeDeltaTable;
            DeltaTables = deltaTables;
            WriteManifest = writeManifest;
        }
    }
}