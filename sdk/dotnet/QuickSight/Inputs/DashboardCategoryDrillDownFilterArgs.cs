// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardCategoryDrillDownFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryValues", required: true)]
        private InputList<string>? _categoryValues;

        /// <summary>
        /// A list of the string inputs that are the values of the category drill down filter.
        /// </summary>
        public InputList<string> CategoryValues
        {
            get => _categoryValues ?? (_categoryValues = new InputList<string>());
            set => _categoryValues = value;
        }

        /// <summary>
        /// The column that the filter is applied to.
        /// </summary>
        [Input("column", required: true)]
        public Input<Inputs.DashboardColumnIdentifierArgs> Column { get; set; } = null!;

        public DashboardCategoryDrillDownFilterArgs()
        {
        }
        public static new DashboardCategoryDrillDownFilterArgs Empty => new DashboardCategoryDrillDownFilterArgs();
    }
}
