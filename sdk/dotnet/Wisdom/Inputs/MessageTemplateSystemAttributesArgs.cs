// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Wisdom.Inputs
{

    /// <summary>
    /// The system attributes that are used with the message template.
    /// </summary>
    public sealed class MessageTemplateSystemAttributesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CustomerEndpoint attribute.
        /// </summary>
        [Input("customerEndpoint")]
        public Input<Inputs.MessageTemplateSystemEndpointAttributesArgs>? CustomerEndpoint { get; set; }

        /// <summary>
        /// The name of the task.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The SystemEndpoint attribute.
        /// </summary>
        [Input("systemEndpoint")]
        public Input<Inputs.MessageTemplateSystemEndpointAttributesArgs>? SystemEndpoint { get; set; }

        public MessageTemplateSystemAttributesArgs()
        {
        }
        public static new MessageTemplateSystemAttributesArgs Empty => new MessageTemplateSystemAttributesArgs();
    }
}
