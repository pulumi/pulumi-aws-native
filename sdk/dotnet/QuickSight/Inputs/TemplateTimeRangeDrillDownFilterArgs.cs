// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTimeRangeDrillDownFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("column", required: true)]
        public Input<Inputs.TemplateColumnIdentifierArgs> Column { get; set; } = null!;

        [Input("rangeMaximum", required: true)]
        public Input<string> RangeMaximum { get; set; } = null!;

        [Input("rangeMinimum", required: true)]
        public Input<string> RangeMinimum { get; set; } = null!;

        [Input("timeGranularity", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.TemplateTimeGranularity> TimeGranularity { get; set; } = null!;

        public TemplateTimeRangeDrillDownFilterArgs()
        {
        }
        public static new TemplateTimeRangeDrillDownFilterArgs Empty => new TemplateTimeRangeDrillDownFilterArgs();
    }
}