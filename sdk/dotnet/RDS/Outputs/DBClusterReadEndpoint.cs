// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS.Outputs
{

    [OutputType]
    public sealed class DBClusterReadEndpoint
    {
        /// <summary>
        /// The reader endpoint for the DB cluster.
        /// </summary>
        public readonly string? Address;

        [OutputConstructor]
        private DBClusterReadEndpoint(string? address)
        {
            Address = address;
        }
    }
}