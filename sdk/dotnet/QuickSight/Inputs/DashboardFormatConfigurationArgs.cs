// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardFormatConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("dateTimeFormatConfiguration")]
        public Input<Inputs.DashboardDateTimeFormatConfigurationArgs>? DateTimeFormatConfiguration { get; set; }

        [Input("numberFormatConfiguration")]
        public Input<Inputs.DashboardNumberFormatConfigurationArgs>? NumberFormatConfiguration { get; set; }

        [Input("stringFormatConfiguration")]
        public Input<Inputs.DashboardStringFormatConfigurationArgs>? StringFormatConfiguration { get; set; }

        public DashboardFormatConfigurationArgs()
        {
        }
        public static new DashboardFormatConfigurationArgs Empty => new DashboardFormatConfigurationArgs();
    }
}