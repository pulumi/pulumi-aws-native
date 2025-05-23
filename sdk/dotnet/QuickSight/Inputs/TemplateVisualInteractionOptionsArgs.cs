// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateVisualInteractionOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The context menu options for a visual.
        /// </summary>
        [Input("contextMenuOption")]
        public Input<Inputs.TemplateContextMenuOptionArgs>? ContextMenuOption { get; set; }

        /// <summary>
        /// The on-visual menu options for a visual.
        /// </summary>
        [Input("visualMenuOption")]
        public Input<Inputs.TemplateVisualMenuOptionArgs>? VisualMenuOption { get; set; }

        public TemplateVisualInteractionOptionsArgs()
        {
        }
        public static new TemplateVisualInteractionOptionsArgs Empty => new TemplateVisualInteractionOptionsArgs();
    }
}
