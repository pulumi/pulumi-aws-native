// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Outputs
{

    /// <summary>
    /// The name of an Amazon Redshift cluster.
    /// </summary>
    [OutputType]
    public sealed class DataSourceRedshiftClusterStorage
    {
        /// <summary>
        /// The name of an Amazon Redshift cluster.
        /// </summary>
        public readonly string ClusterName;

        [OutputConstructor]
        private DataSourceRedshiftClusterStorage(string clusterName)
        {
            ClusterName = clusterName;
        }
    }
}