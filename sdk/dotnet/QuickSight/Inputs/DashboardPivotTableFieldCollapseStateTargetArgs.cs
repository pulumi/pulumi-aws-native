// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPivotTableFieldCollapseStateTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("fieldDataPathValues")]
        private InputList<Inputs.DashboardDataPathValueArgs>? _fieldDataPathValues;

        /// <summary>
        /// The data path of the pivot table's header. Used to set the collapse state.
        /// </summary>
        public InputList<Inputs.DashboardDataPathValueArgs> FieldDataPathValues
        {
            get => _fieldDataPathValues ?? (_fieldDataPathValues = new InputList<Inputs.DashboardDataPathValueArgs>());
            set => _fieldDataPathValues = value;
        }

        /// <summary>
        /// The field ID of the pivot table that the collapse state needs to be set to.
        /// </summary>
        [Input("fieldId")]
        public Input<string>? FieldId { get; set; }

        public DashboardPivotTableFieldCollapseStateTargetArgs()
        {
        }
        public static new DashboardPivotTableFieldCollapseStateTargetArgs Empty => new DashboardPivotTableFieldCollapseStateTargetArgs();
    }
}
