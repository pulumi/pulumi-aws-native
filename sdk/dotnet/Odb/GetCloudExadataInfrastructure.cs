// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Odb
{
    public static class GetCloudExadataInfrastructure
    {
        /// <summary>
        /// The AWS::ODB::CloudExadataInfrastructure resource creates an Exadata Infrastructure
        /// </summary>
        public static Task<GetCloudExadataInfrastructureResult> InvokeAsync(GetCloudExadataInfrastructureArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCloudExadataInfrastructureResult>("aws-native:odb:getCloudExadataInfrastructure", args ?? new GetCloudExadataInfrastructureArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::ODB::CloudExadataInfrastructure resource creates an Exadata Infrastructure
        /// </summary>
        public static Output<GetCloudExadataInfrastructureResult> Invoke(GetCloudExadataInfrastructureInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudExadataInfrastructureResult>("aws-native:odb:getCloudExadataInfrastructure", args ?? new GetCloudExadataInfrastructureInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::ODB::CloudExadataInfrastructure resource creates an Exadata Infrastructure
        /// </summary>
        public static Output<GetCloudExadataInfrastructureResult> Invoke(GetCloudExadataInfrastructureInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCloudExadataInfrastructureResult>("aws-native:odb:getCloudExadataInfrastructure", args ?? new GetCloudExadataInfrastructureInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudExadataInfrastructureArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the Exadata infrastructure.
        /// </summary>
        [Input("cloudExadataInfrastructureArn", required: true)]
        public string CloudExadataInfrastructureArn { get; set; } = null!;

        public GetCloudExadataInfrastructureArgs()
        {
        }
        public static new GetCloudExadataInfrastructureArgs Empty => new GetCloudExadataInfrastructureArgs();
    }

    public sealed class GetCloudExadataInfrastructureInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the Exadata infrastructure.
        /// </summary>
        [Input("cloudExadataInfrastructureArn", required: true)]
        public Input<string> CloudExadataInfrastructureArn { get; set; } = null!;

        public GetCloudExadataInfrastructureInvokeArgs()
        {
        }
        public static new GetCloudExadataInfrastructureInvokeArgs Empty => new GetCloudExadataInfrastructureInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudExadataInfrastructureResult
    {
        /// <summary>
        /// The number of storage servers requested for the Exadata infrastructure.
        /// </summary>
        public readonly int? ActivatedStorageCount;
        /// <summary>
        /// The number of storage servers requested for the Exadata infrastructure.
        /// </summary>
        public readonly int? AdditionalStorageCount;
        /// <summary>
        /// The amount of available storage, in gigabytes (GB), for the Exadata infrastructure.
        /// </summary>
        public readonly int? AvailableStorageSizeInGbs;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the Exadata infrastructure.
        /// </summary>
        public readonly string? CloudExadataInfrastructureArn;
        /// <summary>
        /// The unique identifier for the Exadata infrastructure.
        /// </summary>
        public readonly string? CloudExadataInfrastructureId;
        /// <summary>
        /// The OCI model compute model used when you create or clone an instance: ECPU or OCPU. An ECPU is an abstracted measure of compute resources. ECPUs are based on the number of cores elastically allocated from a pool of compute and storage servers. An OCPU is a legacy physical measure of compute resources. OCPUs are based on the physical core of a processor with hyper-threading enabled.
        /// </summary>
        public readonly string? ComputeModel;
        /// <summary>
        /// The total number of CPU cores that are allocated to the Exadata infrastructure.
        /// </summary>
        public readonly int? CpuCount;
        /// <summary>
        /// The size of the Exadata infrastructure's data disk group, in terabytes (TB).
        /// </summary>
        public readonly double? DataStorageSizeInTbs;
        /// <summary>
        /// The size of the Exadata infrastructure's local node storage, in gigabytes (GB).
        /// </summary>
        public readonly int? DbNodeStorageSizeInGbs;
        /// <summary>
        /// The list of database server identifiers for the Exadata infrastructure.
        /// </summary>
        public readonly ImmutableArray<string> DbServerIds;
        /// <summary>
        /// The software version of the database servers (dom0) in the Exadata infrastructure.
        /// </summary>
        public readonly string? DbServerVersion;
        /// <summary>
        /// The total number of CPU cores available on the Exadata infrastructure.
        /// </summary>
        public readonly int? MaxCpuCount;
        /// <summary>
        /// The total amount of data disk group storage, in terabytes (TB), that's available on the Exadata infrastructure.
        /// </summary>
        public readonly double? MaxDataStorageInTbs;
        /// <summary>
        /// The total amount of local node storage, in gigabytes (GB), that's available on the Exadata infrastructure.
        /// </summary>
        public readonly int? MaxDbNodeStorageSizeInGbs;
        /// <summary>
        /// The total amount of memory, in gigabytes (GB), that's available on the Exadata infrastructure.
        /// </summary>
        public readonly int? MaxMemoryInGbs;
        /// <summary>
        /// The amount of memory, in gigabytes (GB), that's allocated on the Exadata infrastructure.
        /// </summary>
        public readonly int? MemorySizeInGbs;
        /// <summary>
        /// The name of the OCI resource anchor for the Exadata infrastructure.
        /// </summary>
        public readonly string? OciResourceAnchorName;
        /// <summary>
        /// The HTTPS link to the Exadata infrastructure in OCI.
        /// </summary>
        public readonly string? OciUrl;
        /// <summary>
        /// The OCID of the Exadata infrastructure.
        /// </summary>
        public readonly string? Ocid;
        /// <summary>
        /// The software version of the storage servers on the Exadata infrastructure.
        /// </summary>
        public readonly string? StorageServerVersion;
        /// <summary>
        /// Tags to assign to the Exadata Infrastructure.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The total amount of storage, in gigabytes (GB), on the the Exadata infrastructure.
        /// </summary>
        public readonly int? TotalStorageSizeInGbs;

        [OutputConstructor]
        private GetCloudExadataInfrastructureResult(
            int? activatedStorageCount,

            int? additionalStorageCount,

            int? availableStorageSizeInGbs,

            string? cloudExadataInfrastructureArn,

            string? cloudExadataInfrastructureId,

            string? computeModel,

            int? cpuCount,

            double? dataStorageSizeInTbs,

            int? dbNodeStorageSizeInGbs,

            ImmutableArray<string> dbServerIds,

            string? dbServerVersion,

            int? maxCpuCount,

            double? maxDataStorageInTbs,

            int? maxDbNodeStorageSizeInGbs,

            int? maxMemoryInGbs,

            int? memorySizeInGbs,

            string? ociResourceAnchorName,

            string? ociUrl,

            string? ocid,

            string? storageServerVersion,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            int? totalStorageSizeInGbs)
        {
            ActivatedStorageCount = activatedStorageCount;
            AdditionalStorageCount = additionalStorageCount;
            AvailableStorageSizeInGbs = availableStorageSizeInGbs;
            CloudExadataInfrastructureArn = cloudExadataInfrastructureArn;
            CloudExadataInfrastructureId = cloudExadataInfrastructureId;
            ComputeModel = computeModel;
            CpuCount = cpuCount;
            DataStorageSizeInTbs = dataStorageSizeInTbs;
            DbNodeStorageSizeInGbs = dbNodeStorageSizeInGbs;
            DbServerIds = dbServerIds;
            DbServerVersion = dbServerVersion;
            MaxCpuCount = maxCpuCount;
            MaxDataStorageInTbs = maxDataStorageInTbs;
            MaxDbNodeStorageSizeInGbs = maxDbNodeStorageSizeInGbs;
            MaxMemoryInGbs = maxMemoryInGbs;
            MemorySizeInGbs = memorySizeInGbs;
            OciResourceAnchorName = ociResourceAnchorName;
            OciUrl = ociUrl;
            Ocid = ocid;
            StorageServerVersion = storageServerVersion;
            Tags = tags;
            TotalStorageSizeInGbs = totalStorageSizeInGbs;
        }
    }
}
