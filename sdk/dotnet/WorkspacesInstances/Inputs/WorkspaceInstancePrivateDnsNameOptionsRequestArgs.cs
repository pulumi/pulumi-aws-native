// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WorkspacesInstances.Inputs
{

    public sealed class WorkspaceInstancePrivateDnsNameOptionsRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("enableResourceNameDnsARecord")]
        public Input<bool>? EnableResourceNameDnsARecord { get; set; }

        [Input("enableResourceNameDnsAaaaRecord")]
        public Input<bool>? EnableResourceNameDnsAaaaRecord { get; set; }

        [Input("hostnameType")]
        public Input<Pulumi.AwsNative.WorkspacesInstances.WorkspaceInstancePrivateDnsNameOptionsRequestHostnameType>? HostnameType { get; set; }

        public WorkspaceInstancePrivateDnsNameOptionsRequestArgs()
        {
        }
        public static new WorkspaceInstancePrivateDnsNameOptionsRequestArgs Empty => new WorkspaceInstancePrivateDnsNameOptionsRequestArgs();
    }
}
