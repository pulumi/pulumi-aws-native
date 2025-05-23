// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardAxisScaleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The linear axis scale setup.
        /// </summary>
        [Input("linear")]
        public Input<Inputs.DashboardAxisLinearScaleArgs>? Linear { get; set; }

        /// <summary>
        /// The logarithmic axis scale setup.
        /// </summary>
        [Input("logarithmic")]
        public Input<Inputs.DashboardAxisLogarithmicScaleArgs>? Logarithmic { get; set; }

        public DashboardAxisScaleArgs()
        {
        }
        public static new DashboardAxisScaleArgs Empty => new DashboardAxisScaleArgs();
    }
}
