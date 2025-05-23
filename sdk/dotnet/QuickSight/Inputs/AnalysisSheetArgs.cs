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
    /// &lt;p&gt;A &lt;i&gt;sheet&lt;/i&gt;, which is an object that contains a set of visuals that
    ///             are viewed together on one page in Amazon QuickSight. Every analysis and dashboard
    ///             contains at least one sheet. Each sheet contains at least one visualization widget, for
    ///             example a chart, pivot table, or narrative insight. Sheets can be associated with other
    ///             components, such as controls, filters, and so on.&lt;/p&gt;
    /// </summary>
    public sealed class AnalysisSheetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The name of a sheet. This name is displayed on the sheet's tab in the Amazon QuickSight
        ///             console.&lt;/p&gt;
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// &lt;p&gt;The unique identifier associated with a sheet.&lt;/p&gt;
        /// </summary>
        [Input("sheetId")]
        public Input<string>? SheetId { get; set; }

        public AnalysisSheetArgs()
        {
        }
        public static new AnalysisSheetArgs Empty => new AnalysisSheetArgs();
    }
}
