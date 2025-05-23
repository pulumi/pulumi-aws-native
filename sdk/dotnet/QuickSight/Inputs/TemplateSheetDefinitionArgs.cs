// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateSheetDefinitionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The layout content type of the sheet. Choose one of the following options:
        /// 
        /// - `PAGINATED` : Creates a sheet for a paginated report.
        /// - `INTERACTIVE` : Creates a sheet for an interactive dashboard.
        /// </summary>
        [Input("contentType")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateSheetContentType>? ContentType { get; set; }

        /// <summary>
        /// A description of the sheet.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("filterControls")]
        private InputList<Inputs.TemplateFilterControlArgs>? _filterControls;

        /// <summary>
        /// The list of filter controls that are on a sheet.
        /// 
        /// For more information, see [Adding filter controls to analysis sheets](https://docs.aws.amazon.com/quicksight/latest/user/filter-controls.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public InputList<Inputs.TemplateFilterControlArgs> FilterControls
        {
            get => _filterControls ?? (_filterControls = new InputList<Inputs.TemplateFilterControlArgs>());
            set => _filterControls = value;
        }

        [Input("images")]
        private InputList<Inputs.TemplateSheetImageArgs>? _images;

        /// <summary>
        /// A list of images on a sheet.
        /// </summary>
        public InputList<Inputs.TemplateSheetImageArgs> Images
        {
            get => _images ?? (_images = new InputList<Inputs.TemplateSheetImageArgs>());
            set => _images = value;
        }

        [Input("layouts")]
        private InputList<Inputs.TemplateLayoutArgs>? _layouts;

        /// <summary>
        /// Layouts define how the components of a sheet are arranged.
        /// 
        /// For more information, see [Types of layout](https://docs.aws.amazon.com/quicksight/latest/user/types-of-layout.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public InputList<Inputs.TemplateLayoutArgs> Layouts
        {
            get => _layouts ?? (_layouts = new InputList<Inputs.TemplateLayoutArgs>());
            set => _layouts = value;
        }

        /// <summary>
        /// The name of the sheet. This name is displayed on the sheet's tab in the Amazon QuickSight console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameterControls")]
        private InputList<Inputs.TemplateParameterControlArgs>? _parameterControls;

        /// <summary>
        /// The list of parameter controls that are on a sheet.
        /// 
        /// For more information, see [Using a Control with a Parameter in Amazon QuickSight](https://docs.aws.amazon.com/quicksight/latest/user/parameters-controls.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        public InputList<Inputs.TemplateParameterControlArgs> ParameterControls
        {
            get => _parameterControls ?? (_parameterControls = new InputList<Inputs.TemplateParameterControlArgs>());
            set => _parameterControls = value;
        }

        [Input("sheetControlLayouts")]
        private InputList<Inputs.TemplateSheetControlLayoutArgs>? _sheetControlLayouts;

        /// <summary>
        /// The control layouts of the sheet.
        /// </summary>
        public InputList<Inputs.TemplateSheetControlLayoutArgs> SheetControlLayouts
        {
            get => _sheetControlLayouts ?? (_sheetControlLayouts = new InputList<Inputs.TemplateSheetControlLayoutArgs>());
            set => _sheetControlLayouts = value;
        }

        /// <summary>
        /// The unique identifier of a sheet.
        /// </summary>
        [Input("sheetId", required: true)]
        public Input<string> SheetId { get; set; } = null!;

        [Input("textBoxes")]
        private InputList<Inputs.TemplateSheetTextBoxArgs>? _textBoxes;

        /// <summary>
        /// The text boxes that are on a sheet.
        /// </summary>
        public InputList<Inputs.TemplateSheetTextBoxArgs> TextBoxes
        {
            get => _textBoxes ?? (_textBoxes = new InputList<Inputs.TemplateSheetTextBoxArgs>());
            set => _textBoxes = value;
        }

        /// <summary>
        /// The title of the sheet.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("visuals")]
        private InputList<Inputs.TemplateVisualArgs>? _visuals;

        /// <summary>
        /// A list of the visuals that are on a sheet. Visual placement is determined by the layout of the sheet.
        /// </summary>
        public InputList<Inputs.TemplateVisualArgs> Visuals
        {
            get => _visuals ?? (_visuals = new InputList<Inputs.TemplateVisualArgs>());
            set => _visuals = value;
        }

        public TemplateSheetDefinitionArgs()
        {
        }
        public static new TemplateSheetDefinitionArgs Empty => new TemplateSheetDefinitionArgs();
    }
}
