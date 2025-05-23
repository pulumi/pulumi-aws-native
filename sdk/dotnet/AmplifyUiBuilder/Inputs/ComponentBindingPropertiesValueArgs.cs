// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class ComponentBindingPropertiesValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the properties to customize with data at runtime.
        /// </summary>
        [Input("bindingProperties")]
        public Input<Inputs.ComponentBindingPropertiesValuePropertiesArgs>? BindingProperties { get; set; }

        /// <summary>
        /// The default value of the property.
        /// </summary>
        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        /// <summary>
        /// The property type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ComponentBindingPropertiesValueArgs()
        {
        }
        public static new ComponentBindingPropertiesValueArgs Empty => new ComponentBindingPropertiesValueArgs();
    }
}
