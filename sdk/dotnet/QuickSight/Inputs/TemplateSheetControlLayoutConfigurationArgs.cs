// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateSheetControlLayoutConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration that determines the elements and canvas size options of sheet control.
        /// </summary>
        [Input("gridLayout")]
        public Input<Inputs.TemplateGridLayoutConfigurationArgs>? GridLayout { get; set; }

        public TemplateSheetControlLayoutConfigurationArgs()
        {
        }
        public static new TemplateSheetControlLayoutConfigurationArgs Empty => new TemplateSheetControlLayoutConfigurationArgs();
    }
}
