// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisInnerFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryInnerFilter")]
        public Input<Inputs.AnalysisCategoryInnerFilterArgs>? CategoryInnerFilter { get; set; }

        public AnalysisInnerFilterArgs()
        {
        }
        public static new AnalysisInnerFilterArgs Empty => new AnalysisInnerFilterArgs();
    }
}