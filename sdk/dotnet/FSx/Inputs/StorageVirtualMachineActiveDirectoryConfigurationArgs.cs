// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Inputs
{

    public sealed class StorageVirtualMachineActiveDirectoryConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("netBiosName")]
        public Input<string>? NetBiosName { get; set; }

        [Input("selfManagedActiveDirectoryConfiguration")]
        public Input<Inputs.StorageVirtualMachineSelfManagedActiveDirectoryConfigurationArgs>? SelfManagedActiveDirectoryConfiguration { get; set; }

        public StorageVirtualMachineActiveDirectoryConfigurationArgs()
        {
        }
        public static new StorageVirtualMachineActiveDirectoryConfigurationArgs Empty => new StorageVirtualMachineActiveDirectoryConfigurationArgs();
    }
}