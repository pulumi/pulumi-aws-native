// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms.Inputs
{

    /// <summary>
    /// The property identifies the exact type of settings for the data provider.
    /// </summary>
    public sealed class SettingsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// DocDbSettings property identifier.
        /// </summary>
        [Input("docDbSettings")]
        public Input<Inputs.SettingsPropertiesDocDbSettingsPropertiesArgs>? DocDbSettings { get; set; }

        /// <summary>
        /// IbmDb2LuwSettings property identifier.
        /// </summary>
        [Input("ibmDb2LuwSettings")]
        public Input<Inputs.SettingsPropertiesIbmDb2LuwSettingsPropertiesArgs>? IbmDb2LuwSettings { get; set; }

        /// <summary>
        /// IbmDb2zOsSettings property identifier.
        /// </summary>
        [Input("ibmDb2zOsSettings")]
        public Input<Inputs.SettingsPropertiesIbmDb2zOsSettingsPropertiesArgs>? IbmDb2zOsSettings { get; set; }

        /// <summary>
        /// MariaDbSettings property identifier.
        /// </summary>
        [Input("mariaDbSettings")]
        public Input<Inputs.SettingsPropertiesMariaDbSettingsPropertiesArgs>? MariaDbSettings { get; set; }

        /// <summary>
        /// MicrosoftSqlServerSettings property identifier.
        /// </summary>
        [Input("microsoftSqlServerSettings")]
        public Input<Inputs.SettingsPropertiesMicrosoftSqlServerSettingsPropertiesArgs>? MicrosoftSqlServerSettings { get; set; }

        /// <summary>
        /// MongoDbSettings property identifier.
        /// </summary>
        [Input("mongoDbSettings")]
        public Input<Inputs.SettingsPropertiesMongoDbSettingsPropertiesArgs>? MongoDbSettings { get; set; }

        /// <summary>
        /// MySqlSettings property identifier.
        /// </summary>
        [Input("mySqlSettings")]
        public Input<Inputs.SettingsPropertiesMySqlSettingsPropertiesArgs>? MySqlSettings { get; set; }

        /// <summary>
        /// OracleSettings property identifier.
        /// </summary>
        [Input("oracleSettings")]
        public Input<Inputs.SettingsPropertiesOracleSettingsPropertiesArgs>? OracleSettings { get; set; }

        /// <summary>
        /// PostgreSqlSettings property identifier.
        /// </summary>
        [Input("postgreSqlSettings")]
        public Input<Inputs.SettingsPropertiesPostgreSqlSettingsPropertiesArgs>? PostgreSqlSettings { get; set; }

        /// <summary>
        /// RedshiftSettings property identifier.
        /// </summary>
        [Input("redshiftSettings")]
        public Input<Inputs.SettingsPropertiesRedshiftSettingsPropertiesArgs>? RedshiftSettings { get; set; }

        public SettingsPropertiesArgs()
        {
        }
        public static new SettingsPropertiesArgs Empty => new SettingsPropertiesArgs();
    }
}
