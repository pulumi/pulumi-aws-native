// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;Display options related to tiles on a sheet.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class ThemeTileStyle
    {
        /// <summary>
        /// The border around a tile.
        /// </summary>
        public readonly Outputs.ThemeBorderStyle? Border;

        [OutputConstructor]
        private ThemeTileStyle(Outputs.ThemeBorderStyle? border)
        {
            Border = border;
        }
    }
}
