// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PcaConnectorAd.Outputs
{

    [OutputType]
    public sealed class TemplateSubjectNameFlagsV4
    {
        public readonly bool? RequireCommonName;
        public readonly bool? RequireDirectoryPath;
        public readonly bool? RequireDnsAsCn;
        public readonly bool? RequireEmail;
        public readonly bool? SanRequireDirectoryGuid;
        public readonly bool? SanRequireDns;
        public readonly bool? SanRequireDomainDns;
        public readonly bool? SanRequireEmail;
        public readonly bool? SanRequireSpn;
        public readonly bool? SanRequireUpn;

        [OutputConstructor]
        private TemplateSubjectNameFlagsV4(
            bool? requireCommonName,

            bool? requireDirectoryPath,

            bool? requireDnsAsCn,

            bool? requireEmail,

            bool? sanRequireDirectoryGuid,

            bool? sanRequireDns,

            bool? sanRequireDomainDns,

            bool? sanRequireEmail,

            bool? sanRequireSpn,

            bool? sanRequireUpn)
        {
            RequireCommonName = requireCommonName;
            RequireDirectoryPath = requireDirectoryPath;
            RequireDnsAsCn = requireDnsAsCn;
            RequireEmail = requireEmail;
            SanRequireDirectoryGuid = sanRequireDirectoryGuid;
            SanRequireDns = sanRequireDns;
            SanRequireDomainDns = sanRequireDomainDns;
            SanRequireEmail = sanRequireEmail;
            SanRequireSpn = sanRequireSpn;
            SanRequireUpn = sanRequireUpn;
        }
    }
}