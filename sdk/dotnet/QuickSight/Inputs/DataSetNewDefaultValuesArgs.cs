// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;The configuration that overrides the existing default values for a dataset parameter that is inherited from another dataset.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetNewDefaultValuesArgs : global::Pulumi.ResourceArgs
    {
        [Input("dateTimeStaticValues")]
        private InputList<string>? _dateTimeStaticValues;

        /// <summary>
        /// &lt;p&gt;A list of static default values for a given date time parameter.&lt;/p&gt;
        /// </summary>
        public InputList<string> DateTimeStaticValues
        {
            get => _dateTimeStaticValues ?? (_dateTimeStaticValues = new InputList<string>());
            set => _dateTimeStaticValues = value;
        }

        [Input("decimalStaticValues")]
        private InputList<double>? _decimalStaticValues;

        /// <summary>
        /// &lt;p&gt;A list of static default values for a given decimal parameter.&lt;/p&gt;
        /// </summary>
        public InputList<double> DecimalStaticValues
        {
            get => _decimalStaticValues ?? (_decimalStaticValues = new InputList<double>());
            set => _decimalStaticValues = value;
        }

        [Input("integerStaticValues")]
        private InputList<double>? _integerStaticValues;

        /// <summary>
        /// &lt;p&gt;A list of static default values for a given integer parameter.&lt;/p&gt;
        /// </summary>
        public InputList<double> IntegerStaticValues
        {
            get => _integerStaticValues ?? (_integerStaticValues = new InputList<double>());
            set => _integerStaticValues = value;
        }

        [Input("stringStaticValues")]
        private InputList<string>? _stringStaticValues;

        /// <summary>
        /// &lt;p&gt;A list of static default values for a given string parameter.&lt;/p&gt;
        /// </summary>
        public InputList<string> StringStaticValues
        {
            get => _stringStaticValues ?? (_stringStaticValues = new InputList<string>());
            set => _stringStaticValues = value;
        }

        public DataSetNewDefaultValuesArgs()
        {
        }
        public static new DataSetNewDefaultValuesArgs Empty => new DataSetNewDefaultValuesArgs();
    }
}