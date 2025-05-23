// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;The default values of a date time parameter.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetDateTimeDatasetParameterDefaultValuesArgs : global::Pulumi.ResourceArgs
    {
        [Input("staticValues")]
        private InputList<string>? _staticValues;

        /// <summary>
        /// &lt;p&gt;A list of static default values for a given date time parameter.&lt;/p&gt;
        /// </summary>
        public InputList<string> StaticValues
        {
            get => _staticValues ?? (_staticValues = new InputList<string>());
            set => _staticValues = value;
        }

        public DataSetDateTimeDatasetParameterDefaultValuesArgs()
        {
        }
        public static new DataSetDateTimeDatasetParameterDefaultValuesArgs Empty => new DataSetDateTimeDatasetParameterDefaultValuesArgs();
    }
}
