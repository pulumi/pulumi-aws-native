// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateFieldSortOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sort configuration for a column that is not used in a field well.
        /// </summary>
        [Input("columnSort")]
        public Input<Inputs.TemplateColumnSortArgs>? ColumnSort { get; set; }

        /// <summary>
        /// The sort configuration for a field in a field well.
        /// </summary>
        [Input("fieldSort")]
        public Input<Inputs.TemplateFieldSortArgs>? FieldSort { get; set; }

        public TemplateFieldSortOptionsArgs()
        {
        }
        public static new TemplateFieldSortOptionsArgs Empty => new TemplateFieldSortOptionsArgs();
    }
}
