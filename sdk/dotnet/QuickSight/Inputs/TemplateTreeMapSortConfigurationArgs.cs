// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTreeMapSortConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The limit on the number of groups that are displayed.
        /// </summary>
        [Input("treeMapGroupItemsLimitConfiguration")]
        public Input<Inputs.TemplateItemsLimitConfigurationArgs>? TreeMapGroupItemsLimitConfiguration { get; set; }

        [Input("treeMapSort")]
        private InputList<Inputs.TemplateFieldSortOptionsArgs>? _treeMapSort;

        /// <summary>
        /// The sort configuration of group by fields.
        /// </summary>
        public InputList<Inputs.TemplateFieldSortOptionsArgs> TreeMapSort
        {
            get => _treeMapSort ?? (_treeMapSort = new InputList<Inputs.TemplateFieldSortOptionsArgs>());
            set => _treeMapSort = value;
        }

        public TemplateTreeMapSortConfigurationArgs()
        {
        }
        public static new TemplateTreeMapSortConfigurationArgs Empty => new TemplateTreeMapSortConfigurationArgs();
    }
}
