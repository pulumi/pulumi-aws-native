// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateDefaultFilterListControlOptions
    {
        /// <summary>
        /// The display options of a control.
        /// </summary>
        public readonly Outputs.TemplateListControlDisplayOptions? DisplayOptions;
        /// <summary>
        /// A list of selectable values that are used in a control.
        /// </summary>
        public readonly Outputs.TemplateFilterSelectableValues? SelectableValues;
        /// <summary>
        /// The type of the `DefaultFilterListControlOptions` . Choose one of the following options:
        /// 
        /// - `MULTI_SELECT` : The user can select multiple entries from the list.
        /// - `SINGLE_SELECT` : The user can select a single entry from the list.
        /// </summary>
        public readonly Pulumi.AwsNative.QuickSight.TemplateSheetControlListType? Type;

        [OutputConstructor]
        private TemplateDefaultFilterListControlOptions(
            Outputs.TemplateListControlDisplayOptions? displayOptions,

            Outputs.TemplateFilterSelectableValues? selectableValues,

            Pulumi.AwsNative.QuickSight.TemplateSheetControlListType? type)
        {
            DisplayOptions = displayOptions;
            SelectableValues = selectableValues;
            Type = type;
        }
    }
}