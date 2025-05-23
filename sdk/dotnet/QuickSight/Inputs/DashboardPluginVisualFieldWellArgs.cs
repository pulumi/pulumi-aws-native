// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardPluginVisualFieldWellArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The semantic axis name for the field well.
        /// </summary>
        [Input("axisName")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardPluginVisualAxisName>? AxisName { get; set; }

        [Input("dimensions")]
        private InputList<Inputs.DashboardDimensionFieldArgs>? _dimensions;

        /// <summary>
        /// A list of dimensions for the field well.
        /// </summary>
        public InputList<Inputs.DashboardDimensionFieldArgs> Dimensions
        {
            get => _dimensions ?? (_dimensions = new InputList<Inputs.DashboardDimensionFieldArgs>());
            set => _dimensions = value;
        }

        [Input("measures")]
        private InputList<Inputs.DashboardMeasureFieldArgs>? _measures;

        /// <summary>
        /// A list of measures that exist in the field well.
        /// </summary>
        public InputList<Inputs.DashboardMeasureFieldArgs> Measures
        {
            get => _measures ?? (_measures = new InputList<Inputs.DashboardMeasureFieldArgs>());
            set => _measures = value;
        }

        [Input("unaggregated")]
        private InputList<Inputs.DashboardUnaggregatedFieldArgs>? _unaggregated;

        /// <summary>
        /// A list of unaggregated fields that exist in the field well.
        /// </summary>
        public InputList<Inputs.DashboardUnaggregatedFieldArgs> Unaggregated
        {
            get => _unaggregated ?? (_unaggregated = new InputList<Inputs.DashboardUnaggregatedFieldArgs>());
            set => _unaggregated = value;
        }

        public DashboardPluginVisualFieldWellArgs()
        {
        }
        public static new DashboardPluginVisualFieldWellArgs Empty => new DashboardPluginVisualFieldWellArgs();
    }
}
