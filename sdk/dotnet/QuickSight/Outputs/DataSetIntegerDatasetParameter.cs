// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;A parameter created in the dataset of integer data type.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetIntegerDatasetParameter
    {
        public readonly Outputs.DataSetIntegerDatasetParameterDefaultValues? DefaultValues;
        public readonly string Id;
        public readonly string Name;
        public readonly Pulumi.AwsNative.QuickSight.DataSetDatasetParameterValueType ValueType;

        [OutputConstructor]
        private DataSetIntegerDatasetParameter(
            Outputs.DataSetIntegerDatasetParameterDefaultValues? defaultValues,

            string id,

            string name,

            Pulumi.AwsNative.QuickSight.DataSetDatasetParameterValueType valueType)
        {
            DefaultValues = defaultValues;
            Id = id;
            Name = name;
            ValueType = valueType;
        }
    }
}