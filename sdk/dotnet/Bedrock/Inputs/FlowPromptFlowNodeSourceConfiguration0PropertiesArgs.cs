// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Bedrock.Inputs
{

    /// <summary>
    /// Prompt source configuration for prompt node
    /// </summary>
    public sealed class FlowPromptFlowNodeSourceConfiguration0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("resource", required: true)]
        public Input<Inputs.FlowPromptFlowNodeResourceConfigurationArgs> Resource { get; set; } = null!;

        public FlowPromptFlowNodeSourceConfiguration0PropertiesArgs()
        {
        }
        public static new FlowPromptFlowNodeSourceConfiguration0PropertiesArgs Empty => new FlowPromptFlowNodeSourceConfiguration0PropertiesArgs();
    }
}
