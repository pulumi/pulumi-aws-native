// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisImageInteractionOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The menu options for the image.
        /// </summary>
        [Input("imageMenuOption")]
        public Input<Inputs.AnalysisImageMenuOptionArgs>? ImageMenuOption { get; set; }

        public AnalysisImageInteractionOptionsArgs()
        {
        }
        public static new AnalysisImageInteractionOptionsArgs Empty => new AnalysisImageInteractionOptionsArgs();
    }
}
