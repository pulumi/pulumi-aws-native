// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardSheetDefinitionArgs : global::Pulumi.ResourceArgs
    {
        [Input("contentType")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardSheetContentType>? ContentType { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("filterControls")]
        private InputList<Inputs.DashboardFilterControlArgs>? _filterControls;
        public InputList<Inputs.DashboardFilterControlArgs> FilterControls
        {
            get => _filterControls ?? (_filterControls = new InputList<Inputs.DashboardFilterControlArgs>());
            set => _filterControls = value;
        }

        [Input("layouts")]
        private InputList<Inputs.DashboardLayoutArgs>? _layouts;
        public InputList<Inputs.DashboardLayoutArgs> Layouts
        {
            get => _layouts ?? (_layouts = new InputList<Inputs.DashboardLayoutArgs>());
            set => _layouts = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameterControls")]
        private InputList<Inputs.DashboardParameterControlArgs>? _parameterControls;
        public InputList<Inputs.DashboardParameterControlArgs> ParameterControls
        {
            get => _parameterControls ?? (_parameterControls = new InputList<Inputs.DashboardParameterControlArgs>());
            set => _parameterControls = value;
        }

        [Input("sheetControlLayouts")]
        private InputList<Inputs.DashboardSheetControlLayoutArgs>? _sheetControlLayouts;
        public InputList<Inputs.DashboardSheetControlLayoutArgs> SheetControlLayouts
        {
            get => _sheetControlLayouts ?? (_sheetControlLayouts = new InputList<Inputs.DashboardSheetControlLayoutArgs>());
            set => _sheetControlLayouts = value;
        }

        [Input("sheetId", required: true)]
        public Input<string> SheetId { get; set; } = null!;

        [Input("textBoxes")]
        private InputList<Inputs.DashboardSheetTextBoxArgs>? _textBoxes;
        public InputList<Inputs.DashboardSheetTextBoxArgs> TextBoxes
        {
            get => _textBoxes ?? (_textBoxes = new InputList<Inputs.DashboardSheetTextBoxArgs>());
            set => _textBoxes = value;
        }

        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("visuals")]
        private InputList<Inputs.DashboardVisualArgs>? _visuals;
        public InputList<Inputs.DashboardVisualArgs> Visuals
        {
            get => _visuals ?? (_visuals = new InputList<Inputs.DashboardVisualArgs>());
            set => _visuals = value;
        }

        public DashboardSheetDefinitionArgs()
        {
        }
        public static new DashboardSheetDefinitionArgs Empty => new DashboardSheetDefinitionArgs();
    }
}