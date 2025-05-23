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
    public sealed class TemplateFieldSortOptions
    {
        /// <summary>
        /// The sort configuration for a column that is not used in a field well.
        /// </summary>
        public readonly Outputs.TemplateColumnSort? ColumnSort;
        /// <summary>
        /// The sort configuration for a field in a field well.
        /// </summary>
        public readonly Outputs.TemplateFieldSort? FieldSort;

        [OutputConstructor]
        private TemplateFieldSortOptions(
            Outputs.TemplateColumnSort? columnSort,

            Outputs.TemplateFieldSort? fieldSort)
        {
            ColumnSort = columnSort;
            FieldSort = fieldSort;
        }
    }
}
