// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTablePinnedFieldOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("pinnedLeftFields")]
        private InputList<string>? _pinnedLeftFields;
        public InputList<string> PinnedLeftFields
        {
            get => _pinnedLeftFields ?? (_pinnedLeftFields = new InputList<string>());
            set => _pinnedLeftFields = value;
        }

        public TemplateTablePinnedFieldOptionsArgs()
        {
        }
        public static new TemplateTablePinnedFieldOptionsArgs Empty => new TemplateTablePinnedFieldOptionsArgs();
    }
}