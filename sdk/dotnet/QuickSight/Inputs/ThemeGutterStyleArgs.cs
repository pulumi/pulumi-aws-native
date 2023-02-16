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
    /// &lt;p&gt;The display options for gutter spacing between tiles on a sheet.&lt;/p&gt;
    /// </summary>
    public sealed class ThemeGutterStyleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;This Boolean value controls whether to display a gutter space between sheet tiles.
        ///         &lt;/p&gt;
        /// </summary>
        [Input("show")]
        public Input<bool>? Show { get; set; }

        public ThemeGutterStyleArgs()
        {
        }
        public static new ThemeGutterStyleArgs Empty => new ThemeGutterStyleArgs();
    }
}