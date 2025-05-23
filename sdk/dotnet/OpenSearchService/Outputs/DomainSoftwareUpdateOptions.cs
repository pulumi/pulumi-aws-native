// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Outputs
{

    [OutputType]
    public sealed class DomainSoftwareUpdateOptions
    {
        /// <summary>
        /// Specifies whether automatic service software updates are enabled for the domain.
        /// </summary>
        public readonly bool? AutoSoftwareUpdateEnabled;

        [OutputConstructor]
        private DomainSoftwareUpdateOptions(bool? autoSoftwareUpdateEnabled)
        {
            AutoSoftwareUpdateEnabled = autoSoftwareUpdateEnabled;
        }
    }
}
