// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Outputs
{

    [OutputType]
    public sealed class DomainAdvancedSecurityOptionsInput
    {
        public readonly string? AnonymousAuthDisableDate;
        public readonly bool? AnonymousAuthEnabled;
        public readonly bool? Enabled;
        public readonly bool? InternalUserDatabaseEnabled;
        public readonly Outputs.DomainMasterUserOptions? MasterUserOptions;
        public readonly Outputs.DomainSAMLOptions? SAMLOptions;

        [OutputConstructor]
        private DomainAdvancedSecurityOptionsInput(
            string? anonymousAuthDisableDate,

            bool? anonymousAuthEnabled,

            bool? enabled,

            bool? internalUserDatabaseEnabled,

            Outputs.DomainMasterUserOptions? masterUserOptions,

            Outputs.DomainSAMLOptions? sAMLOptions)
        {
            AnonymousAuthDisableDate = anonymousAuthDisableDate;
            AnonymousAuthEnabled = anonymousAuthEnabled;
            Enabled = enabled;
            InternalUserDatabaseEnabled = internalUserDatabaseEnabled;
            MasterUserOptions = masterUserOptions;
            SAMLOptions = sAMLOptions;
        }
    }
}