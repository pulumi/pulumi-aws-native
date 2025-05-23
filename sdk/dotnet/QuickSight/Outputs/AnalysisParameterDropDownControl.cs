// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// The values that are displayed in a control can be configured to only show values that are valid based on what's selected in other controls.
        /// </summary>
        public readonly Outputs.AnalysisCascadingControlConfiguration? CascadingControlConfiguration;
        /// <summary>
        /// The visibility configuration of the Apply button on a `ParameterDropDownControl` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisCommitMode? CommitMode;
        /// <summary>
        /// The display options of a control.
        /// </summary>
        public readonly Outputs.AnalysisDropDownControlDisplayOptions? DisplayOptions;
        /// <summary>
        /// The ID of the `ParameterDropDownControl` .
        /// </summary>
        public readonly string ParameterControlId;
        /// <summary>
        /// A list of selectable values that are used in a control.
        /// </summary>
        public readonly Outputs.AnalysisParameterSelectableValues? SelectableValues;
        /// <summary>
        /// The source parameter name of the `ParameterDropDownControl` .
        /// </summary>
        public readonly string SourceParameterName;
        /// <summary>
        /// The title of the `ParameterDropDownControl` .
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// The type parameter name of the `ParameterDropDownControl` .
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisSheetControlListType? Type;

        [OutputConstructor]
        private AnalysisParameterDropDownControl(
            Outputs.AnalysisCascadingControlConfiguration? cascadingControlConfiguration,

            Pulumi.AwsNative.QuickSight.AnalysisCommitMode? commitMode,

            Outputs.AnalysisDropDownControlDisplayOptions? displayOptions,

            string parameterControlId,

            Outputs.AnalysisParameterSelectableValues? selectableValues,

            string sourceParameterName,

            string title,

            Pulumi.AwsNative.QuickSight.AnalysisSheetControlListType? type)
        {
            CascadingControlConfiguration = cascadingControlConfiguration;
            CommitMode = commitMode;
            DisplayOptions = displayOptions;
            ParameterControlId = parameterControlId;
            SelectableValues = selectableValues;
            SourceParameterName = sourceParameterName;
            Title = title;
            Type = type;
        }
    }
}
