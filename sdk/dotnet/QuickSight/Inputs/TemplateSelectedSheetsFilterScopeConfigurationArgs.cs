// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateSelectedSheetsFilterScopeConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("sheetVisualScopingConfigurations")]
        private InputList<Inputs.TemplateSheetVisualScopingConfigurationArgs>? _sheetVisualScopingConfigurations;

        /// <summary>
        /// The sheet ID and visual IDs of the sheet and visuals that the filter is applied to.
        /// </summary>
        public InputList<Inputs.TemplateSheetVisualScopingConfigurationArgs> SheetVisualScopingConfigurations
        {
            get => _sheetVisualScopingConfigurations ?? (_sheetVisualScopingConfigurations = new InputList<Inputs.TemplateSheetVisualScopingConfigurationArgs>());
            set => _sheetVisualScopingConfigurations = value;
        }

        public TemplateSelectedSheetsFilterScopeConfigurationArgs()
        {
        }
        public static new TemplateSelectedSheetsFilterScopeConfigurationArgs Empty => new TemplateSelectedSheetsFilterScopeConfigurationArgs();
    }
}
