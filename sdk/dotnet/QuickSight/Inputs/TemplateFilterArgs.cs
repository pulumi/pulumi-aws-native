// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `CategoryFilter` filters text values.
        /// 
        /// For more information, see [Adding text filters](https://docs.aws.amazon.com/quicksight/latest/user/add-a-text-filter-data-prep.html) in the *Amazon QuickSight User Guide* .
        /// </summary>
        [Input("categoryFilter")]
        public Input<Inputs.TemplateCategoryFilterArgs>? CategoryFilter { get; set; }

        /// <summary>
        /// A `NestedFilter` filters data with a subset of data that is defined by the nested inner filter.
        /// </summary>
        [Input("nestedFilter")]
        public Input<Inputs.TemplateNestedFilterArgs>? NestedFilter { get; set; }

        /// <summary>
        /// A `NumericEqualityFilter` filters numeric values that equal or do not equal a given numeric value.
        /// </summary>
        [Input("numericEqualityFilter")]
        public Input<Inputs.TemplateNumericEqualityFilterArgs>? NumericEqualityFilter { get; set; }

        /// <summary>
        /// A `NumericRangeFilter` filters numeric values that are either inside or outside a given numeric range.
        /// </summary>
        [Input("numericRangeFilter")]
        public Input<Inputs.TemplateNumericRangeFilterArgs>? NumericRangeFilter { get; set; }

        /// <summary>
        /// A `RelativeDatesFilter` filters date values that are relative to a given date.
        /// </summary>
        [Input("relativeDatesFilter")]
        public Input<Inputs.TemplateRelativeDatesFilterArgs>? RelativeDatesFilter { get; set; }

        /// <summary>
        /// A `TimeEqualityFilter` filters date-time values that equal or do not equal a given date/time value.
        /// </summary>
        [Input("timeEqualityFilter")]
        public Input<Inputs.TemplateTimeEqualityFilterArgs>? TimeEqualityFilter { get; set; }

        /// <summary>
        /// A `TimeRangeFilter` filters date-time values that are either inside or outside a given date/time range.
        /// </summary>
        [Input("timeRangeFilter")]
        public Input<Inputs.TemplateTimeRangeFilterArgs>? TimeRangeFilter { get; set; }

        /// <summary>
        /// A `TopBottomFilter` filters data to the top or bottom values for a given column.
        /// </summary>
        [Input("topBottomFilter")]
        public Input<Inputs.TemplateTopBottomFilterArgs>? TopBottomFilter { get; set; }

        public TemplateFilterArgs()
        {
        }
        public static new TemplateFilterArgs Empty => new TemplateFilterArgs();
    }
}
