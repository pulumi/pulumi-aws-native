// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisFieldSortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sort direction. Choose one of the following options:
        /// 
        /// - `ASC` : Ascending
        /// - `DESC` : Descending
        /// </summary>
        [Input("direction", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisSortDirection> Direction { get; set; } = null!;

        /// <summary>
        /// The sort configuration target field.
        /// </summary>
        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        public AnalysisFieldSortArgs()
        {
        }
        public static new AnalysisFieldSortArgs Empty => new AnalysisFieldSortArgs();
    }
}
