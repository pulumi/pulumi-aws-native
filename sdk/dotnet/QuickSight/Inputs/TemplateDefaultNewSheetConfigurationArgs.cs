// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateDefaultNewSheetConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("interactiveLayoutConfiguration")]
        public Input<Inputs.TemplateDefaultInteractiveLayoutConfigurationArgs>? InteractiveLayoutConfiguration { get; set; }

        [Input("paginatedLayoutConfiguration")]
        public Input<Inputs.TemplateDefaultPaginatedLayoutConfigurationArgs>? PaginatedLayoutConfiguration { get; set; }

        [Input("sheetContentType")]
        public Input<Pulumi.AwsNative.QuickSight.TemplateSheetContentType>? SheetContentType { get; set; }

        public TemplateDefaultNewSheetConfigurationArgs()
        {
        }
        public static new TemplateDefaultNewSheetConfigurationArgs Empty => new TemplateDefaultNewSheetConfigurationArgs();
    }
}