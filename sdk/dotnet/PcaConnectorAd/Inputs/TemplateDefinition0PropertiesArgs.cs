// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PcaConnectorAd.Inputs
{

    public sealed class TemplateDefinition0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("templateV2", required: true)]
        public Input<Inputs.TemplateV2Args> TemplateV2 { get; set; } = null!;

        public TemplateDefinition0PropertiesArgs()
        {
        }
        public static new TemplateDefinition0PropertiesArgs Empty => new TemplateDefinition0PropertiesArgs();
    }
}
