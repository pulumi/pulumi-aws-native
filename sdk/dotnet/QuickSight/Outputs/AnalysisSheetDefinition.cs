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
    public sealed class AnalysisSheetDefinition
    {
        public readonly Pulumi.AwsNative.QuickSight.AnalysisSheetContentType? ContentType;
        public readonly string? Description;
        public readonly ImmutableArray<Outputs.AnalysisFilterControl> FilterControls;
        public readonly ImmutableArray<Outputs.AnalysisLayout> Layouts;
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.AnalysisParameterControl> ParameterControls;
        public readonly ImmutableArray<Outputs.AnalysisSheetControlLayout> SheetControlLayouts;
        public readonly string SheetId;
        public readonly ImmutableArray<Outputs.AnalysisSheetTextBox> TextBoxes;
        public readonly string? Title;
        public readonly ImmutableArray<Outputs.AnalysisVisual> Visuals;

        [OutputConstructor]
        private AnalysisSheetDefinition(
            Pulumi.AwsNative.QuickSight.AnalysisSheetContentType? contentType,

            string? description,

            ImmutableArray<Outputs.AnalysisFilterControl> filterControls,

            ImmutableArray<Outputs.AnalysisLayout> layouts,

            string? name,

            ImmutableArray<Outputs.AnalysisParameterControl> parameterControls,

            ImmutableArray<Outputs.AnalysisSheetControlLayout> sheetControlLayouts,

            string sheetId,

            ImmutableArray<Outputs.AnalysisSheetTextBox> textBoxes,

            string? title,

            ImmutableArray<Outputs.AnalysisVisual> visuals)
        {
            ContentType = contentType;
            Description = description;
            FilterControls = filterControls;
            Layouts = layouts;
            Name = name;
            ParameterControls = parameterControls;
            SheetControlLayouts = sheetControlLayouts;
            SheetId = sheetId;
            TextBoxes = textBoxes;
            Title = title;
            Visuals = visuals;
        }
    }
}