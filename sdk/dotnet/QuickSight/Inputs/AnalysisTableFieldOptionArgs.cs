// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisTableFieldOptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The custom label for a table field.
        /// </summary>
        [Input("customLabel")]
        public Input<string>? CustomLabel { get; set; }

        /// <summary>
        /// The field ID for a table field.
        /// </summary>
        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        /// <summary>
        /// The URL configuration for a table field.
        /// </summary>
        [Input("urlStyling")]
        public Input<Inputs.AnalysisTableFieldUrlConfigurationArgs>? UrlStyling { get; set; }

        /// <summary>
        /// The visibility of a table field.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisVisibility>? Visibility { get; set; }

        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public AnalysisTableFieldOptionArgs()
        {
        }
        public static new AnalysisTableFieldOptionArgs Empty => new AnalysisTableFieldOptionArgs();
    }
}
