// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QBusiness.Inputs
{

    public sealed class PluginCustomPluginConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains either details about the S3 object containing the OpenAPI schema for the action group or the JSON or YAML-formatted payload defining the schema.
        /// </summary>
        [Input("apiSchema", required: true)]
        public InputUnion<Inputs.PluginApiSchema0PropertiesArgs, Inputs.PluginApiSchema1PropertiesArgs> ApiSchema { get; set; } = null!;

        /// <summary>
        /// The type of OpenAPI schema to use.
        /// </summary>
        [Input("apiSchemaType", required: true)]
        public Input<Pulumi.AwsNative.QBusiness.PluginApiSchemaType> ApiSchemaType { get; set; } = null!;

        /// <summary>
        /// A description for your custom plugin configuration.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        public PluginCustomPluginConfigurationArgs()
        {
        }
        public static new PluginCustomPluginConfigurationArgs Empty => new PluginCustomPluginConfigurationArgs();
    }
}
