// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateWordCloudSortConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("categoryItemsLimit")]
        public Input<Inputs.TemplateItemsLimitConfigurationArgs>? CategoryItemsLimit { get; set; }

        [Input("categorySort")]
        private InputList<Inputs.TemplateFieldSortOptionsArgs>? _categorySort;
        public InputList<Inputs.TemplateFieldSortOptionsArgs> CategorySort
        {
            get => _categorySort ?? (_categorySort = new InputList<Inputs.TemplateFieldSortOptionsArgs>());
            set => _categorySort = value;
        }

        public TemplateWordCloudSortConfigurationArgs()
        {
        }
        public static new TemplateWordCloudSortConfigurationArgs Empty => new TemplateWordCloudSortConfigurationArgs();
    }
}