// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SecurityHub.Inputs
{

    /// <summary>
    /// A date filter for querying findings.
    /// </summary>
    public sealed class InsightDateFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("dateRange")]
        public Input<Inputs.InsightDateRangeArgs>? DateRange { get; set; }

        [Input("end")]
        public Input<string>? End { get; set; }

        [Input("start")]
        public Input<string>? Start { get; set; }

        public InsightDateFilterArgs()
        {
        }
        public static new InsightDateFilterArgs Empty => new InsightDateFilterArgs();
    }
}