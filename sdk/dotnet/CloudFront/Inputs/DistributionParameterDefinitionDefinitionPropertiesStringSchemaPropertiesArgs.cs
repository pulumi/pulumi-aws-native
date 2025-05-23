// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class DistributionParameterDefinitionDefinitionPropertiesStringSchemaPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        [Input("required", required: true)]
        public Input<bool> Required { get; set; } = null!;

        public DistributionParameterDefinitionDefinitionPropertiesStringSchemaPropertiesArgs()
        {
        }
        public static new DistributionParameterDefinitionDefinitionPropertiesStringSchemaPropertiesArgs Empty => new DistributionParameterDefinitionDefinitionPropertiesStringSchemaPropertiesArgs();
    }
}
