// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardDefaultDateTimePickerControlOptions
    {
        public readonly Outputs.DashboardDateTimePickerControlDisplayOptions? DisplayOptions;
        public readonly Pulumi.AwsNative.QuickSight.DashboardSheetControlDateTimePickerType? Type;

        [OutputConstructor]
        private DashboardDefaultDateTimePickerControlOptions(
            Outputs.DashboardDateTimePickerControlDisplayOptions? displayOptions,

            Pulumi.AwsNative.QuickSight.DashboardSheetControlDateTimePickerType? type)
        {
            DisplayOptions = displayOptions;
            Type = type;
        }
    }
}