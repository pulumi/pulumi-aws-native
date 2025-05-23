// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateInnerFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A `CategoryInnerFilter` filters text values for the `NestedFilter` .
        /// </summary>
        [Input("categoryInnerFilter")]
        public Input<Inputs.TemplateCategoryInnerFilterArgs>? CategoryInnerFilter { get; set; }

        public TemplateInnerFilterArgs()
        {
        }
        public static new TemplateInnerFilterArgs Empty => new TemplateInnerFilterArgs();
    }
}
