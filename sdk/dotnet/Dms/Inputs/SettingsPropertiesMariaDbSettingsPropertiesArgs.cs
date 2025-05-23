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
    /// MariaDbSettings property identifier.
    /// </summary>
    public sealed class SettingsPropertiesMariaDbSettingsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        [Input("port", required: true)]
        public Input<int> Port { get; set; } = null!;

        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        [Input("sslMode", required: true)]
        public Input<Pulumi.AwsNative.Dms.DataProviderDmsSslModeValue> SslMode { get; set; } = null!;

        public SettingsPropertiesMariaDbSettingsPropertiesArgs()
        {
        }
        public static new SettingsPropertiesMariaDbSettingsPropertiesArgs Empty => new SettingsPropertiesMariaDbSettingsPropertiesArgs();
    }
}
