// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms.Outputs
{

    [OutputType]
    public sealed class Settings1PropertiesMySqlSettingsProperties
    {
        public readonly string? CertificateArn;
        public readonly int? Port;
        public readonly string? ServerName;
        public readonly Pulumi.AwsNative.Dms.DataProviderDmsSslModeValue? SslMode;

        [OutputConstructor]
        private Settings1PropertiesMySqlSettingsProperties(
            string? certificateArn,

            int? port,

            string? serverName,

            Pulumi.AwsNative.Dms.DataProviderDmsSslModeValue? sslMode)
        {
            CertificateArn = certificateArn;
            Port = port;
            ServerName = serverName;
            SslMode = sslMode;
        }
    }
}