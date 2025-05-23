// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisVisualInteractionOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The context menu options for a visual.
        /// </summary>
        [Input("contextMenuOption")]
        public Input<Inputs.AnalysisContextMenuOptionArgs>? ContextMenuOption { get; set; }

        /// <summary>
        /// The on-visual menu options for a visual.
        /// </summary>
        [Input("visualMenuOption")]
        public Input<Inputs.AnalysisVisualMenuOptionArgs>? VisualMenuOption { get; set; }

        public AnalysisVisualInteractionOptionsArgs()
        {
        }
        public static new AnalysisVisualInteractionOptionsArgs Empty => new AnalysisVisualInteractionOptionsArgs();
    }
}
