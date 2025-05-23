// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;A decimal parameter for a dataset.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetDecimalDatasetParameter
    {
        /// <summary>
        /// A list of default values for a given decimal parameter. This structure only accepts static values.
        /// </summary>
        public readonly Outputs.DataSetDecimalDatasetParameterDefaultValues? DefaultValues;
        /// <summary>
        /// &lt;p&gt;An identifier for the decimal parameter created in the dataset.&lt;/p&gt;
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// &lt;p&gt;The name of the decimal parameter that is created in the dataset.&lt;/p&gt;
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value type of the dataset parameter. Valid values are `single value` or `multi value` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.DataSetDatasetParameterValueType ValueType;

        [OutputConstructor]
        private DataSetDecimalDatasetParameter(
            Outputs.DataSetDecimalDatasetParameterDefaultValues? defaultValues,

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
