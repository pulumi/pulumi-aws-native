// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisWordCloudOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("cloudLayout")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisWordCloudCloudLayout>? CloudLayout { get; set; }

        [Input("maximumStringLength")]
        public Input<double>? MaximumStringLength { get; set; }

        [Input("wordCasing")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisWordCloudWordCasing>? WordCasing { get; set; }

        [Input("wordOrientation")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisWordCloudWordOrientation>? WordOrientation { get; set; }

        [Input("wordPadding")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisWordCloudWordPadding>? WordPadding { get; set; }

        [Input("wordScaling")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisWordCloudWordScaling>? WordScaling { get; set; }

        public AnalysisWordCloudOptionsArgs()
        {
        }
        public static new AnalysisWordCloudOptionsArgs Empty => new AnalysisWordCloudOptionsArgs();
    }
}