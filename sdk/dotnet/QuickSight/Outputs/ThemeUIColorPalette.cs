// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The theme colors that apply to UI and to charts, excluding data colors. The colors
    ///             description is a hexadecimal color code that consists of six alphanumerical characters,
    ///             prefixed with &lt;code&gt;#&lt;/code&gt;, for example #37BFF5. For more information, see &lt;a href="https://docs.aws.amazon.com/quicksight/latest/user/themes-in-quicksight.html"&gt;Using Themes in Amazon QuickSight&lt;/a&gt; in the &lt;i&gt;Amazon QuickSight User
    ///                 Guide.&lt;/i&gt;
    ///         &lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class ThemeUIColorPalette
    {
        /// <summary>
        /// &lt;p&gt;This color is that applies to selected states and buttons.&lt;/p&gt;
        /// </summary>
        public readonly string? Accent;
        /// <summary>
        /// &lt;p&gt;The foreground color that applies to any text or other elements that appear over the
        ///             accent color.&lt;/p&gt;
        /// </summary>
        public readonly string? AccentForeground;
        /// <summary>
        /// &lt;p&gt;The color that applies to error messages.&lt;/p&gt;
        /// </summary>
        public readonly string? Danger;
        /// <summary>
        /// &lt;p&gt;The foreground color that applies to any text or other elements that appear over the
        ///             error color.&lt;/p&gt;
        /// </summary>
        public readonly string? DangerForeground;
        /// <summary>
        /// &lt;p&gt;The color that applies to the names of fields that are identified as
        ///             dimensions.&lt;/p&gt;
        /// </summary>
        public readonly string? Dimension;
        /// <summary>
        /// &lt;p&gt;The foreground color that applies to any text or other elements that appear over the
        ///             dimension color.&lt;/p&gt;
        /// </summary>
        public readonly string? DimensionForeground;
        /// <summary>
        /// &lt;p&gt;The color that applies to the names of fields that are identified as measures.&lt;/p&gt;
        /// </summary>
        public readonly string? Measure;
        /// <summary>
        /// &lt;p&gt;The foreground color that applies to any text or other elements that appear over the
        ///             measure color.&lt;/p&gt;
        /// </summary>
        public readonly string? MeasureForeground;
        /// <summary>
        /// &lt;p&gt;The background color that applies to visuals and other high emphasis UI.&lt;/p&gt;
        /// </summary>
        public readonly string? PrimaryBackground;
        /// <summary>
        /// &lt;p&gt;The color of text and other foreground elements that appear over the primary
        ///             background regions, such as grid lines, borders, table banding, icons, and so on.&lt;/p&gt;
        /// </summary>
        public readonly string? PrimaryForeground;
        /// <summary>
        /// &lt;p&gt;The background color that applies to the sheet background and sheet controls.&lt;/p&gt;
        /// </summary>
        public readonly string? SecondaryBackground;
        /// <summary>
        /// &lt;p&gt;The foreground color that applies to any sheet title, sheet control text, or UI that
        ///             appears over the secondary background.&lt;/p&gt;
        /// </summary>
        public readonly string? SecondaryForeground;
        /// <summary>
        /// &lt;p&gt;The color that applies to success messages, for example the check mark for a
        ///             successful download.&lt;/p&gt;
        /// </summary>
        public readonly string? Success;
        /// <summary>
        /// &lt;p&gt;The foreground color that applies to any text or other elements that appear over the
        ///             success color.&lt;/p&gt;
        /// </summary>
        public readonly string? SuccessForeground;
        /// <summary>
        /// &lt;p&gt;This color that applies to warning and informational messages.&lt;/p&gt;
        /// </summary>
        public readonly string? Warning;
        /// <summary>
        /// &lt;p&gt;The foreground color that applies to any text or other elements that appear over the
        ///             warning color.&lt;/p&gt;
        /// </summary>
        public readonly string? WarningForeground;

        [OutputConstructor]
        private ThemeUIColorPalette(
            string? accent,

            string? accentForeground,

            string? danger,

            string? dangerForeground,

            string? dimension,

            string? dimensionForeground,

            string? measure,

            string? measureForeground,

            string? primaryBackground,

            string? primaryForeground,

            string? secondaryBackground,

            string? secondaryForeground,

            string? success,

            string? successForeground,

            string? warning,

            string? warningForeground)
        {
            Accent = accent;
            AccentForeground = accentForeground;
            Danger = danger;
            DangerForeground = dangerForeground;
            Dimension = dimension;
            DimensionForeground = dimensionForeground;
            Measure = measure;
            MeasureForeground = measureForeground;
            PrimaryBackground = primaryBackground;
            PrimaryForeground = primaryForeground;
            SecondaryBackground = secondaryBackground;
            SecondaryForeground = secondaryForeground;
            Success = success;
            SuccessForeground = successForeground;
            Warning = warning;
            WarningForeground = warningForeground;
        }
    }
}