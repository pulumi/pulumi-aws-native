// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisStringFormatConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("nullValueFormatConfiguration")]
        public Input<Inputs.AnalysisNullValueFormatConfigurationArgs>? NullValueFormatConfiguration { get; set; }

        [Input("numericFormatConfiguration")]
        public Input<Inputs.AnalysisNumericFormatConfigurationArgs>? NumericFormatConfiguration { get; set; }

        public AnalysisStringFormatConfigurationArgs()
        {
        }
        public static new AnalysisStringFormatConfigurationArgs Empty => new AnalysisStringFormatConfigurationArgs();
    }
}