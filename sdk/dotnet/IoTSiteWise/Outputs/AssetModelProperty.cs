// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Outputs
{

    /// <summary>
    /// Contains information about an asset model property.
    /// </summary>
    [OutputType]
    public sealed class AssetModelProperty
    {
        /// <summary>
        /// The data type of the asset model property.
        /// </summary>
        public readonly Pulumi.AwsNative.IoTSiteWise.AssetModelDataType DataType;
        /// <summary>
        /// The data type of the structure for this property.
        /// </summary>
        public readonly Pulumi.AwsNative.IoTSiteWise.AssetModelDataTypeSpec? DataTypeSpec;
        /// <summary>
        /// Customer provided ID for property.
        /// </summary>
        public readonly string LogicalId;
        /// <summary>
        /// The name of the asset model property.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The property type
        /// </summary>
        public readonly Outputs.AssetModelPropertyType Type;
        /// <summary>
        /// The unit of the asset model property, such as Newtons or RPM.
        /// </summary>
        public readonly string? Unit;

        [OutputConstructor]
        private AssetModelProperty(
            Pulumi.AwsNative.IoTSiteWise.AssetModelDataType dataType,

            Pulumi.AwsNative.IoTSiteWise.AssetModelDataTypeSpec? dataTypeSpec,

            string logicalId,

            string name,

            Outputs.AssetModelPropertyType type,

            string? unit)
        {
            DataType = dataType;
            DataTypeSpec = dataTypeSpec;
            LogicalId = logicalId;
            Name = name;
            Type = type;
            Unit = unit;
        }
    }
}