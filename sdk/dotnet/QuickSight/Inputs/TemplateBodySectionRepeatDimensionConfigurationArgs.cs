// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateBodySectionRepeatDimensionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the *Category* dataset column and constraints around the dynamic values that will be used in repeating the section contents.
        /// </summary>
        [Input("dynamicCategoryDimensionConfiguration")]
        public Input<Inputs.TemplateBodySectionDynamicCategoryDimensionConfigurationArgs>? DynamicCategoryDimensionConfiguration { get; set; }

        /// <summary>
        /// Describes the *Numeric* dataset column and constraints around the dynamic values used to repeat the contents of a section.
        /// </summary>
        [Input("dynamicNumericDimensionConfiguration")]
        public Input<Inputs.TemplateBodySectionDynamicNumericDimensionConfigurationArgs>? DynamicNumericDimensionConfiguration { get; set; }

        public TemplateBodySectionRepeatDimensionConfigurationArgs()
        {
        }
        public static new TemplateBodySectionRepeatDimensionConfigurationArgs Empty => new TemplateBodySectionRepeatDimensionConfigurationArgs();
    }
}
