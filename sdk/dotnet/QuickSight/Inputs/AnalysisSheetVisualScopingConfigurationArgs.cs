// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisSheetVisualScopingConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("scope", required: true)]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisFilterVisualScope> Scope { get; set; } = null!;

        [Input("sheetId", required: true)]
        public Input<string> SheetId { get; set; } = null!;

        [Input("visualIds")]
        private InputList<string>? _visualIds;
        public InputList<string> VisualIds
        {
            get => _visualIds ?? (_visualIds = new InputList<string>());
            set => _visualIds = value;
        }

        public AnalysisSheetVisualScopingConfigurationArgs()
        {
        }
        public static new AnalysisSheetVisualScopingConfigurationArgs Empty => new AnalysisSheetVisualScopingConfigurationArgs();
    }
}