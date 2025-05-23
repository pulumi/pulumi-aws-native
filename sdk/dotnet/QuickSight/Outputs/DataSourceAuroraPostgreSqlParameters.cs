// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;Parameters for Amazon Aurora PostgreSQL-Compatible Edition.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSourceAuroraPostgreSqlParameters
    {
        /// <summary>
        /// &lt;p&gt;The Amazon Aurora PostgreSQL database to connect to.&lt;/p&gt;
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// &lt;p&gt;The Amazon Aurora PostgreSQL-Compatible host to connect to.&lt;/p&gt;
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// &lt;p&gt;The port that Amazon Aurora PostgreSQL is listening on.&lt;/p&gt;
        /// </summary>
        public readonly double Port;

        [OutputConstructor]
        private DataSourceAuroraPostgreSqlParameters(
            string database,

            string host,

            double port)
        {
            Database = database;
            Host = host;
            Port = port;
        }
    }
}
