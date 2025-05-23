// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchService.Inputs
{

    public sealed class DomainSoftwareUpdateOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether automatic service software updates are enabled for the domain.
        /// </summary>
        [Input("autoSoftwareUpdateEnabled")]
        public Input<bool>? AutoSoftwareUpdateEnabled { get; set; }

        public DomainSoftwareUpdateOptionsArgs()
        {
        }
        public static new DomainSoftwareUpdateOptionsArgs Empty => new DomainSoftwareUpdateOptionsArgs();
    }
}
