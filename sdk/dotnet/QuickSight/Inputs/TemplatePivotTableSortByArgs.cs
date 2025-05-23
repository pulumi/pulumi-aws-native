// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplatePivotTableSortByArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The column sort (field id, direction) for the pivot table sort by options.
        /// </summary>
        [Input("column")]
        public Input<Inputs.TemplateColumnSortArgs>? Column { get; set; }

        /// <summary>
        /// The data path sort (data path value, direction) for the pivot table sort by options.
        /// </summary>
        [Input("dataPath")]
        public Input<Inputs.TemplateDataPathSortArgs>? DataPath { get; set; }

        /// <summary>
        /// The field sort (field id, direction) for the pivot table sort by options.
        /// </summary>
        [Input("field")]
        public Input<Inputs.TemplateFieldSortArgs>? Field { get; set; }

        public TemplatePivotTableSortByArgs()
        {
        }
        public static new TemplatePivotTableSortByArgs Empty => new TemplatePivotTableSortByArgs();
    }
}
