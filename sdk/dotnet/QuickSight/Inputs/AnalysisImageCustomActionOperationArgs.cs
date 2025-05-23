// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisImageCustomActionOperationArgs : global::Pulumi.ResourceArgs
    {
        [Input("navigationOperation")]
        public Input<Inputs.AnalysisCustomActionNavigationOperationArgs>? NavigationOperation { get; set; }

        [Input("setParametersOperation")]
        public Input<Inputs.AnalysisCustomActionSetParametersOperationArgs>? SetParametersOperation { get; set; }

        [Input("urlOperation")]
        public Input<Inputs.AnalysisCustomActionUrlOperationArgs>? UrlOperation { get; set; }

        public AnalysisImageCustomActionOperationArgs()
        {
        }
        public static new AnalysisImageCustomActionOperationArgs Empty => new AnalysisImageCustomActionOperationArgs();
    }
}
