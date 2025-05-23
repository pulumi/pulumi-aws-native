// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    /// <summary>
    /// A structure containing the configName and corresponding configValue for configuring DeviceCertExpirationCheck.
    /// </summary>
    [OutputType]
    public sealed class AccountAuditConfigurationDeviceCertExpirationAuditCheckConfiguration
    {
        /// <summary>
        /// Configuration settings for the device certificate expiration check, including the threshold in days before expiration. This configuration is of type `CertExpirationCheckCustomConfiguration`
        /// </summary>
        public readonly Outputs.AccountAuditConfigurationCertExpirationCheckCustomConfiguration? Configuration;
        /// <summary>
        /// True if the check is enabled.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private AccountAuditConfigurationDeviceCertExpirationAuditCheckConfiguration(
            Outputs.AccountAuditConfigurationCertExpirationCheckCustomConfiguration? configuration,

            bool? enabled)
        {
            Configuration = configuration;
            Enabled = enabled;
        }
    }
}
