// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateTableFieldUrlConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The image configuration of a table field URL.
        /// </summary>
        [Input("imageConfiguration")]
        public Input<Inputs.TemplateTableFieldImageConfigurationArgs>? ImageConfiguration { get; set; }

        /// <summary>
        /// The link configuration of a table field URL.
        /// </summary>
        [Input("linkConfiguration")]
        public Input<Inputs.TemplateTableFieldLinkConfigurationArgs>? LinkConfiguration { get; set; }

        public TemplateTableFieldUrlConfigurationArgs()
        {
        }
        public static new TemplateTableFieldUrlConfigurationArgs Empty => new TemplateTableFieldUrlConfigurationArgs();
    }
}
