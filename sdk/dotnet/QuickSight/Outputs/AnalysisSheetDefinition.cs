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
    public sealed class AnalysisSheetDefinition
    {
        /// <summary>
        /// The layout content type of the sheet. Choose one of the following options:
        /// 
        /// - `PAGINATED` : Creates a sheet for a paginated report.
        /// - `INTERACTIVE` : Creates a sheet for an interactive dashboard.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.AnalysisSheetContentType? ContentType;
        /// <summary>
        /// A description of the sheet.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The list of filter controls that are on a sheet.
        /// 
        /// For more information, see [Adding filter controls to analysis sheets](https://docs.aws.amazon.com/quicksight/latest/user/filter-controls.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisFilterControl> FilterControls;
        /// <summary>
        /// A list of images on a sheet.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisSheetImage> Images;
        /// <summary>
        /// Layouts define how the components of a sheet are arranged.
        /// 
        /// For more information, see [Types of layout](https://docs.aws.amazon.com/quicksight/latest/user/types-of-layout.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisLayout> Layouts;
        /// <summary>
        /// The name of the sheet. This name is displayed on the sheet's tab in the Amazon QuickSight console.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The list of parameter controls that are on a sheet.
        /// 
        /// For more information, see [Using a Control with a Parameter in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-controls.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisParameterControl> ParameterControls;
        /// <summary>
        /// The control layouts of the sheet.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisSheetControlLayout> SheetControlLayouts;
        /// <summary>
        /// The unique identifier of a sheet.
        /// </summary>
        public readonly string SheetId;
        /// <summary>
        /// The text boxes that are on a sheet.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisSheetTextBox> TextBoxes;
        /// <summary>
        /// The title of the sheet.
        /// </summary>
        public readonly string? Title;
        /// <summary>
        /// A list of the visuals that are on a sheet. Visual placement is determined by the layout of the sheet.
        /// </summary>
        public readonly ImmutableArray<Outputs.AnalysisVisual> Visuals;

        [OutputConstructor]
        private AnalysisSheetDefinition(
            Pulumi.AwsNative.QuickSight.AnalysisSheetContentType? contentType,

            string? description,

            ImmutableArray<Outputs.AnalysisFilterControl> filterControls,

            ImmutableArray<Outputs.AnalysisSheetImage> images,

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
            Images = images;
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
