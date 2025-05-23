// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// Specifies which audit checks are enabled and disabled for this account.
    /// </summary>
    public sealed class AccountAuditConfigurationAuditCheckConfigurationsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Checks the permissiveness of an authenticated Amazon Cognito identity pool role. For this check, AWS IoT Device Defender audits all Amazon Cognito identity pools that have been used to connect to the AWS IoT message broker during the 31 days before the audit is performed.
        /// </summary>
        [Input("authenticatedCognitoRoleOverlyPermissiveCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? AuthenticatedCognitoRoleOverlyPermissiveCheck { get; set; }

        /// <summary>
        /// Checks if a CA certificate is expiring. This check applies to CA certificates expiring within 30 days or that have expired.
        /// </summary>
        [Input("caCertificateExpiringCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? CaCertificateExpiringCheck { get; set; }

        /// <summary>
        /// Checks the quality of the CA certificate key. The quality checks if the key is in a valid format, not expired, and if the key meets a minimum required size. This check applies to CA certificates that are `ACTIVE` or `PENDING_TRANSFER` .
        /// </summary>
        [Input("caCertificateKeyQualityCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? CaCertificateKeyQualityCheck { get; set; }

        /// <summary>
        /// Checks if multiple devices connect using the same client ID.
        /// </summary>
        [Input("conflictingClientIdsCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? ConflictingClientIdsCheck { get; set; }

        /// <summary>
        /// Checks when a device certificate has been active for a number of days greater than or equal to the number you specify.
        /// </summary>
        [Input("deviceCertificateAgeCheck")]
        public Input<Inputs.AccountAuditConfigurationDeviceCertAgeAuditCheckConfigurationArgs>? DeviceCertificateAgeCheck { get; set; }

        /// <summary>
        /// Checks if a device certificate is expiring. By default, this check applies to device certificates expiring within 30 days or that have expired. You can modify this threshold by configuring the DeviceCertExpirationAuditCheckConfiguration.
        /// </summary>
        [Input("deviceCertificateExpiringCheck")]
        public Input<Inputs.AccountAuditConfigurationDeviceCertExpirationAuditCheckConfigurationArgs>? DeviceCertificateExpiringCheck { get; set; }

        /// <summary>
        /// Checks the quality of the device certificate key. The quality checks if the key is in a valid format, not expired, signed by a registered certificate authority, and if the key meets a minimum required size.
        /// </summary>
        [Input("deviceCertificateKeyQualityCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? DeviceCertificateKeyQualityCheck { get; set; }

        /// <summary>
        /// Checks if multiple concurrent connections use the same X.509 certificate to authenticate with AWS IoT .
        /// </summary>
        [Input("deviceCertificateSharedCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? DeviceCertificateSharedCheck { get; set; }

        /// <summary>
        /// Checks if device certificates are still active despite being revoked by an intermediate CA.
        /// </summary>
        [Input("intermediateCaRevokedForActiveDeviceCertificatesCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? IntermediateCaRevokedForActiveDeviceCertificatesCheck { get; set; }

        /// <summary>
        /// Checks if an AWS IoT policy is potentially misconfigured. Misconfigured policies, including overly permissive policies, can cause security incidents like allowing devices access to unintended resources. This check is a warning for you to make sure that only intended actions are allowed before updating the policy.
        /// </summary>
        [Input("ioTPolicyPotentialMisConfigurationCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? IoTPolicyPotentialMisConfigurationCheck { get; set; }

        /// <summary>
        /// Checks the permissiveness of a policy attached to an authenticated Amazon Cognito identity pool role.
        /// </summary>
        [Input("iotPolicyOverlyPermissiveCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? IotPolicyOverlyPermissiveCheck { get; set; }

        /// <summary>
        /// Checks if a role alias has access to services that haven't been used for the AWS IoT device in the last year.
        /// </summary>
        [Input("iotRoleAliasAllowsAccessToUnusedServicesCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? IotRoleAliasAllowsAccessToUnusedServicesCheck { get; set; }

        /// <summary>
        /// Checks if the temporary credentials provided by AWS IoT role aliases are overly permissive.
        /// </summary>
        [Input("iotRoleAliasOverlyPermissiveCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? IotRoleAliasOverlyPermissiveCheck { get; set; }

        /// <summary>
        /// Checks if AWS IoT logs are disabled.
        /// </summary>
        [Input("loggingDisabledCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? LoggingDisabledCheck { get; set; }

        /// <summary>
        /// Checks if a revoked CA certificate is still active.
        /// </summary>
        [Input("revokedCaCertificateStillActiveCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? RevokedCaCertificateStillActiveCheck { get; set; }

        /// <summary>
        /// Checks if a revoked device certificate is still active.
        /// </summary>
        [Input("revokedDeviceCertificateStillActiveCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? RevokedDeviceCertificateStillActiveCheck { get; set; }

        /// <summary>
        /// Checks if policy attached to an unauthenticated Amazon Cognito identity pool role is too permissive.
        /// </summary>
        [Input("unauthenticatedCognitoRoleOverlyPermissiveCheck")]
        public Input<Inputs.AccountAuditConfigurationAuditCheckConfigurationArgs>? UnauthenticatedCognitoRoleOverlyPermissiveCheck { get; set; }

        public AccountAuditConfigurationAuditCheckConfigurationsArgs()
        {
        }
        public static new AccountAuditConfigurationAuditCheckConfigurationsArgs Empty => new AccountAuditConfigurationAuditCheckConfigurationsArgs();
    }
}
