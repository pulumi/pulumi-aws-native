// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cassandra.Outputs
{

    [OutputType]
    public sealed class KeyspaceReplicationSpecification
    {
        public readonly ImmutableArray<Pulumi.AwsNative.Cassandra.KeyspaceRegionListItem> RegionList;
        public readonly Pulumi.AwsNative.Cassandra.KeyspaceReplicationSpecificationReplicationStrategy? ReplicationStrategy;

        [OutputConstructor]
        private KeyspaceReplicationSpecification(
            ImmutableArray<Pulumi.AwsNative.Cassandra.KeyspaceRegionListItem> regionList,

            Pulumi.AwsNative.Cassandra.KeyspaceReplicationSpecificationReplicationStrategy? replicationStrategy)
        {
            RegionList = regionList;
            ReplicationStrategy = replicationStrategy;
        }
    }
}