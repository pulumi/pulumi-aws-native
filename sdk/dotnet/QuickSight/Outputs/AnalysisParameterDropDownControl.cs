// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisParameterDropDownControl
    {
        public readonly Outputs.AnalysisCascadingControlConfiguration? CascadingControlConfiguration;
        public readonly Outputs.AnalysisDropDownControlDisplayOptions? DisplayOptions;
        public readonly string ParameterControlId;
        public readonly Outputs.AnalysisParameterSelectableValues? SelectableValues;
        public readonly string SourceParameterName;
        public readonly string Title;
        public readonly Pulumi.AwsNative.QuickSight.AnalysisSheetControlListType? Type;

        [OutputConstructor]
        private AnalysisParameterDropDownControl(
            Outputs.AnalysisCascadingControlConfiguration? cascadingControlConfiguration,

            Outputs.AnalysisDropDownControlDisplayOptions? displayOptions,

            string parameterControlId,

            Outputs.AnalysisParameterSelectableValues? selectableValues,

            string sourceParameterName,

            string title,

            Pulumi.AwsNative.QuickSight.AnalysisSheetControlListType? type)
        {
            CascadingControlConfiguration = cascadingControlConfiguration;
            DisplayOptions = displayOptions;
            ParameterControlId = parameterControlId;
            SelectableValues = selectableValues;
            SourceParameterName = sourceParameterName;
            Title = title;
            Type = type;
        }
    }
}