// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Inputs
{

    /// <summary>
    /// When enabled, Network Firewall checks if the server certificate presented by the server in the SSL/TLS connection has a revoked or unkown status. If the certificate has an unknown or revoked status, you must specify the actions that Network Firewall takes on outbound traffic. To check the certificate revocation status, you must also specify a `CertificateAuthorityArn` in [ServerCertificateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkfirewall-servercertificateconfiguration.html) .
    /// </summary>
    public sealed class TlsInspectionConfigurationServerCertificateConfigurationCheckCertificateRevocationStatusPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("revokedStatusAction")]
        public Input<Pulumi.AwsNative.NetworkFirewall.TlsInspectionConfigurationRevokedStatusAction>? RevokedStatusAction { get; set; }

        [Input("unknownStatusAction")]
        public Input<Pulumi.AwsNative.NetworkFirewall.TlsInspectionConfigurationUnknownStatusAction>? UnknownStatusAction { get; set; }

        public TlsInspectionConfigurationServerCertificateConfigurationCheckCertificateRevocationStatusPropertiesArgs()
        {
        }
        public static new TlsInspectionConfigurationServerCertificateConfigurationCheckCertificateRevocationStatusPropertiesArgs Empty => new TlsInspectionConfigurationServerCertificateConfigurationCheckCertificateRevocationStatusPropertiesArgs();
    }
}
