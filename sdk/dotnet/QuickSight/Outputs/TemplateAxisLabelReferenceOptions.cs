// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateAxisLabelReferenceOptions
    {
        /// <summary>
        /// The column that the axis label is targeted to.
        /// </summary>
        public readonly Outputs.TemplateColumnIdentifier Column;
        /// <summary>
        /// The field that the axis label is targeted to.
        /// </summary>
        public readonly string FieldId;

        [OutputConstructor]
        private TemplateAxisLabelReferenceOptions(
            Outputs.TemplateColumnIdentifier column,

            string fieldId)
        {
            Column = column;
            FieldId = fieldId;
        }
    }
}
