// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;The theme display options for sheets. &lt;/p&gt;
    /// </summary>
    public sealed class ThemeSheetStyleArgs : global::Pulumi.ResourceArgs
    {
        [Input("tile")]
        public Input<Inputs.ThemeTileStyleArgs>? Tile { get; set; }

        [Input("tileLayout")]
        public Input<Inputs.ThemeTileLayoutStyleArgs>? TileLayout { get; set; }

        public ThemeSheetStyleArgs()
        {
        }
        public static new ThemeSheetStyleArgs Empty => new ThemeSheetStyleArgs();
    }
}