// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisCustomActionFilterOperationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration that chooses the fields to be filtered.
        /// </summary>
        [Input("selectedFieldsConfiguration", required: true)]
        public Input<Inputs.AnalysisFilterOperationSelectedFieldsConfigurationArgs> SelectedFieldsConfiguration { get; set; } = null!;

        /// <summary>
        /// The configuration that chooses the target visuals to be filtered.
        /// </summary>
        [Input("targetVisualsConfiguration", required: true)]
        public Input<Inputs.AnalysisFilterOperationTargetVisualsConfigurationArgs> TargetVisualsConfiguration { get; set; } = null!;

        public AnalysisCustomActionFilterOperationArgs()
        {
        }
        public static new AnalysisCustomActionFilterOperationArgs Empty => new AnalysisCustomActionFilterOperationArgs();
    }
}
