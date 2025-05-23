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
    /// &lt;p&gt;Dataset configuration.&lt;/p&gt;
    /// </summary>
    public sealed class TemplateDataSetConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("columnGroupSchemaList")]
        private InputList<Inputs.TemplateColumnGroupSchemaArgs>? _columnGroupSchemaList;

        /// <summary>
        /// &lt;p&gt;A structure containing the list of column group schemas.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.TemplateColumnGroupSchemaArgs> ColumnGroupSchemaList
        {
            get => _columnGroupSchemaList ?? (_columnGroupSchemaList = new InputList<Inputs.TemplateColumnGroupSchemaArgs>());
            set => _columnGroupSchemaList = value;
        }

        /// <summary>
        /// Dataset schema.
        /// </summary>
        [Input("dataSetSchema")]
        public Input<Inputs.TemplateDataSetSchemaArgs>? DataSetSchema { get; set; }

        /// <summary>
        /// &lt;p&gt;Placeholder.&lt;/p&gt;
        /// </summary>
        [Input("placeholder")]
        public Input<string>? Placeholder { get; set; }

        public TemplateDataSetConfigurationArgs()
        {
        }
        public static new TemplateDataSetConfigurationArgs Empty => new TemplateDataSetConfigurationArgs();
    }
}
