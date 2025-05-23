// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateNumericalDimensionFieldArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The column that is used in the `NumericalDimensionField` .
        /// </summary>
        [Input("column", required: true)]
        public Input<Inputs.TemplateColumnIdentifierArgs> Column { get; set; } = null!;

        /// <summary>
        /// The custom field ID.
        /// </summary>
        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        /// <summary>
        /// The format configuration of the field.
        /// </summary>
        [Input("formatConfiguration")]
        public Input<Inputs.TemplateNumberFormatConfigurationArgs>? FormatConfiguration { get; set; }

        /// <summary>
        /// The custom hierarchy ID.
        /// </summary>
        [Input("hierarchyId")]
        public Input<string>? HierarchyId { get; set; }

        public TemplateNumericalDimensionFieldArgs()
        {
        }
        public static new TemplateNumericalDimensionFieldArgs Empty => new TemplateNumericalDimensionFieldArgs();
    }
}
