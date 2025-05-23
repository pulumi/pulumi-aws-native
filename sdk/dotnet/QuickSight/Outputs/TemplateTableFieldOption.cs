// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateTableFieldOption
    {
        /// <summary>
        /// The custom label for a table field.
        /// </summary>
        public readonly string? CustomLabel;
        /// <summary>
        /// The field ID for a table field.
        /// </summary>
        public readonly string FieldId;
        /// <summary>
        /// The URL configuration for a table field.
        /// </summary>
        public readonly Outputs.TemplateTableFieldUrlConfiguration? UrlStyling;
        /// <summary>
        /// The visibility of a table field.
        /// </summary>
        public readonly object? Visibility;
        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        public readonly string? Width;

        [OutputConstructor]
        private TemplateTableFieldOption(
            string? customLabel,

            string fieldId,

            Outputs.TemplateTableFieldUrlConfiguration? urlStyling,

            object? visibility,

            string? width)
        {
            CustomLabel = customLabel;
            FieldId = fieldId;
            UrlStyling = urlStyling;
            Visibility = visibility;
            Width = width;
        }
    }
}
