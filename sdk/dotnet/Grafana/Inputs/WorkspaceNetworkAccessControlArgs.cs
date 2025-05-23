// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Grafana.Inputs
{

    /// <summary>
    /// The configuration settings for Network Access Control.
    /// </summary>
    public sealed class WorkspaceNetworkAccessControlArgs : global::Pulumi.ResourceArgs
    {
        [Input("prefixListIds")]
        private InputList<string>? _prefixListIds;

        /// <summary>
        /// The list of prefix list IDs. A prefix list is a list of CIDR ranges of IP addresses. The IP addresses specified are allowed to access your workspace. If the list is not included in the configuration then no IP addresses will be allowed to access the workspace.
        /// </summary>
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        [Input("vpceIds")]
        private InputList<string>? _vpceIds;

        /// <summary>
        /// The list of Amazon VPC endpoint IDs for the workspace. If a NetworkAccessConfiguration is specified then only VPC endpoints specified here will be allowed to access the workspace.
        /// </summary>
        public InputList<string> VpceIds
        {
            get => _vpceIds ?? (_vpceIds = new InputList<string>());
            set => _vpceIds = value;
        }

        public WorkspaceNetworkAccessControlArgs()
        {
        }
        public static new WorkspaceNetworkAccessControlArgs Empty => new WorkspaceNetworkAccessControlArgs();
    }
}
