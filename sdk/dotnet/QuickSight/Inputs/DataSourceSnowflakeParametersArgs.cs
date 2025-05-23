// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;The parameters for Snowflake.&lt;/p&gt;
    /// </summary>
    public sealed class DataSourceSnowflakeParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication type that you want to use for your connection. This parameter accepts OAuth and non-OAuth authentication types.
        /// </summary>
        [Input("authenticationType")]
        public Input<Pulumi.AwsNative.QuickSight.DataSourceAuthenticationType>? AuthenticationType { get; set; }

        /// <summary>
        /// &lt;p&gt;Database.&lt;/p&gt;
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The database access control role.
        /// </summary>
        [Input("databaseAccessControlRole")]
        public Input<string>? DatabaseAccessControlRole { get; set; }

        /// <summary>
        /// &lt;p&gt;Host.&lt;/p&gt;
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// An object that contains information needed to create a data source connection between an Amazon QuickSight account and Snowflake.
        /// </summary>
        [Input("oAuthParameters")]
        public Input<Inputs.DataSourceOAuthParametersArgs>? OAuthParameters { get; set; }

        /// <summary>
        /// &lt;p&gt;Warehouse.&lt;/p&gt;
        /// </summary>
        [Input("warehouse", required: true)]
        public Input<string> Warehouse { get; set; } = null!;

        public DataSourceSnowflakeParametersArgs()
        {
        }
        public static new DataSourceSnowflakeParametersArgs Empty => new DataSourceSnowflakeParametersArgs();
    }
}
