// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDefaultTextAreaControlOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The delimiter that is used to separate the lines in text.
        /// </summary>
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        /// <summary>
        /// The display options of a control.
        /// </summary>
        [Input("displayOptions")]
        public Input<Inputs.DashboardTextAreaControlDisplayOptionsArgs>? DisplayOptions { get; set; }

        public DashboardDefaultTextAreaControlOptionsArgs()
        {
        }
        public static new DashboardDefaultTextAreaControlOptionsArgs Empty => new DashboardDefaultTextAreaControlOptionsArgs();
    }
}