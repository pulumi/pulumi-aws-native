// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTableSortConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("paginationConfiguration")]
        public Input<Inputs.TemplatePaginationConfigurationArgs>? PaginationConfiguration { get; set; }

        [Input("rowSort")]
        private InputList<Inputs.TemplateFieldSortOptionsArgs>? _rowSort;
        public InputList<Inputs.TemplateFieldSortOptionsArgs> RowSort
        {
            get => _rowSort ?? (_rowSort = new InputList<Inputs.TemplateFieldSortOptionsArgs>());
            set => _rowSort = value;
        }

        public TemplateTableSortConfigurationArgs()
        {
        }
        public static new TemplateTableSortConfigurationArgs Empty => new TemplateTableSortConfigurationArgs();
    }
}