// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PcaConnectorAd.Inputs
{

    public sealed class TemplateKeyUsageProperty1PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("propertyFlags", required: true)]
        public Input<Inputs.TemplateKeyUsagePropertyFlagsArgs> PropertyFlags { get; set; } = null!;

        public TemplateKeyUsageProperty1PropertiesArgs()
        {
        }
        public static new TemplateKeyUsageProperty1PropertiesArgs Empty => new TemplateKeyUsageProperty1PropertiesArgs();
    }
}
