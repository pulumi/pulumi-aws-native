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
    /// Node configuration in a flow
    /// </summary>
    public sealed class FlowNodeConfiguration0PropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("input", required: true)]
        public Input<Inputs.FlowInputFlowNodeConfigurationArgs> Input { get; set; } = null!;

        public FlowNodeConfiguration0PropertiesArgs()
        {
        }
        public static new FlowNodeConfiguration0PropertiesArgs Empty => new FlowNodeConfiguration0PropertiesArgs();
    }
}
