// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx
{
    public static class GetStorageVirtualMachine
    {
        /// <summary>
        /// Resource Type definition for AWS::FSx::StorageVirtualMachine
        /// </summary>
        public static Task<GetStorageVirtualMachineResult> InvokeAsync(GetStorageVirtualMachineArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStorageVirtualMachineResult>("aws-native:fsx:getStorageVirtualMachine", args ?? new GetStorageVirtualMachineArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::FSx::StorageVirtualMachine
        /// </summary>
        public static Output<GetStorageVirtualMachineResult> Invoke(GetStorageVirtualMachineInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageVirtualMachineResult>("aws-native:fsx:getStorageVirtualMachine", args ?? new GetStorageVirtualMachineInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStorageVirtualMachineArgs : global::Pulumi.InvokeArgs
    {
        [Input("storageVirtualMachineId", required: true)]
        public string StorageVirtualMachineId { get; set; } = null!;

        public GetStorageVirtualMachineArgs()
        {
        }
        public static new GetStorageVirtualMachineArgs Empty => new GetStorageVirtualMachineArgs();
    }

    public sealed class GetStorageVirtualMachineInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("storageVirtualMachineId", required: true)]
        public Input<string> StorageVirtualMachineId { get; set; } = null!;

        public GetStorageVirtualMachineInvokeArgs()
        {
        }
        public static new GetStorageVirtualMachineInvokeArgs Empty => new GetStorageVirtualMachineInvokeArgs();
    }


    [OutputType]
    public sealed class GetStorageVirtualMachineResult
    {
        public readonly Outputs.StorageVirtualMachineActiveDirectoryConfiguration? ActiveDirectoryConfiguration;
        public readonly string? ResourceARN;
        public readonly string? StorageVirtualMachineId;
        public readonly string? SvmAdminPassword;
        public readonly ImmutableArray<Outputs.StorageVirtualMachineTag> Tags;
        public readonly string? UUID;

        [OutputConstructor]
        private GetStorageVirtualMachineResult(
            Outputs.StorageVirtualMachineActiveDirectoryConfiguration? activeDirectoryConfiguration,

            string? resourceARN,

            string? storageVirtualMachineId,

            string? svmAdminPassword,

            ImmutableArray<Outputs.StorageVirtualMachineTag> tags,

            string? uUID)
        {
            ActiveDirectoryConfiguration = activeDirectoryConfiguration;
            ResourceARN = resourceARN;
            StorageVirtualMachineId = storageVirtualMachineId;
            SvmAdminPassword = svmAdminPassword;
            Tags = tags;
            UUID = uUID;
        }
    }
}