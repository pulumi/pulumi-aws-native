// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Inputs
{

    /// <summary>
    /// Redshift Properties Input
    /// </summary>
    public sealed class ConnectionRedshiftPropertiesInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("credentials")]
        public InputUnion<Inputs.ConnectionRedshiftCredentials0PropertiesArgs, Inputs.ConnectionRedshiftCredentials1PropertiesArgs>? Credentials { get; set; }

        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("lineageSync")]
        public Input<Inputs.ConnectionRedshiftLineageSyncConfigurationInputArgs>? LineageSync { get; set; }

        [Input("port")]
        public Input<double>? Port { get; set; }

        [Input("storage")]
        public InputUnion<Inputs.ConnectionRedshiftStorageProperties0PropertiesArgs, Inputs.ConnectionRedshiftStorageProperties1PropertiesArgs>? Storage { get; set; }

        public ConnectionRedshiftPropertiesInputArgs()
        {
        }
        public static new ConnectionRedshiftPropertiesInputArgs Empty => new ConnectionRedshiftPropertiesInputArgs();
    }
}
