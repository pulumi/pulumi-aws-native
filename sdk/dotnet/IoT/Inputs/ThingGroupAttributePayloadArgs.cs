// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class ThingGroupAttributePayloadArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<string>? _attributes;

        /// <summary>
        /// A JSON string containing up to three key-value pair in JSON format. For example:
        /// 
        /// `{\"attributes\":{\"string1\":\"string2\"}}`
        /// </summary>
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        public ThingGroupAttributePayloadArgs()
        {
        }
        public static new ThingGroupAttributePayloadArgs Empty => new ThingGroupAttributePayloadArgs();
    }
}
