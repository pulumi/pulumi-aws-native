// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RolesAnywhere.Outputs
{

    [OutputType]
    public sealed class ProfileAttributeMapping
    {
        public readonly Pulumi.AwsNative.RolesAnywhere.ProfileCertificateField CertificateField;
        public readonly ImmutableArray<Outputs.ProfileMappingRule> MappingRules;

        [OutputConstructor]
        private ProfileAttributeMapping(
            Pulumi.AwsNative.RolesAnywhere.ProfileCertificateField certificateField,

            ImmutableArray<Outputs.ProfileMappingRule> mappingRules)
        {
            CertificateField = certificateField;
            MappingRules = mappingRules;
        }
    }
}