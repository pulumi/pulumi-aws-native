// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTimeRangeFilterValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("parameter")]
        public Input<string>? Parameter { get; set; }

        [Input("rollingDate")]
        public Input<Inputs.AnalysisRollingDateConfigurationArgs>? RollingDate { get; set; }

        [Input("staticValue")]
        public Input<string>? StaticValue { get; set; }

        public AnalysisTimeRangeFilterValueArgs()
        {
        }
        public static new AnalysisTimeRangeFilterValueArgs Empty => new AnalysisTimeRangeFilterValueArgs();
    }
}