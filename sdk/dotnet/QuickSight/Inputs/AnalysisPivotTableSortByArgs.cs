// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisPivotTableSortByArgs : global::Pulumi.ResourceArgs
    {
        [Input("column")]
        public Input<Inputs.AnalysisColumnSortArgs>? Column { get; set; }

        [Input("dataPath")]
        public Input<Inputs.AnalysisDataPathSortArgs>? DataPath { get; set; }

        [Input("field")]
        public Input<Inputs.AnalysisFieldSortArgs>? Field { get; set; }

        public AnalysisPivotTableSortByArgs()
        {
        }
        public static new AnalysisPivotTableSortByArgs Empty => new AnalysisPivotTableSortByArgs();
    }
}