// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WorkspacesInstances.Inputs
{

    public sealed class WorkspaceInstanceInstanceMaintenanceOptionsRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("autoRecovery")]
        public Input<Pulumi.AwsNative.WorkspacesInstances.WorkspaceInstanceInstanceMaintenanceOptionsRequestAutoRecovery>? AutoRecovery { get; set; }

        public WorkspaceInstanceInstanceMaintenanceOptionsRequestArgs()
        {
        }
        public static new WorkspaceInstanceInstanceMaintenanceOptionsRequestArgs Empty => new WorkspaceInstanceInstanceMaintenanceOptionsRequestArgs();
    }
}
