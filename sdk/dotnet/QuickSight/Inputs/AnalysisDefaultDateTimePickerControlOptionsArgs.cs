// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisDefaultDateTimePickerControlOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("displayOptions")]
        public Input<Inputs.AnalysisDateTimePickerControlDisplayOptionsArgs>? DisplayOptions { get; set; }

        [Input("type")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisSheetControlDateTimePickerType>? Type { get; set; }

        public AnalysisDefaultDateTimePickerControlOptionsArgs()
        {
        }
        public static new AnalysisDefaultDateTimePickerControlOptionsArgs Empty => new AnalysisDefaultDateTimePickerControlOptionsArgs();
    }
}