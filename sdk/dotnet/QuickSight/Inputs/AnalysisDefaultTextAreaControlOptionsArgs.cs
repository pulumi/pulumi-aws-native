// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisDefaultTextAreaControlOptionsArgs : global::Pulumi.ResourceArgs
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
        public Input<Inputs.AnalysisTextAreaControlDisplayOptionsArgs>? DisplayOptions { get; set; }

        public AnalysisDefaultTextAreaControlOptionsArgs()
        {
        }
        public static new AnalysisDefaultTextAreaControlOptionsArgs Empty => new AnalysisDefaultTextAreaControlOptionsArgs();
    }
}